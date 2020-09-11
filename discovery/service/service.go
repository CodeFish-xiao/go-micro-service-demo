package service

import (
	"context"
	"errors"
	"discovery/config"
	"discovery/discover"
)

type Service interface {


	// 健康检查接口
	HealthCheck() bool

	// sayHelloService
	SayHello() string

	//  服务发现接口
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)

}


var ErrNotServiceInstances = errors.New("instances are not existed")


//服务发现接口具体实现
type DiscoveryServiceImpl struct {
	discoveryClient discover.DiscoveryClient//依赖于DiscoveryClient服务注册接口
}

func NewDiscoveryServiceImpl(discoveryClient discover.DiscoveryClient) Service  {
	return &DiscoveryServiceImpl{
		discoveryClient:discoveryClient,
	}
}

func (*DiscoveryServiceImpl) SayHello() string {
	return "Hello World!"
}

func (service *DiscoveryServiceImpl) DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)  {
	//从consul中根据服务名获取实例列表
	instances := service.discoveryClient.DiscoverServices(serviceName, config.Logger)

	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}
	return instances, nil
}


// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true
func (*DiscoveryServiceImpl) HealthCheck() bool {
	return true
}

