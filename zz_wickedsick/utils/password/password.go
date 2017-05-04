package password

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"strconv"
)

// salting passwords
/*
process of securing password hasshes from a Rainbow Table attrack
- non-salted passwords don't have a property that is unique to themselves
- rainbow table attack is possible because output of a hash function is same as the input
- use a unique "salt" and combine it with the password
Hash(password + salt)
https://crackstation.net/hashing-security.htm
*/

func GenerateHash(input string) string {

	hash := sha256.New()
	hash.Write([]byte(input))
	var outputHash = base64.URLEncoding.EncodeToString(hash.Sum(nil))

	return outputHash

}

func GenerateSalt() string {

	random := rand.Intn(10000000) + rand.Intn(10000000) // could use a MUCH BETTER rand algo
	salt := strconv.Itoa(random)
	hashedSalt := GenerateHash(salt)

	// we COULD take the first x amount of digits but no
	return hashedSalt
}

func EncryptPassword(password string, salt string) string {

	return GenerateHash(password + salt)
}

// func EncryptPassword2(password string) string {

// 	return GenerateHash(password + GenerateSalt())
// }

// when user logs in
func ValidatePassword(password string) string {
	//database query to retrieve the user's name, password and salt
	return "hello"
}
