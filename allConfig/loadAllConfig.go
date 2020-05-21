package allConfig

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"os"
)

func LoadAllConfig() (err error) {

	err = config.LoadDbConfig(os.Getenv("dbconf"))
	if err != nil {
		return fmt.Errorf("config.LoadDbConfig error: %v", err)
	}
	return nil
}
