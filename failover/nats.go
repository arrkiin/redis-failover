package failover

import (
	nats "github.com/nats-io/nats.go"
)

// Nats connection structure
type Nats struct {
	conn    *nats.Conn
	subject string
}

func newNats(server, subject string) (*Nats, error) {
	conn, err := nats.Connect(server)
	if err != nil {
		return nil, err
	}
	return &Nats{
		conn:    conn,
		subject: subject,
	}, nil
}

// SignalFailOver a failover message
func (n *Nats) SignalFailOver(downMaster, newMaster string) error {
	if err := n.conn.Publish(n.subject, []byte("")); err != nil {
		return err
	}
	n.conn.Flush()
	return nil
}
