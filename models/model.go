package models

type Todo struct {
	Id          int    `json:"id"`
	Title       string `gorm:"size:1000;not null; json:"title" binding:required`
	Description string `gorm:"size:1000;not null;json:"description" binding:required`
	Completed   bool   `json:"completed"`
}

func (b *Todo) TableName() string {
	return "todo"
}
