package database

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

const ConfigFile = "settings.yaml"

var (
	DB *gorm.DB
)

func init() {
	DB = InitGorm()
}

// InitGorm gorm连接数据库
func InitGorm() *gorm.DB {
	MysqlConf := &Mysql{}
	file, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}
	err = json.Unmarshal(file, MysqlConf)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}
	if MysqlConf.Host == "" {
		log.Fatalln("未配置Mysql_conf，取消gorm连接")
		return nil
	}

	// 获取dsn语句
	dsn := MysqlConf.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ // gorm 标准的打开方式,这里没有额外控制要求
	})
	if err != nil {
		log.Fatalln("[%s] Mysql_conf连接失败", dsn)
	}

	//返回golang原生的 *sql.DB 对象，用于进一步控制数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳的并发连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过Mysql_conf的wait_timeout

	return db
}
