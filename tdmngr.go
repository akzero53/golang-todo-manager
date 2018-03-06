package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "Todo Manager"
  app.Usage = "タスク管理ツール"
  app.Action = func(c *cli.Context) error {
    fmt.Println("hogehoge")
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
