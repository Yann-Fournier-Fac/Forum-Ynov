package database

import (
	"database/sql"
	"log"
	"strconv"
)

func Database() {
	databaseAll := InitDatabase("forum.db")
	defer databaseAll.Close()
	sqlStmt := `

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		email TEXT NOT NULL, 
		username TEXT NOT NULL, 
		password TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS reponses (
		id INTEGER NOT NULL,
		idPost INTEGER NOT NULL,
		userName TEXT NOT NULL,
		contenu TEXT NOT NULL,
		date TEXT NOT NULL,
		PRIMARY KEY (id)
	);

	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL,
		userName Text NOT NULL,
		tag TEXT NOT NULL,
		titre TEXT NOT NULL,
		description TEXT NOT NULL,
		nbrLikes INTEGER,
		nbrDislikes INTEGER,
		date TEXT NOT NULL,
		PRIMARY KEY (id)
	);

	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER NOT NULL,
		email TEXT NOT NULL,
		uuid TEXT NOT NULL,
		PRIMARY KEY (id)
	);
		`

	_, err := databaseAll.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	// _, err = databaseAll.Exec("PRAGMA foreign_keys = ON;") // doit tjs effacer db d'abord avant de run SI db actuel n'a pas foreign key activée + doit le run en dernier
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func DatabaseAndUsers(values []string) {
	db := InitDatabase("forum.db")
	defer db.Close()
	sqlStmtInsertUsers := `
		INSERT INTO users (email, username, password) VALUES (?, ?, ?);
		`
	// TODO remplacer par les valeurs get du web
	InsertIntoRow(db, values, sqlStmtInsertUsers)
	// rowsUser := selectAllFrom(databaseAll, "users")
	// displayUserRow(rowsUser)
}

func DatabaseAndReponse(values []string) {
	db := InitDatabase("forum.db")
	defer db.Close()
	sqlStmtInsertReponses := `
		INSERT INTO reponses (idPost, userName,contenu, date) VALUES (?, ?, ?, ?);
		`
	InsertIntoRow(db, values, sqlStmtInsertReponses)
	// rowsReponse := selectAllFrom(databaseAll, "reponses")
	// displayReponseRow(rowsReponse)
}

func DatabaseAndPost(values []string) {
	db := InitDatabase("forum.db")
	defer db.Close()
	sqlStmtInsertPosts := `
		INSERT INTO posts (userName, tag, titre, description, nbrLikes, nbrDislikes, date) VALUES (?, ?, ?, ?, ?, ?, ?);
		`
	// TODO remplacer par valeurs ?
	InsertIntoRow(db, values, sqlStmtInsertPosts)
	// Toutes les fonctions possibles ( telles ou telles s'activent avec des if relier au boutons )
	// rowsPost := selectAllFrom(databaseAll, "posts")
	// rowsPost = selectByAscending(databaseAll, "posts", "nbrLikes")
	// rowsPost = selectByDescending(databaseAll, "posts", "nbrLikes")
	// rowsPost = updateNbr(databaseAll, "posts", "nbrLikes", 3) // 3 à remplacer par la valeur javascript du site ( doit actuellement être en int )
	// displayPostRow(rowsPost)
}

func DatabaseAndSession(values []string) {
	db := InitDatabase("forum.db")
	defer db.Close()
	sqlStmtInsertPosts := `
		INSERT INTO sessions (email, uuid) VALUES (?, ?);
		`
	// TODO remplacer par valeurs ?
	InsertIntoRow(db, values, sqlStmtInsertPosts)
}

func InitDatabase(dataBaseName string) *sql.DB {
	// ./ + database pr le mettre au bon endroit
	database, err := sql.Open("sqlite3", "./database/"+dataBaseName)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func InsertIntoRow(db *sql.DB, values []string, stmt string) (int64, error) {
	// Use of interface for go's variadic parameters feature
	args := make([]interface{}, len(values))
	for i, v := range values {
		args[i] = v
	}
	row, err := db.Exec(stmt, args...)
	if err != nil {
		log.Fatal(err)
	}
	return row.LastInsertId()
}

// func DisplayUserRow(rows *sql.Rows) {
// 	// tableau d'objet, à tester si ça ira
// 	var htmlTable []User
// 	for rows.Next() {
// 		// debug console
// 		var user User
// 		err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// %s = %v
// 		htmlTable = append(htmlTable, user)
// 	}
// 	fmt.Printf("web user : %v ", htmlTable)
// 	fmt.Println()
// }

// func DisplayReponseRow(rows *sql.Rows) {
// 	// tableau d'objet, à tester si ça ira
// 	var htmlTable []Reponse
// 	for rows.Next() {
// 		// debug console
// 		var reponse Reponse
// 		err := rows.Scan(&reponse.Id, &reponse.IdUser, &reponse.IdPost, &reponse.Contenu)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// %s = %v
// 		htmlTable = append(htmlTable, reponse)
// 	}
// 	fmt.Printf("web reponse : %v ", htmlTable)
// 	fmt.Println()
// }

// func DisplayPostRow(rows *sql.Rows) {
// 	// tableau d'objet, à tester si ça ira
// 	var htmlTable []Post
// 	for rows.Next() {
// 		// debug console
// 		var post Post
// 		err := rows.Scan(&post.Id, &post.IdUser, &post.Tag, &post.Titre, &post.Description, &post.NbrLikes, &post.NbrDislikes)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// %s = %v
// 		htmlTable = append(htmlTable, post)
// 	}
// 	fmt.Printf("web post : %v ", htmlTable)
// 	fmt.Println()
// }

// Function Gets $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
func GetAllPost() []Post {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "SELECT * FROM posts;"
	// rendu de la requête, recup info
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var res []Post
	for result.Next() {
		// debug console
		var post Post
		err := result.Scan(&post.Id, &post.UserName, &post.Tag, &post.Titre, &post.Description, &post.NbrLikes, &post.NbrDislikes, &post.Date)
		if err != nil {
			log.Fatal(err)
		}
		// %s = %v
		res = append(res, post)
	}

	return res
}

func GetOnePost(id string) Post {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "SELECT * FROM posts WHERE id= '" + id + "';"
	// rendu de la requête, recup info
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var post Post
	for result.Next() {
		err2 := result.Scan(&post.Id, &post.UserName, &post.Tag, &post.Titre, &post.Description, &post.NbrLikes, &post.NbrDislikes, &post.Date)
		if err2 != nil {
			log.Fatal(err)
		}
	}
	return post
}

func GetResponses(id string) []Reponse {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "SELECT * FROM reponses WHERE idPost = '" + id + "';"
	// rendu de la requête, recup info
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var res []Reponse
	for result.Next() {
		var reponse Reponse
		err := result.Scan(&reponse.Id, &reponse.IdPost, &reponse.UserName, &reponse.Contenu, &reponse.Date)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, reponse)
	}
	return res
}

func GetUser(email string) User {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "SELECT * FROM users WHERE email= '" + email + "';"
	// rendu de la requête, recup info
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var user User
	for result.Next() {
		err := result.Scan(&user.Id, &user.Email, &user.Username, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
	return user
}

func GetEmail(email string) bool {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "SELECT * FROM users WHERE email= '" + email + "';"
	// rendu de la requête, recup info
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var res []User
	for result.Next() {
		var user User
		err := result.Scan(&user.Id, &user.Email, &user.Username, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		// %s = %v
		res = append(res, user)
	}

	if len(res) != 0 {
		return true
	} else {
		return false
	}
	// return result
}

func GetTagFilm() []Post {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "SELECT * FROM posts WHERE tag='film';"
	// rendu de la requête, recup info
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var res []Post
	for result.Next() {
		// debug console
		var post Post
		err := result.Scan(&post.Id, &post.UserName, &post.Tag, &post.Titre, &post.Description, &post.NbrLikes, &post.NbrDislikes, &post.Date)
		if err != nil {
			log.Fatal(err)
		}
		// %s = %v
		res = append(res, post)
	}

	return res
}

func GetTagSerie() []Post {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "SELECT * FROM posts WHERE tag='serie';"
	// rendu de la requête, recup info
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var res []Post
	for result.Next() {
		// debug console
		var post Post
		err := result.Scan(&post.Id, &post.UserName, &post.Tag, &post.Titre, &post.Description, &post.NbrLikes, &post.NbrDislikes, &post.Date)
		if err != nil {
			log.Fatal(err)
		}
		// %s = %v
		res = append(res, post)
	}

	return res
}

func GetSession(mail string) bool {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "SELECT * FROM sessions  WHERE email='" + mail + "';"
	// rendu de la requête, recup info
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var res []Session
	for result.Next() {
		// debug console
		var session Session
		err := result.Scan(&session.Id, &session.Email, &session.Uuid)
		if err != nil {
			log.Fatal(err)
		}
		// %s = %v
		res = append(res, session)
	}

	if len(res) != 0 {
		return false
	} else {
		return true
	}
}

func DeleteSession(uuid string) {
	db := InitDatabase("forum.db")
	defer db.Close()

	query := "DELETE FROM sessions WHERE uuid = '" + uuid + "';"
	// rendu de la requête, recup info
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	// var res []Session
	// for result.Next() {
	//     // debug console
	//     var session Session
	//     err := result.Scan(&session.Id, &session.Email, &session.Uuid)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     // %s = %v
	//     res = append(res, session)
	// }
}

func SelectByAscending(filter string) []Post {
	db := InitDatabase("forum.db")
	defer db.Close()
	query := "SELECT * FROM posts ORDER BY " + filter + " ASC;"
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var res []Post
	for result.Next() {
		var post Post
		err := result.Scan(&post.Id, &post.UserName, &post.Tag, &post.Titre, &post.Description, &post.NbrLikes, &post.NbrDislikes, &post.Date)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, post)
	}
	return res
}

func SelectByDescending(filter string) []Post {
	db := InitDatabase("forum.db")
	defer db.Close()
	query := "SELECT * FROM posts ORDER BY " + filter + " DESC;"
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var res []Post
	for result.Next() {
		var post Post
		err := result.Scan(&post.Id, &post.UserName, &post.Tag, &post.Titre, &post.Description, &post.NbrLikes, &post.NbrDislikes, &post.Date)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, post)
	}
	return res
}

func RecupNbr(data string, id string) int {
	db := InitDatabase("forum.db")
	defer db.Close()

	quer := "SELECT " + data + " FROM posts WHERE id=" + id + ";"
	resul, e := db.Query(quer)
	if e != nil {
		log.Fatal(e)
	}
	var nbrLikes int
	for resul.Next() {
		er := resul.Scan(&nbrLikes)
		if er != nil {
			log.Fatal(er)
		}
		//fmt.Println(nbrLikes)
	}
	return nbrLikes
}

func UpdateNbr(data string, nbr int, id string) {
	db := InitDatabase("forum.db")
	defer db.Close()
	query := "UPDATE posts SET " + data + " = " + strconv.Itoa(nbr) + " WHERE id= " + id + ";"
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Result update NBR : ", result)
}
