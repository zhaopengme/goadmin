package admin

import (
	"github.com/gin-gonic/gin"
	"x_admin/admin/sys_job"
	"x_admin/middleware"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入SysJobRoute注册路由

3. 后台手动添加菜单和按钮
sys_job:add
sys_job:edit
sys_job:del
sys_job:list
sys_job:listAll
sys_job:detail
*/

// SysJobRoute(rg)
func SysJobRoute(rg *gin.RouterGroup) {
	handle := sys_job.SysJobHandler{}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/sys_job/list", handle.List)
	rg.GET("/sys_job/listAll", handle.ListAll)
	rg.GET("/sys_job/detail", handle.Detail)
	rg.POST("/sys_job/add", handle.Add)
	rg.POST("/sys_job/edit", handle.Edit)
	rg.POST("/sys_job/del", handle.Del)
}
