package logMap

import "testing"

func TestLoadLogConfig(t *testing.T) {
	filepath := `test.yaml`

	theLogConfig, err := LoadLogConfig(filepath)
	if err != nil {
		t.Fatalf("LoadLogConfig error: %v", err)
	} else {
		t.Logf("Load LogConfig success = %v", theLogConfig)

		normal, okNormal := theLogConfig.LogFile["normal"]
		if !okNormal || normal != "/log/normal" {
			t.Errorf("Verify normal LogConfig failed,")
		}

		login, okLogin := theLogConfig.LogFile["login"]
		if !okLogin || login != "/log/login" {
			t.Errorf("Verify login LogConfig failed,")
		}

		db, okDb := theLogConfig.LogFile["db"]
		if !okDb || db != "/log/db" {
			t.Errorf("Verify db LogConfig failed,")
		}
		t.Log("Verify LogConfig value success")
	}

}
