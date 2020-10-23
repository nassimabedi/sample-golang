package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

var router *gin.Engine






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
  doc, err := goquery.NewDocument(pageURL)
  if err != nil {
    log.Fatal(err)
  }

  var pageTitle string

  // use CSS selector found with the browser inspector
  // for each, use index and item
  pageTitle = doc.Find("title").Contents().Text()
  fmt.Println(pageTitle)


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
        //"payload": articles,
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
