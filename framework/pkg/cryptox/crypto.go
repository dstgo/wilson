package cryptox

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func EncodePwd(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func CompareHashPwd(p1, p2 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)) == nil
}

func Sha256Hex(in []byte) string {
	sum256 := sha256.Sum256(in)
	return hex.EncodeToString(sum256[:])
}

func Sha256HexUpper(in []byte) string {
	return strings.ToUpper(Sha256Hex(in))
}
