package base

type Packet struct {
	ID   UUID
	Data interface{}
}

func NewPacket(data interface{}) *Packet {
	return &Packet{
		ID:   NewUUID(),
		Data: data}
}
