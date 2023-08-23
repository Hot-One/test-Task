package models

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserPrimaryKey struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type UserCreate struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type UserUpdate struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type UserGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type UserGetListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
