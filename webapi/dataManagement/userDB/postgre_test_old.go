package userDB

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/dkzhang/RmsGo/datebaseCommon/postgreOps"
	"github.com/dkzhang/RmsGo/dbManage/pgManage"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"os"
	"testing"
)

func TestUserDB(t *testing.T) {
	os.Setenv("DbConf", "./../../../Configuration/Security/database.yaml")
	pgManage.CreateAllTable()

	//allConfig.LoadAllConfig()
	t.Logf("%v", config.TheDbConfig)

	db, err := postgreOps.ConnectToDatabase(config.TheDbConfig.ThePgConfig)
	if err != nil {
		t.Errorf("postgreSQL.ConnectToDatabase error: %v", err)
	} else {
		t.Logf("postgreSQL.ConnectToDatabase success: %v", db)
	}
	defer db.Close()

	udm := NewUserInPostgre(db)

	err = udm.InsertUser(user.UserInfo{
		UserName:       "zhang001",
		ChineseName:    "张三1",
		Department:     "计服中心",
		DepartmentCode: "JF",
		Section:        "信息技术室",
		Mobile:         "15383021234",
		Role:           1,
		Status:         2,
		Remarks:        "haha",
	})

	if err != nil {
		t.Errorf("InsertUser error: %v", err)
	}
	t.Logf("InsertUser sucess")

	err = udm.InsertUser(user.UserInfo{
		UserName:       "zhang002",
		ChineseName:    "张三2",
		Department:     "计服中心",
		DepartmentCode: "JF",
		Section:        "信息技术室",
		Mobile:         "15383021234",
		Role:           1,
		Status:         2,
		Remarks:        "haha",
	})

	if err != nil {
		t.Errorf("InsertUser error: %v", err)
	}
	t.Logf("InsertUser sucess")

	err = udm.UpdateUser(user.UserInfo{
		UserID:      1,
		UserName:    "zhang",
		ChineseName: "新张三",
		Section:     "newSection",
		Mobile:      "123456",
		Status:      222,
		Remarks:     "newRemarks",
	})

	if err != nil {
		t.Errorf("Update error: %v", err)
	}
	t.Logf("UpdateUser sucess")

	user1, err := udm.QueryUserByID(1)
	if err != nil {
		t.Errorf("QueryUserByID error: %v", err)
	}
	t.Logf("QueryUserByID sucess user = %v", user1)

	user2, err := udm.QueryUserByName("zhang")
	if err != nil {
		t.Errorf("QueryUserByName error: %v", err)
	}
	t.Logf("QueryUserByName sucess user = %v", user2)

	err = udm.UpdateUserDepartment("JF", "新计服", "NJF")
	if err != nil {
		t.Errorf("UpdateUserDepartment error: %v", err)
	}
	t.Logf("UpdateUserDepartment sucess")

	users, err := udm.GetAllUserInfo()
	if err != nil {
		t.Errorf("GetAllUserInfo error: %v", err)
	}
	t.Logf("GetAllUserInfo sucess users = %v", users)
}