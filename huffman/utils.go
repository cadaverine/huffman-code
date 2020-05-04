package huffman

import "container/heap"

// считаем частоты
func getRunesFrequencies(message string) map[rune]int {
	frequencies := make(map[rune]int)

	for _, char := range []rune(message) {
		frequencies[char]++
	}

	return frequencies
}

func getQueue(frequencies map[rune]int) PriorityQueue {
	pq := make(PriorityQueue, len(frequencies))

	i := 0
	for value, priority := range frequencies {
		node := Node{Rune: value, Weight: priority}

		pq[i] = &Item{
			value:    node,
			priority: priority,
			index:    i,
		}
		i++
	}

	heap.Init(&pq)

	return pq
}

func createHuffmanTree(queue PriorityQueue) *Node {
	for queue.Len() != 1 {
		l := heap.Pop(&queue).(*Item)
		r := heap.Pop(&queue).(*Item)

		leftNode := l.value.(Node)
		rightNode := r.value.(Node)
		weight := l.priority + r.priority

		node := Node{Left: &leftNode, Right: &rightNode, Weight: weight}

		heap.Push(&queue, &Item{value: node, priority: weight})
	}

	root := heap.Pop(&queue).(*Item)
	rootNode := root.value.(Node)

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

func getFromTree(tree *Node, code string) rune {
	node := tree

	for _, char := range code {
		if string(char) == "0" {
			node = node.Left
		} else if string(char) == "1" {
			node = node.Right
		}
	}

	return node.Rune
}
