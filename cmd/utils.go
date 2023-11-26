package cmd

import (
	"fmt"
	"log"
	"strings"

	sitter "github.com/manyids2/go-tree-sitter-with-markdown"
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func CheckErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

// Container for node props in bidirectional tree
type NodeTree struct {
	Level                  int
	Text                   string
	StartByte, EndByte     int
	StartRow, EndRow       int
	StartColumn, EndColumn int
	Parent                 *NodeTree
	Children               []*NodeTree
}

func WalkWithIndent(n *sitter.Node, indent int, callback func(n *sitter.Node, indent int)) {
	for i := 0; i < int(n.NamedChildCount()); i++ {
		child := n.NamedChild(i)
		callback(child, indent)
		WalkWithIndent(child, indent+2, callback)
	}
}

func CallbackPrintNodeRange(n *sitter.Node, indent int) {
	s := strings.Repeat(" ", indent)
	fmt.Printf(
		"%s%s [%d,%d] - [%d,%d]\n",
		s, n.Type(),
		n.StartPoint().Row, n.StartPoint().Column,
		n.EndPoint().Row, n.EndPoint().Column,
	)
}

func removeDuplicate(sliceList []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func WalkNodeTree(n *NodeTree, indent int, callback func(n *NodeTree, indent int)) {
	for _, child := range n.Children {
		callback(child, indent)
		WalkNodeTree(child, indent+2, callback)
	}
}

func CallbackPrintRange(n *NodeTree, indent int) {
	s := strings.Repeat(" ", indent)
	fmt.Printf(
		"%s%d [%d,%d] - [%d,%d] (%d, %d)\n",
		s, n.Level,
		n.StartRow, n.StartColumn,
		n.EndRow, n.EndColumn,
		n.StartByte, n.EndByte,
	)
}

func CallbackPrintText(n *NodeTree, indent int) {
	s := strings.Repeat(" ", indent)
	fmt.Printf("%s%s ", s, n.Text)
	fmt.Printf(
		"|> %d [%d, %d] (%d, %d)\n",
		n.Level,
		n.StartRow, n.EndRow,
		n.StartByte, n.EndByte,
	)
}

func MaxOverArray(a []int) int {
	mx := MinInt
	for _, x := range a {
		if x > mx {
			mx = x
		}
	}
	return mx
}
