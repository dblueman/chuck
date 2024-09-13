package main

import (
   "encoding/json"
   "fmt"
   "net/http"
)

func top() error {
   rsp, err := http.Get(`https://api.chucknorris.io/jokes/random`)
   if err != nil {
      return err
   }

   defer rsp.Body.Close()

   var joke struct {
      Value string
   }

   err = json.NewDecoder(rsp.Body).Decode(&joke)
   if err != nil {
      return err
   }

   fmt.Println(joke.Value)
   return nil
}

func main() {
   err := top()
   if err != nil {
      fmt.Println(err.Error())
   }
}
