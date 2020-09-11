package service

import (
	"errors"
	"strings"
)

// Service constants：服务的常量
const (
	StrMaxSize = 1024
)

// Service errors：服务错误
var (
	ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")

	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

// Service Define a service interface：定义一个服务接口
type Service interface {
	// Concat a and b：连接字符串a,b
	Concat(a, b string) (string, error)

	// a,b pkg string value:获取a，b的值
	Diff(a, b string) (string, error)

	// HealthCheck check service health status：健康检查
	HealthCheck() bool
}

//ArithmeticService implement Service interface：实现服务接口
type StringService struct {
}

func (s StringService) Concat(a, b string) (string, error) {
	// test for length overflow：长度溢出试验
	if len(a)+len(b) > StrMaxSize {
		return "", ErrMaxSize
	}
	return a + b, nil
}

func (s StringService) Diff(a, b string) (string, error) {
	if len(a) < 1 || len(b) < 1 {
		return "", nil
	}
	res := ""
	if len(a) >= len(b) {
		for _, char := range b {
			if strings.Contains(a, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range a {
			if strings.Contains(b, string(char)) {
				res = res + string(char)
			}
		}
	}
	return res, nil
}

// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true。
func (s StringService) HealthCheck() bool {
	return true
}

// ServiceMiddleware define service middleware：定义服务中间件
type ServiceMiddleware func(Service) Service
