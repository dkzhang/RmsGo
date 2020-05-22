package allConfig

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"log"
	"os"
)

func LoadAllConfig() (err error) {

	theDbConfig, err := config.LoadDbConfig(os.Getenv("dbconf"))
	if err != nil {
		return fmt.Errorf("config.LoadDbConfig error: %v", err)
	}
	log.Printf("%v", theDbConfig)
	return nil
}
