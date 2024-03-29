package tcp

import (
	"bufio"
	"io"
	"io/ioutil"
	"kiv_ups_server/net/tcp/protocol"
	"math/rand"
	"syscall"
)

// Unique identifier of client. This can be used as map key.
type UID int

// Client structure stores basic data about TCP client
type Client struct {
	Server            *Server
	TCP               *TCP
	UID               UID
	clientMessageChan chan ClientMessage
	MessageChan       chan *protocol.ProtoMessage
	Protocol          protocol.GameProtocol
	failCounter       int
	decodeReader      io.ReadCloser
	decodeWriter      io.WriteCloser
}

// newClient initializes new client
func newClient(s *Server, clientMessageChan chan ClientMessage, fd int, sockaddr syscall.Sockaddr, protocol protocol.GameProtocol,
	messageChan chan *protocol.ProtoMessage) Client {
	return Client{
		Server:            s,
		clientMessageChan: clientMessageChan,
		TCP: &TCP{
			FD:       fd,
			Sockaddr: sockaddr,
		},
		UID:         UID(rand.Int()),
		MessageChan: messageChan,
		Protocol:    protocol,
	}
}

// Send encodes given message and sends it to client
func (c Client) Send(message protocol.ProtoMessage) (err error) {
	reader, writer := io.Pipe()

	go c.Protocol.Encode(message, writer)
	r := bufio.NewReader(reader)

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	_, err = syscall.Write(c.TCP.FD, bytes)

	return
}

// SendBytes allows to send binary data to client.
// You shouldn't need to do it, this function is primarily for TCP server.
func (c *Client) SendBytes(message []byte) (err error) {
	_, err = syscall.Write(c.TCP.FD, message)

	return
}

type ClientMessage struct {
	protocol.Message
	protocol.RequestId
	Sender *Client

	// DisconnectRequest is used to notify master server about
	// client disconnection
	DisconnectRequest bool
}
