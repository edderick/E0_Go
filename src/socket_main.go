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

    return conn
}

func client_main() net.Conn {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")

    if err != nil {
        fmt.Println("net.Dial failed:", err)
    }

    return conn
}

func is_bigger(ours, theirs [6]byte) bool{
    //TODO: This... Note the bit ordering..
    return ours[0] > theirs[0]
}

func receiver(conn io.ReadWriter, s *State) {
  LOOP:
  for {
        packet_type := comms.Recv_packet(conn)

        switch packet_type {
            case 0: 
                OTHER_BD_ADDR := comms.Recv_neg(conn)
                if is_bigger(s.BD_ADDR, OTHER_BD_ADDR) {
                    comms.Send_init(conn, s.clk, [16]byte{}, s.Kc)
                } else {
                    s.BD_ADDR = OTHER_BD_ADDR
                }
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
   
    state.Kc = [16]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    state.clk = 0
    
    if *isServerPtr {
        state.BD_ADDR = [6]byte{255, 255, 255, 255, 255, 255}
        conn = server_main(&state)
        p = ":8888"
    } else {
        state.BD_ADDR = [6]byte{0, 0, 0, 0, 0, 0}
        conn = client_main()
        p = ":6666"
    }

    comms.Send_neg(conn, state.BD_ADDR)
   
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
