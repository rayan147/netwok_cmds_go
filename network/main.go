package main

import (
	"fmt"
	 "os"
)

func main() {
	file,_ := os.Stat("network.go")
	fl := fmt.Println
	fl("File name:",file.Name())
	fl("File in bytes:",file.Size())
	fl("File name:",file.ModTime())
	fl("File name:",file.IsDir())
	
	ff := fmt.Printf
	ff("Perminssion 9-bit:%s\n",file.Mode())
	ff("Perminssion 3-digits:%o\n",file.Mode())
	ff("Perminssion 4-digits:%04o\n",file.Mode())
	

}