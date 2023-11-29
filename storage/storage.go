package storage

import (
	"context"
	pb "editory_submission/genproto/auth_service"
)

type StorageI interface {
	CloseDB()
	Auth() AuthRepoI
}

type AuthRepoI interface {
	User() UserRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *pb.User) (res *pb.User, err error)
	Get(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error)
	GetList(ctx context.Context, req *pb.GetUserListReq) (res *pb.GetUserListRes, err error)
	Update(ctx context.Context, req *pb.User) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *pb.DeleteUserReq) (rowsAffected int64, err error)
}
