package main

import "flag"

import (
    "net"
    "fmt"
    "encoding/binary"
    "time"
) 

func recv_packet(conn net.Conn) uint32 {
    msg := new(uint32)

    conn.SetDeadline(time.Now().Add(100*time.Second)) 
    conn.SetReadDeadline(time.Now().Add(100*time.Second)) 
    err := binary.Read(conn, binary.BigEndian, msg)

    if err != nil {
        if err.Error() == "EOF" {
            return 99
        }
        fmt.Println("binary.Read failed:", err)
    }

    return *msg
}

type NegBody struct {
    ID uint32
}

func send_neg(conn net.Conn, ID uint32) {
    var Type uint32
    Type = 0

    msg := new(NegBody)
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

func recv_neg(conn net.Conn) uint32 {
    msg := new(NegBody)

    err := binary.Read(conn, binary.BigEndian, msg)

    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }

    fmt.Println(msg)
    return msg.ID
}


type InitBody struct {
    Clock uint32
    RAND [16]byte
    Link_key [16]byte
}

func recv_init(conn net.Conn) (clock uint32, RAND, Link_key [16]byte) {
    msg := new(InitBody)

    err := binary.Read(conn, binary.BigEndian, msg)

    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }

    fmt.Println(msg)
        
    clock = msg.Clock
    RAND = msg.RAND
    Link_key = msg.Link_key
    return
}

func send_init(conn net.Conn, clock uint32, RAND, Link_key [16]byte) {
    var Type uint32
    Type = 1

    msg := new(InitBody)

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
    Length uint32
    Data []byte
}

func recv_data(conn net.Conn) (Data []byte) {
    msg := new(DataBody)
   
    err := binary.Read(conn, binary.BigEndian, &msg.Length)
    
    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }
    
    msg.Data = make([]byte, msg.Length)

    err = binary.Read(conn, binary.BigEndian, msg.Data)

    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }

    fmt.Println(msg)

    return msg.Data
}

func send_data(conn net.Conn, Data []byte) {
    var Type uint32
    Type = 2

    msg := new(DataBody)

    msg.Data = Data
    msg.Length = uint32(len(msg.Data))
   
    err := binary.Write(conn, binary.BigEndian, Type)

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



func server_main() {
    ln, err := net.Listen("tcp", ":8080")
    
    if err != nil {
        fmt.Println("net.Listen failed:", err)
    }

    for {
        conn, err := ln.Accept()
        
        
        if err != nil {
            fmt.Println("ln.Accept failed:", err)
            continue
        }

        go func(c net.Conn) {
            fmt.Printf("Established Connection\n")
            packet_type := recv_packet(conn)

            switch packet_type {
                case 0: recv_neg(conn)
                case 1: recv_init(conn)
                case 2: recv_data(conn)
                case 99: break
            }
            c.Close()
        }(conn)
    }
}


func client_main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Println("net.Dial failed:", err)
    }
  
    send_neg(conn, 56)


    conn, err = net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Println("net.Dial failed:", err)
    }
    
    send_init(conn, 1, [16]byte{}, [16]byte{}) 
    
    
    conn, err = net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Println("net.Dial failed:", err)
    }

    send_data(conn, []byte{100, 200, 120})
}


func main() {
    isServerPtr := flag.Bool("server", false, "Run in server mode?")

    flag.Parse()

    fmt.Println("Is Server: ", *isServerPtr)

    if *isServerPtr {
        server_main()
    } else {
        client_main()
    }
    
}
