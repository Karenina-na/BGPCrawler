package exception

import "time"

// SystemError 系统错误
type SystemError struct {
	Name    string
	Time    time.Time
	Message string
}

// Error
// @Description: 获取错误信息
// @receiver     e      系统错误
// @return       string 错误信息
func (e SystemError) Error() string {
	return e.Name + "-" + e.Message + e.Time.Format("2006-01-02 15:04:05")
}

// NewSystemError
// @Description: 创建系统错误
// @param        name    错误名称
// @param        SyncMessage 错误信息
// @return       *SystemError 系统错误
func NewSystemError(name string, Message string) *SystemError {
	return &SystemError{
		Name:    name,
		Message: Message,
		Time:    time.Now(),
	}
}
