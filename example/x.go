package main

import (
	"fmt"

	"github.com/d1y/dutils/dstring"
	"github.com/d1y/dutils/fs"
)

func main() {
	dstringDemo()
	fsDemo()
}

func dstringDemo() {
	var y = "how are you"
	var xs = dstring.Padstart(y, "?", 2)
	var xe = dstring.Padend(y, "#", 2)
	var xa = dstring.Padding(y, "#", 8)
	var fill = dstring.Fill("#", 24)
	var total = dstring.PaddingTotalWidth(y, "#", 24)
	var p = dstring.PaddingLength(y, "$", 5)
	fmt.Println(p)
	fmt.Println(total)
	fmt.Println(xs)
	fmt.Println(xe)
	fmt.Println(xa)
	fmt.Println(fill)
}

func fsDemo() {
	var flag = fs.Exists("test")
	fmt.Println("flag: ", flag)
	fs.CreateIfNotExists("dev", 0755)
	fs.CreateIfNotExists("dev2", 0755)
	fs.Copy("hello.txt", "dev/hello.txt")
	fs.CopyDirectory("dev", "dev2")
	var isDir = fs.IsDir("dev")
	fmt.Println("dev is dir: ", isDir)
	str, _ := fs.ReadFile2String("hello.txt")
	fmt.Println("read string: ", str)
}
