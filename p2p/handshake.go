package p2p

// HandshakeFunc does nothing for now
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
