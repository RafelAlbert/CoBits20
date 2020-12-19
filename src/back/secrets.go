package main

import (
//  "net/http"
  "log"
  "os"
  "path/filepath"
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
}

