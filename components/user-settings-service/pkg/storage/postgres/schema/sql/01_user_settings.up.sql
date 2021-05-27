CREATE TABLE IF NOT EXISTS user_settings
(
    id SERIAL PRIMARY KEY,
    user_name TEXT NOT NULL DEFAULT '',
    connector TEXT NOT NULL DEFAULT 'local',
    settings JSON  NOT NULL DEFAULT '[]'::json,
    CONSTRAINT user_settings_user_name_and_connector UNIQUE (user_name, connector)
);

INSERT INTO user_settings (user_name, connector, settings)
VALUES ('_default', 'local',
        '{
          "date_format": {
            "value": "EEEE, MMMM d, y, h:mm:ss a zzzz",
            "enabled": true
          }
        }')
ON CONFLICT ON CONSTRAINT user_settings_user_name_and_connector
    DO UPDATE
    SET settings='{
      "date_format": {
        "value": "EEEE, MMMM d, y, h:mm:ss a zzzz",
        "enabled": true
      }
    }';