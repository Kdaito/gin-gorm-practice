package main

import (
	"example/api-template/db"
	"example/api-template/server"
)

func main() {
	db.Init()

	server.Init()
}

// 参考
// https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4

//
// https://qiita.com/knsh14/items/8b73b31822c109d4c497#receiver-type
