package main

import (
	"fmt"
	"net/http"
)

func main() {
	rs:=http.FileServer(http.Dir("AVL"))
	http.StripPrefix()
	fmt.Printf("%T",rs)

}
