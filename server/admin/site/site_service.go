package site

import (
	"gorm.io/gorm"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
)

type ISiteService interface {
	List(page request.PageReq, listReq SiteListReq) (res response.PageResp, e error)
	ListAll() (res []SiteResp, e error)

	Detail(id int) (res SiteResp, e error)
	Add(addReq SiteAddReq) (e error)
	Edit(editReq SiteEditReq) (e error)
	Del(id int) (e error)
}

var Service = NewSiteService()

// NewSiteService 初始化
func NewSiteService() *siteService {
	db := core.GetDB()
	return &siteService{db: db}
}

// siteService 站点服务实现类
type siteService struct {
	db *gorm.DB
}

// List 站点列表
func (service siteService) List(page request.PageReq, listReq SiteListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.Site{})
	if listReq.Name != "" {
		dbModel = dbModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	dbModel = dbModel.Where("is_delete = ?", 0)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.Site
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []SiteResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// ListAll 站点列表
func (service siteService) ListAll() (res []SiteResp, e error) {
	var objs []model.Site

	err := service.db.Find(&objs).Error
	if e = response.CheckErr(err, "ListAll Find err"); e != nil {
		return
	}
	response.Copy(&res, objs)
	return res, nil
}

// Detail 站点详情
func (service siteService) Detail(id int) (res SiteResp, e error) {
	var obj model.Site
	err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, obj)
	return
}

// Add 站点新增
func (service siteService) Add(addReq SiteAddReq) (e error) {
	var obj model.Site
	response.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 站点编辑
func (service siteService) Edit(editReq SiteEditReq) (e error) {
	var obj model.Site
	err := service.db.Where("id = ? AND is_delete = ?", editReq.Id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = service.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

// Del 站点删除
func (service siteService) Del(id int) (e error) {
	var obj model.Site
	err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 删除
	obj.IsDelete = 1
	err = service.db.Save(&obj).Error
	e = response.CheckErr(err, "Del Save err")
	return
}
