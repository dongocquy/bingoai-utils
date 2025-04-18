package templatehelper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// RenderRemoteTemplate tải layout từ layoutURL và inject title + nội dung HTML từ file cục bộ
func RenderRemoteTemplate(layoutURL, htmlFilePath, pageTitle string) (string, error) {
	resp, err := http.Get(layoutURL)
	if err != nil {
		return "", fmt.Errorf("❌ Không thể lấy layout từ %s: %w", layoutURL, err)
	}
	defer resp.Body.Close()

	layoutBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("❌ Không đọc được nội dung layout: %w", err)
	}
	layout := string(layoutBytes)

	contentBytes, err := os.ReadFile(htmlFilePath)
	if err != nil {
		return "", fmt.Errorf("❌ Không đọc được file nội dung %s: %w", htmlFilePath, err)
	}
	content := string(contentBytes)

	layout = strings.Replace(layout, "{{ .Title }}", pageTitle, 1)
	layout = strings.Replace(layout, "{{ template \"content\" . }}", content, 1)

	return layout, nil
}
