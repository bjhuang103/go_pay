package logic

import (
	"context"
	"github.com/bjhuang103/go_pay/common/logs"

	"github.com/bjhuang103/go_pay/union_gateway/internal/svc"
	"github.com/bjhuang103/go_pay/union_gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (*types.AddResp, error) {
	// todo: add your logic here and delete this line
	ctx := l.ctx
	logs.CtxInfo(ctx, "enter add...., req:%v", req)

	return &types.AddResp{}, nil
}
