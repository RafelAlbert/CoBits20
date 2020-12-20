package main

import (
//  "net/http"
  "log"
  "os"
	"golang.org/x/oauth2"
  "path/filepath"
  "io/ioutil"
  "encoding/json"
)

var PROJECT_FOLDER string = ""

func initConstants(){
  if PROJECT_FOLDER == "" {
    dir, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), ".."))
    if err != nil {
      log.Fatal(err)
    }

    PROJECT_FOLDER=dir
  }

  data, err := ioutil.ReadFile("./secrets.json")
  if err != nil {
    log.Fatal("Secrets not found.")
  }

  var OAUTH_CONF_AUX oauth2.Config
  err = json.Unmarshal(data, &OAUTH_CONF_AUX)
  if err != nil {
    log.Fatal("error: ", err)
  }

  OAUTH_CONF=&OAUTH_CONF_AUX

}


