package tblsources

import (
	"context"
	"github.com/taskfactory/admin/entity"
	"github.com/taskfactory/admin/repo/mysql"
	"xorm.io/builder"
)

// Model 模型结构
type Model struct{}

// NewModel 构造函数
func NewModel() *Model {
	return new(Model)
}

func (m *Model) TableName() string {
	return "tbl_sources"
}

func (m *Model) DBName() string {
	return `db_admin`
}

// FindOne 单条查询
func (m *Model) FindOne(ctx context.Context, st *mysql.Statement) (*entity.Source, error) {
	clt, err := mysql.NewSession(m.DBName())
	if err != nil {
		return nil, err
	}
	defer clt.Close()
	var row entity.Source
	clt = st.BuildSelect(clt.Table(m.TableName())).Limit(1).Context(ctx)
	has, err := clt.Get(&row)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, mysql.ErrNotFound
	}

	return &row, nil
}

// FindAll 多条查询
func (m *Model) FindAll(ctx context.Context, st *mysql.Statement) ([]*entity.Source, error) {
	clt, err := mysql.NewSession(m.DBName())
	if err != nil {
		return nil, err
	}
	defer clt.Close()

	if st.GetLimit() == 0 || st.GetLimit() > mysql.MaxPageLimit {
		st.Limit(mysql.MaxPageLimit)
	}
	var rows []*entity.Source
	clt = st.BuildSelect(clt.Table(m.TableName())).Context(ctx)
	err = clt.Find(&rows)
	if err != nil {
		return nil, err
	}
	clt.MustCols()

	return rows, nil
}

// Insert 插入单条记录
func (m *Model) Insert(ctx context.Context, row *entity.Source) error {
	clt, err := mysql.NewSession(m.DBName())
	if err != nil {
		return err
	}
	defer clt.Close()

	_, err = clt.Table(m.TableName()).
		Context(ctx).
		Insert(row)
	return err
}

// Update 更新单条记录
func (m *Model) Update(ctx context.Context, st *mysql.Statement, upMap map[string]interface{}) (int64, error) {
	clt, err := mysql.NewSession(m.DBName())
	if err != nil {
		return 0, err
	}
	defer clt.Close()
	clt = st.BuildCond(clt.Table(m.TableName())).Context(ctx)

	return clt.Update(upMap)
}

// FindMaxSID 获取最大SID
func (m *Model) FindMaxSID(ctx context.Context) (uint16, error) {
	clt, err := mysql.NewSession(m.DBName())
	if err != nil {
		return 0, err
	}
	defer clt.Close()
	sql, params, err := builder.Select("MAX(`sid`) AS `sid`").From(m.TableName()).ToSQL()
	if err != nil {
		return 0, err
	}
	var row entity.Source
	_, err = clt.SQL(sql, params...).Context(ctx).Get(&row)
	if err != nil {
		return 0, err
	}

	return row.SID, nil
}
