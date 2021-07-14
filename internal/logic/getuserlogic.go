package logic

import (
  "context"

  "github.com/pz2147/mall-rpc-github/internal/svc"
  "github.com/pz2147/mall-rpc-github/user"

  "github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
  ctx    context.Context
  svcCtx *svc.ServiceContext
  logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
  return &GetUserLogic{
      ctx:    ctx,
      svcCtx: svcCtx,
      Logger: logx.WithContext(ctx),
  }
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
  return &user.UserResponse{
      Id:   "1",
      Name: "test",
  }, nil
}