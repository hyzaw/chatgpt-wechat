// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameKnowledge = "knowledge"

// Knowledge mapped from table <knowledge>
type Knowledge struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	UserID    int64     `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`          // 用户ID
	Name      string    `gorm:"column:name;not null;comment:知识库名称" json:"name"`               // 知识库名称
	Avatar    string    `gorm:"column:avatar;not null;comment:知识库头像" json:"avatar"`           // 知识库头像
	Desc      string    `gorm:"column:desc;not null;comment:知识库描述" json:"desc"`               // 知识库描述
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Knowledge's table name
func (*Knowledge) TableName() string {
	return TableNameKnowledge
}