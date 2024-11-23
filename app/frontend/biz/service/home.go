package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "goods-01", "Price": 99.99, "Picture": "/static/image/bear01.png"},
		{"Name": "goods-02", "Price": 98.99, "Picture": "/static/image/head1.jpg"},
		{"Name": "goods-03", "Price": 99.99, "Picture": "/static/image/bear01.png"},
		{"Name": "goods-04", "Price": 98.99, "Picture": "/static/image/head1.jpg"},
		{"Name": "goods-05", "Price": 99.99, "Picture": "/static/image/bear01.png"},
		{"Name": "goods-06", "Price": 98.99, "Picture": "/static/image/head1.jpg"},
		{"Name": "goods-07", "Price": 99.99, "Picture": "/static/image/bear01.png"},
		{"Name": "goods-08", "Price": 98.99, "Picture": "/static/image/head1.jpg"},
		{"Name": "goods-09", "Price": 99.99, "Picture": "/static/image/bear01.png"},
		{"Name": "goods-10", "Price": 98.99, "Picture": "/static/image/head1.jpg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
