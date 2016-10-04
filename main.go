package main

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "net"
   "time"
   "bytes"
   "bufio"
)

func proxyHandler(conn net.Conn) {
   defer conn.Close()
   conn.SetReadDeadline(time.Now().Add(10 * time.Second))
   fmt.Println("client accept!")
   messageBuf := make([]byte, 1024)
   messageLen, _ := conn.Read(messageBuf)

   message := messageBuf[:messageLen]

   buf := bufio.NewReader(bytes.NewReader(message))
   fmt.Println(buf)
   req, _ := http.ReadRequest(buf)
   fmt.Println(req)
   client := new(http.Client)
   req.RequestURI = ""
   resp, _ := client.Do(req)

   defer resp.Body.Close()
   fmt.Println(resp)

   dumpResp, _ := httputil.DumpResponse(resp, true)

   resp_message := string(dumpResp)
   conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
   conn.Write([]byte(resp_message))

}

func main() {
   port := ":8888"
   tcpAddr, _ := net.ResolveTCPAddr("tcp", port)
   listner, _ := net.ListenTCP("tcp", tcpAddr)
   for {
      conn, err := listner.Accept()
         if err != nil {
            continue
         }

      go proxyHandler(conn)
   }
//   http.HandleFunc("/", proxyHandler)
//   http.ListenAndServe("localhost:8888", nil)
}
