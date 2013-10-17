package ExeContext

/*
	@author sunyuw <leentingOne@gmail.com>
	@date 2013-9-24
	@attention 这暂时还不是一个开源项目,如果您需要使用本项目的代码帮助开发请咨询作者，在取得书面同意后方可使用.
*/
import (
	"lib/user_cast"
	"lib/user_error"
	"net"
)

const (
	PRO_NOMAL    = 0
	PRO_INIT_ERR = 1
	PRO_ANAL_ERR = 2
	PRO_UNKNOWN  = 3

	MAGIC = 0xf0f0

	BUFFER_SIZE = 512
)

/*
	@author sunyuw <leentingOne@gmail.com>
	@date 2013-9-24

	第一层协议格式描述.
	byte[0:4] 长4byte的magic段.
	byte[4:7] dummy.
	byte[7:10] 通信协议版本号,版本为3部分,通信协议主版本号byte[0],通信协议次版本号byte[1],通信协议子版本号byte[2].
	byte[10:13] 客户端版本号,版本为3部分,客户端主版本号byte[3],客户端次版本号[4],客户端子版本号[5].
	byte[13:16] 客户端对应服务器的版本号,格式如上.byte[6],byte[7],byte[8].
	byte[16:20] 用户名加密串长度UserNameLen.
	byte[20:24] 密码加密串长度PasswdLen.
	byte[24:28] 凭证长度CertificateLen.
	byte[28:36] 内容长度ContentLen.
	byte[36:36 + UserNameLen] 用户名加密信息.
	byte[36+UserNameLen:36+UserNameLen + PasswdLen] 密码加密内容.
	byte[36+ UserNameLen + PasswdLen:36+UserNameLen + PasswdLen + CertificateLen] 凭证内容.
	byte[36+ UserNameLen + PasswdLen:36+UserNameLen + PasswdLen + CertificateLen + 4] CRC32
	设:baseInfo = 36+UserNameLen + PasswdLen + CertificateLen;
	byte[baseInfo:baseInfo + ContentLen] 内容.此内容交由二层协议处理.
*/
type ProtocalContext struct {
	Magic int64

	MsgLength int64

	err error

	ByteBuffer []byte

	ByteContent []byte

	Link LinkContext

	ProtocalMainVersion int64

	ProtocalSecondaryVersion int64

	ProtocalChildVersion int64

	RemoteMainVersion int64

	RemoteSecondaryVersion int64

	RemoteChildVersion int64

	LocalMainVersion int64

	LocalSecondaryVersion int64

	LocalChildVersion int64

	UserNameLen int64

	PasswdLen int64

	CertifyLen int64

	ContenLength int64

	UserNameCrypt []byte

	PasswdCrypt []byte

	CertificationCrypt []byte

	CRC32 []byte
}

func (self ProtocalContext) Constructor(conn net.Conn) {
	self.ByteBuffer = make([]byte, BUFFER_SIZE)
	length, err := conn.Read(self.ByteBuffer[0:])
	if err != nil {
		user_error.Check_error(err, user_error.FATAL)
	}
	self.err = err
	self.MsgLength = length
	self.ByteContent = self.ByteBuffer[0:self.MsgLength]
	// Magic Code
	self.Magic = self.ByteBuffer[0:4]
	// Protocal Version
	self.ProtocalMainVersion = self.ByteBuffer[7:8]
	self.ProtocalSecondaryVersion = self.ByteBuffer[8:9]
	self.ProtocalChildVersion = self.ByteBuffer[9:10]
	// Remote Version
	self.RemoteMainVersion = self.ByteBuffer[10:11]
	self.RemoteSecondaryVersion = self.ByteBuffer[11:12]
	self.RemoteChildVersion = self.ByteBuffer[12:13]
	// Local Version
	self.LocalMainVersion = self.ByteBuffer[13:14]
	self.LocalSecondaryVersion = self.ByteBuffer[14:15]
	self.LocalChildVersion = self.ByteBuffer[15:16]
	// UserNameLength
	self.UserNameLen = self.ByteBuffer[16:20]
	// PasswdLen
	self.PasswdLen = self.ByteBuffer[20:24]
	// CertifyLen
	self.CertifyLen = self.ByteBuffer[24:28]
	// ContentLength
	self.ContenLength = self.ByteBuffer[28:32]
	// UserName

	// Passwd
}
