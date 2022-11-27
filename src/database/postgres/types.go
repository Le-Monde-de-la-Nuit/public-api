package postgres

type User struct {
	ID      int `json:"id"`
	RoleID  int `json:"role_id"`
	Discord int `json:"discord"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Place struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Type struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Action struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Place Place    `json:"place"`
	Type  Type     `json:"type"`
	Goal  string   `json:"goal"`
	Time  NullTime `json:"time"`
}

type Posts struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	CreatedAt NullTime `json:"created_at"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PostTag struct {
	ID     int `json:"id"`
	PostID int `json:"post"`
	TagID  int `json:"tag"`
}
