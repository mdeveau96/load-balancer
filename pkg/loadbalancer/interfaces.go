package loadbalancer

import (
	"context"
)

type LoadBalancer interface {
	SelectServer(ctx context.Context) (Server, error)
	AddServer(server Server) error
	RemoveServer(server Server) error
	GetServers() []Server
	GetHealthyServers() []Server
	Algorithm() string
}

type Server interface {
	ID() string
	URL() string
	IsHealthy() bool
	SetHealthy(healthy bool)
}

type HealthChecker interface {
	CheckHealth(server Server) bool
}
