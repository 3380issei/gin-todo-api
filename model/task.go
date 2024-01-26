package model

type Task struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Completed bool   `json:"compketed"`

	User   User `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	UserID uint `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
