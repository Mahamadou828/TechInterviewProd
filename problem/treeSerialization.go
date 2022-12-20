package problem

import (
	"fmt"
	"strconv"
	"strings"

	dts "github.com/Mahamadou828/coderpro/datastruct"
)

// TreeSerialization serializes a given bst
// Time Complexity O(n)
// Space complexity O(n)
func TreeSerialization(tree *dts.Tree) string {
	var str strings.Builder
	treeSerializationHelper(tree.Root, &str)
	return str.String()
}

// TreeDeserialization deserializes a given string to a bst
// Time Complexity O(n)
// Space complexity O(n)
func TreeDeserialization(str string) *dts.Tree {
	root, _ := treeDeserializationHelper(strings.Split(str, " "))
	return &dts.Tree{Root: root}
}

func treeSerializationHelper(leaf *dts.Leaf, b *strings.Builder) *strings.Builder {
	b.WriteString(fmt.Sprintf("%v ", leaf.Value))
	switch leaf.Left {
	case nil:
		b.WriteString("# ")
	default:
		treeSerializationHelper(leaf.Left, b)
	}
	switch leaf.Right {
	case nil:
		b.WriteString("# ")
	default:
		treeSerializationHelper(leaf.Right, b)
	}
	return b
}

func treeDeserializationHelper(strs []string) (*dts.Leaf, []string) {
	valueStr, strs := strs[0], strs[1:]
	if valueStr == "#" {
		return nil, strs
	}
	value, _ := strconv.Atoi(valueStr)
	leaf := &dts.Leaf{Value: float64(value)}
	leaf.Left, strs = treeDeserializationHelper(strs)
	leaf.Right, strs = treeDeserializationHelper(strs)

	return leaf, strs
}
