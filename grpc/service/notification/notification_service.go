package notification

import (
	"context"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	pb "editory_submission/genproto/notification_service"
	"editory_submission/grpc/client"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/logger"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/url"
	"time"
)

const (
	workers = 2
)

type notificationService struct {
	cfg        config.Config
	log        logger.LoggerI
	strg       storage.StorageI
	services   client.ServiceManagerI
	notifyChan chan string
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *notificationService {
	notifyChan := make(chan string, 50)

	for i := 0; i < workers; i++ {
		go SendMessage(notifyChan, svcs)
	}

	return &notificationService{
		cfg:        cfg,
		log:        log,
		strg:       strg,
		services:   svcs,
		notifyChan: notifyChan,
	}
}

func SendMessage(ch <-chan string, services client.ServiceManagerI) {
	cfg := config.Load()
	fmt.Println("SendMessage")
	for v := range ch {
		fmt.Println("SendMessage")
		var err error
		defer func() {
			if err != nil {
				ctx := context.Background()
				_, err = services.NotificationService().UpdateNotification(ctx, &pb.UpdateNotificationReq{
					Id:     v,
					Status: config.EMAIL_STATUS_FAILED,
				})
				if err != nil {
					fmt.Printf("!!!SendMessage---> %s", err.Error())
				}
			}
		}()

		ctx, finish := context.WithTimeout(context.Background(), 15*time.Second)
		defer finish()
		res, err := services.NotificationService().GetNotification(ctx, &pb.GetNotificationReq{
			Id: v,
		})
		if err != nil {
			fmt.Printf("!!!SendMessage---> %s", err.Error())
			continue
		}

		_, err = services.NotificationService().UpdateNotification(ctx, &pb.UpdateNotificationReq{
			Id:     v,
			Status: config.EMAIL_STATUS_PENDING,
		})
		if err != nil {
			fmt.Printf("!!!SendMessage---> %s", err.Error())
			continue
		}

		err = helper.GoMessageSend(helper.SendMessageByEmail{
			From: helper.EmailInfo{
				Username: cfg.EmailUsername,
				Password: cfg.EmailPassword,
			},
			To:      res.GetEmail(),
			Subject: res.Subject,
			Message: res.Text,
		})
		if err != nil {
			fmt.Printf("!!!SendMessage---> %s", err.Error())
			continue
		}

		_, err = services.NotificationService().UpdateNotification(ctx, &pb.UpdateNotificationReq{
			Id:     v,
			Status: config.EMAIL_STATUS_SENT,
		})
		if err != nil {
			fmt.Printf("!!!SendMessage---> %s", err.Error())
			continue
		}
	}
}

func (s *notificationService) CreateNotification(ctx context.Context, req *pb.CreateNotificationReq) (res *pb.CreateNotificationRes, err error) {
	s.log.Info("---CreateNotification--->", logger.Any("req", req))

	res, err = s.strg.Notification().Notification().Create(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateNotification--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *notificationService) GetNotification(ctx context.Context, req *pb.GetNotificationReq) (res *pb.GetNotificationRes, err error) {
	s.log.Info("---GetNotification--->", logger.Any("req", req))

	res, err = s.strg.Notification().Notification().Get(ctx, req)
	if err != nil {
		s.log.Error("!!!GetNotification--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *notificationService) GetNotificationList(ctx context.Context, req *pb.GetNotificationListReq) (res *pb.GetNotificationListRes, err error) {
	s.log.Info("---GetNotificationList--->", logger.Any("req", req))

	res, err = s.strg.Notification().Notification().GetList(ctx, req)
	if err != nil {
		s.log.Error("!!!GetNotificationList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *notificationService) UpdateNotification(ctx context.Context, req *pb.UpdateNotificationReq) (res *pb.UpdateNotificationRes, err error) {
	s.log.Info("---UpdateNotification--->", logger.Any("req", req))

	res, err = s.strg.Notification().Notification().Update(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateNotification--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *notificationService) DeleteNotification(ctx context.Context, req *pb.DeleteNotificationReq) (res *emptypb.Empty, err error) {
	s.log.Info("---DeleteNotification--->", logger.Any("req", req))

	res = &emptypb.Empty{}

	rowsAffected, err := s.strg.Notification().Notification().Delete(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteNotification--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *notificationService) GenerateMailMessage(ctx context.Context, req *pb.GenerateMailMessageReq) (*pb.GenerateMailMessageRes, error) {
	s.log.Info("---GenerateMailMessage--->", logger.Any("req", req))

	mailData := make(map[string]string)

	if !util.IsValidUUID(req.UserId) {
		err := errors.New("invalid user id")
		return nil, err
	}

	user, err := s.services.UserService().GetUser(
		ctx,
		&auth_service.GetUserReq{
			Id: req.UserId,
		},
	)
	if err != nil {
		return nil, err
	}

	mailData["first_name"] = user.FirstName
	mailData["last_name"] = user.LastName
	mailData["email"] = user.Email
	mailData["phone"] = user.Phone

	token, err := s.services.UserService().GenerateEmailVerificationToken(ctx, &auth_service.GenerateEmailVerificationTokenReq{
		Email:  user.GetEmail(),
		UserId: user.GetId(),
	})
	if err != nil {
		return nil, err
	}

	redirectUrl, err := url.Parse(req.RedirectLink)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("token", token.Token)
	values.Add("email", token.Email)

	redirectUrl.RawQuery = values.Encode()

	mailData["link"] = redirectUrl.String()

	tmp, err := s.services.EmailTmpService().GetEmailTmpList(
		ctx,
		&pb.GetEmailTmpListReq{
			Type: req.Type,
		},
	)
	if err != nil {
		return nil, err
	}

	if len(tmp.EmailTmps) == 0 {
		err := errors.New("cant find suitable mail template")
		return nil, err
	}

	subject, mailBody := helper.MakeEmailMessage(
		mailData,
		tmp.EmailTmps[0].GetTitle(),
		tmp.EmailTmps[0].GetText(),
	)

	res, err := s.services.NotificationService().CreateNotification(
		ctx,
		&pb.CreateNotificationReq{
			Subject: subject,
			Text:    mailBody,
			Email:   user.Email,
			Status:  config.EMAIL_STATUS_NEW,
		},
	)
	if err != nil {
		return nil, err
	}

	//err = helper.GoMessageSend(helper.SendMessageByEmail{
	//	From: helper.EmailInfo{
	//		Username: h.cfg.EmailUsername,
	//		Password: h.cfg.EmailPassword,
	//	},
	//	To:      res.GetEmail(),
	//	Subject: "Email Verification",
	//	Message: message,
	//})
	//if err != nil {
	//	return err
	//}

	fmt.Println("here")
	s.notifyChan <- res.Id
	fmt.Println("here")

	return &pb.GenerateMailMessageRes{
		Id:        res.Id,
		Subject:   res.Subject,
		Text:      res.Text,
		Email:     res.Email,
		Status:    res.Status,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}
