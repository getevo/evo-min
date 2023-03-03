package main

import "github.com/getevo/evo-min"

func main() {
	evo.Setup()
	var db = evo.GetDBO()
	_ = db
	evo.Run()
}
