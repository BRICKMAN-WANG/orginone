package marketappgroupdistribution

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDistributedGroupIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDistributedGroupIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDistributedGroupIdListLogic {
	return GetDistributedGroupIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDistributedGroupIdListLogic) GetDistributedGroupIdList(req types.GetDistributedGroupIdListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
