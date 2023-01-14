package main

import (
    "bytes"
    "fmt"
    "net/http"
    "net/url"
)

func main() {
    notifyURL := "https://notify-api.line.me/api/notify"
    token := "HSjKlkm6ylCcZoHgtaAbk37EZqOdW52lrXuIw1CFqNb"
    message := "Hello, World!"
    imageURL := "" // optional image URL

    data := url.Values{}
    data.Set("message", message)
    if imageURL != "" {
        data.Add("imageThumbnail", imageURL)
        data.Add("imageFullsize", imageURL)
    }

    client := &http.Client{}
    req, _ := http.NewRequest("POST", notifyURL, bytes.NewBufferString(data.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Authorization", "Bearer "+token)

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Println("Failed to send notification")
    } else {
        fmt.Println("Notification sent successfully")
    }
}
