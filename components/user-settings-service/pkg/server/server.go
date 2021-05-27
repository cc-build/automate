package server

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/chef/automate/api/external/lib/errorutils"

	pb "github.com/golang/protobuf/ptypes/empty"

	automate_event "github.com/chef/automate/api/interservice/event"
	"github.com/chef/automate/api/interservice/user_settings"
	"github.com/chef/automate/components/compliance-service/inspec-agent/scheduler"
	"github.com/chef/automate/components/user-settings-service/pkg/storage"
	"github.com/chef/automate/components/user-settings-service/pkg/storage/postgres"
	"github.com/chef/automate/lib/grpc/health"
	"github.com/chef/automate/lib/grpc/secureconn"
	"github.com/sirupsen/logrus"
)

var empty = pb.Empty{}

//UserSettingsServer is a map of user Id to map of
type UserSettingsServer struct {
	db            *postgres.DB
	storageClient storage.Client
	health        *health.Service
	mutex         sync.Mutex
	vals          map[string]map[string]*user_settings.UserSettingValue
}

type Server struct {
	db              *postgres.DB
	connFactory     *secureconn.Factory
	eventsClient    automate_event.EventServiceClient
	schedulerServer *scheduler.Scheduler
}

func New(db *postgres.DB, storage storage.Client) *UserSettingsServer {
	return &UserSettingsServer{
		db:            db,
		storageClient: storage,
		health:        health.NewService(),
		vals:          make(map[string]map[string]*user_settings.UserSettingValue),
	}
}

func (s *UserSettingsServer) GetUserSettings(ctx context.Context, req *user_settings.GetUserSettingsRequest) (*user_settings.GetUserSettingsResponse, error) {
	//first try to get the default settings from the database.
	//all user settings that we support will have an entry in the db held in user: _default connector: local
	getDefaultUserSettingsResp, err := s.db.GetUserSettings("_default", "local")
	logrus.Debugf("GetUserSettings for _default user - response: %v", getDefaultUserSettingsResp)
	if err != nil {
		return nil, errorutils.FormatErrorMsg(err, "Could not get default user settings")
	}

	//The call to get _default user settings has succeeded so we may now get the user settings for this user
	name := req.GetUser().GetName()
	if name == "" {
		return nil, errors.New("cannot get user settings.. no user name was provided")
	}
	connector := req.GetUser().GetConnector()
	if connector == "" {
		return nil, errors.New(fmt.Sprintf("cannot get user settings for user: %s.  no connector was provided", name))
	}

	logrus.Debugf("Get settings for connector: %s user: %s", connector, name)
	getUserSettingsResp, err := s.db.GetUserSettings(name, connector)
	if err != nil && getUserSettingsResp != nil {
		//merge the settings
		for keyFromUser, valFromUser := range getUserSettingsResp.Settings {
			if _, ok := getDefaultUserSettingsResp.Settings[keyFromUser]; ok {
				//default has this key so it's a supported setting.  overwrite the default with that of this user.
				getDefaultUserSettingsResp.Settings[keyFromUser] = valFromUser
			}
		}
	}

	//now that we've overridden any defaults with the user settings, let's assign them to the user so that the user
	//has all defaults plus their settings
	getUserSettingsResp.Settings = getDefaultUserSettingsResp.Settings

	return getUserSettingsResp, nil
}

//	rpc UpdateUserSettings(PutUserSettings) returns (UpdateUserSettingsResponse) {
func (s *UserSettingsServer) PutUserSettings(ctx context.Context, req *user_settings.PutUserSettingsRequest) (*user_settings.PutUserSettingsResponse, error) {
	logrus.Debugf("PutUserSettings from server.go")
	name := req.GetUser().GetName()
	if name == "" {
		return nil, errors.New("invalid user Id")
	}
	connector := req.GetUser().GetConnector()
	if connector == "" {
		return nil, errors.New("invalid connector")
	}
	logrus.Debugf("Put settings for connector: %s user: %s", connector, name)
	var userSettings map[string]*user_settings.UserSettingValue
	userSettings = req.GetSettings()

	getDefaultUserSettingsResp, err := s.db.GetUserSettings("_default", "local")
	if err != nil {
		return nil, errorutils.FormatErrorMsg(err,
			fmt.Sprintf("for connector: %s and user: %s", connector, name))
	}

	logrus.Debugf("----->Checking default settings to make sure we're putting in supported settings")
	for keyFromUser, _ := range userSettings {
		if _, ok := getDefaultUserSettingsResp.Settings[keyFromUser]; !ok {
			//it's not in the default map so delete it from user settings before we save to db
			logrus.Debugf("----->Removing %s from map because it's not a supported setting", keyFromUser)
			delete(userSettings, keyFromUser)
		}
	}

	putUserSettingsResp, err := s.db.PutUserSettings(name, connector, userSettings)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("PutUserSettings response: %v", putUserSettingsResp)
	return putUserSettingsResp, nil
}

//rpc DeleteUserSettings(DeleteUserSettingsRequest) returns (DeleteUserSettingsResponse){
func (s *UserSettingsServer) DeleteUserSettings(ctx context.Context, req *user_settings.DeleteUserSettingsRequest) (*pb.Empty, error) {
	logrus.Infof("DeleteUserSettings from server.go")
	name := req.GetUser().GetName()
	if name == "" {
		return nil, errors.New("invalid user Id")
	}
	connector := req.GetUser().GetConnector()
	if connector == "" {
		return nil, errors.New("invalid connector")
	}
	logrus.Infof("Put settings for connector: %s user: %s", connector, name)
	// delete the user settings for this id
	err := s.db.DeleteUserSettings(name, connector)
	if err != nil {
		return nil, err
	}

	return &empty, nil
}
