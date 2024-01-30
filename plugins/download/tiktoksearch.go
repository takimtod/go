package download

import (
  "inc/lib"
  "fmt"
  "net/http"
   "encoding/json"
  "io/ioutil"
)

func init() {
  lib.NewCommands(
    &lib.ICommand{
    Name:     "(ttsearch|tiktoksearch)",
    As:       []string{"ttsearch"},
    Tags:     "downloader",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      // Mengambil informasi tentang video
      url := "https://skizo.tech/api/ttsearch?search=" + m.Querry + "&apikey=kertas"
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

      var res struct {
        Title      string `json:"title"`
        Region     string `json:"region"`
        Music      string `json:"music"`
        MusicInfo struct {
          Title string `json:"title"`
          Play   string `json:"play"`
          Author string `json:"author"`
        } `json:"music_info"`
        Play string `json:"play"`
      }

      err = json.Unmarshal(body, &res)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      caption := fmt.Sprintf(`*TIKTOK SEARCH*

        *𖦹 Judul:* %s
        *𖦹 Region:* %s
        *𖦹 Musik:* %s
       *- Musik Info:*
              *• Judul:* %s
              *• Link:* %s
              *• Author:* %s
        `, res.Title, res.Region, res.Music, res.MusicInfo.Title, res.MusicInfo.Play, res.MusicInfo.Author)

      // Simulasi mengirim gambar
      bytes, err := client.GetBytes(res.Play)
      if err != nil {
        m.Reply(err.Error())
        return
      }
      client.SendVideo(m.From, bytes, caption, m.ID)
    },
  })
}