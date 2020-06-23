package myPgBytea

import (
	"github.com/dkzhang/RmsGo/DK_Experiments/myGob"
	"log"
	"testing"
)

func TestCreateTable(t *testing.T) {
	conf := `C:\Users\dkzhang\go\src\github.com\dkzhang\RmsGo\Configuration\Security\db41.yaml`
	db := ConnectToDatabase(conf)
	CreateTable(db)
	defer DropTable(db)

	data := []int{1, 3, 5, 7, 9}

	err := InsertUser(db, UserInfo{
		UserName: myGob.Store(data),
		Remarks:  "",
	})
	if err != nil {
		t.Fatalf("InsertUser error: %v", err)
	}

	var dataRead = []int{}
	u, err := RetrieveUser(db, 1)
	if err != nil {
		t.Fatalf("RetrieveUser error: %v", err)
	}
	myGob.Load(&dataRead, u.UserName)

	log.Printf("dataRead = %v", dataRead)

}
