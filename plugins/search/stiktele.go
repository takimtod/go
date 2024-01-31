package search

import (
  "inc/lib" 
 "fmt"
  "inc/lib/api"
  "inc/lib/typings"
  "time"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "stiktele",
    As:       []string{"stiktele"},
    Tags:     "search",
    IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      resp, err := lib.Stiktele(m.Querry)
      if err != nil {
        fmt.Println(err)
      }

    for _, sticker := range resp["sticker"].([]string) {
      time.Sleep(5 * time.Second)

    bytes, err := client.GetBytes(sticker)
    if err != nil {
       fmt.Println("Error:", err)
      return
    }
      
      s := api.StickerApi(&typings.Sticker{
        File: bytes,
        Tipe: func() typings.MediaType {
          if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
            return typings.IMAGE
          } else if m.IsVideo || m.IsQuotedVideo {
            return typings.VIDEO
          } else {
            return typings.TEKS
          }

        }(),
      }, &typings.MetadataSticker{
        Author:    m.PushName,
        Pack:      "https://s.id/ryuubot",
        KeepScale: true,
        Removebg:  "true",
        Circle: func() bool {
          if m.Querry == "-c" {
            return true
          } else {
            return false
          }
        }(),
      })

      client.SendSticker(m.From, s.Build(), m.ID)
      
    }
  
    },
  })
}
