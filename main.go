package main

import (
    "fmt"
    "net/http"
   "log"

    // import third party libraries
    "github.com/PuerkitoBio/goquery"
    "strings"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/meta", metaScrape)
    http.ListenAndServe(":8080", nil)
}


func metaScrape(w http.ResponseWriter, r *http.Request) {
    //pageURL := "https://www.w3schools.com/html/tryit.asp?filename=tryhtml_headings"
    //pageURL := "http://jonathanmh.com"
    pageURL := "https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements" //links
    // pageURL := "http://metalsucks.net" //inaccible code 403
    //error handing
    res, err := http.Get(pageURL)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    if res.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }
    doc, err := goquery.NewDocument(pageURL)
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

    fmt.Println("================heading")
    //find heading : six level
    doc.Find("h1").Each(func(i int, h1 *goquery.Selection) {
         fmt.Println(i)
    })

    fmt.Println("================link+++++++++++++++++++++++")
    //internal and external links
    doc.Find("body a").Each(func(index int, item *goquery.Selection) {
        linkTag := item
        link, _ := linkTag.Attr("href")
        linkText := linkTag.Text()
	checkExtrernal := strings.HasPrefix(link, "http")
	external := "true"
	domain := ""
	if checkExtrernal {
		external = "true"

	} else {
		external = "false"
		domain = r.URL.Host
		fmt.Println(domain)
		return 

	}

	

	res, err := http.Get(link)
        if err != nil {
            log.Fatal(err)
        }
        defer res.Body.Close()

        fmt.Printf("Link #%d: '%s' - '%s'-  '%s', '%s', '%s' \n", index, linkText, link,external, strings.HasPrefix(link, "http"), res.StatusCode )
    })

    //login
    doc.Find("html body").Each(func(_ int, item *goquery.Selection) {
        // for debug.
        println(item.Size()) // return 1

       // if len(s.Nodes) > 0 && s.Nodes[0].Type == html.ElementNode {
         //   println(s.Nodes[0].Data)
        //}
	if( item.AttrOr("name","") == "password") {
		fmt.Println("here is login form")
	}
    })

}


//func findHeading(doc) {


//}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
