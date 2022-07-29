package main

import (
	"final-project/routers"
)

func main() {
	r := routers.StartApp()

	r.Run((":8001"))
}
