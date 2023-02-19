package p2p

// HandshakeFunc does nothing for now
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
