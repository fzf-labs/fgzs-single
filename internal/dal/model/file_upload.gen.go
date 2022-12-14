// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFileUpload = "file_upload"

// FileUpload mapped from table <file_upload>
type FileUpload struct {
	ID               int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	FileCategory     string         `gorm:"column:file_category" json:"fileCategory"`                     // 文件分类
	FileName         string         `gorm:"column:file_name" json:"fileName"`                             // 文件新名称
	OriginalFileName string         `gorm:"column:original_file_name" json:"originalFileName"`            // 文件原名称
	Storage          string         `gorm:"column:storage" json:"storage"`                                // 存储方式
	Path             string         `gorm:"column:path" json:"path"`                                      // 文件路径
	Ext              string         `gorm:"column:ext" json:"ext"`                                        // 文件类型
	Size             int64          `gorm:"column:size" json:"size"`                                      // 文件大小
	Sha1             string         `gorm:"column:sha1" json:"sha1"`                                      // 文件sha1值
	Status           int32          `gorm:"column:status;not null;default:1" json:"status"`               // 状态(1 正常 2冻结)
	CreatedAt        time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                           // 删除时间
}

// TableName FileUpload's table name
func (*FileUpload) TableName() string {
	return TableNameFileUpload
}
