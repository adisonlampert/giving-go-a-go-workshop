# Giving Go a Go: Build a REST API with Golang!

Welcome to Hack the North 2023! We're Angela and Adison, your workshop leads. Get a head start on your hackathon project by building a REST API with Go! In this workshop, you will gain experience creating web applications that can communicate with other services and exchange data—a fundamental skill—all while learning Golang, a very powerful and in-demand programming language.

![image](https://hackthenorth.com/preview_img.png)

- REST is a widely used architectural style for creating web applications that can communicate with other services and exchange data.
- Go has gained significant popularity in recent years due to its simplicity, performance, and built-in support for concurrency.
- By participating in this workshop, hackers can learn the fundamentals of Go by using it to build a REST API and understand its unique features, enabling them to work with the language effectively in their hackathon projects or other tech endeavors.

# Table of Contents
- [Before You Start](#before-you-start)
- [Code from Slides](#code-from-slides)
  - [Read Part 1](#read-part-1)
  - [Read Part 2](#read-part-2)
  - [Create](#create)
  - [Update](#update)
  - [Delete](#delete)
- [Closing Words](#closing-words)

# Before You Start
Please make sure that you've checked out the [Hack Pack](https://docs.google.com/document/d/19IOBkdTcl-_0GgUKi_pFwVHpY6XimzM6wjDOcOw3hbQ/edit#heading=h.5c4h3x37oxvr) for our workshop and follow through the doc to get yourself set up for success.

For easy access, here are the [slides](https://docs.google.com/presentation/d/1sZ5KwUzFqdukH_RCzL2q1imVodvokXQYFII_cc73t-8/edit#slide=id.g2764ce0426f_3_218)!

# Code from Slides
## Read Part 1

``` golang
func getHackathons(c *gin.Context) {
    var hackathons []Hackathon
    db.Find(&hackathons)
    c.IndentedJSON(http.StatusOK, gin.H{"data": hackathons})
}
```
[(Back to top)](#table-of-contents)

## Read Part 2

``` golang
func getHackathonById(c *gin.Context) {
    var hackathon Hackathon

    err := db.Where("id = ?", c.Param("id")).First(&hackathon).Error
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Hackathon not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": hackathon})
}
```
[(Back to top)](#table-of-contents)

## Create
Create this `struct` underneath `Hackathon`:
``` golang
type HackathonInput struct {
    Name     string `json:"name" validate:"required"`
    Date     string `json:"date" validate:"required"`
    Url      string `json:"url" validate:"required"`
    Location string `json:"location" validate:"required"`
}

```

In `createHackathon`, write the code below
``` golang
func createHackathon(c *gin.Context) {
    var input HackathonInput

    err := c.ShouldBindJSON(&input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = validate.Struct(input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var hackathons []Hackathon
    hackathon := Hackathon{
        Id:       int(db.Find(&hackathons).RowsAffected) + 1,
        Name:     input.Name,
        Date:     input.Date,
        Url:      input.Url,
        Location: input.Location,
    }
    db.Create(&hackathon)
    c.JSON(http.StatusOK, gin.H{"data": hackathon})
}

```
[(Back to top)](#table-of-contents)

## Update

``` golang
func updateHackathon(c *gin.Context) {
    var hackathon Hackathon
    err := db.Where("id = ?", c.Param("id")).First(&hackathon).Error
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Hackathon not found!"})
        return
    }
    var input HackathonInput

    err = c.ShouldBindJSON(&input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = validate.Struct(input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Model(&hackathon).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": hackathon})
}

```
[(Back to top)](#table-of-contents)

## Delete

``` golang
func deleteHackathon(c *gin.Context) {
    var hackathon Hackathon

    err := db.Where("id = ?", c.Param("id")).First(&hackathon).Error
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Hackathon not found!"})
        return
    }

    db.Delete(&hackathon)
    c.JSON(http.StatusOK, gin.H{"data": true})
}

```
[(Back to top)](#table-of-contents)

# Closing Words
Thank you so much for checking us out at Hack The North 2023, or on the interwebs. We hope you had as much fun as we did making this workshop. We'd love to see what you end up coming up with on your own!

With love,
Angela and Adison

[(Back to top)](#table-of-contents)
