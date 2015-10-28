// Package list implements a doubly linked list. This is a practice exercise.
// Find the official Go version at:
// https://golang.org/src/container/list/list.go
package "list"

// Doubly linked list
type List struct {
    // Root Node of the list; sentinel, so that even empty lists have Nodes
    root *Node
    // Number of elements in the list
    length int
}

// Node of a List
type Node struct {
    // List to which the Node belongs
    list *List
    // Pointers to previous and next Nodes
    next, prev *Node
    // Value of the node
    Value interface{}
}

func (n *Node) Next() *Node {
    // Check if prev is a valid, non-sentinel Node
    if x := n.next; x.list != nil && x != n.list.root {
        return x
    }
    return nil
}

func (n *Node) Prev() *Node {
    // Check if prev is a valid, non-sentinel Node
    if x := n.prev; x.list != nil && x != n.list.root {
        return x
    }
    return nil
}

// TODO Fundamentals
// func (l *List) Pop() {} (n *Node, err error) {}
// func (l *List) Push() (len int, err error) {}
// func (l *List) Remove(n *Node) {} (val interface{}, err error) {}
// func (l *List) InsertBefore(v interface{}, n *Node) err error {}
// func (l *List) InsertAfter(v interface{}, n *Node) err error {}

// TODO Utilities
// func SliceToList(s []interface) List {}
