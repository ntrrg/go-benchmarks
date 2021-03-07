package pkg

type Tree struct{
	Left, Right *Tree
	Name, Type, Family string
	Height, Age int
}

func (t Tree) HasChildren() bool {
	return t.Left != nil || t.Right != nil
}

func (t *Tree) HasChildrenP() bool {
	return t.Left != nil || t.Right != nil
}

func (t *Tree) SetName(name string) {
	t.Name = name
}
