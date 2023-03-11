package entity

type Pager struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	PagerName   string `gorm:"type:varchar(255)" json:"pagername"`
	PagerStatus bool   `gorm:"type:boolean" json:"xstatus"`
	SSID        string `gorm:"type:varchar(255)" json:"ssid"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	SSIDPass    string `gorm:"type:varchar(255)" json:"ssidpass"`
}
