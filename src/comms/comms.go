package comms

import (
    "fmt"
    "encoding/binary"
    "io"
) 

func Recv_packet(conn io.Reader) (Type uint32) {
    err := binary.Read(conn, binary.BigEndian, &Type)

    if err != nil {
        return 99
    }

    return
}


type NegBody struct {
    BD_ADDR [6]byte
}

func Recv_neg(conn io.Reader) (BD_ADDR [6]byte) {
    var msg NegBody
    err := binary.Read(conn, binary.BigEndian, &msg)

    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }

    fmt.Println(msg)
    return msg.BD_ADDR
}

func Send_neg(conn io.Writer, BD_ADDR [6]byte) {
    Type := uint32(0)
    var msg NegBody
    
    msg.BD_ADDR = BD_ADDR

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

func Recv_init(conn io.Reader) (clock uint32, RAND, Link_key [16]byte) {
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

func Send_init(conn io.Writer, clock uint32, RAND, Link_key [16]byte) {
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

func Recv_data(conn io.Reader) (clock uint32, Data []byte) {
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

    clock, Data = msg.Clock, msg.Data
    return
}

func Send_data(conn io.Writer, clock uint32, Data []byte) {
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

