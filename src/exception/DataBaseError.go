package exception

import "time"

// DataBaseError 数据库错误
type DataBaseError struct {
	Name    string
	Time    time.Time
	Message string
}

// Error
// @Description: 实现error接口
// @receiver     e      错误对象
// @return       string 错误信息
func (e DataBaseError) Error() string {
	return e.Name + "-" + e.Message + e.Time.Format("2006-01-02 15:04:05")
}

// NewDataBaseError
// @Description: 创建数据库错误对象
// @param        name    错误名称
// @param        SyncMessage 错误信息
// @return       *DataBaseError 错误对象
func NewDataBaseError(name string, Message string) *DataBaseError {
	return &DataBaseError{
		Name:    name,
		Message: Message,
		Time:    time.Now(),
	}
}
