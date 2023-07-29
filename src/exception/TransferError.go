package exception

import (
	"time"
)

// TransferError 转换错误
type TransferError struct {
	Name    string
	Time    time.Time
	Message string
}

// Error
// @Description: 实现error接口
// @receiver     e      TransferError
// @return       string 错误信息
func (e TransferError) Error() string {
	return e.Name + "-" + e.Message + e.Time.Format("2006-01-02 15:04:05")
}

// NewTransferError
// @Description: 创建用户错误
// @param        name    string 错误名称
// @param        SyncMessage string 错误信息
// @return       *TransferError 转换错误
func NewTransferError(name string, Message string) *TransferError {
	return &TransferError{
		Name:    name,
		Message: Message,
		Time:    time.Now(),
	}
}
