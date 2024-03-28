package sys_job

import "x_admin/core"

// SysJobListReq 定时任务调度列表参数
type SysJobListReq struct {
	Name   string `form:"name"`   // 调度名称
	Status string `form:"status"` // 状态（0正常 1暂停）
}

// SysJobDetailReq 定时任务调度详情参数
type SysJobDetailReq struct {
	Id int `form:"id"` // 主键ID
}

// SysJobAddReq 定时任务调度新增参数
type SysJobAddReq struct {
	Name           string `form:"name"`           // 调度名称
	InvokeTarget   string `form:"invokeTarget"`   // 调用目标字符串
	CronExpression string `form:"cronExpression"` // cron执行表达式
	Concurrent     string `form:"concurrent"`     // 是否并发执行（0允许 1禁止）
	Status         string `form:"status"`         // 状态（0正常 1暂停）
	Remark         string `form:"remark"`         // 备注信息
}

// SysJobEditReq 定时任务调度编辑参数
type SysJobEditReq struct {
	Id             int    `form:"id"`             // 主键ID
	Name           string `form:"name"`           // 调度名称
	InvokeTarget   string `form:"invokeTarget"`   // 调用目标字符串
	CronExpression string `form:"cronExpression"` // cron执行表达式
	Concurrent     string `form:"concurrent"`     // 是否并发执行（0允许 1禁止）
	Status         string `form:"status"`         // 状态（0正常 1暂停）
	Remark         string `form:"remark"`         // 备注信息
}

// SysJobDelReq 定时任务调度新增参数
type SysJobDelReq struct {
	Id int `form:"id"` // 主键ID
}

// SysJobResp 定时任务调度返回信息
type SysJobResp struct {
	Id             int         `json:"id" structs:"id"`                         // 主键ID
	Name           string      `json:"name" structs:"name"`                     // 调度名称
	InvokeTarget   string      `json:"invokeTarget" structs:"invokeTarget"`     // 调用目标字符串
	CronExpression string      `json:"cronExpression" structs:"cronExpression"` // cron执行表达式
	Concurrent     string      `json:"concurrent" structs:"concurrent"`         // 是否并发执行（0允许 1禁止）
	Status         string      `json:"status" structs:"status"`                 // 状态（0正常 1暂停）
	IsDelete       int         `json:"isDelete" structs:"isDelete"`             // 是否删除: 0=否, 1=是
	CreateTime     core.TsTime `json:"createTime" structs:"createTime"`         // 创建时间
	CreateBy       string      `json:"createBy" structs:"createBy"`             // 创建者
	UpdateTime     core.TsTime `json:"updateTime" structs:"updateTime"`         // 更新时间
	UpdateBy       string      `json:"updateBy" structs:"updateBy"`             // 更新者
	Remark         string      `json:"remark" structs:"remark"`                 // 备注信息
}
