package main

import (
	"github.com/gogf/gf/frame/g"
	_ "github.com/stardemo/go-similar-keywords/src/router"
)

func main() {
	s := g.Server()
	s.Run()
}
