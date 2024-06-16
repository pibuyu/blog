package logic

import (
	"blog/rpc/internal/models"
	"blog/rpc/internal/svc"
	"context"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	mux.Handle(models.USER_REGISTER_JOB, NewUserRegisterHandler(l.svcCtx))

	return mux
}
