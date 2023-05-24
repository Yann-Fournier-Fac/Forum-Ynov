package main

import (
	"Josh/database"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type PageHome struct {
	Posts           []database.Post
	IsConnecter     bool
	ConnectUserInfo string
	Try             bool
}

type PagePost struct {
	OnePost     database.Post
	Responses   []database.Reponse
	IsConnecter bool
}

func ResetDB() {
	database.Database()

	database.DatabaseAndUsers([]string{"yann@ynov.com", "Yann", HashPassword("yann")})
	database.DatabaseAndUsers([]string{"elisa@ynov.com", "Elisa", HashPassword("elisa")})
	database.DatabaseAndUsers([]string{"kevin@ynov.com", "Kévin", HashPassword("kevin")})
	database.DatabaseAndUsers([]string{"liliane@ynov.com", "Liliane", HashPassword("liliane")})
	database.DatabaseAndUsers([]string{"joshua@ynov.com", "Joshua", HashPassword("joshua")})

	database.DatabaseAndPost([]string{strconv.Itoa(1), "film", "First Post", "Moi j'adore ET", strconv.Itoa(55), strconv.Itoa(3), strconv.Itoa(0), strconv.Itoa(0)})
	database.DatabaseAndPost([]string{strconv.Itoa(1), "serie", "Second Post", "Moi j'adore GOT", strconv.Itoa(2), strconv.Itoa(33)})
	database.DatabaseAndPost([]string{strconv.Itoa(2), "film", "First Post", "Moi j'adore ET", strconv.Itoa(55), strconv.Itoa(3)})
	database.DatabaseAndPost([]string{strconv.Itoa(2), "serie", "Second Post", "Moi j'adore Got", strconv.Itoa(2), strconv.Itoa(33)})
	database.DatabaseAndPost([]string{strconv.Itoa(3), "film", "First Post", "Moi j'adore ET", strconv.Itoa(55), strconv.Itoa(3)})
	database.DatabaseAndPost([]string{strconv.Itoa(3), "serie", "Second Post", "Moi j'adore Got", strconv.Itoa(2), strconv.Itoa(33)})
	database.DatabaseAndPost([]string{strconv.Itoa(4), "film", "First Post", "Moi j'adore ET", strconv.Itoa(55), strconv.Itoa(3)})
	database.DatabaseAndPost([]string{strconv.Itoa(4), "serie", "Second Post", "Moi j'adore Got", strconv.Itoa(2), strconv.Itoa(33)})
	database.DatabaseAndPost([]string{strconv.Itoa(5), "film", "First Post", "Moi j'adore ET", strconv.Itoa(55), strconv.Itoa(3)})
	database.DatabaseAndPost([]string{strconv.Itoa(5), "serie", "Second Post", "Moi j'adore Got", strconv.Itoa(2), strconv.Itoa(33)})

	database.DatabaseAndReponse([]string{strconv.Itoa(5), strconv.Itoa(9), "Moi aussi !!!!!"})
	database.DatabaseAndReponse([]string{strconv.Itoa(5), strconv.Itoa(9), "Moi aussi !!!!!"})
	database.DatabaseAndReponse([]string{strconv.Itoa(5), strconv.Itoa(9), "Moi aussi !!!!!"})
	database.DatabaseAndReponse([]string{strconv.Itoa(5), strconv.Itoa(9), "Moi aussi !!!!!"})

	database.DatabaseAndSession([]string{"aze"})
	database.DatabaseAndSession([]string{"qsd"})
}

func initStruct() (PageHome, PagePost) {
	var home PageHome
	home.Posts = database.GetAllPost()
	home.IsConnecter = false

	var post PagePost
	post.IsConnecter = false

	return home, post
}

var tmplHome = template.Must(template.ParseFiles("./html/home.html"))
var tmplPost = template.Must(template.ParseFiles("./html/post.html"))
var tmplNewPost = template.Must(template.ParseFiles("./html/newpost.html"))
var HomeStruct, PostStruct = initStruct()

// var HomeStruct PageHome

func main() {

	// ResetDB()

	fmt.Printf("\n")
	fmt.Println("http://localhost:8080/")
	fmt.Printf("\n")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/newpost", newPostHandler)

	// http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css"))))
	// http.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("assets/images"))))

	http.ListenAndServe(":8080", nil)
}

var truc = "home"

func filter(HomeStruct *PageHome) {
	if truc == "home" {
		HomeStruct.Posts = database.GetAllPost()
	} else if truc == "films" {
		HomeStruct.Posts = database.GetTagFilm()
	} else if truc == "series" {
		HomeStruct.Posts = database.GetTagSerie()
	} else if truc == "Plikes" {
		HomeStruct.Posts = database.SelectByDescending("nbrLikes")
	} else if truc == "Pdislikes" {
		HomeStruct.Posts = database.SelectByDescending("NbrDislikes")
	} else if truc == "Mlikes" {
		HomeStruct.Posts = database.SelectByAscending("nbrLikes")
	} else if truc == "Mdislikes" {
		HomeStruct.Posts = database.SelectByAscending("NbrDislikes")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	buLikesDislikes := r.FormValue("bulike/dislike")
	if buLikesDislikes != "" {
		idBu := strings.Split(buLikesDislikes, ",")
		if idBu[1] == "like" {
			nbr := database.RecupNbr("nbrLikes", idBu[0])
			nbr += 1
			database.UpdateNbr("nbrLikes", nbr, idBu[0])
			//fmt.Println(idBu[0], "Like")
		} else if idBu[1] == "dislike" {
			nbr := database.RecupNbr("nbrDislikes", idBu[0])
			nbr += 1
			database.UpdateNbr("nbrDislikes", nbr, idBu[0])
			fmt.Println(idBu[0], "Dislike")
		}
	}

	// Connection -----------------------------------------------------------------
	connectionEmail := r.FormValue("ConnectionEmail")
	connectionMdp := r.FormValue("ConnectionMdp")
	if connectionEmail != "" {
		user := database.GetUser(connectionEmail)
		fmt.Println(uuid.NewRandom())
		fmt.Print(user)
		if CheckPasswordHash(connectionMdp, user.Password) && user.Email == connectionEmail {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}

	// Inscription ----------------------------------------------------------------
	inscriptionName := r.FormValue("InscriptionName")
	inscriptionEmail := r.FormValue("InscriptionEmail")
	inscriptionMdp := r.FormValue("InscriptionMdp")
	fmt.Println(inscriptionName)

	if inscriptionName != "" {
		if !database.GetEmail(inscriptionEmail) {
			database.DatabaseAndUsers([]string{inscriptionEmail, inscriptionName, HashPassword(inscriptionMdp)})
		} else {
			fmt.Println("veuillez entrer une autre adresse mail. Celle-ci est déjà prise.")
		}
	}

	headerLinks := r.FormValue("link")
	if headerLinks != "" {
		truc = headerLinks
	}

	BuMenuDeroulant := r.FormValue("BuMenuDeroulant")
	if BuMenuDeroulant != "" {
		truc = BuMenuDeroulant
	}

	filter(&HomeStruct)

	err := tmplHome.Execute(w, HomeStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	err := tmplPost.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {
	err := tmplNewPost.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
