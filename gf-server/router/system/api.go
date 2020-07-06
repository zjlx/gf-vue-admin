package system

import (
	v1 "gf-server/app/api/v1"
	"gf-server/global"
	"gf-server/middleware"
)

// InitApiRouter 注册功能api路由
func InitApiRouter() {
	// TODO 缺少CasbinHandler中间件
	ApiRouter := global.GFVA_SERVER.Group("api").Middleware(middleware.JwtAuth)
	{
		ApiRouter.POST("createApi", v1.CreateApi)  // 创建Api
		ApiRouter.POST("deleteApi", v1.DeleteApi)  // 删除Api
		ApiRouter.POST("getApiList", v1.GetApiList) // 获取Api列表
		ApiRouter.POST("getApiById", v1.GetApiById) // 获取单条Api消息
		ApiRouter.POST("updateApi", v1.UpdateApi)  // 更新api
		ApiRouter.POST("getAllApis", v1.GetAllApis) // 获取所有api
	}
}
