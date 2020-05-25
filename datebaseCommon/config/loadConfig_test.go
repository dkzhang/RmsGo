package config

import "testing"

func TestLoadDbConfig(t *testing.T) {
	filepath := `./../../Configuration/Security/database.yaml`

	err := LoadDbConfig(filepath)
	if err != nil {
		t.Fatalf("LoadDbConfig error: %v", err)
	} else {
		t.Logf("Load DbConfig success = %v", TheDbConfig)
		if TheDbConfig.ThePgConfig.Host == "rms-pg" &&
			TheDbConfig.ThePgConfig.User == "rms" &&
			TheDbConfig.ThePgConfig.Password == "123456" &&
			TheDbConfig.ThePgConfig.DbName == "rms_db" &&
			TheDbConfig.ThePgConfig.Sslmode == "disable" {
			t.Logf("Verify PgConfig success: %v", TheDbConfig.ThePgConfig)
		} else {
			t.Errorf("Verify PgConfig failed: %v", TheDbConfig.ThePgConfig)
		}

		if TheDbConfig.TheRedisConfig.Host == "rms-redis:6379" &&
			TheDbConfig.TheRedisConfig.Password == "234567" {
			t.Logf("Verify RedisConfig success: %v", TheDbConfig.TheRedisConfig)
		} else {
			t.Errorf("Verify RedisConfig failed: %v", TheDbConfig.TheRedisConfig)
		}
	}

}
