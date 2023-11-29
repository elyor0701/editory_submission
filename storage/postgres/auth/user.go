package auth

import (
	"context"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.UserRepoI {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(ctx context.Context, req *pb.User) (res *pb.User, err error) {
	return nil, nil
}

func (u *UserRepo) Get(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error) {
	return nil, nil
}

func (u *UserRepo) GetList(ctx context.Context, req *pb.GetUserListReq) (res *pb.GetUserListRes, err error) {
	return nil, nil
}
func (u *UserRepo) Update(ctx context.Context, req *pb.User) (rowsAffected int64, err error) {
	return 0, nil
}
func (u *UserRepo) Delete(ctx context.Context, req *pb.DeleteUserReq) (rowsAffected int64, err error) {
	return 0, nil
}
