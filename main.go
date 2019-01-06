package main

import (
    "net/http"
    f "fmt"
    "go-tax/controller"
)

func main() {

    http.HandleFunc("/tax", func(w http.ResponseWriter, r *http.Request) {
      switch r.Method {
        case http.MethodPost:
          controller.TaxControllerStore(w, r)
        case http.MethodGet:
          controller.TaxControllerGet(w, r)
        default:
          f.Println("error")
      }
    })

    server := new(http.Server)
    server.Addr = ":9000"

    f.Println("server started at localhost:9000")
    server.ListenAndServe()
}
