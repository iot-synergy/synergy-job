package task

import (
	"context"

	"github.com/iot-synergy/synergy-job/internal/svc"
	"github.com/iot-synergy/synergy-job/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-job/types/job"

	"github.com/iot-synergy/synergy-common/i18n"

	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskLogic) UpdateTask(in *job.TaskInfo) (*job.BaseResp, error) {
	err := l.svcCtx.DB.Task.UpdateOneID(*in.Id).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilName(in.Name).
		SetNotNilTaskGroup(in.TaskGroup).
		SetNotNilCronExpression(in.CronExpression).
		SetNotNilPattern(in.Pattern).
		SetNotNilPayload(in.Payload).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
