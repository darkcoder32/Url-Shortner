/*
Copyright © 2022 Sumit Kumar <sk424527@gmail.com>

*/
package main

import (
	"gin/shorti/cmd"
	"gin/shorti/myconfig"
)

func main() {
	myconfig.Init()
	cmd.Execute()
}
