package ExeContext

/*
	@author sunyuw <leentingOne@gmail.com>
	@date 2013-9-24
	@attention 这暂时还不是一个开源项目,如果您需要使用本项目的代码帮助开发请咨询作者，在取得书面同意后方可使用.
*/
import (
	"net"
	"strconv"
	"strings"
)

const (
	NORMAL    = 0
	INIT_ERR  = 1
	CONN_ROOR = 2
)

type LinkContext struct {
	localAddress string

	localPort int

	connection net.Conn

	remoteAddress string

	remotePort int

	status int
}

func (self LinkContext) Constructor(conn net.Conn, err error) {
	self.SetConnection(conn)
	addressUrl := strings.Split(conn.RemoteAddr().String(), ":")
	address := addressUrl[0]
	port, _ := strconv.Atoi(addressUrl[1])
	self.SetRemoteAddress(address)
	self.SetRemotePort(port)
	localUrl := strings.Split(conn.LocalAddr().String(), ":")
	localAddress := localUrl[0]
	localPort, _ := strconv.Atoi(localUrl[1])
	self.SetLocalAddress(localAddress)
	self.SetLocalPort(localPort)
}

func (self LinkContext) SetStatus(status int) {
	self.status = status
}

func (self LinkContext) SetRemotePort(port int) {
	self.remotePort = port
}

func (self LinkContext) SetRemoteAddress(address string) {
	self.remoteAddress = address
}

func (self LinkContext) SetConnection(conn net.Conn) {
	self.connection = conn
}

func (self LinkContext) SetLocalPort(port int) {
	self.localPort = port
}

func (self LinkContext) SetLocalAddress(address string) {
	self.localAddress = address
}

func (self LinkContext) GetStatus() (status int) {
	return self.status
}

func (self LinkContext) GetRemotePort() (remotePort int) {
	return self.remotePort
}

func (self LinkContext) GetRemoteAddress() (remoteAddress string) {
	return self.remoteAddress
}

func (self LinkContext) GetConnection() (conn net.Conn) {
	return self.connection
}

func (self LinkContext) GetLocalAddress() (localAddress string) {
	return self.localAddress
}

func (self LinkContext) GetLocalPort() (localPort int) {
	return self.localPort
}
