package database

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)
type Company struct {
	ID int `gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt sql.NullTime
}

type Hr struct{
	ID        int
	CompanyID int `json:"company_id,omitempty"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Password sql.NullString `json:"password,omitempty"`
	CreatedAt sql.NullTime `json:"created_at,omitempty"`
}

type JobAdvert struct{
	Company   Company
	CompanyID int `json:"company_id"`
	Hr        Hr
	HrId      int `json:"hr_id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Address string `json:"address"`
	Date sql.NullTime `json:"date"`
}

type InviteHr struct{
	ID int `gorm:"primaryKey"`
	HrID int `json:"hr_id"`
	Name string `json:"name"`
	Link uuid.UUID `json:"link"`
	Email string `json:"email"`
	CompanyID int `json:"company_id"`
	CreatedAt time.Time
}

type Applicants struct{
	ID int `gorm:"primaryKey"`
	HrID int `json:"hr_id"`
	CompanyID int `json:"company_id"`
	Body string `json:"body"`
	Status string `json:"status"`
	Time sql.NullTime
}