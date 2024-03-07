package main

import(
    "E-COMMERS/controls" // Adjust the import path based on your project structure
    "E-COMMERS/middlereware"
	"E-COMMERS/routes"
)

// Declare your variables here
var (
    initializer string
    config      string
)
func init() {
	initializer.LoadEnv()
	var err error
	config.DB, err = config.DBconnect()
	if err != nil {
		panic(err)
	}

	R.LoadHTMLGlob("templates/*.html")
}

var R = gin.Default()

func main() {

	routes.AdminRouts(R)
	routes.UserRouts(R)

	R.Run()
}
