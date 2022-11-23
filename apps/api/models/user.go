package models

type User struct {
	ID       string `gorm:"primaryKey" json:"user_id,omitempty"`
	Username string `gorm:"unique" json:"username,omitempty"`
	Email    string `gorm:"unique" json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Updated  int64  `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	Created  int64  `gorm:"autoCreateTime" json:"created_at,omitempty"`
}
