package curl

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/idio1981/go-pn-public/fo"
	"github.com/idio1981/go-pn-public/logger"
)

func AutoUpload(url string, file string, retry int, options *map[string]string) error {
	curTry := 0
	for curTry <= retry {
		logger.Info("curl upload try: %d, %s", curTry, url)

		curTry++
		_, err := Upload(url, file, options)
		if err == nil {
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	logger.Error("curl upload failed: %s, retry: %d", url, retry)
	return fmt.Errorf("curl upload failed: %s", url)
}

func AutoDownload(url string, savePath string, md5 string, retry int, options *map[string]string) error {
	if md5 != "" {
		if _, err := os.Stat(savePath); err == nil {
			md5file, err := fo.Md5(savePath)
			if err != nil {
				return err
			}
			if md5 == md5file {
				return nil
			}
		}
	}

	curTry := 0
	for curTry <= retry {
		logger.Info("curl download try: %d, %s", curTry, url)

		curTry++
		opts := options
		fi, err := os.Stat(savePath)
		if err == nil && opts != nil {
			(*opts)["-C"] = strconv.FormatInt(fi.Size(), 10)
		}

		_, err = Download(url, savePath, opts)
		if err != nil {
			continue
		}

		if md5 == "" {
			return nil
		}

		md5file, err := fo.Md5(savePath)
		if err != nil {
			return err
		}

		if md5 == md5file {
			return nil
		}

		os.Remove(savePath)
		time.Sleep(1 * time.Second)
	}

	logger.Error("curl download failed: %s, retry: %d", url, retry)
	return fmt.Errorf("curl download failed: %s", url)
}

func Upload(url string, file string, options *map[string]string) (string, error) {
	opts := []string{}
	if options != nil {
		for k, v := range *options {
			opts = append(opts, k, v)
		}
	}

	if file != "" {
		opts = append(opts, "-F", "file=@"+file)
	}

	opts = append(opts, url)
	cmd := exec.Command("curl", opts...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("curl upload failed: %s, %s, %s", url, file, string(output))
		return string(output), err
	}

	logger.Success("curl upload success: %s, %s", url, string(output))
	return string(output), nil
}

func Download(url string, savePath string, options *map[string]string) (string, error) {
	opts := []string{}
	if options != nil {
		for k, v := range *options {
			opts = append(opts, k, v)
		}
	}

	opts = append(opts, "-o", savePath, url)
	cmd := exec.Command("curl", opts...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("curl download failed: %s, %s, %s", url, savePath, string(output))
		return string(output), err
	}

	logger.Success("curl download success: %s, %s", url, string(output))
	return string(output), nil
}

func GetSize(url string) (int64, error) {
	cmd := exec.Command("curl", "-I", url)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("curl get size failed: %s, %s", url, string(output))
		return 0, err
	}

	re := regexp.MustCompile(`content-length: (\d+)`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		logger.Error("curl get size failed: %s, %s", url, string(output))
		return 0, fmt.Errorf("content length not found")
	}

	size, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		logger.Error("curl get size failed: %s, %s", url, string(output))
		return 0, err
	}

	return size, nil
}
