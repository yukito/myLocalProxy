package main

import (
   "net/http"
   "io/ioutil"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {

   client := new(http.Client)
   r.RequestURI = ""
   resp, _ := client.Do(r)

   body, _ := ioutil.ReadAll(resp.Body)
   for key, value := range resp.Header{
      w.Header().Set(key, value[0])
   }
   w.Write(body)
   w.WriteHeader(resp.StatusCode)

}

func main() {
   http.HandleFunc("/", proxyHandler)
   http.ListenAndServe("localhost:8888", nil)
}
