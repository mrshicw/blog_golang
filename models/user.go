package models

import (
	"time"

	"gorm.io/gorm"
)

// 2. 数据库设计与模型定义
// 设计数据库表结构，至少包含以下几个表：

// users 表：存储用户信息，
// 包括 id 、 username 、 password 、 email 等字段。
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:100"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// posts 表：存储博客文章信息，
// 包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
type Posts struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"uniqueIndex;not null;size:50"`
	Content   string         `json:"content" gorm:"not null;size:100"`
	UserID    uint           `json:"user_id" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// comments 表：存储文章评论信息，
// 包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
type Comments struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Content   string         `json:"content" gorm:"not null;size:100"`
	UserID    uint           `json:"user_id" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PostID    uint           `json:"post_id" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

//使用 GORM 定义对应的 Go 模型结构体。
