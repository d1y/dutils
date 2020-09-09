// Author: d1y<chenhonzhou@gmail.com>

package check

import (
	"net/url"
	"os/exec"
)

// CommandAvailableUnix command exists(判断一个命令是否存在)
//
// only support unix
func CommandAvailableUnix(name string) bool {
	cmd := exec.Command("command", "-v", name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// ValidURL check url
func ValidURL(testURL string) bool {
	_, err := url.ParseRequestURI(testURL)
	if err != nil {
		return false
	}

	u, err := url.Parse(testURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
