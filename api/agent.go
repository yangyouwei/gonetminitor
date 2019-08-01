package main

import (
"fmt"
"github.com/shirou/gopsutil/net"
)

func collet() {
	nv, _ := net.IOCounters(true)
	for _,i:= range nv {
		fmt.Printf("        Network %v: %v bytes / %v bytes\n",i.Name, i.BytesRecv, i.BytesSent)
	}

}

func main()  {
	collet()
}
