package auth_service

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/logger"
	"editory_submission/pkg/security"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type sessionService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedSessionServiceServer
}

func NewSessionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *sessionService {
	return &sessionService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *sessionService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	s.log.Info("---Login--->", logger.Any("req", req))
	res := &pb.LoginRes{}

	if !util.IsValidEmail(req.GetEmail()) {
		err := errors.New("invalid username")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if len(req.Password) < 6 {
		err := errors.New("invalid password")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.strg.Auth().User().GetByEmail(ctx, &pb.GetUserReq{
		Email: req.GetEmail(),
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		err := errors.New("invalid username or password")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	match, err := security.ComparePassword(user.Password, req.Password)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !match {
		err := errors.New("username or password is wrong")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if !user.GetEmailVerification() {
		err := errors.New("email is not activated yet")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.UserFound = true
	res.User = user

	roleTypes := make([]string, 0)

	if req.GetXRole() == config.ADMIN {
		roleTypes = []string{config.SUPERADMIN, config.EDITOR}
	} else if req.GetXRole() == config.USER {
		roleTypes = []string{config.AUTHOR}
	} else {
		err := errors.New("not valid user type")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	roles, err := s.strg.Auth().Role().GetList(
		ctx,
		&pb.GetRoleListReq{ // @TODO limit offset
			UserId:    user.GetId(),
			RoleTypes: roleTypes,
		},
	)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.Roles = roles.GetRoles()

	if len(roles.GetRoles()) == 0 {
		err := errors.New("permission denied")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	} else if len(roles.GetRoles()) == 1 {
		res.RoleId = roles.GetRoles()[0].GetId()
	}

	s.log.Info("Login--->STRG: DeleteExpiredUserSessions", logger.Any("user_id", user.Id))
	rowsAffected, err := s.strg.Auth().Session().DeleteExpiredUserSessions(ctx, user.Id)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.log.Info("Login--->DeleteExpiredUserSessions", logger.Any("rowsAffected", rowsAffected))

	userSessionList, err := s.strg.Auth().Session().GetSessionListByUserID(ctx, user.Id)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.Sessions = userSessionList.Sessions

	sessionPKey, err := s.strg.Auth().Session().Create(ctx, &pb.CreateSessionReq{
		UserId:    user.Id,
		RoleId:    res.GetRoleId(),
		Ip:        "0.0.0.0",
		Data:      "additional json data",
		ExpiresAt: time.Now().Add(config.RefreshTokenExpiresInTime).Format(config.DatabaseTimeLayout),
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	session, err := s.strg.Auth().Session().GetByPK(ctx, &pb.SessionPrimaryKey{
		Id: sessionPKey.GetId(),
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	m := map[string]interface{}{
		"id": session.Id,
	}

	accessToken, err := security.GenerateJWT(m, config.AccessTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	refreshToken, err := security.GenerateJWT(m, config.RefreshTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	res.Token = &pb.Token{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.UpdatedAt,
		ExpiresAt:        session.ExpiresAt,
		RefreshInSeconds: int32(config.AccessTokenExpiresInTime.Seconds()),
	}

	return res, nil
}

func (s *sessionService) Logout(ctx context.Context, req *pb.LogoutReq) (*emptypb.Empty, error) {
	s.log.Info("---Logout--->", logger.Any("req", req))
	tokenInfo, err := security.ParseClaims(req.AccessToken, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Logout--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	rowsAffected, err := s.strg.Auth().Session().Delete(ctx, &pb.SessionPrimaryKey{
		Id: tokenInfo.ID,
	})
	if err != nil {
		s.log.Error("!!!Logout--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	s.log.Info("---Logout--->", logger.Any("tokenInfo", tokenInfo))
	s.log.Info("---Logout--->", logger.Any("rowsAffected", rowsAffected))

	return &emptypb.Empty{}, nil
}

func (s *sessionService) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (*pb.RefreshTokenRes, error) {
	res := &pb.RefreshTokenRes{}

	tokenInfo, err := security.ParseClaims(req.RefreshToken, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!RefreshToken--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if tokenInfo.ExpiresAt.Unix() < time.Now().Unix() {
		err := errors.New("token has been expired")
		s.log.Error("!!!HasAccess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	session := &pb.Session{}

	session, err = s.strg.Auth().Session().GetByPK(ctx, &pb.SessionPrimaryKey{
		Id: tokenInfo.ID,
	})
	if err != nil {
		s.log.Error("!!!RefreshToken--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if util.IsValidUUID(req.GetRoleId()) {
		session.RoleId = req.RoleId
	}

	_, err = s.strg.Auth().Session().Update(ctx, &pb.UpdateSessionReq{
		Id:        session.Id,
		UserId:    session.UserId,
		RoleId:    session.RoleId,
		Ip:        session.Ip,
		Data:      session.Data,
		ExpiresAt: session.ExpiresAt,
	})
	if err != nil {
		s.log.Error("!!!RefreshToken--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = s.strg.Auth().User().Get(ctx, &pb.GetUserReq{
		Id: session.UserId,
	})
	if err != nil {
		s.log.Error("!!!RefreshToken--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	m := map[string]interface{}{
		"id": session.Id,
	}

	accessToken, err := security.GenerateJWT(m, config.AccessTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	refreshToken, err := security.GenerateJWT(m, config.RefreshTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	res.Token = &pb.Token{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.UpdatedAt,
		ExpiresAt:        session.ExpiresAt,
		RefreshInSeconds: int32(config.AccessTokenExpiresInTime.Seconds()),
	}

	return res, nil
}

func (s *sessionService) HasAccess(ctx context.Context, req *pb.HasAccessReq) (*pb.HasAccessRes, error) {
	tokenInfo, err := security.ParseClaims(req.AccessToken, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!HasAccess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if tokenInfo.ExpiresAt.Unix() < time.Now().Unix() {
		err := errors.New("token has been expired")
		s.log.Error("!!!HasAccess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	session, err := s.strg.Auth().Session().GetByPK(ctx, &pb.SessionPrimaryKey{
		Id: tokenInfo.ID,
	})
	if err != nil {
		s.log.Error("!!!HasAccess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	fmt.Println(session)

	if !util.IsValidUUID(session.GetRoleId()) {
		err := errors.New("not valid session role")
		s.log.Error("!!!HasAccess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.strg.Auth().User().Get(ctx, &pb.GetUserReq{
		Id: session.UserId,
	})
	if err != nil {
		s.log.Error("!!!HasAccess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if !user.EmailVerification {
		err := errors.New("email hasn't been activated yet")
		s.log.Error("!!!HasAccess--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//_, err = s.strg.Scope().Upsert(ctx, &pb.UpsertScopeRequest{
	//	ClientPlatformId: req.ClientPlatformId,
	//	Path:             req.Path,
	//	Method:           req.Method,
	//})
	//if err != nil {
	//	s.log.Error("!!!HasAccess--->", logger.Error(err))
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	// DONT FORGET TO UNCOMMENT THIS!!!

	// hasAccess, err := s.strg.PermissionScope().HasAccess(ctx, user.RoleId, req.ClientPlatformId, req.Path, req.Method)
	// if err != nil {
	// 	s.log.Error("!!!HasAccess--->", logger.Error(err))
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }

	// if !hasAccess {
	// 	err = errors.New("access denied")
	// 	s.log.Error("!!!HasAccess--->", logger.Error(err))
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }

	return &pb.HasAccessRes{
		Id:        session.Id,
		UserId:    session.UserId,
		RoleId:    session.RoleId,
		Ip:        session.Ip,
		Data:      session.Data,
		ExpiresAt: session.ExpiresAt,
		CreatedAt: session.CreatedAt,
		UpdatedAt: session.UpdatedAt,
	}, nil
}
