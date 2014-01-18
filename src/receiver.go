package main 

import (
    "net/http"
    "fmt"
    "log"
)

func main() {
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Ooh a packet!")
        if r.Method == "POST" {
            r.ParseForm()
            fmt.Println("Received: ", r.PostForm["msg"][0])
        }

    })
    log.Fatal(http.ListenAndServe(":9999", nil))
}
