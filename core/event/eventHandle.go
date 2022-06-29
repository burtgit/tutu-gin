package event

import "log"

func EventHandler(eventInterface EventInterface) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("异步任务发现错误", err)
		}
	}()
	eventInterface.PublishEvent()
}
