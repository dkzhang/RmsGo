package config

import "testing"

func TestLoadDbConfig(t *testing.T) {
	filepath := `test.yaml`

	theDbConfig, err := LoadDbConfig(filepath)
	if err != nil {
		t.Fatalf("LoadDbConfig error: %v", err)
	} else {
		t.Logf("Load DbConfig success = %v", theDbConfig)
		if theDbConfig.ThePgConfig.Host == "rms-pg" &&
			theDbConfig.ThePgConfig.User == "rms" &&
			theDbConfig.ThePgConfig.Password == "123456" &&
			theDbConfig.ThePgConfig.DbName == "rms_db" &&
			theDbConfig.ThePgConfig.Sslmode == "disable" {
			t.Logf("Verify PgConfig success: %v", theDbConfig.ThePgConfig)
		} else {
			t.Errorf("Verify PgConfig failed: %v", theDbConfig.ThePgConfig)
		}

		if theDbConfig.TheRedisConfig.Host == "rms-redis:6379" &&
			theDbConfig.TheRedisConfig.Password == "234567" {
			t.Logf("Verify RedisConfig success: %v", theDbConfig.TheRedisConfig)
		} else {
			t.Errorf("Verify RedisConfig failed: %v", theDbConfig.TheRedisConfig)
		}
	}

}
