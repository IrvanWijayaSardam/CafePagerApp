package dto

type PagerUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	PagerName   string `json:"pagername" form:"pagername" `
	PagerStatus bool   `json:"xstatus" form:"xstatus"`
	UserID      uint64 `json:"userid" form:"userid" `
	SSID        string `json:"ssid" form:"ssid"`
	SSIDPass    string `json:"ssidpass" form:"ssidpass" `
}

type PagerCreateDTO struct {
	PagerName   string `json:"pagername" form:"pagername" binding:"required"`
	PagerStatus bool   `json:"xstatus" form:"xstatus" binding:"required"`
	UserID      uint64 `json:"userid" form:"userid" binding:"required"`
	SSID        string `json:"ssid" form:"ssid" binding:"required"`
	SSIDPass    string `json:"ssidpass" form:"ssidpass" binding:"required"`
}
