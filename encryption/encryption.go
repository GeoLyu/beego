package encryption

import(
	"fmt"
	"golang.org/x/crypto/scrypt"
	"crypto/sha512"
	"crypto/rand"
	"encoding/base64"
)

func encrypt(pwd string, keyLen int) (string, string) {
	salt := salt(keyLen)
	encryption := fmt.Sprintf("%x", sha512.Sum512([]byte(pwd)))
	val, _ := scrypt.Key([]byte(encryption),[]byte(salt), 16384, 8, 1, keyLen)
	encrypted := fmt.Sprintf("%x",val)
	return encrypted, salt
}

func salt(size int) string {
	randLen := make([]byte, size)
	_, err := rand.Read(randLen)
	if err != nil {
		fmt.Println(err)
	}

	result := base64.URLEncoding.EncodeToString(randLen)
	return result
}

func errorHandler() {

}