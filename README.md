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
    config := logger.Config{
        LogDir:         "logs",
        SentryDSN:      "your_sentry_dsn",
        TelegramToken:  "your_telegram_token",
        TelegramChatID: "your_chat_id",
        Environment:    "production",
        Fatal:          false,
    }
    err := errors.New("test error")
    logger.HandleError(err, "Test", "Lỗi thử nghiệm", config)
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