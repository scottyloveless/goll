/*Package goll*/
package goll

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	Value T
}

func (n *Node[T]) setNext(node *Node[T]) {
}
