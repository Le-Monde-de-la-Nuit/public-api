package mariadb

type User struct {
	ID      int64 `json:"id"`
	RoleID  int64 `json:"role_id"`
	Discord int64 `json:"discord"`
}

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Place struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Type struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Action struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	Place Place    `json:"place"`
	Type  Type     `json:"type"`
	Goal  string   `json:"goal"`
	Time  NullTime `json:"time"`
}

type RawPost struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	CreatedAt NullTime `json:"created_at"`
}

type Post struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	CreatedAt NullTime `json:"created_at"`
	Tag       []Tag    `json:"tag"`
}

type NewPost struct {
	ID      int64   `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Tag     []int64 `json:"tag"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PostTag struct {
	ID     int64 `json:"id"`
	PostID int64 `json:"post"`
	TagID  int64 `json:"tag"`
}

func (p *RawPost) ToPost() Post {
	return Post{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
	}
}

func (p *NewPost) ToPost() Post {
	return Post{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
	}
}
