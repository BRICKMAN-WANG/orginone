package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencydeletedinnernodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencydeletedinnernodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencydeletedinnernodeLogic {
	return AgencydeletedinnernodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencydeletedinnernodeLogic) Agencydeletedinnernode(req types.AgencyDeletedInnerNodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
