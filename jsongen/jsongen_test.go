package jsongen

import (
	"testing"
)

func TestGenClazz(t *testing.T) {
	//genClazz()
}
func TestGenRootClass(t *testing.T) {
	userJson := "{\"username\":\"system\",\"password\":\"123456\"}"
	genRootClass(userJson, "com.qb.monetization.demo", "Json")
}
