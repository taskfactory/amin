package entity

import (
	"github.com/taskfactory/admin/common/constant"
	proto "github.com/taskfactory/admin/tars-protocol/admin"
	"time"
)

// Source 数据来源模型结构定义
// SID 数据来源ID
// SName 数据来源名称
// Desc 数据来源描述信息
// ConfVer 当前生效的配置版本
// CreatedAt 创建时间
// UpdatedAt 更新时间
type Source struct {
	ID        uint64    `xorm:"'id' pk autoincr bigint"`
	SID       uint16    `xorm:"'sid' smallint"`
	SName     string    `xorm:"'sname' varchar"`
	Desc      string    `xorm:"'desc' varchar"`
	ConfVer   uint16    `xorm:"'conf_ver' smallint"`
	CreatedAt time.Time `xorm:"'created_at' created"`
	UpdatedAt time.Time `xorm:"'updated_at' updated"`
}

// ToProto 转化为proto的source对象
func (s Source) ToProto() proto.Source {
	return proto.Source{
		Id:        int64(s.ID),
		Sid:       s.SID,
		Sname:     s.SName,
		Desc:      s.Desc,
		Confver:   s.ConfVer,
		CreatedAt: s.CreatedAt.Format(constant.TimeLayoutSec),
		UpdatedAt: s.UpdatedAt.Format(constant.TimeLayoutSec),
	}
}
