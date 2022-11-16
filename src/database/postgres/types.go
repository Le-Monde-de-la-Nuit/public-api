package postgres

type User struct {
	ID      int  `json:"id"`
	RoleID  Role `json:"role_id"`
	Discord int  `json:"discord"`
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
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Place Place  `json:"place"`
	Type  Type   `json:"type"`
	Goal  string `json:"goal"`
	Time  string `json:"time"`
}
