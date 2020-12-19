package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
  log.Print("Init server...")

  initConstants()

  router := mux.NewRouter()

  router.HandleFunc("/admin", apiAdminEndpoint).Methods("GET")


  log.Fatal(http.ListenAndServe(":3000", router))
}
