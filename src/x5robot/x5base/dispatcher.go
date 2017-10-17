package x5base

type MsgHandler func([]interface{})

type EventMsg struct {
	msgHandler    MsgHandler
}

func Register(id interface{}, f interface{}) {
	switch f.(type) {
	case func(interface{},interface{}):
		default:
		panic("不支持的注册类型函数，函数因为2个参数，并且第一个可以是Robot,第二个消息回复结构")
	}
}