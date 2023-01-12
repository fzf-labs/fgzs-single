// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	dashboard "fgzs-single/internal/app/admin/internal/handler/dashboard"
	file "fgzs-single/internal/app/admin/internal/handler/file"
	sysadmin "fgzs-single/internal/app/admin/internal/handler/sys/admin"
	sysdept "fgzs-single/internal/app/admin/internal/handler/sys/dept"
	sysjob "fgzs-single/internal/app/admin/internal/handler/sys/job"
	syslog "fgzs-single/internal/app/admin/internal/handler/sys/log"
	sysmanage "fgzs-single/internal/app/admin/internal/handler/sys/manage"
	syspermmenu "fgzs-single/internal/app/admin/internal/handler/sys/permmenu"
	sysrole "fgzs-single/internal/app/admin/internal/handler/sys/role"
	system "fgzs-single/internal/app/admin/internal/handler/system"
	"fgzs-single/internal/app/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: system.PingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/stat",
					Handler: system.StatHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/system"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/speech",
					Handler: dashboard.DashboardSpeechHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/dashboard"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: sysdept.SysDeptListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: sysdept.SysDeptInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/store",
					Handler: sysdept.SysDeptStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/del",
					Handler: sysdept.SysDeptDelHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/sys/dept"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: sysjob.SysJobListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: sysjob.SysJobInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/store",
					Handler: sysjob.SysJobStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/del",
					Handler: sysjob.SysJobDelHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/sys/job"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: syspermmenu.SysPermMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: syspermmenu.SysPermMenuInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/store",
					Handler: syspermmenu.SysPermMenuStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/del",
					Handler: syspermmenu.SysPermMenuDelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/status",
					Handler: syspermmenu.SysPermMenuStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/sys/permmenu"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: sysrole.SysRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: sysrole.SysRoleInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/store",
					Handler: sysrole.SysRoleStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/del",
					Handler: sysrole.SysRoleDelHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/sys/role"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/login/captcha",
				Handler: sysadmin.SysAdminLoginCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: sysadmin.SysAdminLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logout",
				Handler: sysadmin.SysAdminLogoutHandler(serverCtx),
			},
		},
		rest.WithPrefix("/sys/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: sysadmin.SysAdminInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info/update",
					Handler: sysadmin.SysAdminInfoUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/permmenu",
					Handler: sysadmin.SysAdminPermMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/avatar/generate",
					Handler: sysadmin.SysAdminGenerateAvatarHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/sys/admin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: sysmanage.SysManageListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: sysmanage.SysManageInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/store",
					Handler: sysmanage.SysManageStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/del",
					Handler: sysmanage.SysManageDelHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/sys/manage"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/ownlist",
					Handler: syslog.SysLogOwnListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: syslog.SysLogListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: syslog.SysLogInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/sys/log"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/static/:p1/:p2/:p3",
				Handler: file.FileStaticHandler(serverCtx),
			},
		},
		rest.WithPrefix("/file"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware, serverCtx.AuthMiddleware, serverCtx.SysLogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: file.FileUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: file.FileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/info",
					Handler: file.FileInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/del",
					Handler: file.FileDelHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/file"),
	)
}
