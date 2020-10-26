package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "io/ioutil"
  "log"
  "net"
  "net/http"
  "net/url"
  "strings"
  "sync"
  "time"

  "github.com/gin-gonic/gin"
)

var router *gin.Engine


type pageInfo struct {
  HTMLVersion string
  pageTitle string
  Heading1Count int
  Heading2Count int
  Heading3Count int
  Heading4Count int
  Heading5Count int
  Heading6Count int
  AmountInternalLinks int
  AmountExternalLinks int
  AmountInaccessibleLinks int
  LoginForm bool

}

var linkInfo pageInfo

func fetch(link , host string, wg *sync.WaitGroup) {
  time.Sleep(1 * time.Second)
  defer wg.Done()
  defer func() {
      if err := recover(); err != nil {
        log.Println("panic occurred:", err)
      }
     }()

  checkExternal := strings.HasPrefix(link, "http")

  if checkExternal {
    linkInfo.AmountExternalLinks = linkInfo.AmountExternalLinks + 1

  } else {
    linkInfo.AmountInternalLinks = linkInfo.AmountInternalLinks + 1
    //fmt.Println(host)
    //fmt.Println(host+link)
    link = "https://"+host+link
  }

  defer func() {
    res, err := http.Get(link)
    if err != nil {
      fmt.Println("err1111111111111")
      fmt.Println(err)
      //log.Fatal(err)
      if e, ok := err.(net.Error); ok && e.Timeout() {
        // This was a timeout
        fmt.Println(">>>>>>>>>>>>>>>>>>>>>>timeout error <<<<<<<<<<<<<<<<<<<<<<<")
        fmt.Println(linkInfo.AmountInternalLinks)
        fmt.Println(linkInfo.AmountExternalLinks)
        fmt.Println(linkInfo.AmountInaccessibleLinks)
        fmt.Println(linkInfo.Heading1Count)
        linkInfo.AmountInaccessibleLinks = linkInfo.AmountInaccessibleLinks + 1
        log.Println("there is error")
      } else if err != nil {
        fmt.Println(">>>>>>>>>>>>>>>>>>>>>>timeout not error <<<<<<<<<<<<<<<<<<<<<<<")
        // This was an error, but not a timeout
      }

      if res.StatusCode != 200 {
        fmt.Println("-----------------------------not status-----------------------------")
        linkInfo.Heading1Count = linkInfo.Heading1Count + 1
      }
    }
  }()



}


func search (c *gin.Context) {
   //link := c.Param("q")
   //link ,err :=  c.Params.Get("q")
  pageURL := c.Query("q")

  fmt.Println("==================================")
    fmt.Println(pageURL)
   //fmt.Println(err)


  res, err := http.Get(pageURL)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  body := res.Body


  //html, err := ioutil.ReadAll(res.Body)
  //if err != nil {
  //  fmt.Println("errrr")
  //  panic(err)
  //}
  ////fmt.Println("11111111=================")
  //if strings.Contains(string(html[:]),"<!DOCTYPE html>") {
  //  //fmt.Println("2222222=================")
  //  linkInfo.HTMLVersion = "5"
  //}
  //fmt.Println(html)
  //fmt.Println("*******************************")
  //fmt.Println(string(html[:]))
  //fmt.Println("333333333=================")



  // show the HTML code as a string %s
  //fmt.Printf("%s\n", html)


  u, err := url.Parse(pageURL)
  if err != nil {
   panic(err)
  }
  host := u.Host
  //doc, err := goquery.NewDocument(pageURL)
  fmt.Println(body)
  //doc, err := goquery.NewDocumentFromReader(body)

  doc, err := goquery.NewDocument(pageURL)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("======================>>>>>>")
  fmt.Println(body)
  html, err := ioutil.ReadAll(body)
  if err != nil {
    fmt.Println("errrr")
    panic(err)
  }
  if strings.Contains(string(html[:]),"<!DOCTYPE html>") {
    //  //fmt.Println("2222222=================")
     linkInfo.HTMLVersion = "5"
    }
  fmt.Println(html)
  fmt.Println("======================>>>>>>...")

  var pageTitle string

  // use CSS selector found with the browser inspector
  // for each, use index and item
  pageTitle = doc.Find("title").Contents().Text()
  fmt.Println(pageTitle)

  doc.Find("h1").Each(func(i int, s *goquery.Selection) {
    linkInfo.Heading1Count = linkInfo.Heading1Count + 1
  })

  doc.Find("h2").Each(func(i int, s *goquery.Selection) {
    linkInfo.Heading2Count = linkInfo.Heading2Count + 1
  })

  doc.Find("h3").Each(func(i int, s *goquery.Selection) {
    linkInfo.Heading3Count = linkInfo.Heading3Count + 1
  })

  doc.Find("h4").Each(func(i int, s *goquery.Selection) {
    linkInfo.Heading4Count = linkInfo.Heading4Count + 1
  })

  doc.Find("h5").Each(func(i int, s *goquery.Selection) {
    linkInfo.Heading5Count = linkInfo.Heading5Count + 1
  })

  doc.Find("h6").Each(func(i int, s *goquery.Selection) {
    linkInfo.Heading6Count = linkInfo.Heading6Count + 1
  })

  doc.Find("body input").Each(func(_ int, item *goquery.Selection) {
    itemId, _ := item.Attr("id")
    if itemId == "password" {
      linkInfo.LoginForm = true
    }
  })

  login, _ := doc.Find("form").Attr("id")
  checkLogin := strings.Contains(login, "login")
  if checkLogin {
    linkInfo.LoginForm = true

  }

  var wg sync.WaitGroup
  doc.Find("body a").Each(func(index int, item *goquery.Selection) {
   wg.Add(1)
   linkTag := item
   link, _ := linkTag.Attr("href")
   go fetch(link, host, &wg)
  })

  wg.Wait()



  // Call the HTML method of the Context to render a template
    c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses
      gin.H{
        "title":   "Home Page",
        "pageTitle" : pageTitle,
        "linkInfo":linkInfo,
      },
    )
}

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("templates/*")

  // Define the route for the index page and display the index.html template
  // To start with, we'll use an inline route handler. Later on, we'll create
  // standalone functions that will be used as route handlers.
  router.GET("/", func(c *gin.Context) {

    // Call the HTML method of the Context to render a template
    c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
        "title": "Home Page",
      },
    )

  })

  router.GET("/search", search )

  // Start serving the application
  router.Run()

}
