# bingoai-utils

A utility library for Go projects.

## Installation

```bash
go get github.com/dongocquy/bingoai-utils@v0.3.2
```

## Packages

### logger

The `logger` package provides `HandleError` to handle errors with file logging, Telegram notifications, and Sentry reporting.

#### Usage

```go
import "github.com/dongocquy/bingoai-utils/logger"

func main() {
    // LoggerConfig là cấu hình logger dùng chung
    var LoggerConfig = logger.Config{
        LogDir:         os.Getenv("LOG_DIR", "logs"),
        SentryDSN:      os.Getenv("SENTRY_DSN"),
        TelegramToken:  os.Getenv("TELEGRAM_BOT_TOKEN"),
        TelegramChatID: os.Getenv("TELEGRAM_ADMIN_ID"),
        Environment:    os.Getenv("APP_ENV", "development"),
        Fatal:          false,
    }
    err := errors.New("test error")
    logger.HandleError(err, "Test", "Lỗi thử nghiệm", LoggerConfig)
}
```

#### Config Fields
- `LogDir`: Directory for log files (default: `logs`).
- `SentryDSN`: Sentry DSN for error reporting.
- `TelegramToken`, `TelegramChatID`: Telegram bot token and chat ID.
- `Environment`: Application environment (e.g., dev, staging, prod).
- `Fatal`: If true, stops the application on error.

### crypto

[Existing crypto package documentation...]