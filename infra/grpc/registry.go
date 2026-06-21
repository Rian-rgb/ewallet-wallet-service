package grpc

import (
	"github.com/Rian-rgb/ewallet-common-lib/config"
)

type ConnRegistry struct {
	UmsConn *ClientConnWrapper
}

type RegistryConfig struct {
	UmsTarget string
}

func NewConnRegistry() (*ConnRegistry, func()) {
	registry := &ConnRegistry{}
	registryCFG := &RegistryConfig{
		UmsTarget: config.GetEnv("UMS_GRPC_HOST", ""),
	}

	if registryCFG.UmsTarget != "" {
		umsWrapper, err := NewClientConn(Config{Target: registryCFG.UmsTarget})
		if err != nil {
			return nil, nil
		}
		registry.UmsConn = umsWrapper
	}

	cleanup := func() {
		registry.CloseAll()
	}

	return registry, cleanup
}

func (r *ConnRegistry) CloseAll() {
	if r.UmsConn != nil {
		r.UmsConn.Close()
	}
}
