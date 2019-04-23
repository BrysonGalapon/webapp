package server

import (
	"github.com/BrysonGalapon/webapp/database"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var dbHandler *database.DBHandler

func LaunchServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/view/", viewHandler)
	mux.HandleFunc("/insert/", insertHandler)
	mux.HandleFunc("/delete/", deleteHandler)

	server := &http.Server{
		Addr:         ":" + SERVER_PORT,
		Handler:      mux,
		ReadTimeout:  READ_TIMEOUT * time.Second,
		WriteTimeout: WRITE_TIMEOUT * time.Second,
	}

	dbHandler = database.LaunchDB()

	log.Printf("Launched server on port %v\n", SERVER_PORT)
	log.Fatal(server.ListenAndServe())
}

//////////////// HANDLER FUNCTIONS ///////////////////

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		t   *template.Template
	)

	filename := path.Join(STATIC_FOLDER, HTML_FOLDER, "index.html")

	t, err = template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	type ViewCapture struct {
		Username string
		Password string
	}

	rows := dbHandler.View([]string{"username", "password"}, "users")

	for rows.Next() {
		var v ViewCapture

		err := rows.Scan(&v.Username, &v.Password)

		if err != nil {
			log.Fatal(err)
		}

		log.Println(v)
	}
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	dbHandler.Insert([]string{"username", "password", "usertype"}, []string{`"test"`, `"test"`, `"test"`}, "users")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	dbHandler.Delete([]string{"username", "password", "usertype"}, []string{`"test"`, `"test"`, `"test"`}, "users")
}

//////////////////////////////////////////////////////
