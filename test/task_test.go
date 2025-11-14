package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/idio1981/go-pn-public/curl"
	"github.com/idio1981/go-pn-public/logger"
)

func TestCURLUpload(t *testing.T) {
	perpare()
	ctx, _ := context.WithCancel(context.Background())
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	cancel()
	// }()

	uri := "http://apps.17995.com/upload.php"
	file := "/Users/muzi/Downloads/Cathub.17128.1712890024.52814.ipa"
	output, err := curl.Upload(ctx, uri, file, nil)
	if err != nil {
		logger.Error("upload failed: %s", err)
		return
	}

	logger.Success("upload success: %s", output)
}

func TestCURLGetSize(t *testing.T) {
	perpare()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	size, err := curl.GetSize(ctx, "https://ndog6appskf8kdfn.17995cdn.com/appcenter/uploads/launcherrele.1762769572.72825.aab")
	if err != nil {
		logger.Error("get size failed: %s", err)
		return
	}

	logger.Success("size: %d", size)
}

func TestCURLDownload(t *testing.T) {
	perpare()
	ctx, _ := context.WithCancel(context.Background())
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	cancel()
	// }()

	uri := "https://ndog6appskf8kdfn.17995cdn.com/appcenter/uploads/launcherrele.1762769572.72825.aab"
	savePath := "launcherrele.52814.ipa"
	_, err := curl.Download(ctx, uri, savePath, nil)
	if err != nil {
		logger.Error("download failed: %s", err)
		return
	}

	logger.Success("download success: %s", savePath)

}

func TestCURLDownload2(t *testing.T) {
	perpare()
	ctx, _ := context.WithCancel(context.Background())
	uri := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/financeReports?filter[reportDate]=%s&filter[reportType]=FINANCE_DETAIL&filter[regionCode]=Z1&filter[vendorNumber]=%s", "2025-09", "87349754")
	err := curl.AutoDownload(ctx, uri, "/Users/muzi/Downloads/financeReport.zip", "", 3, &map[string]string{
		"--location": "",
		"--globoff":  "",
		"--header":   fmt.Sprintf("Authorization: Bearer %s", "eyJhbGciOiJFUzI1NiIsImtpZCI6IjQ0Wlo1NlY2RlUiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhcHBzdG9yZWNvbm5lY3QtdjEiLCJleHAiOjE3NjMwMDMzMTgsImlhdCI6MTc2MzAwMjExOCwiaXNzIjoiNjlhNmRlOTAtNWE0ZC00N2UzLWUwNTMtNWI4YzdjMTFhNGQxIn0.9e0nCqMlAWhnmxMdyxtaXPepyiLx4yhjt5hkmxoPygHqCbXp7iXWyDWHREtNFYzv5W2-5nON0cFEnVUjoSa0cA"),
	})

	if err != nil {
		logger.Error("download failed: %v", err)
		return
	}

	logger.Success("download success: %s", "/Users/muzi/Downloads/financeReport.zip")

}

func TestColors(t *testing.T) {
	perpare()
	logger.Success("Success: %s", "Success")
	logger.Info("Info: %s", "Info")
	logger.Warn("Warn: %s", "Warn")
	logger.Error("Error: %s", "Error")
	logger.Debug("Debug: %s", "Debug")
}
