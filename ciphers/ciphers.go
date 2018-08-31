package ciphers

import (
	"github.com/sipt/shuttle"
	"github.com/sipt/shuttle/ciphers/ssstream"
	"fmt"
)

type ConnDecorate func(password string, conn shuttle.IConn) (shuttle.IConn, error)

//加密装饰
func CipherDecorate(password, method string, conn shuttle.IConn) (shuttle.IConn, error) {
	d := ssstream.GetStreamChiphers(method)
	if d != nil {
		return d(password, conn)
	}
	return nil, fmt.Errorf("[SS Cipher] not support : %s", method)
}
