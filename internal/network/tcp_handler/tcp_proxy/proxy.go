package tcp_proxy

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"go.uber.org/zap"
	"io"
	"net"
)

func EstablishProxy(dstAddresses []string) error {
	logger := log.Logger()
	proxyServer, err := net.Listen("tcp", configs.Config.TcpProxyServer.ProxyAddress)
	if err != nil {
		logger.Error("listen proxyServer failed", zap.Error(err))
		return err
	}

	go serveProxy(proxyServer, dstAddresses)

	return nil
}

func serveProxy(proxyServer net.Listener, dstAddresses []string) {
	logger := log.Logger()
	for {
		conn, err := proxyServer.Accept()
		if err != nil {
			logger.Error("accept proxy conn failed", zap.Error(err))
			return
		}
		logger.Info("get a new connection", zap.String("address", conn.RemoteAddr().String()))

		go func(from net.Conn) {
			defer from.Close()

			for _, dstAddr := range dstAddresses {
				to, err := net.Dial("tcp", dstAddr)
				if err != nil {
					logger.Error("proxy dial failed", zap.Error(err))
					return
				}
				defer to.Close()

				err = proxy(from, to)
				if err != nil && err != io.EOF {
					logger.Error("error while serving proxy", zap.Error(err))
				}
			}
		}(conn)
	}
}

func proxy(from io.Reader, to io.Writer) error {
	fromWriter, fromIsWriter := from.(io.Writer)
	toReader, toIsReader := to.(io.Reader)

	if toIsReader && fromIsWriter {
		go func() { io.Copy(fromWriter, toReader) }()
	}
	_, err := io.Copy(to, from)
	return err
}
