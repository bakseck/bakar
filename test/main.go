package main

import (
	"fmt"
	"bakar"
)

func main() {
	root := &bakar.TreeNode{Data: "4"}
	bakar.BTreeInsertData(root, "1")
	bakar.BTreeInsertData(root, "7")
	bakar.BTreeInsertData(root, "5")
	fmt.Println(root.Left.Data)
	fmt.Println(root.Data)
	fmt.Println(root.Right.Left.Data)
	fmt.Println(root.Right.Data)
}
