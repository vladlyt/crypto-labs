package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

const (
	MIN_RANDOM_NUMBER_LEN = 5
	MAX_RANDOM_NUMBER_LEN = 14
	LOWERCASE             = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE             = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS               = "1234567890"
	SYMBOLS               = "!@#$%&"
	PASSWORD_CHARS        = LOWERCASE + UPPERCASE + NUMBERS + SYMBOLS
)

var (
	TOP100     = make([]string, 0)
	TOP1000000 = make([]string, 0)
)

func init() {
	TOP100 = LoadPasswords("top-100-passwords.txt")
	TOP1000000 = LoadPasswords("top-1000000-passwords.txt")
}

func LoadPasswords(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	passwords := make([]string, 0)

	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}
	return passwords
}

type PasswordGenerator interface {
	GeneratePassword() string
	GetPasswordCoef() float64
}

type Common100PasswordGenerator struct{}

func (Common100PasswordGenerator) GeneratePassword() string {
	return TOP100[rand.Int()%len(TOP100)]
}

func (Common100PasswordGenerator) GetPasswordCoef() float64 {
	return 0.05
}

type Common1000000PasswordGenerator struct{}

func (Common1000000PasswordGenerator) GeneratePassword() string {
	return TOP1000000[rand.Int()%len(TOP1000000)]
}

func (Common1000000PasswordGenerator) GetPasswordCoef() float64 {
	return 0.8
}

type RandomPasswordGenerator struct{}

func (RandomPasswordGenerator) GeneratePassword() string {
	password := make([]string, rand.Intn(MAX_RANDOM_NUMBER_LEN-MIN_RANDOM_NUMBER_LEN)+MIN_RANDOM_NUMBER_LEN)
	for i := 0; i < len(password); i++ {
		password[i] = string(PASSWORD_CHARS[rand.Intn(len(PASSWORD_CHARS))])
	}
	return strings.Join(password, "")
}

func (RandomPasswordGenerator) GetPasswordCoef() float64 {
	return 0.05
}

type Rule interface {
	Generate(password string) string
}

// Rules
type Reverse struct{}

func (Reverse) Generate(password string) string {
	reversed := ""
	for _, c := range password {
		reversed = string(c) + reversed
	}
	return reversed
}

type Upper struct{}

func (Upper) Generate(password string) string {
	return strings.ToUpper(password)
}

type Lower struct{}

func (Lower) Generate(password string) string {
	return strings.ToLower(password)
}

type AddNumbersStartOrEnd struct{}

func (AddNumbersStartOrEnd) Generate(password string) string {
	isAddToEnd := rand.Int()%2 == 0
	result := ""
	if isAddToEnd {
		result += password
	}

	for i := 0; i < rand.Intn(5)+1; i++ {
		result += string(NUMBERS[rand.Int()%len(NUMBERS)])
	}

	if !isAddToEnd {
		result += password
	}
	return result

}

type Capitalize struct{}

func (Capitalize) Generate(password string) string {
	return string(unicode.ToUpper(rune(password[0]))) + password[1:]
}

type RulePasswordGenerator struct {
	rules []Rule
}

func (c RulePasswordGenerator) GeneratePassword() string {
	password := TOP1000000[rand.Int()%len(TOP1000000)]
	rule := c.rules[rand.Int()%len(c.rules)]
	return rule.Generate(password)
}

func (c RulePasswordGenerator) GetPasswordCoef() float64 {
	return 0.1
}
