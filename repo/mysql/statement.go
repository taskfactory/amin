package mysql

import (
	"strings"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// Statement 查询构建工具
type Statement struct {
	fields        string
	where         builder.Cond
	orderBy       string
	limit, offset int
	groupBy       string
}

// NewStatement 构造函数
func NewStatement() *Statement {
	st := new(Statement)
	st.where = builder.NewCond()

	return st
}

// Select 查询字段
func (st *Statement) Select(fields ...string) *Statement {
	if len(fields) == 0 {
		return st
	}
	st.fields = "`" + strings.Join(fields, "`,`") + "`"

	return st
}

// AndEqual 等于
func (st *Statement) AndEqual(key string, val interface{}) *Statement {
	st.where = st.where.And(builder.Eq{key: val})

	return st
}

// AndIn in查询
func (st *Statement) AndIn(key string, val ...interface{}) *Statement {
	st.where = st.where.And(builder.In(key, val...))

	return st
}

// AndLt < 查询
func (st *Statement) AndLt(key string, val interface{}) *Statement {
	st.where = st.where.And(builder.Lt{key: val})

	return st
}

// AndLte <= 查询
func (st *Statement) AndLte(key string, val interface{}) *Statement {
	st.where = st.where.And(builder.Lte{key: val})

	return st
}

// AndGt > 查询
func (st *Statement) AndGt(key string, val interface{}) *Statement {
	st.where = st.where.And(builder.Gt{key: val})

	return st
}

// AndGte >= 查询
func (st *Statement) AndGte(key string, val interface{}) *Statement {
	st.where = st.where.And(builder.Gte{key: val})

	return st
}

// AndLike like 查询
func (st *Statement) AndLike(key string, val string) *Statement {
	st.where = st.where.And(builder.Like{key, val})

	return st
}

// OrderBy 查询
func (st *Statement) OrderBy(order string) *Statement {
	st.orderBy = order

	return st
}

// Limit limit 查询
func (st *Statement) Limit(limit int, offset ...int) *Statement {
	st.limit = limit
	if len(offset) > 0 {
		st.offset = offset[0]
	}

	return st
}

// GroupBy group by查询
func (st *Statement) GroupBy(groupBy ...string) *Statement {
	if len(groupBy) == 0 {
		return st
	}
	st.groupBy = "`" + strings.Join(groupBy, "`,`") + "`"

	return st
}

// GetLimit 获取limit
func (st *Statement) GetLimit() int {
	return st.limit
}

// BuildSelect 构建查询
func (st *Statement) BuildSelect(clt *xorm.Session) *xorm.Session {
	if st.fields != "" {
		clt.Select(st.fields)
	}
	if st.where != nil {
		clt.Where(st.where)
	}
	if st.orderBy != "" {
		clt.OrderBy(st.orderBy)
	}
	if st.limit > 0 {
		clt.Limit(st.limit, st.offset)
	}
	if st.groupBy != "" {
		clt.GroupBy(st.groupBy)
	}

	return clt
}

// BuildCond 仅构建条件，用于更新或删除等场景
func (st *Statement) BuildCond(clt *xorm.Session) *xorm.Session {
	if st.where != nil {
		clt.Where(st.where)
	}
	if st.limit > 0 {
		clt.Limit(st.limit, st.offset)
	}

	return clt
}
