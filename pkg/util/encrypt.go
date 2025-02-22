package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

var Encrypt *Encryption

//AES 对称加密

// Encryption 加密
type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

//PadPwd 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// AesEncoding 加密
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		return src
	}
	NewSrcByte := PadPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	// base64 编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

//UnPadPwd 去掉填充的部分
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) < 1 {
		return dst, errors.New("长度有误")
	}
	// 去掉的长度
	unpadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:(len(dst) - unpadNum)]
	return str, nil
}

// AesDecoding 解密
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd
	}
	block, errBlock := aes.NewCipher([]byte(k.key))
	if errBlock != nil {
		return pwd
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, err = UnPadPwd(dst) // 填充的要去掉
	if err != nil {
		return "0"
	}
	return string(dst)
}

// SetKey 方法
func (k *Encryption) SetKey(key string) {
	k.key = key
}
