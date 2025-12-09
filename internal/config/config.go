package config

import (
	"os"
)

type Config struct {
	ListenAddr string
	MgmtDBUrl  string
	JWTSecret  string
	//AuthDBURL  string
	//MnmBaseURL string
	//AccessTTL  time.Duration
	//RefreshTTL time.Duration

}

func loadConfig() Config {
	//ttl := func(key string, defaultValue time.Duration) time.Duration {
	//	if value := os.Getenv(key); value != "" {
	//		d, err := time.ParseDuration(value)
	//		if err != nil {
	//			log.Fatalf("Error parsing environment variable `%s`: %v", key, err)
	//		} else {
	//			return d
	//		}
	//	}
	//	return defaultValue
	//}
	return Config{
		ListenAddr: getEnv("LISTEN_ADDR", ":8080"),
		MgmtDBUrl:  getEnv("MGMT_DB_URL", "postgres://mgmt_user:mgmt_pass@db-mgmt:5432/mgmt_db?sslmode=disable"),
		//AuthDBURL:  getEnv("AUTH_URL", "http://localhost:8080"),
		//MnmBaseURL: getEnv("MNM_URL", "http://localhost:8080"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		//AccessTTL:  ttl("ACCESS_TTL", 24*time.Hour),
		//RefreshTTL: ttl("REFRESH_TTL", 24*time.Hour),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
