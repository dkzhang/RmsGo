package login

import "testing"

func TestLoadLoginConfig(t *testing.T) {
	filepath := `./../../Configuration/Parameter/login.yaml`

	err := LoadLoginConfig(filepath)
	if err != nil {
		t.Fatalf("LoadDbConfig error: %v", err)
	} else {
		t.Logf("LoadDbConfig success: %v", TheLoadLoginConfig)
	}

}
