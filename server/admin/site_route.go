package admin

import (
	"github.com/gin-gonic/gin"
	"x_admin/admin/site"
	"x_admin/middleware"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入SiteRoute注册路由

3. 后台手动添加菜单和按钮
site:add
site:edit
site:del
site:list
site:listAll
site:detail
*/

// SiteRoute(rg)
func SiteRoute(rg *gin.RouterGroup) {
	handle := site.SiteHandler{}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/site/list", handle.List)
	rg.GET("/site/listAll", handle.ListAll)
	rg.GET("/site/detail", handle.Detail)
	rg.POST("/site/add", handle.Add)
	rg.POST("/site/edit", handle.Edit)
	rg.POST("/site/del", handle.Del)
}
