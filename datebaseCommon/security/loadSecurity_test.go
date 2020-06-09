package security

import "testing"

func TestLoadDbConfig(t *testing.T) {
	filepath := `./../../Configuration/Security/database.yaml`

	err := LoadDbSecurity(filepath)
	if err != nil {
		t.Fatalf("LoadDbSecurity error: %v", err)
	} else {
		t.Logf("Load DbSecurity success = %v", TheDbConfig)
		if TheDbConfig.ThePgConfig.Host == "rms-pg" &&
			TheDbConfig.ThePgConfig.User == "rms" &&
			TheDbConfig.ThePgConfig.Password == "123456" &&
			TheDbConfig.ThePgConfig.DbName == "rms_db" &&
			TheDbConfig.ThePgConfig.Sslmode == "disable" {
			t.Logf("Verify PgSecurity success: %v", TheDbConfig.ThePgConfig)
		} else {
			t.Errorf("Verify PgSecurity failed: %v", TheDbConfig.ThePgConfig)
		}

		if TheDbConfig.TheRedisConfig.Host == "rms-redis:6379" &&
			TheDbConfig.TheRedisConfig.Password == "234567" {
			t.Logf("Verify RedisSecurity success: %v", TheDbConfig.TheRedisConfig)
		} else {
			t.Errorf("Verify RedisSecurity failed: %v", TheDbConfig.TheRedisConfig)
		}
	}

}
