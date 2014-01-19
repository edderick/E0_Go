package main 

import (
        "./comms"
        "./EncryptionEngine"

        "flag"

        "net"
        "fmt"
        "io"

        "net/http"
        "net/url"
        "html"
        "log"
       )

type State struct {
    Kc [16]byte
    clk uint32
    BD_ADDR [6]byte
}


func server_main(s *State) net.Conn {
    ln, err := net.Listen("tcp", ":8080")
   
    if err != nil {
        fmt.Println("net.Listen failed:", err)
    }

    conn, err := ln.Accept()
    
    fmt.Printf("Accepted Connection\n")
    
    if err != nil {
        fmt.Println("ln.Accept failed:", err)
    }

    s.BD_ADDR = [6]byte{}
    s.Kc = [16]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    s.clk = 0

    comms.Send_neg(conn, s.BD_ADDR)
    comms.Send_init(conn, s.clk, [16]byte{}, s.Kc)

    return conn
}

func client_main() net.Conn {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")

    if err != nil {
        fmt.Println("net.Dial failed:", err)
    }

    return conn
}


func receiver(conn io.Reader, s *State) {
  LOOP:
  for {
        packet_type := comms.Recv_packet(conn)

        switch packet_type {
            case 0: 
                s.BD_ADDR = comms.Recv_neg(conn)
            case 1: 
                s.clk, _, s.Kc = comms.Recv_init(conn)
            case 2: 
                var msg []byte
                s.clk, msg = comms.Recv_data(conn)

                fmt.Println("Recieved: ", string(msg))

                keyStream := EncryptionEngine.GetKeyStream(s.Kc, s.BD_ADDR, s.clk, len(msg))  
                msg = EncryptionEngine.Encrypt(msg, keyStream) 

                fmt.Println("Decypted as: ", string(msg))

                _, err := http.PostForm("http://127.0.0.1:9999/bar", 
                    url.Values{"msg": {string(msg)}})
        
                if err != nil {
                    fmt.Println("There was an http error: ", err)
                }

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

    var state State
    
    if *isServerPtr {
        conn = server_main(&state)
        p = ":8888"
    } else {
        conn = client_main()
        p = ":6666"
    }


    go receiver(conn, &state)
   
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        
        if r.Method == "POST" {
            state.clk++
            r.ParseForm()
             
            fmt.Println("Sending: ", r.PostForm["msg"][0])
            
            fmt.Fprintf(w, html.EscapeString("Sending packet"))

            pt := []byte(r.PostForm["msg"][0]) 
    
            keyStream := EncryptionEngine.GetKeyStream(
                state.Kc, state.BD_ADDR, state.clk, len(pt))  
            msg := EncryptionEngine.Encrypt(pt, keyStream) 

            comms.Send_data(conn, state.clk, msg)
            
            fmt.Println("Encrypted as: ", string(msg))
        }
    })

    log.Fatal(http.ListenAndServe(p, nil))

    conn.Close()
}
