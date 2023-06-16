package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func Md5ByString(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// Md5BySalt 先写盐值
func Md5BySalt(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// EncryptPassword 将密码加密，需要传入密码返回的是加密后的密码
func EncryptPassword(password string) (string, error) {
	// 加密密码，使用 bcrypt 包当中的 GenerateFromPassword 方法，bcrypt.DefaultCost 代表使用默认加密成本
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 如果有错误则返回异常，加密后的空字符串返回为空字符串，因为加密失败
		return "", err
	} else {
		// 返回加密后的密码和空异常
		return string(encryptPassword), nil
	}
}

// EqualsPassword 对比密码是否正确
func EqualsPassword(password, encryptPassword string) bool {
	// 使用 bcrypt 当中的 CompareHashAndPassword 对比密码是否正确，第一个参数为加密后的密码，第二个参数为未加密的密码
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	// 对比密码是否正确会返回一个异常，按照官方的说法是只要异常是 nil 就证明密码正确
	return err == nil
}
