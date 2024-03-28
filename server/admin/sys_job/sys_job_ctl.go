package sys_job

import (
	"github.com/gin-gonic/gin"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
)

type SysJobHandler struct{}

// @Summary	定时任务调度列表
// @Tags		sys_job-定时任务调度
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		PageNo		query		int					true	"页码"
// @Param		PageSize	query		int					true	"每页数量"
// @Param		name		query		string				false	"调度名称."
// @Param		status		query		string				false	"状态（0正常 1暂停）."
// @Success	200			{object}	[]SysJobResp	"成功"
// @Failure	400			{object}	string				"请求错误"
// @Router		/api/admin/sys_job/list [get]
func (hd SysJobHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq SysJobListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//		@Summary	定时任务调度列表-所有
//		@Tags		sys_job-定时任务调度
//	 @Produce	json
//		@Success	200			{object}	[]SysJobResp	"成功"
//		@Router		/api/admin/sys_job/listAll [get]
func (hd SysJobHandler) ListAll(c *gin.Context) {
	res, err := Service.ListAll()
	var listReq SysJobListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	定时任务调度详情
// @Tags		sys_job-定时任务调度
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		id		query		int				false	"主键ID."
// @Success	200			{object}	SysJobResp	"成功"
// @Router		/api/admin/sys_job/detail [get]
func (hd SysJobHandler) Detail(c *gin.Context) {
	var detailReq SysJobDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	定时任务调度新增
// @Tags		sys_job-定时任务调度
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		name		body		string				false	"调度名称."
// @Param		invokeTarget		body		string				false	"调用目标字符串."
// @Param		cronExpression		body		string				false	"cron执行表达式."
// @Param		concurrent		body		string				false	"是否并发执行（0允许 1禁止）."
// @Param		status		body		string				false	"状态（0正常 1暂停）."
// @Param		remark		body		string				false	"备注信息."
// @Success	200			{object}	response.RespType	"成功"
// @Router		/api/admin/sys_job/add [post]
func (hd SysJobHandler) Add(c *gin.Context) {
	var addReq SysJobAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	e := Service.Add(addReq)
	if e == nil {
		AddJob(addReq.Name, addReq.CronExpression, addReq.InvokeTarget)
	}
	response.CheckAndResp(c, e)
}

// @Summary	定时任务调度编辑
// @Tags		sys_job-定时任务调度
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		id		body		int				false	"主键ID."
// @Param		name		body		string				false	"调度名称."
// @Param		invokeTarget		body		string				false	"调用目标字符串."
// @Param		cronExpression		body		string				false	"cron执行表达式."
// @Param		concurrent		body		string				false	"是否并发执行（0允许 1禁止）."
// @Param		status		body		string				false	"状态（0正常 1暂停）."
// @Param		remark		body		string				false	"备注信息."
// @Success	200			{object}	response.RespType	"成功"
// @Router		/api/admin/sys_job/edit [post]
func (hd SysJobHandler) Edit(c *gin.Context) {
	var editReq SysJobEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	e := Service.Edit(editReq)
	if e == nil {
		if editReq.Status == "0" {
			AddJob(editReq.Name, editReq.CronExpression, editReq.InvokeTarget)
		} else {
			DelJob(editReq.Name)
		}
	}
	response.CheckAndResp(c, e)
}

// @Summary	定时任务调度删除
// @Tags		sys_job-定时任务调度
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		id		body		int				false	"主键ID."
// @Success	200			{object}	response.RespType	"成功"
// @Router		/api/admin/sys_job/del [post]
func (hd SysJobHandler) Del(c *gin.Context) {
	var delReq SysJobDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	detail, _ := Service.Detail(delReq.Id)
	DelJob(detail.Name)
	e := Service.Del(delReq.Id)
	response.CheckAndResp(c, e)
}
