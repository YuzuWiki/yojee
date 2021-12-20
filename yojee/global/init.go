package global

import "sync"

var (
	once  = sync.Once{}
)

func init()  {
	InitLogger()
	InitDB()
}