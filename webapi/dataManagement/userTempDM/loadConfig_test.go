package userTempDM

import "testing"

func TestLoadLoginConfig(t *testing.T) {
	filepath := "../../../Configuration/Parameter/login.yaml"

	cfg, err := LoadLoginConfig(filepath)
	if err != nil {
		t.Errorf("LoadLoginConfig error: %v", err)
	} else {
		t.Logf("LoadLoginConfig success, cfg = %v", cfg)
	}
}
