package main

import "github.com/zenazn/goji"

func main() {
	// flag.Set("bind", ":other_port_here") - the default is 8000
	goji.Serve()
}
