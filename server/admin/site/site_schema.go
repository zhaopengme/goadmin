package site

import "x_admin/core"

// SiteListReq 站点列表参数
type SiteListReq struct {
	Name string `form:"name"` // 名称
}

// SiteDetailReq 站点详情参数
type SiteDetailReq struct {
	Id int `form:"id"` // 主键ID
}

// SiteAddReq 站点新增参数
type SiteAddReq struct {
	Name string `form:"name"` // 名称
	Url  string `form:"url"`  // 站点url
}

// SiteEditReq 站点编辑参数
type SiteEditReq struct {
	Id   int    `form:"id"`   // 主键ID
	Name string `form:"name"` // 名称
	Url  string `form:"url"`  // 站点url
}

// SiteDelReq 站点新增参数
type SiteDelReq struct {
	Id int `form:"id"` // 主键ID
}

// SiteResp 站点返回信息
type SiteResp struct {
	Id         int         `json:"id" structs:"id"`                 // 主键ID
	Name       string      `json:"name" structs:"name"`             // 名称
	Url        string      `json:"url" structs:"url"`               // 站点url
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
