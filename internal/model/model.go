package model

    type Task struct {
		ID 				string `json:"id"`
		Title			string `json:"title"` 
		Description		string `json:"description"` 
		Status			string `json:"status"` 
		CreatedAt		string `json:"created_at"`
		UpdatedAt		string `json:"updated_at"`
	}


	const (
		New 		string = "new"
		InProgress	string = "in progress"
		Done 		string = "done"
		Cancelled	string = "cancelled"

	)