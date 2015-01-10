package gocontrol

func ip(headers map[string]string) string {
	ip := headers["IP"]
	return Render("ip", ip)
}

func browser(headers map[string]string)  string {
	browser := headers["User-Agent"]
	return Render("browser", browser)
}