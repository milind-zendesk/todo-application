package model

type Todos struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	UserID   int    `json:"user_id"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type UserTodoDetails struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Todos      []Todos
	TotalCount int
	Priorities map[string]int
}
