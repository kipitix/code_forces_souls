package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
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
		result := Solve(root, nodeCount)
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
func Solve(root *Node, nodeCount int) []int {
	result := make([]int, nodeCount)

	// root.PrintTree()
	groups := InitGroups(root)

	// for i, g := range groups {
	// 	g.PrintGroup(i + 1)
	// }

	startGroupsCount := len(groups)

	for i := range startGroupsCount - 1 {
		result[i] = -1
	}
	result[startGroupsCount-1] = SumGroups(groups)

	for i := startGroupsCount; i < nodeCount; i++ {
		groups2, diff := SplitGroups(groups)
		groups = groups2

		result[i] = result[i-1] + diff

		// for i, g := range groups {
		// 	g.PrintGroup(i + 1)
		// }

	}

	return result
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
	groups, _ := findSubgroups(root)
	// Нужно отсортировать элементы в группах.
	// Сортировать буду по возрастанию, чтобы убирать последний.
	for _, g := range groups {
		slices.SortFunc(g.nodes, func(a, b *Node) int {
			return cmp.Compare(a.weight, b.weight)
		})
	}
	return groups
}

// Поиск подгрупп, включая корень.
// После поиска, корень окажется в одной из подгрупп.
// Корень попадает в подгруппу для которой он может иметь значение.
// Если корень перевешивает вес всей подгруппы, то он добавляется в неё.
func findSubgroups(node *Node) ([]*Group, *Group) {
	// Если дочерних элементов нет, то создаём новую группу
	// с одним собой.
	if len(node.children) == 0 {
		newGroup := &Group{
			nodes:     []*Node{node},
			maxWeight: node.weight,
		}
		// node.group = newGroup
		return []*Group{newGroup}, newGroup
	}
	// Собираем информацию о подгруппах у дочерних элементов.
	groups := make([]*Group, 0)
	chainedGroups := make([]*Group, 0)
	for _, childNode := range node.children {
		allChildGroups, childChainedGroup := findSubgroups(childNode)
		groups = append(groups, allChildGroups...)
		chainedGroups = append(chainedGroups, childChainedGroup)
	}
	// Находим в какую группу нужно включить себя.
	slices.SortFunc(chainedGroups, func(a, b *Group) int {
		return cmp.Compare(a.maxWeight, b.maxWeight)
	})
	chainedGroup := chainedGroups[0]
	chainedGroup.nodes = append(chainedGroup.nodes, node)
	// node.group = chainedGroup
	if node.weight > chainedGroup.maxWeight {
		chainedGroup.maxWeight = node.weight
	}

	return groups, chainedGroup
}

// Суммируем значения максимальных весов всех групп.
func SumGroups(groups []*Group) int {
	sum := 0
	for _, g := range groups {
		sum += g.maxWeight
	}
	return sum
}

// Поиск группы, которую наиболее эффективно было бы разбить.
// Число групп должно увеличиться на 1.
// Вход: Текущие группы.
// Выход: Новые группы, на сколько изменилась сумма весов групп.
func SplitGroups(groupsInput []*Group) ([]*Group, int) {
	activeGroups := []*Group{}
	// Удаление из оценки групп с одним элементом.
	for _, g := range groupsInput {
		if len(g.nodes) > 1 {
			activeGroups = append(activeGroups, g)
		}
	}
	// Нужно отсортировать группы.
	// Сортируем группы так, чтобы первой оказалась
	// группа с максимальным узлом среди всех вторых элементов
	// всех групп.
	slices.SortFunc(activeGroups, func(a, b *Group) int {
		return cmp.Compare(a.nodes[len(a.nodes)-1].weight, b.nodes[len(b.nodes)-1].weight)
	})
	// После такой сортировки берём последнюю группу.
	// Убираем в ней последний элемент и считаем что вес всех
	// групп увеличился на значение веса предпоследнего элемента.
	lastGroupIndex := len(activeGroups) - 1
	activeGroups[lastGroupIndex].nodes = activeGroups[lastGroupIndex].nodes[:len(activeGroups[lastGroupIndex].nodes)-1]
	sumWeightDiff := activeGroups[lastGroupIndex].nodes[len(activeGroups[lastGroupIndex].nodes)-1].weight
	return activeGroups, sumWeightDiff
}

// Узел дерева.
type Node struct {
	index    int
	weight   int
	children []*Node
	parent   *Node
	// group    *Group
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
	// Определяем символы для визуализации.
	var connector, childPrefix string
	if isLast {
		connector = "└── "
		childPrefix = prefix + "    "
	} else {
		connector = "├── "
		childPrefix = prefix + "│   "
	}
	// Выводим информацию об узле.
	parentInfo := "nil"
	if n.parent != nil {
		parentInfo = fmt.Sprintf("%d", n.parent.index+1)
	}
	fmt.Printf("%s%s Node %d (weight=%d, parent=%s, children=%d)\n",
		prefix, connector, n.index+1, n.weight, parentInfo, len(n.children))
	// Рекурсивно выводим детей.
	for i, child := range n.children {
		child.printTreeDetailed(childPrefix, i == len(n.children)-1)
	}
}

// Группа.
type Group struct {
	nodes     []*Node
	maxWeight int
}

// Функция печати информации по каждой отдельной группе.
func (g *Group) PrintGroup(groupIndex int) {
	if g == nil {
		fmt.Println("Group is nil")
		return
	}
	fmt.Printf("Group %d (nodes count: %d, max weight: %d)\n", groupIndex, len(g.nodes), g.maxWeight)
	if len(g.nodes) == 0 {
		fmt.Println("  (empty group)")
		return
	}
	// Выводим все узлы группы.
	for i, node := range g.nodes {
		var prefix string
		if i == len(g.nodes)-1 {
			prefix = "└── "
		} else {
			prefix = "├── "
		}
		// Информация о родителе узла.
		parentInfo := "nil"
		if node.parent != nil {
			parentInfo = fmt.Sprintf("%d", node.parent.index+1)
		}
		// Информация о детях узла.
		childrenInfo := "none"
		if len(node.children) > 0 {
			childrenInfo = fmt.Sprintf("%v", getChildrenIndexes(node.children))
		}
		fmt.Printf("  %sNode %d (weight=%d, parent=%s, children=[%s])\n",
			prefix, node.index+1, node.weight, parentInfo, childrenInfo)
	}
}

// Вспомогательная функция для получения индексов детей узлов.
func getChildrenIndexes(children []*Node) string {
	if len(children) == 0 {
		return ""
	}
	indexes := make([]string, len(children))
	for i, child := range children {
		indexes[i] = fmt.Sprintf("%d", child.index+1)
	}
	return strings.Join(indexes, ", ")
}
