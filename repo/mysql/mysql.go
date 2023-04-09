package mysql

import (
	"errors"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/yaml.v3"
	"xorm.io/xorm"
)

// engines 全局mysql连接的clients实例
var engines *Engines
var cfgs map[string]config
var ErrNotFound = errors.New("record not found")

// Engines mysql client 结构
type Engines struct {
	sync.RWMutex
	Inst map[string]*xorm.Engine
}

type config struct {
	DBName   string `yaml:"db_name"`  // 数据库名称
	User     string `yaml:"user"`     // 用户名称
	Password string `yaml:"password"` // 用户密码
	Address  string `yaml:"address"`  // 数据库连接地址
	Args     string `yaml:"args"`     // 连接参数
}

func init() {
	once := new(sync.Once)
	once.Do(func() {
		cfgs = make(map[string]config)
		if engines == nil {
			fmt.Println("init")
			engines = new(Engines)
			engines.Inst = make(map[string]*xorm.Engine)
		}
	})
}

func Init(cfgStr string) error {
	err := yaml.Unmarshal([]byte(cfgStr), &cfgs)
	if err != nil {
		return err
	}
	engines.Lock()
	defer engines.Unlock()
	for instName, cfg := range cfgs {
		engines.Inst[instName], err = xorm.NewEngine("mysql", cfg.toDSN())
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *config) toDSN() string {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s", c.User, c.Password, c.Address, c.DBName)
	if len(c.Args) > 0 {
		dsn += "?" + c.Args
	}

	return dsn
}

// NewSession 创建一个新的MySQL 连接会话
func NewSession(instName string) (*xorm.Session, error) {
	engines.RLock()
	engine := engines.Inst[instName]
	if engine == nil {
		return nil, errors.New("no mysql instance found for " + instName)
	}
	err := engine.Ping()
	if err != nil {
		engine, err := refreshEngine(instName)
		if err != nil {
			return nil, errors.New("cannot connect to mysql instance " + instName)
		}
		engines.RUnlock()
		engines.Lock()
		engines.Inst[instName] = engine
		engines.Unlock()
		return engine.NewSession(), nil
	}
	engines.RUnlock()

	return engine.NewSession(), nil
}

func refreshEngine(instName string) (*xorm.Engine, error) {
	cfg := cfgs[instName]
	return xorm.NewEngine("mysql", cfg.toDSN())
}
