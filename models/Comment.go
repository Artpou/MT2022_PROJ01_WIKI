package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Comment struct {
	ID           uint    `gorm:"primaryKey"`
	AuthorID     uint    `gorm:"not null"`
	User         User    `gorm:"foreignKey:AuthorID" json:"-"`
	ArticleID    uint    `gorm:"not null"`
	Article      Article `gorm:"foreignKey:ArticleID" json:"-"`
	Content      string  `gorm:"not null;size:500"`
	CreationDate JSONTime
	LatestUpdate JSONTime
}

func NewComment(articleID uint, content string) *Comment {
	comment := Comment{Content: content, ArticleID: articleID}
	comment.AuthorID = 1
	comment.CreationDate = JSONTime(time.Now())
	comment.LatestUpdate = JSONTime(time.Now())
	return &comment
}

func UpdateComment(article Comment, content string) *Comment {
	if content != "" {
		article.Content = content
	}
	article.LatestUpdate = JSONTime(time.Now())
	return &article
}
