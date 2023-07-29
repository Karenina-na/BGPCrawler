package exception

import "time"

// ConfigurationError 配置错误
type ConfigurationError struct {
	Name    string
	Time    time.Time
	Message string
}

// Error
// @Description: 获取错误信息
// @receiver     e      ConfigurationError
// @return       string 错误信息
func (e ConfigurationError) Error() string {
	return e.Name + "-" + e.Message + e.Time.Format("2006-01-02 15:04:05")
}

// NewConfigurationError
// @Description: 创建配置错误
// @param        name    错误名称
// @param        SyncMessage 错误信息
// @return       *ConfigurationError 配置错误
func NewConfigurationError(name string, Message string) *ConfigurationError {
	return &ConfigurationError{
		Name:    name,
		Message: Message,
		Time:    time.Now(),
	}
}
