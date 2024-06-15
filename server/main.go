package main

import (
	server "github.com/uncleTen276/attendance.git/server/app"
)

// @title Fiber Go API
// @version 1.0
// @description employee attendance
// @contact.name Nguyen Tri
// @contact.email tringuyen2762001@gmail.com
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	server.Serve()
}
