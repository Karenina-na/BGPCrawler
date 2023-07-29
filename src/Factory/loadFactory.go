package Factory

import (
	"BGP/src/config"
	"BGP/src/exception"
	"BGP/src/pool"
	"BGP/src/util"
	"github.com/spf13/viper"
	"os"
)

// LoadConfigFactory
//
//	@Description: 加载配置文件工厂
//	@param arg	加载模式
func LoadConfigFactory(arg *string) {
	defer func() {
		r := recover()
		if r != nil {
			exception.HandleException(exception.NewSystemError("LoadConfigFactory", util.Strval(r)))
		}
	}()
	if *arg == "debug" {
		util.LoggerInit(func(r any) {
			util.Loglevel(util.Error, "main", util.Strval(r))
		}, util.Debug)
		util.Loglevel(util.Info, "main", "debug mode")
	} else {
		util.LoggerInit(func(r any) {
			util.Loglevel(util.Error, "main", util.Strval(r))
		}, util.Info)
		util.Loglevel(util.Info, "main", "release mode")
	}

	// 初始化配置文件
	util.Loglevel(util.Debug, "LoadConfigFactory", "初始化配置文件")
	if err := loadConfigFactory(); err != nil {
		exception.HandleException(err)
	}

	//初始化协程池
	if err := pool.InitRoutinePool(); err != nil {
		exception.HandleException(err)
	}
	util.Loglevel(util.Debug, "LoadConfigFactory", "初始化协程池")
}

func loadConfigFactory() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("loadConfigFactory", util.Strval(r))
		}
	}()

	// config
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err == nil {
		// goroutine pool
		config.Goroutine.MaxRoutineNum = viper.GetInt("goroutine.max-goroutine")
		if !config.VerifyReg(config.PositiveReg, util.Strval(config.Goroutine.MaxRoutineNum)) {
			return exception.NewConfigurationError("LoadConfigFactory-config", "goroutine.max-goroutine 非法")
		}
		config.Goroutine.CoreRoutineNum = viper.GetInt("goroutine.core-goroutine")
		if !config.VerifyReg(config.PositiveReg, util.Strval(config.Goroutine.CoreRoutineNum)) {
			return exception.NewConfigurationError("LoadConfigFactory-config", "goroutine.core-goroutine 非法")
		}
		if config.Goroutine.CoreRoutineNum > config.Goroutine.MaxRoutineNum {
			return exception.NewConfigurationError("LoadConfigFactory-config", "goroutine.core-goroutine 大于 goroutine.max-goroutine")
		}

		// database
		config.Database.DatabaseType = viper.GetString("database.type")
		if !config.VerifyReg(config.DatabaseTypeReg, config.Database.DatabaseType) {
			return exception.NewConfigurationError("LoadConfigFactory-config", "database.type 非法")
		}
		config.Database.DatabaseHost = viper.GetString("database.host")
		if !config.VerifyReg(config.IpReg, config.Database.DatabaseHost) {
			return exception.NewConfigurationError("LoadConfigFactory-config", "database.host 非法")
		}
		config.Database.DatabasePort = viper.GetString("database.port")
		if !config.VerifyReg(config.PortReg, util.Strval(config.Database.DatabasePort)) {
			return exception.NewConfigurationError("LoadConfigFactory-config", "database.port 非法")
		}
		config.Database.DatabaseUser = viper.GetString("database.user")
		config.Database.DatabasePassword = viper.GetString("database.password")

		// BGP
		config.BGP.Frequency = viper.GetInt("BGP.frequency")
		if !config.VerifyReg(config.PositiveReg, util.Strval(config.BGP.Frequency)) {
			return exception.NewConfigurationError("LoadConfigFactory-config", "BGP.frequency 非法")
		}
		config.BGP.StoragePath = viper.GetString("BGP.storage-path")
		if e := os.MkdirAll(config.BGP.StoragePath, os.ModePerm); e != nil {
			return exception.NewSystemError("LoadConfigFactory-config", "创建路径失败 "+e.Error())
		}
		config.BGP.ProcessPath = viper.GetString("BGP.processed-path")
		if e := os.MkdirAll(config.BGP.ProcessPath, os.ModePerm); e != nil {
			return exception.NewSystemError("LoadConfigFactory-config", "创建路径失败 "+e.Error())
		}
		config.BGP.StorageTime = viper.GetInt("BGP.storage-time")
		if !config.VerifyReg(config.PositiveReg, util.Strval(config.BGP.StorageTime)) {
			return exception.NewConfigurationError("LoadConfigFactory-config", "BGP.storage-time 非法")
		}
		config.BGP.ProcessTime = viper.GetInt("BGP.processed-time")
		if !config.VerifyReg(config.PositiveReg, util.Strval(config.BGP.ProcessTime)) {
			return exception.NewConfigurationError("LoadConfigFactory-config", "BGP.processed-time 非法")
		}

	} else {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if err == configFileNotFoundError {
			return exception.NewConfigurationError("LoadConfigFactory-config", "配置文件不存在")
		} else {
			return exception.NewSystemError("LoadConfigFactory-config", "读取配置文件失败 "+err.Error())
		}
	}
	return nil
}
