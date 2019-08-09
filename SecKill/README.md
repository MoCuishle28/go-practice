秒杀系统练习

Note：
	conf文件下的 app.conf 配置文件会被 beego 自动读取，但是beego的启动程序 main.exe 要和conf文件在同一目录下

	app.conf 文件中 "${ProRunMode||dev}" 是自动选择开发环境配置 [dev] 下的； 还是正式环境下配置 [prod] 下的

	编译：
		I:\Go_WorkSpace\src\Go-practice\SecKill\SecProxy>go build ./main