package base

type PacketWithSender struct {
	Packet *Packet
	Sender *Node
}
