package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/itchyny/base58-go"
	"github.com/joho/godotenv"
)

var (
	err         = godotenv.Load()
	keyGenerate = os.Getenv("KEY_GENERATE_SHORT_LINK")
)

func GenerateShortLink(initialLink string, userid string) string {
	// create url hashing using func sha256Of with parameter of long link, user_id and secret key generate API
	urlHashBytes := sha256Of(initialLink + userid + keyGenerate)

	// generate number using setBytes with return type of uint64
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// creeate string using func base58Encoding to get the string of encoded
	finalString := base58Encoding([]byte(fmt.Sprintf("%d", generatedNumber)))

	// return value of finalString with slice of 8 character first
	return finalString[:8]
}

func sha256Of(input string) []byte {
	// initial new sha356
	algorithm := sha256.New()

	// write new algorithm using sha256 with byte data of input
	algorithm.Write([]byte(input))

	// return result write sha356 with byte data
	return algorithm.Sum(nil)
}

func base58Encoding(bytes []byte) string {
	// initial encoding using package base58 with hashing BitcoinEncoding
	encoding := base58.BitcoinEncoding

	// create / write encoding with data of bytes
	encoded, err := encoding.Encode(bytes)

	// handle error when encoded the bytes
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// return string value of encoded
	return string(encoded)
}

// checking valid link with www and https://
func CheckingValidLink(link string) string {
	checkWww := strings.Contains(link, "www.")
	checkHttp := strings.Contains(link, "http://")
	checkHttps := strings.Contains(link, "https://")

	if !checkWww && !checkHttp && !checkHttps {
		link = "http://www." + link
	} else if checkWww && !checkHttp && !checkHttps {
		link = "http://" + link
	} else if !checkWww && (!checkHttp || !checkHttps) {
		split := strings.Split(link, "//")
		link = split[0] + "//www." + split[1]
	}

	return link
}
