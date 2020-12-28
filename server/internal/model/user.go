package model

import "time"

type User struct {
	UserId        int
	NickName      string
	LoginName     string
	PasswordMd5   string
	IntroduceSign string
	IsDeleted     int
	LockedFlag    int
	CreateTime    time.Time
}

func (User) TableName() string {
	return "tb_newbee_mall_user"
}

func (this *User) CheckPassword(password string) bool {
	return password == this.PasswordMd5
}
