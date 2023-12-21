package notification

import (
	"editory_submission/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type notificationRepo struct {
	db       *pgxpool.Pool
	emailTmp storage.EmailTemplateRepoI
	notify   storage.NotifyRepoI
}

func NewNotificationRepo(db *pgxpool.Pool) storage.NotificationRepoI {
	return &notificationRepo{
		db: db,
	}
}

func (n notificationRepo) Notification() storage.NotifyRepoI {
	if n.notify == nil {
		n.notify = NewNotifyRepo(n.db)
	}

	return n.notify
}

func (n notificationRepo) EmailTemplate() storage.EmailTemplateRepoI {
	if n.emailTmp == nil {
		n.emailTmp = NewEmailTmpRepo(n.db)
	}

	return n.emailTmp
}
