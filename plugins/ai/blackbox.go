package ai

import (
  "inc/lib"
  "fmt"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(bb|blackbox)",
    As:       []string{"blackbox"},
    Tags:     "ai",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      res, err := lib.Blackbox(m.Querry)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
         m.Reply(res)
 
    },
  })
}