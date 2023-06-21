package models

// FadeBackModel 用户反馈表
type FadeBackModel struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ApplyContent string `gorm:"size:128" json:"apply_content"` //回复的内容
	IsApply      bool   `json:"isApply"`                       //是否回复
}
