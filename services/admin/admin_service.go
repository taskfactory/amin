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
	var rsp admin.GetSourcesRsp
	count, data, err := sources.List(ctx, req.Sname, req.Page, req.PageSize)
	if err != nil {
		rsp.Code, rsp.Msg = errs.Code(err), errs.Msg(err)
	}
	rsp.Data.Page = req.Page
	rsp.Data.PageSize = req.PageSize
	rsp.Data.Total = count
	rsp.Data.Sources = data

	return rsp, err
}

// UpsertSource 更新或创建来源
func (s *Serv) UpsertSource(ctx context.Context, req *admin.UpsertSourceReq) (admin.UpsertSourceRsp, error) {
	var rsp admin.UpsertSourceRsp
	source, err := sources.Upsert(ctx, req.Id, req.Sname, req.Desc)
	if err != nil {
		rsp.Code, rsp.Msg = errs.Code(err), errs.Msg(err)
	}
	if source != nil {
		rsp.Source = *source
	}

	return rsp, err
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
