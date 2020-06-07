package main

import (
	"GoTools/mylogger"
)

func main() {

	log := mylogger.NewFileLogger("info", "./", "micor.log", 100*1024)
	for {
		id := 36
		name := "加油，一定要争气！"
		log.Debug("这是一条Debug级别日志,id=%d, name=%s\n", id, name)
		log.Info("这是一条Info级别日志,id=%d, name=%s\n", id, name)

		log.Warning("这是一条Warning级别日志,id=%d, name=%s\n", id, name)
		log.Error("这是一条Error级别日志,id=%d, name=%s\n", id, name)
		log.Fatal("这是一条Fatal级别日志,id=%d, name=%s\n", id, name)
		// time.Sleep(2 * time.Second)
	}

}
