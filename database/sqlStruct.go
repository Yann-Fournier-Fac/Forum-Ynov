package database

type User struct {
	Id       int64
	Email    string
	Username string
	Password string
}

type Reponse struct {
	Id      int64
	IdUser  int64
	IdPost  int64
	Contenu string
}

type Post struct {
	Id          int64
	IdUser      int64
	Tag         string
	Titre       string
	Description string
	NbrLikes    int64
	NbrDislikes int64
	// IsLike      bool
	// IsDislike   bool
}

type Session struct {
	Id   int64
	Uuid string
}
