package exception

import (
	"time"
)

// DeleteError 删除错误
type DeleteError struct {
	Name    string
	Time    time.Time
	Message string
}

// Error
// @Description: 实现error接口
// @receiver     e      DeleteError
// @return       string 错误信息
func (e DeleteError) Error() string {
	return e.Name + "-" + e.Message + e.Time.Format("2006-01-02 15:04:05")
}

// NewDeleteError
// @Description: 创建用户错误
// @param        name    string 错误名称
// @param        SyncMessage string 错误信息
// @return       *DeleteError 删除错误
func NewDeleteError(name string, Message string) *DeleteError {
	return &DeleteError{
		Name:    name,
		Message: Message,
		Time:    time.Now(),
	}
}
