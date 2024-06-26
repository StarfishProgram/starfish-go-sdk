package sdkdb

import (
	"fmt"
	"time"

	"github.com/StarfishProgram/starfish-go-sdk/sdk"
	"github.com/StarfishProgram/starfish-go-sdk/sdklog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Config 数据源配置
type Config struct {
	Host        string `toml:"host" yaml:"host"`               // 主机
	Port        *int   `toml:"port" yaml:"port"`               // 端口
	Username    string `toml:"username" yaml:"username"`       // 用户名
	Password    string `toml:"password" yaml:"password"`       // 密码
	Database    string `toml:"database" yaml:"database"`       // 数据库
	Params      string `toml:"params" yaml:"params"`           // 连接属性
	MaxIdleConn *int   `toml:"maxIdleConn" yaml:"maxIdleConn"` // 最大空闲连接数
	MaxOpenConn *int   `toml:"maxOpenConn" yaml:"maxOpenConn"` // 最大连接数
	MaxLifetime *int64 `toml:"maxLifetime" yaml:"maxLifetime"` // 最大超时时间(秒)
	ShowSql     bool   `toml:"showSql" yaml:"showSql"`         // 显示执行SQL
	SlowTime    *int64 `toml:"slowTime" yaml:"slowTime"`       // 慢查询时间(毫秒)
}

type _LogWriter struct{}

func (*_LogWriter) Printf(format string, v ...interface{}) {
	sdklog.Infof(format, v...)
}

var ins map[string]*gorm.DB

func init() {
	ins = make(map[string]*gorm.DB)
}

// Init 数据源初始化
func Init(config *Config, key ...string) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v",
		config.Username,
		config.Password,
		config.Host,
		sdk.IfNil(config.Port, 3306),
		config.Database,
		config.Params,
	)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         1000,
		SkipInitializeWithVersion: false,
	}

	dbLog := logger.New(
		&_LogWriter{},
		logger.Config{
			SlowThreshold:             time.Millisecond * time.Duration(sdk.IfNil(config.SlowTime, 200)),
			LogLevel:                  sdk.If(config.ShowSql, logger.Info, logger.Error),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
			// ParameterizedQueries:      true,
		},
	)
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		CreateBatchSize:        1000,
		Logger:                 dbLog,
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
	})

	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if config.MaxIdleConn != nil {
		sqlDB.SetMaxIdleConns(*config.MaxIdleConn)
	}
	if config.MaxOpenConn != nil {
		sqlDB.SetMaxOpenConns(*config.MaxOpenConn)
	}
	if config.MaxLifetime != nil {
		sqlDB.SetConnMaxIdleTime(time.Duration(*config.MaxLifetime) * time.Second)
	}

	if len(key) == 0 {
		ins[""] = db
	} else {
		ins[key[0]] = db
	}
}

// Ins 获取数据源
func Ins(key ...string) *gorm.DB {
	if len(key) == 0 {
		return ins[""]
	} else {
		return ins[key[0]]
	}
}

func LockingForUpdate() clause.Locking {
	return clause.Locking{Strength: "UPDATE"}
}
