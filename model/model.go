package model

type Blog struct {
	Title string
	Link  string
}

type INI struct {
	SENDER_EMAIL    string
	SENDER_PASSWORD string
	RECEIVER_EMAIL  string
	SMTP_SERVER     string
	SMTP_PORT       int
}

type Post struct {
	ID          uint   `gorm:"primary_key"`
	BlogTitle   string `gorm:"column:blog_title"`
	Link        string `gorm:"column:link"`
	IsEmailSent bool   `gorm:"column:is_email_sent"`
}

func (Post) TableName() string {
	return "posts"
}
