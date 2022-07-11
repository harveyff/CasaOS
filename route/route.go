// 这是一个用来反馈 API 设计的 PR，不要 merge

package route

import (
	"github.com/IceWhaleTech/CasaOS/middleware"
	"github.com/IceWhaleTech/CasaOS/pkg/config"
	jwt2 "github.com/IceWhaleTech/CasaOS/pkg/utils/jwt"
	v1 "github.com/IceWhaleTech/CasaOS/route/v1"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var OnlineDemo bool = false

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.WriteLog())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	gin.SetMode(config.ServerInfo.RunMode)

	r.GET("/v1/debug", v1.GetSystemConfigDebug) // 这里改成 /v1/sys/debug

	// @tiger - 不要把同一个词汇按单词来分割，改成 /v1/sys/socket_port
	r.GET("/v1/sys/socket/port", v1.GetSystemSocketPort)

	// @tiger - （nice-to-have）开源项目应该删除所有注释代码，增加代码整洁性。或者增加注释说明
	//r.POST("/v1/user/refresh/token", v1.PostUserRefreshToken)

	v1Group := r.Group("/v1")

	v1Group.Use(jwt2.JWT())
	{
		v1AppGroup := v1Group.Group("/app")
		v1AppGroup.Use()
		{
			// @tiger - 按照 RESTFul 规范，改成 GET /v1/apps?installed=true
			//获取我的已安装的列表
			v1AppGroup.GET("/my/list", v1.MyAppList)

			// @tiger - 按照 RESTFul 规范，改成 GET /v1/apps/usage
			v1AppGroup.GET("/usage", v1.AppUsageList)

			// @tiger - 按照 RESTFul 规范，改成 GET /v1/app/:id
			//app详情
			v1AppGroup.GET("/appinfo/:id", v1.AppInfo)

			// @tiger - 按照 RESTFul 规范，改成 GET /v1/apps?installed=false
			//获取未安装的列表
			v1AppGroup.GET("/list", v1.AppList)

			// @tiger - 这个信息和应用无关，应该挪到 /v1/sys/port/avaiable
			//获取端口
			v1AppGroup.GET("/port", v1.GetPort)

			// @tiger - RESTFul 路径中尽量不要有动词，同时这个信息和应用无关，应该挪到 /v1/sys/port/:port
			//检查端口
			v1AppGroup.GET("/check/:port", v1.PortCheck)

			// @tiger - 应用分类和应用不是一类资源，应该挪到 GET /v1/app_categories
			v1AppGroup.GET("/category", v1.CategoryList)

			// @tiger - Docker Terminal 和应用不是一类资源，应该挪到 GET /v1/container/:id/terminal
			//          另外这个返回的不是一个 HTTP 响应，应该返回一个 wss://... 的 URL给前端，由前端另行处理
			v1AppGroup.GET("/terminal/:id", v1.DockerTerminal)

			// @tiger - 所有跟 Docker 有关的 API，应该挪到 /v1/container 下
			//app容器详情
			v1AppGroup.GET("/info/:id", v1.ContainerInfo) // 改成 GET /v1/container/:id
			//app容器日志
			v1AppGroup.GET("/logs/:id", v1.ContainerLog) // 改成 GET /v1/container/:id/log
			//暂停或启动容器
			v1AppGroup.PUT("/state/:id", v1.ChangAppState) // 改成 PUT /v1/container/:id/state
			//安装app
			v1AppGroup.POST("/install", v1.InstallApp)
			//卸载app
			v1AppGroup.DELETE("/uninstall/:id", v1.UnInstallApp)
			//获取进度
			v1AppGroup.GET("/state/:id", v1.GetContainerState)
			//更新容器配置
			v1AppGroup.PUT("/update/:id/setting", v1.UpdateSetting)
			//获取可能新数据
			v1AppGroup.GET("/update/:id/info", v1.ContainerUpdateInfo)

			// @tiger - rely -> dependency - 依赖是什么意思？
			v1AppGroup.GET("/rely/:id/info", v1.ContainerRelyInfo)

			// @tiger - 按照 RESTFul 规范，改成 GET /v1/container/:id/config
			v1AppGroup.GET("/install/config", v1.GetDockerInstallConfig)

			v1AppGroup.PUT("/update/:id", v1.PutAppUpdate)
			v1AppGroup.POST("/share", v1.ShareAppFile)
		}

		v1SysGroup := v1Group.Group("/sys")
		v1SysGroup.Use()
		{
			v1SysGroup.GET("/version/check", v1.GetSystemCheckVersion)
			v1SysGroup.GET("/hardware/info", v1.GetSystemHardwareInfo)
			v1SysGroup.POST("/update", v1.SystemUpdate)
			v1SysGroup.GET("/wsssh", v1.WsSsh)
			v1SysGroup.GET("/config", v1.GetSystemConfig)
			//v1SysGroup.POST("/config", v1.PostSetSystemConfig)
			v1SysGroup.GET("/error/logs", v1.GetCasaOSErrorLogs)
			v1SysGroup.GET("/widget/config", v1.GetWidgetConfig)
			v1SysGroup.POST("/widget/config", v1.PostSetWidgetConfig)
			v1SysGroup.GET("/port", v1.GetCasaOSPort)
			v1SysGroup.PUT("/port", v1.PutCasaOSPort)
			v1SysGroup.POST("/stop", v1.PostKillCasaOS)
			v1SysGroup.GET("/utilization", v1.GetSystemUtilization)
			v1SysGroup.PUT("/usb/:status", v1.PutSystemUSBAutoMount)
			v1SysGroup.GET("/usb/status", v1.GetSystemUSBAutoMount)
			v1SysGroup.GET("/cpu", v1.GetSystemCupInfo)
			v1SysGroup.GET("/mem", v1.GetSystemMemInfo)
			v1SysGroup.GET("/disk", v1.GetSystemDiskInfo)
			v1SysGroup.GET("/network", v1.GetSystemNetInfo)
		}
		v1FileGroup := v1Group.Group("/file")
		v1FileGroup.Use()
		{
			v1FileGroup.PUT("/rename", v1.RenamePath)
			v1FileGroup.GET("/read", v1.GetFilerContent)
			v1FileGroup.POST("/upload", v1.PostFileUpload)
			v1FileGroup.GET("/upload", v1.GetFileUpload)
			v1FileGroup.GET("/dirpath", v1.DirPath)
			//create folder
			v1FileGroup.POST("/mkdir", v1.MkdirAll)
			v1FileGroup.POST("/create", v1.PostCreateFile)

			v1FileGroup.GET("/download", v1.GetDownloadFile)
			v1FileGroup.GET("/download/*path", v1.GetDownloadSingleFile)
			v1FileGroup.POST("/operate", v1.PostOperateFileOrDir)
			v1FileGroup.DELETE("/delete", v1.DeleteFile)
			v1FileGroup.PUT("/update", v1.PutFileContent)
			v1FileGroup.GET("/image", v1.GetFileImage)
			v1FileGroup.DELETE("/operate/:id", v1.DeleteOperateFileOrDir)
			//v1FileGroup.GET("/download", v1.UserFileDownloadCommonService)
		}
		v1DiskGroup := v1Group.Group("/disk")
		v1DiskGroup.Use()
		{
			v1DiskGroup.GET("/check", v1.GetDiskCheck)

			v1DiskGroup.GET("/list", v1.GetDiskList)

			//获取磁盘详情
			v1DiskGroup.GET("/info", v1.GetDiskInfo)

			//format storage
			v1DiskGroup.POST("/format", v1.PostDiskFormat)

			// add storage
			v1DiskGroup.POST("/storage", v1.PostDiskAddPartition)

			//mount SATA disk
			v1DiskGroup.POST("/mount", v1.PostMountDisk)

			//umount sata disk
			v1DiskGroup.POST("/umount", v1.PostDiskUmount)

			//获取可以格式化的内容
			v1DiskGroup.GET("/type", v1.FormatDiskType)

			//删除分区
			v1DiskGroup.DELETE("/delpart", v1.RemovePartition)
			v1DiskGroup.GET("/usb", v1.GetUSBList)

		}
		v1Group.GET("/sync/config", v1.GetSyncConfig)
	}
	return r
}
