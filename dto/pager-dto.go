package dto

type PagerUpdateDTO struct {
	PagerID     uint64 `json:"id" form:"id" binding:"required"`
	PagerName   string `json:"title" form:"title" binding:"required"`
	PagerStatus bool   `json:"xstatus" form:"xstatus" binding:"required"`
	UserID      uint64 `json:"userid" form:"userid" binding:"required"`
	SSID        string `json:"ssid" form:"ssid" binding:"required"`
	SSIDPass    string `json:"ssidpass" form:"ssidpass" binding:"required"`
}

type PagerCreateDTO struct {
	PagerName   string `json:"title" form:"title" binding:"required"`
	PagerStatus bool   `json:"xstatus" form:"xstatus" binding:"required"`
	UserID      uint64 `json:"userid" form:"userid" binding:"required"`
	SSID        string `json:"ssid" form:"ssid" binding:"required"`
	SSIDPass    string `json:"ssidpass" form:"ssidpass" binding:"required"`
}
