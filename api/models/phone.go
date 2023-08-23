package models

type PhonePrimaryKey struct {
	Id string `json:"id"`
}

type Phone struct {
	Id          string `json:"id"`
	User_id     string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Is_fax      int    `json:"is_fax"`
}

type PhoneCreate struct {
	User_id     string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Is_fax      int    `json:"is_fax"`
}

type PhoneUpdate struct {
	Id          string `json:"id"`
	User_id     string `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Is_fax      int    `json:"is_fax"`
}

type PhoneGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type PhoneGetListResponse struct {
	Count  int      `json:"count"`
	Phones []*Phone `json:"Phones"`
}
