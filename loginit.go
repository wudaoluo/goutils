package goutils

import (
	"github.com/wudaoluo/go-logger/logger"
)



func LogInit(debug bool,logdir, logfile string,nums int64) {
	if debug {
		logger.SetConsole(true)        //指定是否控制台打印，默认为true
		logger.SetLevel(logger.DEBUG)  //ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF
	}else {
		logger.SetConsole(false)        //指定是否控制台打印，默认为true
		logger.SetLevel(logger.WARN)  //ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF
	}


	//指定日志文件备份方式为文件大小的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	//第三个参数为备份文件最大数量
	//第四个参数为备份文件大小
	//第五个参数为文件大小的单位 KB，MB，GB TB
	//logger.SetRollingFile(utils.Config.Logdir, "proxy.log", 7, 10, logger.KB)

	//指定日志文件备份方式为日期的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	logger.SetRollingDaily(logdir, logfile,nums)

}
