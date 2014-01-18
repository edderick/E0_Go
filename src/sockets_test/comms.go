package main

import "flag"

import (
    "net"
    "fmt"
    "encoding/binary"
    "io"
) 

func recv_packet(conn io.Reader) uint32 {
    msg := new(uint32)

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

func recv_neg(conn io.Reader) uint32 {
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

func recv_init(conn io.Reader) (clock uint32, RAND, Link_key [16]byte) {
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
    Clock uint32
    Length uint32
    Data []byte
}

func recv_data(conn io.Reader) (clock uint32, Data []byte) {
    msg := new(DataBody)
   
    err := binary.Read(conn, binary.BigEndian, &msg.Clock)
    
    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }
    
    err = binary.Read(conn, binary.BigEndian, &msg.Length)
    
    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }
    
    msg.Data = make([]byte, msg.Length)

    err = binary.Read(conn, binary.BigEndian, msg.Data)

    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }

    fmt.Println(msg)

    clock, Data = msg.Clock, msg.Data
    return
}

func send_data(conn net.Conn, clock uint32, Data []byte) {
    var Type uint32
    Type = 2

    msg := new(DataBody)

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


func main() {
    isServerPtr := flag.Bool("server", false, "Run in server mode?")

    flag.Parse()

    fmt.Println("Is Server: ", *isServerPtr)

    var conn net.Conn

    if *isServerPtr {
        conn = server_main()
    } else {
        conn = client_main()
    }

    for i := uint32(0); i < 20; i++ {

        fmt.Println(i)

        send_neg(conn, 56)
        send_init(conn, i, [16]byte{}, [16]byte{}) 
        send_data(conn, i, []byte{100, 200, 120})
       
        for j := 0; j < 3; j++ {
            packet_type := recv_packet(conn)

            switch packet_type {
                case 0: recv_neg(conn)
                case 1: recv_init(conn)
                case 2: recv_data(conn)
                case 99: break
                
            }
        }
    }

    conn.Close()

}
