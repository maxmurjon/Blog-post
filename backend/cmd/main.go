package main

import (
	"fmt"

	"github.com/maxmurjon/blog_post/config"
)

func main() {
	cfg:=config.Load()
	fmt.Println(cfg)
}