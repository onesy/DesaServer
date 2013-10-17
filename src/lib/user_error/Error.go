package user_error

/*
	@author sunyuw <leentingOne@gmail.com>
	@date 2013-9-24
	@attention 这暂时还不是一个开源项目,如果您需要使用本项目的代码帮助开发请咨询作者，在取得书面同意后方可使用.
*/
import (
	"log"
)

type callback interface{}

const (
	NORMAL  = 0
	FATAL   = 1
	ERROR   = 2
	WARNING = 3
	NOTICE  = 4
	UNKNOWN = 5
)

func Check_error(err error, level int) {
	if err != nil {
		switch level {
		case FATAL:
			log.Println("Event Class:FATAL!", err)
		case ERROR:
			log.Println("Event Class:ERROR!", err)
		case WARNING:
			log.Println("Event Class:WAUNING!", err)
		case NOTICE:
			log.Println("Event Class:NOTICE!", err)
		default:
			log.Println("Event Class:UNKNOWN!", err)
		}
	}
}

func WipeAss(err error, level int, callback func(...interface{}), Parameters4CallBack ...interface{}) {
	Check_error(err, level)
	callback(Parameters4CallBack)
}
