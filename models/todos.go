package models

type Todos struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"type:varchar(300)" json:"title"`
	Completed bool   `gorm:"type:boolean" json:"completed"`
}

type CreateTodos struct {
	Title     string `gorm:"type:varchar(300)" json:"title"`
	Completed bool   `gorm:"type:boolean" json:"completed"`
}
