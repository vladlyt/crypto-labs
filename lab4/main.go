package main

import (
	"crypto/md5"
	cryptoRand "crypto/rand"
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	PASSWORD_GENERATION_COUNT = 200000
	PW_SALT_BYTES             = 16
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type HashFunction func(password string) string

type Hash struct {
	function       HashFunction
	outputFilepath string
}

func GeneratePasswords(outFilepath string, passwordsCount int, generators []PasswordGenerator, hashFunc HashFunction) {
	file, err := os.Create(outFilepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, generator := range generators {
		for i := 0; i < int(generator.GetPasswordCoef()*float64(passwordsCount)); i++ {
			err := writer.Write([]string{hashFunc(generator.GeneratePassword())})
			if err != nil {
				panic(err)
			}
		}
	}
}

func CreateSalt() string {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(cryptoRand.Reader, salt)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(salt)
}

func EncryptMD5(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func EncryptSHA1(password string) string {
	salt := CreateSalt()
	hasher := sha1.New()
	hasher.Write([]byte(password + salt))

	return fmt.Sprintf(
		"%s:%s",
		salt,
		hex.EncodeToString(hasher.Sum(nil)),
	)
}

func EncryptBcrypt(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func main() {
	generators := []PasswordGenerator{
		Common100PasswordGenerator{},
		Common1000000PasswordGenerator{},
		RandomPasswordGenerator{},
		RulePasswordGenerator{rules: []Rule{
			Upper{},
			Lower{},
			Reverse{},
			Capitalize{},
			AddNumbersStartOrEnd{},
		}},
	}

	hashes := []Hash{
		{
			function:       EncryptMD5,
			outputFilepath: "generated-md5.csv",
		},
		{
			function:       EncryptSHA1,
			outputFilepath: "generated-sha1.csv",
		},
		{
			function:       EncryptBcrypt,
			outputFilepath: "generated-bcrypt.csv",
		},
	}

	//// pimpalas
	//hash := "d1025ab01b085246e8f6f3294875c175:3fd7820dcbbca2d0cd5a0f476180c358bd5ef49d"
	//
	//hparts := strings.Split(hash, ":")
	//
	//salt := hparts[0]
	//
	//
	//
	//hasher := sha1.New()
	//hasher.Write([]byte("pimpalas" + salt))
	//
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))

	wg := sync.WaitGroup{}
	for _, hash := range hashes {
		wg.Add(1)
		go func(hash Hash) {
			fmt.Println("LOL")
			GeneratePasswords(hash.outputFilepath, PASSWORD_GENERATION_COUNT, generators, hash.function)
			wg.Done()
			fmt.Println("DONE")
		}(hash)
	}
	wg.Wait()
}
