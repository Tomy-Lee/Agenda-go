package main

import  "github.com/Tomy-Lee/Agenda-golang/cmd"

func init() {
	log.SetFlags(log.Lshortfile)
	cmd.EnsureAgendaDir()
}

func main() {
	cmd.Execute()
}
