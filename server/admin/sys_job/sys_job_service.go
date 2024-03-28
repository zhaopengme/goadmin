package sys_job

import (
	"gorm.io/gorm"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
)

type ISysJobService interface {
	List(page request.PageReq, listReq SysJobListReq) (res response.PageResp, e error)
	ListAll() (res []SysJobResp, e error)

	Detail(id int) (res SysJobResp, e error)
	Add(addReq SysJobAddReq) (e error)
	Edit(editReq SysJobEditReq) (e error)
	Del(id int) (e error)
}

var Service = NewSysJobService()

// NewSysJobService 初始化
func NewSysJobService() *sysJobService {
	db := core.GetDB()
	return &sysJobService{db: db}
}

// sysJobService 定时任务调度服务实现类
type sysJobService struct {
	db *gorm.DB
}

// List 定时任务调度列表
func (service sysJobService) List(page request.PageReq, listReq SysJobListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.SysJob{})
	if listReq.Name != "" {
		dbModel = dbModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Status != "" {
		dbModel = dbModel.Where("status = ?", listReq.Status)
	}
	dbModel = dbModel.Where("is_delete = ?", 0)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.SysJob
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []SysJobResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// ListAll 定时任务调度列表
func (service sysJobService) ListAll() (res []SysJobResp, e error) {
	var objs []model.SysJob

	err := service.db.Find(&objs).Error
	if e = response.CheckErr(err, "ListAll Find err"); e != nil {
		return
	}
	response.Copy(&res, objs)
	return res, nil
}

// Detail 定时任务调度详情
func (service sysJobService) Detail(id int) (res SysJobResp, e error) {
	var obj model.SysJob
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

// Add 定时任务调度新增
func (service sysJobService) Add(addReq SysJobAddReq) (e error) {
	var obj model.SysJob
	response.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 定时任务调度编辑
func (service sysJobService) Edit(editReq SysJobEditReq) (e error) {
	var obj model.SysJob
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

// Del 定时任务调度删除
func (service sysJobService) Del(id int) (e error) {
	var obj model.SysJob
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
