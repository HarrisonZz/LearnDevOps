package main

import (

  "net/http"
  "log"
)

func main() {
     fileServer := http.FileServer(http.Dir("."))
    
    http.Handle("/", fileServer)

    log.Println("伺服器啟動於 http://localhost:8080")
    
    log.Fatal(http.ListenAndServe(":8080", nil))



}
