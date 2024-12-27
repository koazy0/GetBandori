package database

import "strconv"

type Mysql struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	DB     string `yaml:"db"`
	Config string `yaml:"config"` // 高级设置，如charset等
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
}

// Dsn 返回dsn语句
func (m Mysql) Dsn() string {
	return m.User + ":" + m.Passwd + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + m.Config
}
