package main

import (
    "fmt"
    "net/http"
    "log"

    // import third party libraries
    "github.com/PuerkitoBio/goquery"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/meta", metaScrape)
    http.ListenAndServe(":8080", nil)
}


func metaScrape(w http.ResponseWriter, r *http.Request) {
    doc, err := goquery.NewDocument("http://jonathanmh.com")
    if err != nil {
        log.Fatal(err)
    }

    var metaDescription string
    var pageTitle string

    // use CSS selector found with the browser inspector
    // for each, use index and item
    pageTitle = doc.Find("title").Contents().Text()

    doc.Find("meta").Each(func(index int, item *goquery.Selection) {
        if( item.AttrOr("name","") == "description") {
            metaDescription = item.AttrOr("content", "")
        }
    })
    fmt.Printf("Page Title: '%s'\n", pageTitle)
    fmt.Printf("Meta Description: '%s'\n", metaDescription)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
