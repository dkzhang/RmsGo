package postgreOpsGorm

import (
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/security"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(pg security.PgSecurity) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai",
		pg.Host, pg.Port, pg.User, pg.Password, pg.DbName, pg.Sslmode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func DropTable(db *gorm.DB, dst ...interface{}) (err error) {
	err = db.Migrator().DropTable(dst)
	return fmt.Errorf("db.Migrator().DropTable error: %v", err)
}

func ClearTable(db *gorm.DB, dst ...interface{}) {
	db.Where("1 = 1").Delete(dst)
}
