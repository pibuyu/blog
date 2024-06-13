// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"blog/rpc/internal/logic"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServiceServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServiceServer) GetUserRepositoryByUserIdentity(ctx context.Context, in *user.GetUserRepositoryByUserIdentityRequest) (*user.GetUserRepositoryByUserIdentityResponse, error) {
	l := logic.NewGetUserRepositoryByUserIdentityLogic(ctx, s.svcCtx)
	return l.GetUserRepositoryByUserIdentity(in)
}
