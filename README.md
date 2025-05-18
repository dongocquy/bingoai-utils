
```
go get github.com/dongocquy/bingoai-utils@latest
```
## Logger

Gói `logger` cung cấp hàm `HandleError` để xử lý lỗi với ba phương pháp: ghi file log, gửi Telegram, và gửi Sentry.

```go
import "github.com/dongocquy/bingoai-utils/logger"

config := logger.Config{
    LogDir:         "logs",
    SentryDSN:      "your_sentry_dsn",
    TelegramToken:  "your_telegram_token",
    TelegramChatID: "your_chat_id",
    Environment:    "production",
    Fatal:          false,
}
logger.HandleError(err, "Database", "Không kết nối được", config)