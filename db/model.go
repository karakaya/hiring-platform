package db

import "database/sql"
type Company struct {
	ID uint `gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt sql.NullTime
}

type Hr struct{
	ID        int

	CompanyID uint `json:"companyID,omitempty"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Password sql.NullString `json:"password,omitempty"`
	CreatedAt sql.NullTime `json:"created_at,omitempty"`
}

type JobAdvert struct{
	Company   Company
	CompanyID uint
	Hr        Hr
	HrId      uint
	Title string `json:"title"`
	Body string `json:"body"`
	Address string `json:"address"`
	Date sql.NullTime `json:"date"`
}