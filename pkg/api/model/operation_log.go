package model

import "time"

//operation log mapping
type OperationLog struct {
	Id               int       `json:"id"`
	LogNo            string    `form:"log_no" json:"log_no" binding:"required"`
	Module           string    `form:"module" json:"module"`
	RequestUrl       string    `form:"request_url" json:"request_url"`
	OperationMethod  string    `form:"operation_method" json:"operation_method"`
	Params           string    `form:"params" json:"params"`
	ExceptionStack   string    `form:"exception_stack" json:"exception_stack"`
	OperationResult  string    `form:"operation_result" json:"operation_result"`
	OperationSuccess int       `form:"operation_success" json:"operation_success"`
	OperationTime    time.Time `form:"operation_time" json:"operation_time"`
	UserId           string    `form:"user_id" json:"user_id"`
	Ip               string    `form:"ip" json:"ip"`
	IpLocation       string    `form:"ip_location" json:"ip_location"`
	OperationContent string    `form:"operation_content" json:"operation_content"`
	CreateTime       time.Time `gorm:"type:time;column:create_time;not null;default:CURRENT_TIMESTAMP" json:"created_time,omitempty" example:"2019-07-10 0:39"`
	LastUpdateTime   time.Time `gorm:"type:time;column:last_update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE" json:"last_update_time,omitempty" example:"2019-07-10 0:39"`
}

func (log *OperationLog) TableName() string {
	return "operationL_log"
}