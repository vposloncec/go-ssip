package base

type Packet struct {
	ID   PacketUUID
	Data interface{}
}

func NewPacket(data interface{}) *Packet {
	return &Packet{
		ID:   NewUUID(),
		Data: data}
}
