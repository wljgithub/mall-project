package dto

type LoginReq struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}

type LoginToken struct {
	Token string
}

type RegisterReq struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type GetUserInfoResp struct {
	IntroduceSign string `json:"introduceSign"`
	LoginName     string `json:"loginName"`
	NickName      string `json:"nickName"`
}

type UpdateUserInfoReq struct {
	IntroduceSign string `json:"introduceSign"`
	NickName string `json:"nickName"`
	PasswordMd5 string `json:"passwordMd5"`
}