package bakar 

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error))  {
	b := BTreeLevelCount(root)
	for i := 1; i <= b; i++ {
		BtreeCurrentLevel(root, i, f)
	}
}

func BtreeCurrentLevel()  {
	if root == nil {
		return
	}
	if b == 1 {
		f(root.Data)
	} else if b > 1 {
		BtreeCurrentLevel(root.Left, b-1, f)
		BtreeCurrentLevel(root.Right, b-1, f)
	}
}
