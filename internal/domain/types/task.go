package types

import "time"

type Task struct {
	ID 				uint `json:"id" gorm:"primaryKey"`
	Title			string `json:"title" gorm:"not null"` 
	Description		string `json:"description"` 
	Status			string `json:"status"` 
	CreatedAt   time.Time `gorm:"autoCreateTime"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}


const (
	New 		string = "new"
	InProgress	string = "in_progress"
	Done 		string = "done"
	Cancelled	string = "cancelled"
)