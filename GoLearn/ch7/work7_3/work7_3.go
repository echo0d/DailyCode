package main

import "fmt"

// 在gopl.io/ch4/treesort（§4.4）中的*tree类型实现一个String方法去展示tree类型的值序列。

type tree struct {
	value       int
	left, right *tree
}

func (tree *tree) String() string {
	var values []int
	values = appendValues(values, tree)
	return fmt.Sprint(values)
}

// Sort sorts values in place.
func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	var m = []int{1, 4, 7, 9, 3, 5, 7, 2, 4}
	t := Sort(m)
	fmt.Println(m)
	fmt.Println(t.String())
	fmt.Println(t)

}
