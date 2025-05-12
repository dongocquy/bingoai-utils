package userstore

import "os"

var (
	TableUserConfig   = getEnv("TABLE_USER_CONFIG", "SaaS_BingoAI")
	TableClientUsers  = getEnv("TABLE_CLIENT_USERS", "SaaS_BingoAI_client_users")
)

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
