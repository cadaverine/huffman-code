package huffman

type Node struct {
	Left   *Node
	Right  *Node
	Weight int
	Rune   rune
}

// GetTable - получить таблицу Хаффмана
func GetTable(tree *Node) map[rune]string {
	return createHuffmanTable(tree)
}

// GetTree - получить дерево Хаффмана
func GetTree(msg string) *Node {
	frequencies := getRunesFrequencies(msg)
	queue := getQueue(frequencies)

	return createHuffmanTree(queue)
}

// Encode - закодировать сообщение
func Encode(message string, table map[rune]string) []string {
	result := []string{}

	for _, char := range message {
		result = append(result, table[char])
	}

	return result
}

// Decode - раскодировать сообщение
func Decode(encoded []string, tree *Node) string {
	result := []rune{}

	for _, code := range encoded {
		result = append(result, getFromTree(tree, string(code)))
	}

	return string(result)
}
