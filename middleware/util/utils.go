package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var eq = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 13, 05}

// const secretKey string = "abc&1*~#^2^#s0^=)^^7%b34"

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, eq)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encodeBase64(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	cipherText := decodeBase64(text)
	cfb := cipher.NewCFBDecrypter(block, eq)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

type (
	Months struct {
		MonthID int
		Month   string
	}

	Years struct {
		Year int
	}
)

func GetMonths() interface{} {
	listMonths := []Months{
		{
			MonthID: 1,
			Month:   "January",
		},
		{
			MonthID: 2,
			Month:   "February",
		},
		{
			MonthID: 3,
			Month:   "March",
		},
		{
			MonthID: 4,
			Month:   "April",
		},
		{
			MonthID: 5,
			Month:   "May",
		},
		{
			MonthID: 6,
			Month:   "June",
		},
		{
			MonthID: 7,
			Month:   "July",
		},
		{
			MonthID: 8,
			Month:   "August",
		},
		{
			MonthID: 9,
			Month:   "September",
		},
		{
			MonthID: 10,
			Month:   "October",
		},
		{
			MonthID: 11,
			Month:   "November",
		},
		{
			MonthID: 12,
			Month:   "December",
		},
	}

	return listMonths
}

func GetYears() [100]Years {
	currentYear := time.Now().Year()
	years := [100]Years{}

	for ctr := 0; ctr < 100; ctr++ {
		years[ctr].Year = currentYear - ctr
	}
	return years
}

func GetGreetings() string {
	currentTime := time.Now()
	hour := currentTime.Hour()

	var greeting string

	if hour >= 0 && hour < 12 {
		greeting = "Good morning,"
	} else if hour >= 12 && hour < 17 {
		greeting = "Good afternoon,"
	} else {
		greeting = "Good evening,"
	}
	return greeting

}
