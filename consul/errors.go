package consul

import "errors"

var (
	ErrNoAvailableServiceInstance = errors.New("the requested service has no available instances")
)
