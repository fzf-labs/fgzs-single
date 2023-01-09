// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TableNameSysAdmin = "sys_admin"

// SysAdmin mapped from table <sys_admin>
type SysAdmin struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                     // 编号
	Username  string         `gorm:"column:username;not null" json:"username"`                              // 用户名
	Password  string         `gorm:"column:password;not null" json:"password"`                              // 密码
	Nickname  string         `gorm:"column:nickname;not null" json:"nickname"`                              // 昵称
	Avatar    string         `gorm:"column:avatar;not null" json:"avatar"`                                  // 头像
	Gender    int32          `gorm:"column:gender;not null" json:"gender"`                                  // 0=保密 1=女 2=男
	Email     string         `gorm:"column:email;not null" json:"email"`                                    // 邮件
	Mobile    string         `gorm:"column:mobile;not null" json:"mobile"`                                  // 手机号
	JobID     int64          `gorm:"column:job_id" json:"jobId"`                                            // 岗位
	DeptID    int64          `gorm:"column:dept_id" json:"deptId"`                                          // 部门
	RoleIds   datatypes.JSON `gorm:"column:role_ids" json:"roleIds"`                                        // 角色集
	Salt      string         `gorm:"column:salt;not null" json:"salt"`                                      // 盐值
	Status    int32          `gorm:"column:status;not null;default:1" json:"status"`                        // 0=禁用 1=开启
	Sort      int64          `gorm:"column:sort;not null" json:"sort"`                                      // 排序值
	Motto     string         `gorm:"column:motto;not null" json:"motto"`                                    // 个性签名
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`          // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                                    // 删除时间
}

// TableName SysAdmin's table name
func (*SysAdmin) TableName() string {
	return TableNameSysAdmin
}
