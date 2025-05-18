package logger

// Config chứa cấu hình cho HandleError
type Config struct {
    LogDir         string // Thư mục lưu file log (mặc định: "logs")
    SentryDSN      string // DSN cho Sentry
    TelegramToken  string // Token cho Telegram bot
    TelegramChatID string // Chat ID cho Telegram
    Environment    string // Môi trường (dev, staging, prod)
    Fatal          bool   // Có dừng ứng dụng khi gặp lỗi không
}