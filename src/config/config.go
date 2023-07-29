package config

var (
	// Goroutine goroutine参数
	Goroutine struct {
		// MaxRoutineNum goroutine池最大线程数
		MaxRoutineNum int
		// CoreRoutineNum goroutine池核心线程数
		CoreRoutineNum int
		// RoutineTimeOut goroutine池线程超时时间
		RoutineTimeOut int
	}

	// Database 数据库参数
	Database struct {
		// DatabaseType 数据库类型
		DatabaseType string
		// DatabaseHost 数据库地址
		DatabaseHost string
		// DatabasePort 数据库端口
		DatabasePort string
		// DatabaseUser 数据库用户名
		DatabaseUser string
		// DatabasePassword 数据库密码
		DatabasePassword string
	}

	// BGP BGP参数
	BGP struct {
		// Frequency BGP数据采集频率
		Frequency int
		// StoragePath BGP原始数据存储路径
		StoragePath string
		// ProcessPath BGP处理数据存储路径
		ProcessPath string
		// StorageTime BGP原始数据存储时间
		StorageTime int
		// ProcessTime BGP处理数据存储时间
		ProcessTime int
	}
)
