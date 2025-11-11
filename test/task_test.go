package test

import (
	"testing"

	"github.com/idio1981/go-pn-public/curl"
	"github.com/idio1981/go-pn-public/logger"
)

func TestCURLUpload(t *testing.T) {
	perpare()
	uri := "http://apps.17995.com/upload.php"
	file := "/Users/muzi/Downloads/Cathub.17128.1712890024.52814.ipa"
	output, err := curl.Upload(uri, file, nil)
	if err != nil {
		logger.Error("upload failed: %s", err)
		return
	}

	logger.Success("upload success: %s", output)
}

func TestCURLGetSize(t *testing.T) {
	perpare()
	size, err := curl.GetSize("https://ndog6appskf8kdfn.17995cdn.com/appcenter/uploads/Cathub.17128.1712890024.52814.ipa")
	if err != nil {
		logger.Error("get size failed: %s", err)
		return
	}

	logger.Success("size: %d", size)
}

func TestCURLDownload(t *testing.T) {
	perpare()
	uri := "https://ndog6appskf8kdfn.17995cdn.com/appcenter/uploads/Cathub.17128.1712890024.52814.ipa"
	savePath := "Cathub.17128.1712890024.52814.ipa"
	_, err := curl.Download(uri, savePath, nil)
	if err != nil {
		logger.Error("download failed: %s", err)
		return
	}

	logger.Success("download success: %s", savePath)

}

func TestColors(t *testing.T) {
	perpare()
	logger.Success("Success: %s", "Success")
	logger.Info("Info: %s", "Info")
	logger.Warn("Warn: %s", "Warn")
	logger.Error("Error: %s", "Error")
	logger.Debug("Debug: %s", "Debug")
}
