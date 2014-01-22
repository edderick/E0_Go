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
        "encoding/hex"
        "encoding/base64"
        "time"
       )

type State struct {
    Kc [16]byte
    clk uint32
    BD_ADDR [6]byte
    is_master bool
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
                if !is_bigger(s.BD_ADDR, OTHER_BD_ADDR) {
                    s.BD_ADDR = OTHER_BD_ADDR
                    s.is_master = true
                } else {
                    s.is_master = false
                }
            case 1: 
                s.clk, _, s.Kc = comms.Recv_init(conn)
            case 2: 
                var msg []byte
                s.clk, msg = comms.Recv_data(conn)

                fmt.Println("Recieved: ", string(msg))

                keyStream := EncryptionEngine.GetKeyStream(s.Kc, s.BD_ADDR, s.clk, len(msg))  
                decrypted_msg := EncryptionEngine.Encrypt(msg, keyStream) 

                fmt.Println("Decypted as: ", string(decrypted_msg))
               
                ciphertext_b64 := base64.StdEncoding.EncodeToString(msg)
                keystream_b64 := base64.StdEncoding.EncodeToString(keyStream)

                var role string
                
                if s.is_master {
                    role = "master"
                } else {
                    role = "slave"
                }

                _, err := http.PostForm("http://127.0.0.1:8000/log?role=" + role,   
                        url.Values{
                        "CLK" : { string(s.clk) },
                        "is_receiving" : { "true" },
                        "keystream" : { keystream_b64 },
                        "ciphertext" : { ciphertext_b64 },
                        "plaintext" : { string(decrypted_msg) },
                        "timestamp" : { time.Now().Format("Jan _2 15:04:05") },
                        })
                    
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

    fmt.Println("Running on ", p)

    comms.Send_neg(conn, state.BD_ADDR)
   
    go receiver(conn, &state)
   
    http.HandleFunc("/Kc", func(w http.ResponseWriter, r *http.Request) {
        var Kc_str string
        if r.Method == "POST" {
            r.ParseForm()
            Kc_str = r.PostForm["Kc"][0] 
            fmt.Println("Kc String: ", Kc_str)
        }
        
        Kc, err := hex.DecodeString(Kc_str)
        
        if err != nil || len(Kc) < 16 {
            fmt.Println("Invalid key!")
            return 
        }
        
        fmt.Println("Kc: ", Kc)
    
        for i:= 0; i < 16; i++ {
            state.Kc[i] = Kc[i]
        }

    })

    http.HandleFunc("/isMaster", func(w http.ResponseWriter, r *http.Request) {
        if state.is_master {
            fmt.Fprintf(w, "true")
        } else {
            fmt.Fprintf(w, "false")
        }
    })

    http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
        
        if r.Method == "POST" {
            state.clk++
            r.ParseForm()
             
            pt := []byte(r.PostForm["plaintext"][0]) 
            fmt.Println("Sending: ", string(pt))
            
            fmt.Fprintf(w, html.EscapeString("Sending packet"))

            keyStream := EncryptionEngine.GetKeyStream(
                state.Kc, state.BD_ADDR, state.clk, len(pt))  
            msg := EncryptionEngine.Encrypt(pt, keyStream) 

            comms.Send_data(conn, state.clk, msg)
            
            fmt.Println("Encrypted as: ", string(msg))

            keystream_b64 := base64.StdEncoding.EncodeToString(keyStream)
            ciphertext_b64 := base64.StdEncoding.EncodeToString(msg)

            var role string
            
            if state.is_master {
                role = "master"
            } else {
                role = "slave"
            }
    
            _, err := http.PostForm("http://127.0.0.1:8000/log?role=" + role,   
                    url.Values{
                    "CLK" : { string(state.clk) },
                    "is_receiving" : { "false" },
                    "keystream" : { keystream_b64 },
                    "ciphertext" : { ciphertext_b64 },
                    "plaintext" : { string(pt) },
                    "timestamp" : { time.Now().Format("Jan _2 15:04:05") },
                    })   
            
            if err != nil {
                    fmt.Println("There was an http error: ", err)
            }
        }
    })

    log.Fatal(http.ListenAndServe(p, nil))

    conn.Close()
}
