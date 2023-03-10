package p2p

// Peer represents the remote node
type Peer interface {
	Close() error
}

// Transport handles the communication b/w the nodes in the network (Ex: TCP, UDP, Websockets, etc)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
