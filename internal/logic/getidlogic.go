package logic

import (
	"context"

	"github.com/pz2147/mall-rpc-github/internal/svc"
	"github.com/pz2147/mall-rpc-github/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdLogic {
	return &GetIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIdLogic) GetId(in *user.NameRequest) (*user.UserNameResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserNameResponse{}, nil
}
