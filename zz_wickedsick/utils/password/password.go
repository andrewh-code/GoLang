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
*/

func GenerateHash(input string) string {

	hash := sha256.New()
	hash.Write([]byte(input))
	var outputHash = base64.URLEncoding.EncodeToString(hash.Sum(nil))

	return outputHash

}

func GenerateSalt() string {

	random := rand.Intn(10000000) // could use a MUCH BETTER rand algo
	salt := strconv.Itoa(random)
	//hash the salt and take the first 10 digits
	hashedSalt := GenerateHash(salt)

	return hashedSalt[0:10]
}

func EncryptPassword(password string) string {

	return GenerateHash(password + GenerateSalt())
}
