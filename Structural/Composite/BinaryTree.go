package Composite

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}
