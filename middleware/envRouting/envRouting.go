// Package envRouting that provides environment variable access with static and secure bindings/configuration
package envRouting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// Port ...
	Port string
	// SSL
	SSL string
	// Environment
	Environment string
	// SSLCertificate ...
	SSLCertificate string
	// SSLKey ...
	SSLKey string
	// SSLSIgning
	SSLSigning string
	// SSLSigningKey
	SSLSigningKey string
	// RBI PUBKEY
	RBI_PublicKey string
	// Mastercard PUBKEY
	MC_PublicKey string
	// Mastercard Signing
	MC_SSLSigning string
	// For EncryptionDecryption
	SecretKey string
	// POSTGRESQL CONFIG
	PostgreSQLHost    string
	PostgresPort      string
	PostgresSSLMode   string
	PostgresTimeZone  string
	DatabaseName      string
	PostgreslUsername string
	PostgreslPassword string
	//iGATE
	IGateBaseUrl string
)

// LoadEnv Staticly load environment variables
func LoadEnv() {
	Port = getEnv("PORT")
	Environment = getEnv("ENVIRONMENT")
	SSL = getEnv("SSL")
	SSLCertificate = getEnv("SSL_CERTIFICATE")
	SSLKey = getEnv("SSL_KEY")
	SSLSigning = getEnv("SSL_SIGNING")
	SSLSigningKey = getEnv("SSL_SIGNING_KEY")
	RBI_PublicKey = getEnv("RBI_PUBKEY")
	MC_PublicKey = getEnv("MC_PUBKEY")
	MC_SSLSigning = getEnv("MC_SSL_SIGNING")
	SecretKey = getEnv("SECRET_KEY")
	PostgreSQLHost = getEnv("POSTGRES_HOST")
	PostgresPort = getEnv("POSTGRES_PORT")
	PostgresSSLMode = getEnv("POSTGRES_SSL_MODE")
	PostgresTimeZone = getEnv("POSTGRES_TIMEZONE")
	DatabaseName = getEnv("DATABASE_NAME")
	PostgreslUsername = getEnv("POSTGRES_USERNAME")
	PostgreslPassword = getEnv("POSTGRES_PASSWORD")
	IGateBaseUrl = getEnv("IGATE_BASE_URL")
}

func getEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
