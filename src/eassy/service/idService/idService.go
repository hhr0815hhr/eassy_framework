package idService

import (
	"sync"
)

var node *Node
var once sync.Once
var NodeId int64 = 1

func GenerateID() ID {
	once.Do(func() {
		node, _ = NewNode(NodeId)
	})
	return node.Generate()
}
