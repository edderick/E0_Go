package main 

import (
    "net/http"
    "fmt"
    "net/url"
    "encoding/base64"
)

func main() {
    str := base64.StdEncoding.EncodeToString([]byte("Hello, playground"))

    resp, err := http.PostForm("http://127.0.0.1:8000/keyStream?role=master",  
        url.Values{"msg": { str  }})
    
    fmt.Println(resp)
    fmt.Println(err)
}
