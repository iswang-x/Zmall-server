package service

import (
	"context"
	"strconv"

	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	rpccart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	// 展示给用户一个后面需要结算的商品列表
	var items []map[string]string
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}
	return utils.H{
		"title":    "Checkout",
		"items":    items,
		"cart_num": len(items),
		"total":    strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
