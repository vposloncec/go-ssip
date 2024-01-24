package base

import "time"

type Packet struct {
	ID        PacketUUID
	Timestamp time.Time
	Data      interface{}
}

func NewPacket(data interface{}) *Packet {
	return &Packet{
		ID:   NewUUID(),
		Data: data}
}
