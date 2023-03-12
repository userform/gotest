package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/userform/gotest/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	fmt.Println("=============>")
	// 直接打印日志
	i.l.Debug("create host")
	// 带Format的日志打印,fmt.Sprintf()
	i.l.Debugf("create host %s", ins.Name)
	// 携带额外meta数据,常用于Trace系统
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostResquest) (*host.HostSet, error) {
	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostResquest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostResquest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostResquest) (*host.Host, error) {
	return nil, nil
}
