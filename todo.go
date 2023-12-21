package webback

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `jason:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id      int
	User_id int
	ListId  int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `jason:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type LsitsItem struct {
	Id     int
	ListId int
	ItemId int
}
