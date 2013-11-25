package main

import "net"
import "bufio"
import "fmt"

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
        if err != nil {
            // handle error
        }
    
    fmt.Fprintf(conn, "Message from Client\n")
    
    status, err := bufio.NewReader(conn).ReadString('\n')
    
    fmt.Printf(status)

}
