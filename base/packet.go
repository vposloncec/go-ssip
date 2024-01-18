package base

type Packet struct {
	Identifier UUID
	Data       interface{}
}

func NewPacket(data interface{}) *Packet {
	return &Packet{
		Identifier: NewUUID(),
		Data:       data}
}
