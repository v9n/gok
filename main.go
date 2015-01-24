package main

import (
  "fmt"
  "log"
  "os"
  "github.com/codegangsta/cli"
  //"github.com/kureikain/gok"
)

func main() {

  app := cli.NewApp()
  app.Name = "Gok"
  app.Usage = "Hacker News bookmark with full text search in blevesearch"
  app.Action = func(c *cli.Context) {
    fmt.Println("Welcome to " + name)
  }

  println(Version)

  app.Commands = []cli.Command{
    {
      Name:      "version",
      ShortName: "v",
      Usage:     "Version",
      Action: func(c *cli.Context) {
        println("Gok v ", Version)
      },
    },
    {
      Name:      "init",
      ShortName: "i",
      Usage:     "add a link to the list",
      Action: func(c *cli.Context) {
        println("added task: ", c.Args().First())
        _, err := NewStorage("gok.db")
        if err != nil {
          println("Succesfully creae db")
        }
      },
    },
    {
      Name:      "add",
      ShortName: "a",
      Usage:     "add a link to the list",
      Action: func(c *cli.Context) {
        s,_ := NewStorage("gok.db")
        item, err := NewItem(c.Args().First())

        if err != nil {
          log.Fatal(err)
        }
        if item == nil {
          log.Fatal("Invalid URL")
        }

        s.Add(item)
      },
    },
    {
      Name:      "ls",
      ShortName: "l",
      Usage:     "list the link",
      Action: func(c *cli.Context) {
        s,_ := NewStorage("gok.db")
        s.List()
      },
    },
    {
      Name:      "date",
      ShortName: "d",
      Usage: "yesterday, today, last week, last month",
      Action: func(c *cli.Context) {
        println("added task: ", c.Args().First())
      },
    },
  }

  app.Run(os.Args)
}

