package server

import(
  "html/template"
  "log"
  "net/http"
  "path"
  "time"
)

func LaunchServer() {
  mux := http.NewServeMux()

  mux.HandleFunc("/", homeHandler)

  server := &http.Server{
    Addr:         ":"+SERVER_PORT,
    Handler:      mux,
    ReadTimeout:  READ_TIMEOUT * time.Second,
    WriteTimeout: WRITE_TIMEOUT * time.Second,
  }

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

//////////////////////////////////////////////////////

