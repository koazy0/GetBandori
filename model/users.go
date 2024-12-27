package model

import "time"

//todo 用户表
//存储每个用户的成绩截图，和SongExtract作为多对多的关系表进行

type UserModel struct {
	ID       uint      `gorm:"primarykey" json:"id"`            // 主键ID
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"` // 创建时间
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"-"`         // 更新时间
	NickName string    `gorm:"size:36" json:"nick_name"`        // 昵称 用于显示，可以重复
	UserName string    `gorm:"size:36" json:"user_name"`        // 用户名 用于登录,不能重复
	Password string    `gorm:"size:128" json:"password"`        // 密码
	Avatar   string    `gorm:"size:256" json:"avatar_id"`       // 头像url
	//Email         string           `gorm:"size:128" json:"email"`           // 邮箱
	//Tel           string           `gorm:"size:18" json:"tel"`              // 手机号
	//Addr          string           `gorm:"size:64" json:"addr"`             // 地址
	Token string `gorm:"size:64" json:"token"` // 其他平台的唯一id
	//IP            string           `gorm:"size:20" json:"ip"`               // ip地址
	PictureModels []PictureModels `gorm:"foreignKey:UserID" json:"-"` // 发布的图片ID
}
