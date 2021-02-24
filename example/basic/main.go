package main

import (
	"github.com/uberswe/html"
	"log"
)

func main() {
	log.Println(html.Button().Class("test").Link("https://beubo.com").Content("Test").Render())
}
