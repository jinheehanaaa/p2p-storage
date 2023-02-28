package p2p

import (
	"fmt"
	"net"
)

// TCPPeer is the remote node over a TCP established connection
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn
	// outbound = dial & retrieve a conn to peer
	// inbound = accept & retrieve the peer
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Close implements the Peer interface.
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

// TCPTransportOpts holds configuration option
type TCPTransportOpts struct {
	ListenAddr    string // For address: host:port
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener // Dynamic object for listening
	rpcch    chan RPC
}

// Consume imlements the Transport interface, which will return read-only channel
// for reading the incoming messages received from another peer in the network.
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

// startAcceptLoop accepts new connection
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {

			fmt.Printf("TCP accept error: %s\n", err)
		}

		fmt.Printf("new incoming connection: %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct{}

// hnadleConn is responsible for handling communication (HandshakeFunc, OnPeer, & incoming messages)
func (t *TCPTransport) handleConn(conn net.Conn) {

	var err error

	defer func() {
		fmt.Printf("dropping peer connection: %s", err)
		conn.Close()
	}()

	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	// Read loop
	rpc := RPC{}

	for {
		err := t.Decoder.Decode(conn, &rpc)
		if err != nil {
			return
		}

		rpc.From = conn.RemoteAddr()

		fmt.Printf("message: %+v\n", rpc)
		t.rpcch <- rpc
	}
}
