package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// encrypt
func RsaEncrypt(data []byte, publicKey string) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	rsaRawData, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		return nil, err
	}
	return rsaRawData, nil
}

// to string
func RsaEncryptToString(data []byte, publicKey string) string {
	rsaRawData, err := RsaEncrypt(data, publicKey)
	if err != nil || rsaRawData == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(rsaRawData)
}

// decrypt
func RsaDecrypt(data []byte, privateKey string) ([]byte, error) {
	//decode
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}

func RsaDecryptFromString(data string, privateKey string) ([]byte, error) {
	// decode base64
	rawData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return RsaDecrypt(rawData, privateKey)
}
