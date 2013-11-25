package main

import "fmt"
import "net"
import "bufio"

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
            // handle error
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
            continue
        }
        go func(c net.Conn) {
            fmt.Printf("Established Connection\n")
            // Echo all incoming data.

            status, _ := bufio.NewReader(c).ReadString('\n')
            fmt.Printf(status)

            fmt.Fprintf(c, "Thank you for your message\n")

            // Shut down the connection.
            c.Close()
        }(conn)
    }

}
