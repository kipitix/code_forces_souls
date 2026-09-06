package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	caseCount, _ := strconv.Atoi(scanner.Text())

	for range caseCount {
		scanner.Scan()
		nodeCount, _ := strconv.Atoi(scanner.Text())
		nodesWeights := make([]int, nodeCount)
		nodesParents := make([]int, nodeCount-1)
		for i := range nodeCount {
			scanner.Scan()
			nodesWeights[i], _ = strconv.Atoi(scanner.Text())
		}
		for i := range nodeCount - 1 {
			scanner.Scan()
			nodesParents[i], _ = strconv.Atoi(scanner.Text())
		}

		root := CreateTree(nodeCount, nodesWeights, nodesParents)
		result := Solve(root)
		for i, v := range result {
			if i == len(result)-1 {
				fmt.Println(v)
			} else {
				fmt.Printf("%d ", v)
			}
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

type Node struct {
	index    int
	weight   int
	children []*Node
	parent   *Node
	group    *Group
}

func (n *Node) PrintTree() {
	if n == nil {
		fmt.Println("Tree is nil")
		return
	}
	n.printTreeDetailed("", true)
}

func (n *Node) printTreeDetailed(prefix string, isLast bool) {
	// Определяем символы для визуализации
	var connector, childPrefix string
	if isLast {
		connector = "└── "
		childPrefix = prefix + "    "
	} else {
		connector = "├── "
		childPrefix = prefix + "│   "
	}

	// Выводим информацию об узле
	parentInfo := "nil"
	if n.parent != nil {
		parentInfo = fmt.Sprintf("%d", n.parent.index)
	}

	fmt.Printf("%s%s %d (weight=%d, parent=%s, children=%d)\n",
		prefix, connector, n.index, n.weight, parentInfo, len(n.children))

	// Рекурсивно выводим детей
	for i, child := range n.children {
		child.printTreeDetailed(childPrefix, i == len(n.children)-1)
	}
}

type Group struct {
	index int
	nodes []*Node
}

func CreateTree(nodeCount int, nodesWeights []int, nodesParents []int) *Node {
	if nodeCount < 1 {
		return nil
	}
	allNodes := make([]*Node, nodeCount)
	for i := 0; i < nodeCount; i++ {
		allNodes[i] = &Node{
			index:    i,
			weight:   nodesWeights[i],
			children: nil,
			parent:   nil,
		}
	}
	for i := 0; i < nodeCount-1; i++ {
		childNode := allNodes[i+1]
		parentNodeIndex := nodesParents[i] - 1
		parentNode := allNodes[parentNodeIndex]

		parentNode.children = append(parentNode.children, childNode)
		childNode.parent = parentNode
	}
	return allNodes[0]
}

func Solve(root *Node) []int {
	root.PrintTree()
	return nil
}
