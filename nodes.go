package rabbithole

import "encoding/json"

// TODO: this probably should be fixed in RabbitMQ management plugin
type OsPid string

type NameDescriptionEnabled struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type AuthMechanism NameDescriptionEnabled

type ExchangeType NameDescriptionEnabled

type NameDescriptionVersion struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type ErlangApp NameDescriptionVersion

type NodeInfo struct {
	Name      string `json:"name"`
	NodeType  string `json:"type"`
	IsRunning bool   `json:"running"`
	OsPid     OsPid  `json:"os_pid"`

	FdUsed        int  `json:"fd_used"`
	FdTotal       int  `json:"fd_total"`
	SocketsUsed   int  `json:"sockets_used"`
	SocketsTotal  int  `json:"sockets_total"`
	MemUsed       int  `json:"mem_used"`
	MemLimit      int  `json:"mem_limit"`
	MemAlarm      bool `json:"mem_alarm"`
	DiskFreeAlarm bool `json:"disk_free_alarm"`

	ExchangeTypes  []ExchangeType  `json:"exchange_types"`
	AuthMechanisms []AuthMechanism `json:"auth_mechanisms"`
	ErlangApps     []ErlangApp     `json:"applications"`
	Contexts       []BrokerContext `json:"contexts"`
}

//
// GET /api/nodes
//

func (c *Client) ListNodes() ([]NodeInfo, error) {
	var err error
	req, err := NewGETRequest(c, "nodes")
	if err != nil {
		return []NodeInfo{}, err
	}

	res, err := ExecuteHTTPRequest(c, req)
	if err != nil {
		return []NodeInfo{}, err
	}

	var rec []NodeInfo
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&rec)

	return rec, nil
}
