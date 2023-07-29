package exception

import (
	"time"
)

// DownloadError 下载错误
type DownloadError struct {
	Name    string
	Time    time.Time
	Message string
}

// Error
// @Description: 实现error接口
// @receiver     e      DownloadError
// @return       string 错误信息
func (e DownloadError) Error() string {
	return e.Name + "-" + e.Message + e.Time.Format("2006-01-02 15:04:05")
}

// NewDownloadError
// @Description: 创建用户错误
// @param        name    string 错误名称
// @param        SyncMessage string 错误信息
// @return       *DownloadError 下载错误
func NewDownloadError(name string, Message string) *DownloadError {
	return &DownloadError{
		Name:    name,
		Message: Message,
		Time:    time.Now(),
	}
}
