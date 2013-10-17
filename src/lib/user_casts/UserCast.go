package user_casts

/*
	@author sunyuw <leentingOne@gmail.com>
	@date 2013-9-24
	@attention 这暂时还不是一个开源项目,如果您需要使用本项目的代码帮助开发请咨询作者，在取得书面同意后方可使用.
*/
import (
	"bytes"
	"encoding/binary"
)

func bytes2int(bytes2Cst []byte) (inty int64, err error) {
	buff := bytes.NewBuffer(bytes2Cst)
	inty, err = binary.ReadVarint(buff)
	return
}
