package security

import (
	"testing"
)

func TestLoadDbConfig(t *testing.T) {
	filepath := `./../../Configuration/Security/database.yaml`

	theDbSecurity, err := LoadDbSecurity(filepath)
	if err != nil {
		t.Fatalf("LoadDbSecurity error: %v", err)
		return
	} else {
		t.Logf("Load DbSecurity success = %v", theDbSecurity)
		if theDbSecurity.ThePgSecurity.Host == "db" &&
			theDbSecurity.ThePgSecurity.Port == 5432 &&
			theDbSecurity.ThePgSecurity.User == "rmsu" &&
			theDbSecurity.ThePgSecurity.Password == "rmsp" &&
			theDbSecurity.ThePgSecurity.DbName == "rms_db" &&
			theDbSecurity.ThePgSecurity.Sslmode == "disable" {
			t.Logf("Verify PgSecurity success: %v", theDbSecurity.ThePgSecurity)
		} else {
			t.Errorf("Verify PgSecurity failed: %v", theDbSecurity.ThePgSecurity)
		}

		if theDbSecurity.TheRedisSecurity.Host == "redis:6379" {
			t.Logf("Verify RedisSecurity success: %v", theDbSecurity.TheRedisSecurity)
		} else {
			t.Errorf("Verify RedisSecurity failed: %v", theDbSecurity.TheRedisSecurity)
		}
	}

}
