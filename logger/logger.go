package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
)

// Config chứa cấu hình cho HandleError
type Config struct {
	LogDir         string
	SentryDSN      string
	TelegramToken  string
	TelegramChatID string
	Environment    string
	Fatal          bool
}

// SentryInitialized là biến toàn cục theo dõi trạng thái khởi tạo Sentry
var SentryInitialized bool

// HandleError xử lý lỗi với ghi log, gửi Telegram, và Sentry
func HandleError(err error, errorType, message string, config Config) {
	// Tạo thông điệp lỗi chi tiết
	errorMsg := fmt.Sprintf("🛑 *LỖI HỆ THỐNG*\n\n📅 %s\n🌍 Môi trường: %s\n📍 *Loại*: %s\n💥 *Lỗi*: %s\n🔍 *Chi tiết*: %v",
		time.Now().Format("2006-01-02 15:04:05"),
		config.Environment,
		errorType,
		message,
		err)

	// Log vào stdout để debug
	log.Println(errorMsg)

	// Sử dụng WaitGroup để đảm bảo tất cả goroutines hoàn thành
	var wg sync.WaitGroup
	wg.Add(3)

	// 1. Ghi vào file log cục bộ
	go func() {
		defer wg.Done()
		logDir := config.LogDir
		if logDir == "" {
			logDir = "logs"
		}
		logFile := filepath.Join(logDir, fmt.Sprintf("error_%s.log", time.Now().Format("2006-01-02")))

		if err := os.MkdirAll(logDir, 0755); err != nil {
			log.Printf("❌ Không thể tạo thư mục logs: %v", err)
			return
		}

		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("❌ Không thể mở file log: %v", err)
			return
		}
		defer f.Close()

		// Xử lý trường hợp err là nil
		errorStr := "No error provided"

		// Ghi lỗi vào file với định dạng JSON
		logEntry := map[string]interface{}{
			"timestamp":  time.Now().Format("2006-01-02 15:04:05"),
			"env":        config.Environment,
			"error_type": errorType,
			"message":    message,
			"error":      errorStr,
		}
		logData, _ := json.Marshal(logEntry)
		if _, err := f.WriteString(string(logData) + "\n"); err != nil {
			log.Printf("❌ Không thể ghi vào file log: %v", err)
		} else {
			log.Println("✅ Đã ghi lỗi vào file log")
		}
	}()

	// 2. Gửi thông báo qua Telegram
	go func() {
		defer wg.Done()
		if config.TelegramToken == "" || config.TelegramChatID == "" {
			log.Println("❌ Thiếu TelegramToken hoặc TelegramChatID")
			return
		}

		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.TelegramToken)
		payload := map[string]interface{}{
			"chat_id":    config.TelegramChatID,
			"text":       errorMsg,
			"parse_mode": "Markdown",
		}
		body, _ := json.Marshal(payload)

		resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
		if err != nil {
			log.Printf("❌ Không thể gửi Telegram: %v", err)
			return
		}
		defer resp.Body.Close()
		log.Println("✅ Đã gửi thông báo Telegram")
	}()

	// 3. Gửi lỗi đến Sentry
	go func() {
		defer wg.Done()
		if config.SentryDSN == "" {
			log.Println("❌ Thiếu SentryDSN")
			return
		}
		if !SentryInitialized {
			log.Println("❌ Sentry chưa được khởi tạo")
			return
		}
		sentry.CaptureMessage(errorMsg)
		log.Println("✅ Đã gửi lỗi đến Sentry")
	}()

	// Đợi tất cả goroutines hoàn thành
	wg.Wait()

	// Dừng ứng dụng nếu fatal được yêu cầu
	if config.Fatal {
		log.Fatal(errorMsg)
	}
}