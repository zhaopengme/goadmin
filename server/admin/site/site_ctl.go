package site

import (
	"github.com/gin-gonic/gin"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
)

type SiteHandler struct{}

// @Summary	站点列表
// @Tags		site-站点
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		PageNo		query		int					true	"页码"
// @Param		PageSize	query		int					true	"每页数量"
// @Param		name		query		string				false	"名称."
// @Success	200			{object}	[]SiteResp	"成功"
// @Failure	400			{object}	string				"请求错误"
// @Router		/api/admin/site/list [get]
func (hd SiteHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq SiteListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//		@Summary	站点列表-所有
//		@Tags		site-站点
//	 @Produce	json
//		@Success	200			{object}	[]SiteResp	"成功"
//		@Router		/api/admin/site/listAll [get]
func (hd SiteHandler) ListAll(c *gin.Context) {
	var listReq SiteListReq
	res, err := Service.ListAll()
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	站点详情
// @Tags		site-站点
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		id		query		int				false	"主键ID."
// @Success	200			{object}	SiteResp	"成功"
// @Router		/api/admin/site/detail [get]
func (hd SiteHandler) Detail(c *gin.Context) {
	var detailReq SiteDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	站点新增
// @Tags		site-站点
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		name		body		string				false	"名称."
// @Param		url		body		string				false	"站点url."
// @Success	200			{object}	response.RespType	"成功"
// @Router		/api/admin/site/add [post]
func (hd SiteHandler) Add(c *gin.Context) {
	var addReq SiteAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// @Summary	站点编辑
// @Tags		site-站点
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		id		body		int				false	"主键ID."
// @Param		name		body		string				false	"名称."
// @Param		url		body		string				false	"站点url."
// @Success	200			{object}	response.RespType	"成功"
// @Router		/api/admin/site/edit [post]
func (hd SiteHandler) Edit(c *gin.Context) {
	var editReq SiteEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// @Summary	站点删除
// @Tags		site-站点
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		id		body		int				false	"主键ID."
// @Success	200			{object}	response.RespType	"成功"
// @Router		/api/admin/site/del [post]
func (hd SiteHandler) Del(c *gin.Context) {
	var delReq SiteDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.Id))
}
