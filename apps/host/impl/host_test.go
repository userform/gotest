package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/userform/gotest/apps/host"
	"github.com/userform/gotest/apps/host/impl"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	fmt.Println(ins.Resource)
	service.CreateHost(context.Background(), ins)
}

func init() {
	// 需要初始化全局Logger
	// 为什么不设计未默认打印，因为性能
	fmt.Println(zap.DevelopmentSetup())

	// host service 的具体实现
	service = impl.NewHostServiceImpl()

}
