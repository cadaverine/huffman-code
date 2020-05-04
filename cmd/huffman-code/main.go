package main

import (
	"fmt"

	"github.com/cadaverine/huffman-code/huffman"
)

func main() {
	msg := "hello, world!"

	tree := huffman.GetTree(msg)
	table := huffman.GetTable(tree)

	encoded := huffman.Encode(msg, table)
	decoded := huffman.Decode(encoded, tree)

	fmt.Printf("Message: %s\n", msg)
	fmt.Printf("Table:   %v\n", table)
	fmt.Printf("Encoded: %v\n", encoded)
	fmt.Printf("Decoded: %s\n", decoded)
}
