// Package consul provides a Consul registry. Implementation https://godoc.org/github.com/micro/go-micro/registry/consul
package consul

import (
	"github.com/asim/go-micro/tree/v1.18.0/registry"
	"github.com/Ovat1975/go-plugins/tree/master/registry/consul"
)

func NewRegistry(opts ...registry.Option) registry.Registry {
	return consul.NewRegistry(opts...)
}
