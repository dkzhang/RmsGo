package user

import (
	"github.com/dkzhang/RmsGo/allConfig"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreSQL"
	"os"
	"testing"
)

func TestGetAllUserInfo(t *testing.T) {
	os.Setenv("DbConf", "./../../../Configuration/Security/database.yaml")

	allConfig.LoadAllConfig()
	t.Logf("%v", config.TheDbConfig)

	db, err := postgreSQL.ConnectToDatabase(config.TheDbConfig.ThePgConfig)
	if err != nil {
		t.Errorf("postgreSQL.ConnectToDatabase error: %v", err)
	} else {
		t.Logf("postgreSQL.ConnectToDatabase success: %v", db)
	}
	defer db.Close()

	err = InsertUser(UserInfo{
		UserName:       "zhang002",
		ChineseName:    "张三",
		Department:     "计服中心",
		DepartmentCode: "JF",
		Section:        "信息技术室",
		Mobile:         "15383021234",
		Role:           1,
		Status:         1,
		Remarks:        "haha",
	}, db)

	if err != nil {
		t.Errorf("InsertUser error: %v", err)
	}
}
