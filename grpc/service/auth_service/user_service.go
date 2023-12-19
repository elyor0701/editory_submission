package auth_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/security"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"editory_submission/storage/postgres/models"
	"fmt"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type userService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *userService {
	return &userService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *pb.User) (res *pb.User, err error) {
	s.log.Info("---CreateUser--->", logger.Any("req", req))

	if ok := util.IsValidEmail(req.Email); !ok {
		err = fmt.Errorf("email is not valid")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}

	if len(req.Password) < 6 {
		err := fmt.Errorf("password must not be less than 6 characters")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	req.Password = hashedPassword

	if !util.IsValidPhone(req.GetPhone()) {
		req.Phone = ""
	}

	if !util.IsValidPhone(req.GetExtraPhone()) {
		req.ExtraPhone = ""
	}

	res, err = s.strg.Auth().User().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	author := false

	for _, v := range req.GetRole() {
		if v.GetRoleType() == "AUTHOR" {
			author = true
		}
		_, err = s.strg.Auth().Role().Create(ctx, &pb.Role{
			UserId:    res.GetId(),
			RoleType:  v.GetRoleType(),
			JournalId: v.GetJournalId(),
		})
		if err != nil {
			s.log.Error("!!!CreateRole--->", logger.Error(err))
			continue
		}
	}

	if !author {
		_, err = s.strg.Auth().Role().Create(ctx, &pb.Role{
			UserId:   res.GetId(),
			RoleType: "AUTHOR",
		})
		if err != nil {
			s.log.Error("!!!CreateRole--->", logger.Error(err))
		}
	}

	return res, nil
}

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserReq) (res *pb.User, err error) {
	s.log.Info("---GetUser--->", logger.Any("req", req))

	res, err = s.strg.Auth().User().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUser--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	role, err := s.strg.Auth().Role().GetList(ctx, &pb.GetRoleListReq{
		UserId: req.GetId(),
		Limit:  100,
	})
	if err != nil {
		s.log.Error("!!!GetRole--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	res.Role = role.GetRoles()

	return res, nil
}

func (s *userService) GetUserList(ctx context.Context, req *pb.GetUserListReq) (res *pb.GetUserListRes, err error) {
	s.log.Info("---GetUserList--->", logger.Any("req", req))

	res, err = s.strg.Auth().User().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUserList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *userService) GetUserListByRole(ctx context.Context, req *pb.GetUserListByRoleReq) (res *pb.GetUserListByRoleRes, err error) {
	s.log.Info("---GetUserListByRole--->", logger.Any("req", req))

	res, err = s.strg.Auth().User().GetListWithRole(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUserListByRole--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *pb.User) (res *pb.User, err error) {
	s.log.Info("---UpdateUser--->", logger.Any("req", req))

	// validate data

	rowsAffected, err := s.strg.Auth().User().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return req, nil
}

func (s *userService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteUser--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Auth().User().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteUser--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *userService) GenerateEmailVerificationToken(ctx context.Context, req *pb.GenerateEmailVerificationTokenReq) (res *pb.GenerateEmailVerificationTokenRes, err error) {
	s.log.Info("---GenerateEmailVerificationToken--->", logger.Any("req", req))

	token, err := security.GenerateRandomString(64)
	if err != nil {
		s.log.Error("!!!GenerateEmailVerificationToken--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")

	emailVerif, err := s.strg.Auth().User().CreateEmailVerification(ctx, &models.CreateEmailVerificationReq{
		Email:     req.GetEmail(),
		Token:     token,
		ExpiresAt: expiresAt,
		UserId:    req.GetUserId(),
	})
	if err != nil {
		s.log.Error("!!!GenerateEmailVerificationToken--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GenerateEmailVerificationTokenRes{
		Email:     emailVerif.Email,
		Token:     emailVerif.Token,
		ExpiresAt: emailVerif.ExpiresAt,
	}, err
}

func (s *userService) EmailVerification(ctx context.Context, req *pb.EmailVerificationReq) (res *pb.EmailVerificationRes, err error) {
	s.log.Info("---EmailVerification--->", logger.Any("req", req))

	tokens, err := s.strg.Auth().User().GetEmailVerificationList(ctx, &models.GetEmailVerificationListReq{
		Email: req.GetEmail(),
	})
	if err != nil {
		s.log.Error("!!!EmailVerification--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	found := false
	userId := ""

	for _, v := range tokens.Tokens {
		expiresAt, _ := time.Parse(time.RFC3339, v.ExpiresAt)
		if v.Token == req.GetToken() && expiresAt.After(time.Now()) {
			found = true
			userId = v.UserId
			break
		}
	}

	if !found {
		return &pb.EmailVerificationRes{
			Status: found,
		}, nil
	}

	_, err = s.strg.Auth().User().UpdateUserEmailVerificationStatus(ctx, &models.UpdateUserEmailVerificationStatusReq{
		Email:              req.GetEmail(),
		VerificationStatus: true,
	})
	if err != nil {
		s.log.Error("!!!EmailVerification--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = s.strg.Auth().User().DeleteEmailVerification(ctx, &models.DeleteEmailVerificationReq{
		Email: req.GetEmail(),
	})
	if err != nil {
		s.log.Error("!!!EmailVerification--->", logger.Error(err))
	}

	return &pb.EmailVerificationRes{
		Status: true,
		UserId: userId,
	}, err
}
