package server

import(
  "net/http"
  "time"
  "log"
)

func LaunchServer() {
  mux := http.NewServeMux()

  mux.HandleFunc("/", myFunc)

  server := &http.Server{
    Addr:         ":"+SERVER_PORT,
    Handler:      mux,
    ReadTimeout:  READ_TIMEOUT * time.Second,
    WriteTimeout: WRITE_TIMEOUT * time.Second,
  }

  log.Printf("Launched server on port %v\n", SERVER_PORT)
  log.Fatal(server.ListenAndServe())
}

func myFunc(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)
  w.Write([]byte("Test"))
}

