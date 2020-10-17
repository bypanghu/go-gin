package config

import "time"

const (
	AppMode = "release" //debug or release
	AppPort = ":8888"
	AppName = "app"

	// 日志文件
	AppAccessLogName = "log/" + AppName + "-access.log"
	AppErrorLogName  = "log/" + AppName + "-error.log"
	AppGrpcLogName   = "log/" + AppName + "-grpc.log"



	//数据库相关
	DateBaseIp = "127.0.0.1"  //数据库运行ip
	DataBasePort = "3306"		//数据库运行端口
	DataBaseRoot = "root"		//数据库账号
	DataBasePassword = "root"		//数据库密码
	DataBaseTable ="school"			//数据表


	//jwt相关
	TokenExpireDuration = time.Hour * 2
	TokenIssuer = "app" //token签发人

)
