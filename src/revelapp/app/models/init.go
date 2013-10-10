package models

import (
	"fmt"
	_ "github.com/coocood/mysql"
	"github.com/coocood/qbs"
	"github.com/robfig/config"
	"path/filepath"
)

func RegisterDb(driver, dbname, user, password, host string) {
	qbs.Register(driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname), dbname, qbs.NewMysql())
	err := CreateTabel()
	if err != nil {
		fmt.Println(err)
	}
}

func CreateTabel() error {
	migration, err := qbs.GetMigration()
	if err != nil {
		return err
	}
	defer migration.Close()

	err = migration.CreateTableIfNotExists(new(User))
	err = migration.CreateTableIfNotExists(new(Category))
	err = migration.CreateTableIfNotExists(new(Topic))
	err = migration.CreateTableIfNotExists(new(Reply))
	err = migration.CreateTableIfNotExists(new(Permissions))

	return err
}

func init() {
	path, _ := filepath.Abs("")
	c, _ := config.ReadDefault(fmt.Sprintf("%s/src/revelapp/conf/my.conf", path))

	driver, _ := c.String("database", "db.driver")
	dbname, _ := c.String("database", "db.dbname")
	user, _ := c.String("database", "db.user")
	password, _ := c.String("database", "db.password")
	host, _ := c.String("database", "db.host")

	RegisterDb(driver, dbname, user, password, host)
}
