package main

import (
  "net/http"
  "fmt"
)

func apiAdminEndpoint(w http.ResponseWriter, r *http.Request){
  fmt.Println(PROJECT_FOLDER +"/front/admin_login.html" )
  http.ServeFile(w, r, PROJECT_FOLDER +  "/front/admin_login.html")  
}
