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
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"uniqueIndex;not null;size:50"`
	Email     string         `gorm:"uniqueIndex;not null;size:100"`
	Password  string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// posts 表：存储博客文章信息，
// 包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
type Posts struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string         `gorm:"uniqueIndex;not null;size:50"`
	Content   string         `gorm:"not null;size:100"`
	UserID    uint           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// comments 表：存储文章评论信息，
// 包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
type Comments struct {
	ID        uint           `gorm:"primaryKey"`
	Content   string         `gorm:"not null;size:100"`
	UserID    uint           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PostID    uint           `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//使用 GORM 定义对应的 Go 模型结构体。
