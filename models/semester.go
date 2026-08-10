package models

import (
	"time"

	"gorm.io/gorm"
)

type Semester struct {
	ID             uint64         `gorm:"primaryKey;column:id" json:"id"`
	TahunAjaranID  uint64         `gorm:"column:tahun_ajaran_id;not null" json:"tahun_ajaran_id"`
	NamaSemester   string         `gorm:"column:nama_semester;size:50;not null" json:"nama_semester"`
	Kode   string         `gorm:"column:kode;size:20;not null" json:"kode"`
	IsActive       bool           `gorm:"column:is_active;not null;default:false" json:"is_active"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	TahunAjaran TahunAjaran `gorm:"foreignKey:TahunAjaranID" json:"tahun_ajaran,omitempty"`
}

func (Semester) TableName() string {
	return "semester"
}
