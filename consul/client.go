package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type ConsulClient interface {
	GetAvailableService(service, tag string) (*api.ServiceEntry, error)
	GetAvailableUrl(service, tag string) (*string, error)
	GetAvailableUrlOrDefault(service, tag, defaultUrl string) string
}

type consulClientImpl struct {
	client *api.Client
}

type ServiceUrl struct {
	Host string
	Port int
}

func NewConsulClient(host string, port int) (ConsulClient, error) {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", host, port)
	client, err := api.NewClient(config)

	if err != nil {
		return nil, err
	}

	c := consulClientImpl{client: client}
	return c, nil
}

func (c consulClientImpl) GetAvailableService(service, tag string) (*api.ServiceEntry, error) {
	addrs, _, err := c.client.Health().Service(service, "", true, nil)
	if err != nil {
		return nil, err
	}
	if len(addrs) == 0 {
		return nil, ErrNoAvailableServiceInstance
	}

	s := addrs[0]
	return s, nil
}

func (c consulClientImpl) GetAvailableUrl(service, tag string) (*string, error) {
	addrs, _, err := c.client.Health().Service(service, "", true, nil)
	if err != nil {
		return nil, err
	}
	if len(addrs) == 0 {
		return nil, ErrNoAvailableServiceInstance
	}

	s := addrs[0]
	url := fmt.Sprintf("http://%s:%d", s.Service.Address, s.Service.Port)
	return &url, nil
}

func (c consulClientImpl) GetAvailableUrlOrDefault(service, tag, defaultUrl string) string {
	url, err := c.GetAvailableUrl(service, tag)
	if err != nil {
		return defaultUrl
	}
	if url == nil {
		return defaultUrl
	}

	return *url
}
