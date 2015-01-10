package gocontrol

func index(headers map[string]string)  string {
	helloString := "Hello World!"
	return Render("index", helloString)
}

func test(headers map[string]string)  string {
	return Render("test")
}