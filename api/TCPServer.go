package api

import "net"

func StartServer() error {
	var TCPServer, TCPErr = net.Listen("tcp", ":49069")

	if TCPErr != nil {
		return TCPErr
	}

	defer TCPServer.Close()

	for {
		var con, conErr = TCPServer.Accept()
		if conErr != nil {
			continue
		}
		var data, dataErr = getImages()
		if dataErr != nil {
			con.Write([]byte("server Error"))
		}
		con.Write(data)
		con.Close()
	}

	return nil
}
