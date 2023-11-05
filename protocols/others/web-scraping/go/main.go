package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()

    titles := make([]string, 0, 10)

    c.OnHTML(".css-1qiat4j", func(e *colly.HTMLElement) {
        if len(titles) < 10 {
            title := e.ChildText(".css-1echdzn")
            titles = append(titles, title)
        }
    })

    c.Visit("https://www.nytimes.com/")

    for i, title := range titles {
        fmt.Printf("%d. %s\n", i+1, title)
    }
}
