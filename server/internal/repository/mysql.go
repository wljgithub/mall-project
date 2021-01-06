package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/wljgithub/mall-project/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewMysql() (*gorm.DB, func(), error) {
	db, err := newMysql(conf.Conf.MySQL)

	return db, func() {}, err
}

func newMysql(config conf.MySQLConfig) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		config.UserName,
		config.Password,
		config.Addr,
		config.Name,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open mysql")
	}
	sql, err := db.DB()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to config mysql")
	}
	sql.SetMaxIdleConns(config.MaxIdleConn)
	sql.SetMaxOpenConns(config.MaxOpenConn)
	sql.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifeTime))

	//db.AutoMigrate(tableToRegister...)

	return db, err

}

//func testDbCon() {
//
//}
