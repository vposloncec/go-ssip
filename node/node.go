package node

type Node struct {
	Subscribers    []Node
	MessageHistory map[UUID]Message
}

func (n *Node) SendMessage(msg string) {

}
