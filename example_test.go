package beeep

func ExampleBeep() {
	Beep(DefaultFreq, DefaultDuration)
}

func ExampleNotify() {
	Notify("Title", "MessageBody", "assets/icon128.png")
}

func ExampleAlert() {
	Alert("Title", "MessageBody", "assets/icon128.png")
}
