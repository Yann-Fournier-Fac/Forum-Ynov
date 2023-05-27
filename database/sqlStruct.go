package database

type User struct {
	Id       int64
	Email    string
	Username string
	Password string
}

type Reponse struct {
	Id      int64
	IdPost  int64
	UserName string
	Contenu string
	Date    string
}

type Post struct {
	Id          int64
	UserName    string
	Tag         string
	Titre       string
	Description string
	NbrLikes    int64
	NbrDislikes int64
	Date        string
	// IsLike      bool
	// IsDislike   bool
}

type Session struct {
	Id    int64
	Email string
	Uuid  string
}
