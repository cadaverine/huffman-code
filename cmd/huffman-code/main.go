package main

import (
	"fmt"

	"github.com/cadaverine/huffman-code/heap"
)

type Node struct {
	Left   *Node
	Right  *Node
	Weight int
	Rune   rune
}

type HuffmanTable map[rune]int

// считаем частоты
func getRunesFrequencies(message string) map[rune]int {
	frequencies := make(map[rune]int)

	for _, char := range []rune(message) {
		frequencies[char]++
	}

	return frequencies
}

func getQueue(frequencies map[rune]int) *heap.Heap {
	queue := heap.Init(heap.Minimum)

	for key, value := range frequencies {
		node := Node{Rune: key, Weight: value}
		queue.Enqueue(node, value)
	}

	return queue
}

func createHuffmanTree(queue *heap.Heap) *Node {
	for queue.Size() != 1 {
		l, lPr := queue.Dequeue()
		r, rPr := queue.Dequeue()

		left := l.(Node)
		right := r.(Node)
		weight := lPr + rPr

		node := Node{Left: &left, Right: &right, Weight: weight}

		queue.Enqueue(node, weight)
	}

	root, _ := queue.Dequeue()
	rootNode := root.(Node)

	return &rootNode
}

func traverse(tree *Node, code string, table map[rune]string) map[rune]string {
	if tree.Left != nil {
		traverse(tree.Left, code+"0", table)
	}
	if tree.Right != nil {
		traverse(tree.Right, code+"1", table)
	}

	if tree.Rune != 0 {
		table[tree.Rune] = code
	}

	return table
}

func createHuffmanTable(tree *Node) map[rune]string {
	table := traverse(tree, "", make(map[rune]string))

	return table
}

func encode(message string, table map[rune]string) []string {
	result := []string{}

	for _, char := range message {
		result = append(result, table[char])
	}

	return result
}

func decode(encoded []string, tree *Node) string {
	return ""
}

func main() {
	msg := "hello, world!"

	frequencies := getRunesFrequencies(msg)
	queue := getQueue(frequencies)

	tree := createHuffmanTree(queue)
	table := createHuffmanTable(tree)

	encoded := encode(msg, table)
	decoded := decode(encoded, tree)

	fmt.Printf("Message: %s\n", msg)
	fmt.Printf("Table: %+v\n", table)
	fmt.Printf("Encoded: %v\n", encoded)
	fmt.Printf("Decoded: %s", decoded)
}
