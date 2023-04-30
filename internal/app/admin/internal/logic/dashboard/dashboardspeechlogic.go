package dashboard

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/fzf-labs/fpkg/third_api/speech"

	"github.com/zeromicro/go-zero/core/logx"
)

type DashboardSpeechLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDashboardSpeechLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DashboardSpeechLogic {
	return &DashboardSpeechLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DashboardSpeechLogic) DashboardSpeech(req *types.DashboardSpeechReq) (*types.DashboardSpeechResp, error) {
	resp := new(types.DashboardSpeechResp)
	word, _ := speech.GetWord()
	resp.Word = word
	return resp, nil
}
