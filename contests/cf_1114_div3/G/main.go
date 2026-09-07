package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Поиск решения кейса.
func Solve(root *Node) []int {
	root.PrintTree()
	groups := InitGroups(root)

	for _, g := range groups {
		g.PrintGroup()
	}

	return nil
}

// Создание дерева на основании входных данных.
// На выходе получаем указатель на вершину.
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

// Инициализация групп.
// Разбиение группы по правилам задачи.
// Создание групп по количеству листьев дерева.
func InitGroups(root *Node) []*Group {
	groups := findSubgroups(root)
	return groups
}

// Поиск подгрупп, включая корень.
// После поиска, корень окажется в одной из подгрупп.
// Корень попадает в подгруппу для которой он может иметь значение.
// Если корень перевешивает вес всей подгруппы, то он добавляется в неё.
func findSubgroups(node *Node) []*Group {
	subGroups := make([]*Group, len(node.children))

	return subGroups
}

// Узел дерева.
type Node struct {
	index    int
	weight   int
	children []*Node
	parent   *Node
	group    *Group
}

// Вывод отладочной информации.
func (n *Node) PrintTree() {
	if n == nil {
		fmt.Println("Tree is nil")
		return
	}
	n.printTreeDetailed("", true)
}

// Дополнительная функция для вывода древовидного представления.
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

	fmt.Printf("%s%s Node %d (weight=%d, parent=%s, children=%d)\n",
		prefix, connector, n.index, n.weight, parentInfo, len(n.children))

	// Рекурсивно выводим детей
	for i, child := range n.children {
		child.printTreeDetailed(childPrefix, i == len(n.children)-1)
	}
}

// Группа.
type Group struct {
	index     int
	nodes     []*Node
	maxWeight int
}

// Функция печати информации по каждой отдельной группе.
func (g *Group) PrintGroup() {
	if g == nil {
		fmt.Println("Group is nil")
		return
	}

	fmt.Printf("Group %d (nodes count: %d)\n", g.index, len(g.nodes))

	if len(g.nodes) == 0 {
		fmt.Println("  (empty group)")
		return
	}

	// Выводим все узлы группы
	for i, node := range g.nodes {
		var prefix string
		if i == len(g.nodes)-1 {
			prefix = "└── "
		} else {
			prefix = "├── "
		}

		// Информация о родителе узла
		parentInfo := "nil"
		if node.parent != nil {
			parentInfo = fmt.Sprintf("%d", node.parent.index)
		}

		// Информация о детях узла
		childrenInfo := "none"
		if len(node.children) > 0 {
			childrenInfo = fmt.Sprintf("%v", getChildrenIndexes(node.children))
		}

		fmt.Printf("  %sNode %d (weight=%d, parent=%s, children=[%s])\n",
			prefix, node.index, node.weight, parentInfo, childrenInfo)
	}
}

// Вспомогательная функция для получения индексов детей узлов.
func getChildrenIndexes(children []*Node) string {
	if len(children) == 0 {
		return ""
	}
	indexes := make([]string, len(children))
	for i, child := range children {
		indexes[i] = fmt.Sprintf("%d", child.index)
	}
	return strings.Join(indexes, ", ")
}
