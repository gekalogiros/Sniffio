package main

import (
	"os"
  "github.com/gekalogiros/sniffio/cmd"
)

func main() {

  root:=cmd.NewSniffioCommand()

  root.AddCommand(
    cmd.NewVersionCommand(os.Stdout),
    cmd.NewArpCommand(os.Stdout),
    cmd.NewIfacesCommand(os.Stdout),
    cmd.NewTCPCommand(os.Stdout),
  )

  root.Execute()
}