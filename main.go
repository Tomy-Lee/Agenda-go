package main

import (
	"cmd"
	"log"
	"model"
)

func init() {
	log.SetFlags(log.Lshortfile)
	model.EnsureAgendaDir()
}

func main() {
	cmd.Execute()
}