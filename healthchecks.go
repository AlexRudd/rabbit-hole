package rabbithole

type HealthCheckStatus struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func (hcs *HealthCheckStatus) Ok() bool {
	return hcs.Status == "ok"
}

//
// GET /api/healthchecks/node
//

//
// {"status":"ok"}
//
// {"status":"failed","reason":"string"}
//

// GetHealthCheckStatus Runs a basic healthchecks in the current node. Checks that the rabbit application
// is running, channels and queues can be listed successfully, and that no alarms are in effect.
func (c *Client) GetHealthCheckStatus() (rec *HealthCheckStatus, err error) {
	req, err := newGETRequest(c, "healthchecks/node")
	if err != nil {
		return nil, err
	}

	if err = executeAndParseRequest(c, req, &rec); err != nil {
		return nil, err
	}

	return rec, nil
}

//
// GET /api/healthchecks/node/{name}
//

//
// {"status":"ok"}
//
// {"status":"failed","reason":"string"}
//

// GetHealthCheckStatusFor Runs a basic healthchecks in the given node. Checks that the rabbit application
// is running, channels and queues can be listed successfully, and that no alarms are in effect.
func (c *Client) GetHealthCheckStatusFor(name string) (rec *HealthCheckStatus, err error) {
	req, err := newGETRequest(c, "healthchecks/node/"+PathEscape(name))
	if err != nil {
		return nil, err
	}

	if err = executeAndParseRequest(c, req, &rec); err != nil {
		return nil, err
	}

	return rec, nil
}
