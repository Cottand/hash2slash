#! /usr/bin/env nix
#! nix shell github:cottand/hash2slash#go-run --quiet --command hash2slash-go-run

package main

import "os"

func main() {
	println("Hello world, I am a script!")
	println(os.Args[1])
}