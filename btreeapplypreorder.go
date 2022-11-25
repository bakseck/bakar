package bakar

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {

	f(root.Data)
	BTreeApplyInorder(root.Left, f)
	BTreeApplyInorder(root.Right, f)
}
