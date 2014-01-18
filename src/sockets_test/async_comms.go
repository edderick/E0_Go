package main

import "flag"

import (
    "net"
    "fmt"
    "encoding/binary"
    "io"

    "net/http"
    "net/url"
    "html"
    "log"
) 

func recv_packet(conn io.Reader) (Type uint32) {
    err := binary.Read(conn, binary.BigEndian, &Type)

    if err != nil {
        return 99
    }

    return
}


type NegBody struct {
    ID uint32
}

func recv_neg(conn io.Reader) (ID uint32) {
    var msg NegBody
    err := binary.Read(conn, binary.BigEndian, &msg)

    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }

    fmt.Println(msg)
    return msg.ID
}

func send_neg(conn io.Writer, ID uint32) {
    Type := uint32(0)
    var msg NegBody
    
    msg.ID = ID

    err := binary.Write(conn, binary.BigEndian, Type)

    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    
    err = binary.Write(conn, binary.BigEndian, msg)

    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
}


type InitBody struct {
    Clock uint32
    RAND [16]byte
    Link_key [16]byte
}

func recv_init(conn io.Reader) (clock uint32, RAND, Link_key [16]byte) {
    var msg InitBody

    err := binary.Read(conn, binary.BigEndian, &msg)

    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }

    fmt.Println(msg)
        
    clock = msg.Clock
    RAND = msg.RAND
    Link_key = msg.Link_key
    return
}

func send_init(conn io.Writer, clock uint32, RAND, Link_key [16]byte) {
    Type := uint32(1)
    var msg InitBody

    msg.Clock = clock
    msg.RAND = RAND
    msg.Link_key = Link_key
        
    err := binary.Write(conn, binary.BigEndian, Type)

    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    
    err = binary.Write(conn, binary.BigEndian, msg)

    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
}


type DataBody struct {
    Clock uint32
    Length uint32
    Data []byte
}

func recv_data(conn io.Reader) (clock uint32, Data []byte) {
    var msg DataBody
   
    err := binary.Read(conn, binary.BigEndian, &msg.Clock)
    
    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }
    
    err = binary.Read(conn, binary.BigEndian, &msg.Length)
    
    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }
    
    msg.Data = make([]byte, msg.Length)

    err = binary.Read(conn, binary.BigEndian, &msg.Data)

    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }

        
    fmt.Print(msg, " -- ")
    fmt.Println(string(msg.Data))
    
    _, err = http.PostForm("http://127.0.0.1:9999/bar", 
            url.Values{"msg": {string(msg.Data)}})
        
    if err != nil {
        fmt.Println("There was an http error: ", err)
    }

    clock, Data = msg.Clock, msg.Data
    return
}

func send_data(conn io.Writer, clock uint32, Data []byte) {
    Type := uint32(2)

    var msg DataBody

    msg.Data = Data
    msg.Length = uint32(len(msg.Data))
    msg.Clock = clock
   
    err := binary.Write(conn, binary.BigEndian, Type)

    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    
    err = binary.Write(conn, binary.BigEndian, msg.Clock)

    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    
    err = binary.Write(conn, binary.BigEndian, msg.Length)
    
    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    
    err = binary.Write(conn, binary.BigEndian, msg.Data)

    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
}



func server_main() net.Conn {
    ln, err := net.Listen("tcp", ":8080")
    
    if err != nil {
        fmt.Println("net.Listen failed:", err)
    }

    conn, err := ln.Accept()
    
    fmt.Printf("Accepted Connection\n")
    
    if err != nil {
        fmt.Println("ln.Accept failed:", err)
    }

    return conn
}

func client_main() net.Conn {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")

    if err != nil {
        fmt.Println("net.Dial failed:", err)
    }

    return conn
}

func receiver(conn io.Reader) {
  LOOP:
  for {
        packet_type := recv_packet(conn)

        switch packet_type {
            case 0: recv_neg(conn)
            case 1: recv_init(conn)
            case 2: recv_data(conn)
            case 99: break LOOP
        }
    }
}

func main() {
    isServerPtr := flag.Bool("server", false, "Run in server mode?")
    flag.Parse()
    fmt.Println("Is Server: ", *isServerPtr)

    var conn net.Conn
    var p string

    if *isServerPtr {
        conn = server_main()
        p = ":8888"
    } else {
        conn = client_main()
        p = ":6666"
    }

    go receiver(conn)
   
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        
        if r.Method == "POST" {
            r.ParseForm()
             
            fmt.Println("Requested to send: ", r.PostForm["msg"][0])
            
            fmt.Fprintf(w, html.EscapeString("Sending packet"))
            send_data(conn, 0, []byte(r.PostForm["msg"][0]))
        }
    })

    log.Fatal(http.ListenAndServe(p, nil))

    conn.Close()
}
