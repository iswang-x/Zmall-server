package service

import (
	"context"

	"testing"

	api "github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
)

func TestEcho_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoService(ctx)
	// init req and assert value

	req := &api.Request{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
