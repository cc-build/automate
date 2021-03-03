package integration_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chef/automate/api/interservice/infra_proxy/request"
	"github.com/stretchr/testify/assert"
)

func TestGetEnvironments(t *testing.T) {
	// rpc GetEnvironments (request.Environments) returns (response.Environments)
	ctx := context.Background()

	// Adds environments
	addEnvironments(ctx, 10)

	req := &request.Environments{
		ServerId: autoDeployedChefServerID,
		OrgId:    autoDeployedChefOrganizationID,
	}

	t.Run("Environments list without a search params", func(t *testing.T) {
		res, err := infraProxy.GetEnvironments(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, int(res.Page))
		assert.GreaterOrEqual(t, int(res.Total), 0)

	})

	t.Run("Environments list with a per_page search param", func(t *testing.T) {
		req.SearchQuery = &request.SearchQuery{
			PerPage: 1,
		}

		res, err := infraProxy.GetEnvironments(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, int(res.Page))
		assert.Equal(t, 1, len(res.Environments))
		assert.GreaterOrEqual(t, int(res.Total), 0)
	})

	t.Run("Environments list with a page search param", func(t *testing.T) {
		req.SearchQuery = &request.SearchQuery{
			PerPage: 1,
			Page:    1,
		}

		res, err := infraProxy.GetEnvironments(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 1, int(res.Page))
		assert.Equal(t, 1, len(res.Environments))
		assert.GreaterOrEqual(t, int(res.Total), 0)
	})

	t.Run("Environments list with an invalid query search param", func(t *testing.T) {
		req.SearchQuery = &request.SearchQuery{
			Q:       "NO_KEY:NO_VALUE",
			Page:    0,
			PerPage: 5,
		}
		res, err := infraProxy.GetEnvironments(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, int(res.Page))
		assert.Equal(t, 0, int(res.Total))
		assert.Equal(t, 0, len(res.Environments))
	})

	t.Run("Environments list with an invalid query search param", func(t *testing.T) {
		req.SearchQuery = &request.SearchQuery{
			Q:       "INVALID_QUERY",
			Page:    0,
			PerPage: 5,
		}
		res, err := infraProxy.GetEnvironments(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, 0, int(res.Page))
		assert.Equal(t, 0, int(res.Total))
		assert.Equal(t, 0, len(res.Environments))
	})

	t.Run("Environments list with a valid query search param", func(t *testing.T) {
		name := fmt.Sprintf("chef-environment-%d", time.Now().Nanosecond())
		createReq := &request.CreateEnvironment{
			ServerId: autoDeployedChefServerID,
			OrgId:    autoDeployedChefOrganizationID,
			Name:     name,
		}
		env, err := infraProxy.CreateEnvironment(ctx, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, env)

		req.SearchQuery = &request.SearchQuery{
			Q:       fmt.Sprintf("name:%s", name),
			Page:    0,
			PerPage: 5,
		}
		res, err := infraProxy.GetEnvironments(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, int(res.Page))
		assert.Equal(t, 1, int(res.Total))
		assert.Equal(t, name, res.Environments[0].GetName())
	})

}

// Adds environments records
func addEnvironments(ctx context.Context, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		req := &request.CreateEnvironment{
			ServerId: autoDeployedChefServerID,
			OrgId:    autoDeployedChefOrganizationID,
			Name:     fmt.Sprintf("chef-environment-%d", time.Now().Nanosecond()),
		}
		_, err := infraProxy.CreateEnvironment(ctx, req)

		if err == nil {
			total++
		}
	}

	return total
}
