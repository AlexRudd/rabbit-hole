package rabbithole

import "net/url"

// Provides information about connection to a RabbitMQ node.
type ConnectionInfo struct {
	// Connection name
	Name     string `json:"name"`
	// Node the client is connected to
	Node     string `json:"node"`
	// Number of open channels
	Channels int    `json:"channels"`
	// Connection state
	State    string `json:"state"`
	// Connection type, network (via AMQP client) or direct (via direct Erlang client)
	Type     string `json:"type"`

	// Server port
	Port     Port `json:"port"`
	// Client port
	PeerPort Port `json:"peer_port"`

	// Server host
	Host     string `json:"host"`
	// Client host
	PeerHost string `json:"peer_host"`

	// Last connection blocking reason, if any
	LastBlockedBy  string `json:"last_blocked_by"`
	// When connection was last blocked
	LastBlockedAge string `json:"last_blocked_age"`

	// True if connection uses TLS/SSL
	UsesTLS          bool   `json:"ssl"`
	PeerCertSubject  string `json:"peer_cert_subject"`
	PeerCertValidity string `json:"peer_cert_validity"`
	PeerCertIssuer   string `json:"peer_cert_issuer"`

	SSLProtocol    string `json:"ssl_protocol"`
	SSLKeyExchange string `json:"ssl_key_exchange"`
	SSLCipher      string `json:"ssl_cipher"`
	SSLHash        string `json:"ssl_hash"`

	Protocol string `json:"protocol"`
	User     string `json:"user"`
	Vhost    string `json:"vhost"`

	Timeout  int `json:"timeout"`
	FrameMax int `json:"frame_max"`

	ClientProperties Properties `json:"client_properties"`

	RecvOct        uint64      `json:"recv_oct"`
	SendOct        uint64      `json:"send_oct"`
	RecvCount      uint64      `json:"recv_cnt"`
	SendCount      uint64      `json:"send_cnt"`
	SendPendi      uint64      `json:"send_pend"`
	RecvOctDetails RateDetails `json:"recv_oct_details"`
	SendOctDetails RateDetails `json:"send_oct_details"`
}

//
// GET /api/connections
//

func (c *Client) ListConnections() (rec []ConnectionInfo, err error) {
	req, err := newGETRequest(c, "connections")
	if err != nil {
		return []ConnectionInfo{}, err
	}

	if err = executeAndParseRequest(req, &rec); err != nil {
		return []ConnectionInfo{}, err
	}

	return rec, nil
}

//
// GET /api/connections/{name}
//

func (c *Client) GetConnection(name string) (rec *ConnectionInfo, err error) {
	req, err := newGETRequest(c, "connections/"+url.QueryEscape(name))
	if err != nil {
		return nil, err
	}

	if err = executeAndParseRequest(req, &rec); err != nil {
		return nil, err
	}

	return rec, nil
}
