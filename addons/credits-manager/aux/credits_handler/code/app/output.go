package app

var caller func(text string)

func _print(text string) {
	print(text)
}

func Print(text string) {
	if nil == caller {
		SetCaller(_print)
	}
	caller(text)
}

func SetCaller(ref func(text string)) {
	caller = ref
}
