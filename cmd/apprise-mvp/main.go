package main

import "apprise-mvp/internal/app"

// @title           Apprise MVP API
// @version         1.0
// @description     REST API for sending emails via Apprise in stateless mode.
// @BasePath        /
// @schemes         http
func main() {
	app.App()
}
