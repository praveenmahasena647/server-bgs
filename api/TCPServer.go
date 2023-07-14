package api

import (
	"net"

	"github.com/praveenmahasena647/bg-server/internal/helpers"
)

func RunServer() error {
	var TCPServer, TCPErr = net.Listen("tcp", ":42069")

	if TCPErr != nil {
		return TCPErr
	}
	defer TCPServer.Close()

	for {
		var con, conErr = TCPServer.Accept()

		if conErr != nil {
			continue
		}

		var byteBuffer, err = helpers.GetImages()

		if err != nil {
			return err
		}
		con.Write(byteBuffer)
		con.Close()
	}
	return nil
}
