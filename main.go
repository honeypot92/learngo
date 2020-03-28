package main

import (
	"github.com/honeypot92/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary.Add("first", "first word")
	dictionary.Update("fist", "new definition")

}
