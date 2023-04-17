package sources

import (
	"context"
	"github.com/taskfactory/admin/common/errs"
	"github.com/taskfactory/admin/entity"
	"github.com/taskfactory/admin/repo/mysql"
	"github.com/taskfactory/admin/repo/mysql/tblsources"
	proto "github.com/taskfactory/admin/tars-protocol/admin"
)

// List 列表查询
func List(ctx context.Context, sname string, page, pageSize int32) (int32, []proto.Source, error) {
	st := mysql.NewStatement()
	if sname != "" {
		st.AndLike("sname", sname)
	}
	st.OrderBy("id asc").Limit(int(pageSize), int(pageSize*(page-1)))
	count, sources, err := tblsources.NewModel().
		FindAndCount(ctx, st)
	if err != nil {
		err = errs.Newf(errs.CodeDBRead, "failed to query db,err:%v", err)
		return 0, nil, err
	}

	var ret []proto.Source
	for _, source := range sources {
		ret = append(ret, source.ToProto())
	}

	return int32(count), ret, nil
}

// Upsert 创建或更新
func Upsert(ctx context.Context, id int64, sname, desc string) (*proto.Source, error) {
	var row *entity.Source
	var err error
	if id == 0 {
		sid, e := New(ctx, sname, desc)
		if e != nil {
			return nil, e
		}
		st := mysql.NewStatement().Limit(1).AndEqual("sid", sid)
		row, err = tblsources.NewModel().FindOne(ctx, st)
	} else {
		err = Update(ctx, id, sname, desc)
		st := mysql.NewStatement().Limit(1).AndEqual("id", id)
		row, err = tblsources.NewModel().FindOne(ctx, st)
	}

	if err != nil {
		return nil, errs.Newf(errs.CodeDBRead, "failed to query record,err:%v", err)
	}
	if row == nil {
		return nil, errs.Newf(errs.CodeDBWrite, "failed to insert new record to db")
	}
	source := row.ToProto()

	return &source, nil
}

// New 创建新数据源
func New(ctx context.Context, sname, desc string) (uint16, error) {
	row := &entity.Source{
		SID:   10000,
		SName: sname,
		Desc:  desc,
	}
	sid, err := tblsources.NewModel().FindMaxSID(ctx)
	if err != nil {
		err = errs.Newf(errs.CodeDBRead, "failed to query db,err:%v", err)
		return 0, err
	}
	if sid > 0 {
		row.SID = sid + 1
	}
	err = tblsources.NewModel().Insert(ctx, row)
	if err != nil {
		err = errs.Newf(errs.CodeDBWrite, "failed to insert record to db,err:%v", err)
	}

	return row.SID, err
}

// Update 更新记录
func Update(ctx context.Context, id int64, sname, desc string) error {
	st := mysql.NewStatement().AndEqual("id", id).Limit(1)
	upMap := map[string]interface{}{
		"sname": sname,
		"desc":  desc,
	}
	_, err := tblsources.NewModel().Update(ctx, st, upMap)
	if err != nil {
		err = errs.Newf(errs.CodeDBWrite, "failed to update record, err:%v", err)
	}

	return err
}
