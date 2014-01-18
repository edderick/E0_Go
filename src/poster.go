package main 

import (
    "net/http"
    "fmt"
    "net/url"
)

func main() {
    resp, err := http.PostForm("http://127.0.0.1:6666/bar",  
        url.Values{"msg": {"Hello, World!"}})
    
    fmt.Println(resp)
    fmt.Println(err)
}
