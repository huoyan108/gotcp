package gotcp

type Packet interface {
	Serialize() []byte
}

type Protocol interface {
	ReadPacket(conn *Conn) (Packet, error)
}
