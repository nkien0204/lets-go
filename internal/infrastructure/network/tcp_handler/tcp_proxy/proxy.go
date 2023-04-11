package tcp_proxy

import (
	"io"
	"net"

	"github.com/nkien0204/lets-go/internal/infrastructure/configs"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

func EstablishProxy(dstAddress string) error {
	logger := rolling.New()
	proxyServer, err := net.Listen("tcp", configs.GetConfigs().TcpProxyServer.ProxyAddress)
	if err != nil {
		logger.Error("listen proxyServer failed", zap.Error(err))
		return err
	}

	go serveProxy(proxyServer, dstAddress)

	return nil
}

func serveProxy(proxyServer net.Listener, dstAddress string) {
	logger := rolling.New()
	for {
		conn, err := proxyServer.Accept()
		if err != nil {
			logger.Error("accept proxy conn failed", zap.Error(err))
			return
		}
		logger.Info("get a new connection", zap.String("address", conn.RemoteAddr().String()))

		go func(from net.Conn) {
			defer from.Close()

			to, err := net.Dial("tcp", dstAddress)
			if err != nil {
				logger.Error("proxy dial failed", zap.Error(err))
				return
			}
			defer to.Close()

			err = proxy(from, to)
			if err != nil && err != io.EOF {
				logger.Error("error while serving proxy", zap.Error(err))
			}
		}(conn)
	}
}

func proxy(from io.Reader, to io.Writer) error {
	fromWriter, fromIsWriter := from.(io.Writer)
	toReader, toIsReader := to.(io.Reader)

	if toIsReader && fromIsWriter {
		go func() {
			io.Copy(fromWriter, toReader)
		}()
	}
	_, err := io.Copy(to, from)
	return err
}
