package database

import "GetBandori/model"

// todo 迁移数据库表
func Makemigrations() {

	// 生成表结构
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(model.SongExtractModel{})
}
