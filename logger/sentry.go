package logger

import (
	"log"

	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"
)

// SentryInitialized theo dõi trạng thái khởi tạo Sentry

// SendToSentry gửi thông điệp hoặc lỗi tới Sentry
func SendToSentry(c *fiber.Ctx, message string, err error, errorType string) {
	// Lấy hub từ context Fiber hoặc hub mặc định
	var hub *sentry.Hub
	if c != nil {
		hub = sentryfiber.GetHubFromContext(c)
	}
	if hub == nil {
		hub = sentry.CurrentHub()
	}

	if hub == nil {
		log.Println("❌ Không tìm thấy Sentry hub, bỏ qua gửi sự kiện")
		return
	}

	// Tạo scope và thêm metadata
	hub.WithScope(func(scope *sentry.Scope) {
		scope.SetTag("error_type", errorType)
		if c != nil {
			scope.SetTag("request_method", c.Method())
			scope.SetTag("request_path", c.Path())
		}
		if err != nil {
			scope.SetExtra("error", err.Error())
		}
		scope.SetExtra("message", message)

		// Gửi thông điệp hoặc lỗi
		if err != nil {
			hub.CaptureException(err)
		} else {
			hub.CaptureMessage(message)
		}
	})
}
