package flag

// todo 做数据库迁移之类的操作
// 先空着
import (
	sys_flag "flag"
)

type Option struct {
	DB bool
}

// Parse 解析命令行参数
func Parse() Option {

	// db默认值为false
	// 在命令行中传入 -db 参数，则 db 的值会变为 true
	db := sys_flag.Bool("db", false, "迁移至数据库")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}
