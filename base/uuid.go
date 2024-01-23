package base

import "github.com/google/uuid"

type PacketUUID string

func NewUUID() PacketUUID {
	return PacketUUID(uuid.New().String())
}
