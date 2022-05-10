package models

type User struct {
	Userid   int64  `json:"user_id,omitempty"`
	Username string `json:"username"`
	Password string `gorm:"-"`
}

//表示配置操作数据库的表名称
func (User) TableName() string {
	return "user"
}
