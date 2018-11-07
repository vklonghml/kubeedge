package modules

import (
	"fmt"

	"kubeedge/beehive/pkg/core"
	"kubeedge/beehive/pkg/core/context"
)

const (
	DestinationGroupModule = "destinationgroupmodule"
)

type testModuleDestGroup struct {
	context *context.Context
}

func init() {
	core.Register(&testModuleDestGroup{})
}

func (*testModuleDestGroup) Name() string {
	return DestinationGroupModule
}

func (*testModuleDestGroup) Group() string {
	return DestinationGroup
}

func (m *testModuleDestGroup) Start(c *context.Context) {
	m.context = c
	message, err := c.Receive(DestinationGroupModule)
	fmt.Printf("destination group module receive message:%v error:%v\n", message, err)
	if message.IsSync() {
		resp := message.NewRespByMessage(&message, "10 years old")
		c.SendResp(*resp)
	}
}

func (m *testModuleDestGroup) Cleanup() {
	m.context.Cleanup(m.Name())
}