package main

import "net"

func runServer() error {
	var conn, connErrconnErr = net.Listen("tcp", ":42069")

	if connErrconnErr != nil {
		return connErrconnErr
	}

	defer conn.Close()

	for {
		var con, conErr = conn.Accept()
		if conErr != nil {
			continue
		}
		var data, dataErr = getImages()
		if dataErr != nil {
			con.Write([]byte("couldnt do it"))
		}
		con.Write(data)
		con.Close()
	}
}
