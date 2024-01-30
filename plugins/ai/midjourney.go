package ai

import (
  "inc/lib"
  "encoding/base64"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "midjourney",
    As:       []string{"midjourney"},
    Tags:     "ai",
    IsPrefix: true,
    IsQuerry: true,
     IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

        m.Reply("Processing 3 minutes...")
        
      type ApiResponse struct {
        Creator string `json:"creator"`
        Status  bool   `json:"status"`
        Base64  string `json:"base64"`
      }
      
      res := "https://skizo.tech/api/midjourney/v2?text="+m.Querry+"&apikey=batu" 

      resp, err := http.Get(res)
      if err != nil {
        fmt.Println("Error sending GET request:", err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error reading response body:", err)
        return
      }

      var apiResponse ApiResponse
      err = json.Unmarshal(body, &apiResponse)

      if err != nil {
          fmt.Println(err)
          return
      }
      
      decoded, err := base64.StdEncoding.DecodeString(apiResponse.Base64)
      if err != nil {
          fmt.Println("Error decoding string:", err)
          return
      }
      
      client.SendImage(m.From, decoded, "nihhh", m.ID)

    },
  })
}
