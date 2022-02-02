package event

var handlers = make(map[string]func(args ...string))

func Register(event string, handler func(args ...string)) {
	handlers[event] = handler
}

func Emit(event string, args ...string) {
	handlers[event](args...)
}
