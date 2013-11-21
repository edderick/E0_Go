package main

import "fmt"
import "net"
import "io"
import "os"

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
            fmt.Printf("Established Connection")
            // Echo all incoming data.
            io.Copy(os.Stdout, c)
            // Shut down the connection.
            c.Close()
        }(conn)
    }

}
