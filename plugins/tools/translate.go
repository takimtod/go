package tools

import (
  "fmt"
  "inc/lib"
  "encoding/json"
  "io/ioutil"
  "net/http"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(tr|translate)",
    As:       []string{"translate"},
    Tags:     "tools",
    IsPrefix: true,
    IsQuerry: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {
     
      type TranslateResponse struct {
        Creator string `json:"creator"`
        Result string `json:"result"`
      }

        url := "https://skizo.tech/api/translate?text=" + m.Querry + "&lang=id&apikey=batu"
        resp, err := http.Get(url)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }

        var result TranslateResponse
        err = json.Unmarshal(body, &result)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }

        m.Reply(result.Result)

    },
  })
}
