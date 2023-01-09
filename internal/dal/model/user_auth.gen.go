// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TableNameUserAuth = "user_auth"

// UserAuth mapped from table <user_auth>
type UserAuth struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID            string         `gorm:"column:uid;default:0" json:"uid"`                              // 用户ID
	IdentityType   int32          `gorm:"column:identity_type;not null" json:"identityType"`            // 1 微信 2 苹果
	IdentityKey    string         `gorm:"column:identity_key;not null" json:"identityKey"`              // 业务登录key
	IdentifierCode string         `gorm:"column:identifier_code;not null" json:"identifierCode"`        // 标识码
	IdentityName   string         `gorm:"column:identity_name" json:"identityName"`                     // 昵称
	Other          datatypes.JSON `gorm:"column:other" json:"other"`                                    // 其他数据
	Status         int32          `gorm:"column:status;not null;default:1" json:"status"`               // 状态 1绑定 0解绑
	CreatedAt      time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt      time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                           // 删除时间
}

// TableName UserAuth's table name
func (*UserAuth) TableName() string {
	return TableNameUserAuth
}
