package types

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type Report struct {
	ID          uint      `gorm:"primaryKey"`
	AdminName   string    `gorm:"not null"`
	RoleID      uint      `gorm:"not null"`
	Role        Role      `gorm:"foreignKey:RoleID"`
	PunishDate  time.Time `gorm:"not null"`
	Description string    `gorm:"type:text; not null"`
	Evidence    string    `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ReportInput struct {
	AdminName   string `json:"admin_name"`
	RoleName    string `json:"role"`
	PunishDate  string `json:"punish_date"`
	Description string `json:"description"`
	Evidence    string `json:"evidence"`
}

type ReportOutput struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

type ReportResponse struct {
	ID          uint   `json:"id"`
	AdminName   string `json:"adminName"`
	RoleName    string `json:"roleName"`
	PunishDate  string `json:"punishDate"`
	Description string `json:"description"`
	Evidence    string `json:"evidence"`
}
