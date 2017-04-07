package main

import (
  "log"
  "time"
  "fmt"
  "github.com/jochasinga/requests"
)

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main(){
    data := map[string]string{
      "res_id":"1",
      "title":"Assistant Restaurant Manager",
      "description":"Managing restaurant for clients and business.",
    }

    fmt.Println(data)
    defer timeTrack(time.Now(), "url")
    res, err := requests.PostJSON("http://52.228.33.214:8888/semantic_match/v1.0/job", data)
    if err != nil {
      panic(err)
    }
    fmt.Println(res.StatusCode)

}
