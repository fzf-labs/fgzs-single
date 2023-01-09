// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAsset = "asset"

// Asset mapped from table <asset>
type Asset struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID       string         `gorm:"column:uid;not null" json:"uid"`                               // 用户id
	AssetType int32          `gorm:"column:asset_type;not null" json:"assetType"`                  // 资产类型
	Amount    int64          `gorm:"column:amount;not null" json:"amount"`                         // 资产值
	Version   int64          `gorm:"column:version" json:"version"`                                // 乐观锁
	Status    int32          `gorm:"column:status;not null;default:1" json:"status"`               // 状态(1 正常 2冻结)
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                           // 删除时间
}

// TableName Asset's table name
func (*Asset) TableName() string {
	return TableNameAsset
}
