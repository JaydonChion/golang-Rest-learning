package main

import(
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
//	"strconv"
//	"math/rand"
	"fmt"
	"database/sql"
  	_"github.com/go-sql-driver/mysql"
)

type Post struct{

	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`

}

var err error
var db *sql.DB


func main(){
	db,err = sql.Open("mysql","root:XXjetherngXX11234@tcp(127.0.0.1:3306)/userpost")


	if err != nil {
		fmt.Println("fail connection")
		panic(err.Error())

	}else{
		fmt.Println("connection is successful")
	}

	defer db.Close()

	//posts = append(posts, Post{ID: "1", Title: "My first post", 
//	Body:      "This is the content of my first post"})
	router := mux.NewRouter()
	router.HandleFunc("/posts",getPosts).Methods("GET")
  	router.HandleFunc("/posts", createPost).Methods("POST")
  	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
  	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
  	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	http.ListenAndServe(":3000",router)
}


func  getPosts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	posts, err := db.Query("SELECT * FROM posts")
	if err != nil{
		fmt.Println("error occurs while quering database")
		panic (err.Error())
	}else{

		fmt.Println("managed to retrieve from database")
	}
	for posts.Next(){

	var post Post


	err := posts.Scan(&post.ID, &post.Title, &post.Body)
	if err != nil{
		fmt.Println("error at storing values")
		panic(err.Error())

	}else{

		fmt.Println("managed to store values")
	}
	json.NewEncoder(w).Encode(post)
	fmt.Println(post)
	}
}

func createPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
//	var post Post
//	_ = json.NewDecoder(r.Body).Decode(post)
//	post.ID = strconv.Itoa(rand.Intn(10000000))
//	posts = append(posts,post)
//	json.NewEncoder(w).Encode(&post)

}


func getPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
//	params := mux.Vars(r)
//	for _,item :=range posts{
//		if item.ID == params["id"]{
//			json.NewEncoder(w).Encode(item)
//		}
//		return
//	}
	json.NewEncoder(w).Encode(&Post{})


}

func updatePost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
//	params := mux.Vars(r)
//	for index, item := range posts{
//		if item.ID == params["id"]{
		//	posts = append(posts[:index],posts[index+1:]...)
//			var post Post
	//		_ = json.NewDecoder(r.Body).Decode(&post)
//			posts = append(posts,post)
//			json.NewEncoder(w).Encode(&post)
//
//			return
//		}
//	}

}




func deletePost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
//	params := mux.Vars(r)
//	for index,item := range posts{
//
//		if item.ID == params["id"]{
//			posts = append(posts[:index],posts[index+1:]...)
//			break
//		}
//	}
//	json.NewEncoder(w).Encode(posts)
}
