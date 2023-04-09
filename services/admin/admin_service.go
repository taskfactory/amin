package admin

import (
	"context"
	"github.com/taskfactory/admin/common/errs"
	"github.com/taskfactory/admin/logic/sources"
	"github.com/taskfactory/admin/tars-protocol/admin"
)

// Serv 服务结构
type Serv struct{}

// NewServ 构造函数
func NewServ() *Serv {
	return new(Serv)
}

// GetSources 获取来源列表
func (s *Serv) GetSources(ctx context.Context, req *admin.GetSourcesReq) (admin.GetSourcesRsp, error) {
	var ret admin.GetSourcesRsp
	data, err := sources.List(ctx, req.Sname, req.Page, req.PageSize)
	if err != nil {
		ret.Code, ret.Msg = errs.Code(err), errs.Msg(err)
	}
	ret.Sources = data

	return ret, err
}

// UpsertSource 更新或创建来源
func (s *Serv) UpsertSource(ctx context.Context, req *admin.UpsertSourceReq) (admin.UpsertSourceRsp, error) {
	var ret admin.UpsertSourceRsp
	source, err := sources.Upsert(ctx, req.Id, req.Sname, req.Desc)
	if err != nil {
		ret.Code, ret.Msg = errs.Code(err), errs.Msg(err)
	}
	if source != nil {
		ret.Source = *source
	}

	return ret, err
}

// Init servant
func (s *Serv) Init(ctx context.Context) error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destroy
func (s *Serv) Destroy(ctx context.Context) {
	//destroy servant here:
	//...
}
