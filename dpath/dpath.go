package dpath

import (
	"errors"
	"os"
	"path"
	"runtime"
)

var bos = runtime.GOOS

// Homedir get user home dir
var Homedir string

const (

	// LinuxDesktopEnv linux `desktop` folder env
	LinuxDesktopEnv = "XDG_DESKTOP_DIR"

	// LinuxAudioEnv linux `audio` folder env
	LinuxAudioEnv = "XDG_MUSIC_DIR"

	// LinuxTemplateEnv linux `template` folder env
	LinuxTemplateEnv = "XDG_TEMPLATES_DIR"

	// LinuxVideosEnv linux `videos` folder env
	LinuxVideosEnv = "XDG_VIDEOS_DIR"

	// LinuxDownloadEnv linux `videos` folder env
	LinuxDownloadEnv = "XDG_DOWNLOAD_DIR"

	// LinuxDocumentEnv linux `document` folder env
	LinuxDocumentEnv = "XDG_DOCUMENTS_DIR"
)

const (

	// WindowsDesktopEnv windows `desktop` folder env
	WindowsDesktopEnv = "FOLDERID_Desktop"

	// WindowsAudioEnv windows `audio` folder env
	WindowsAudioEnv = "FOLDERID_Music"

	// WindowsTemplateEnv windows `template` folder env
	WindowsTemplateEnv = "FOLDERID_Templates"

	// WindowsVideoEnv windows `video` folder env
	WindowsVideoEnv = "FOLDERID_Videos"

	// WindowsDownloadEnv windows `download` folder env
	WindowsDownloadEnv = "FOLDERID_Downloads"

	// WindowsDocumentEnv windows `document` folder env
	WindowsDocumentEnv = "FOLDERID_Documents"
)

// GetDocumentDir get document dir
//
// https://github.com/justjavac/deno_document_dir
func GetDocumentDir() (string, error) {
	switch bos {
	case "windows":
		return os.Getenv(WindowsDocumentEnv), nil
	case "linux":
		return os.Getenv(LinuxDocumentEnv), nil
	case "darwin":
		return path.Join(Homedir, "./Documents"), nil
	}
	return "", errors.New("get documents dir failed")
}

// GetDownloadDir get download dir
//
// https://github.com/justjavac/deno_download_dir
func GetDownloadDir() (string, error) {
	switch bos {
	case "windows":
		return os.Getenv(WindowsDownloadEnv), nil
	case "linux":
		return os.Getenv(LinuxDownloadEnv), nil
	case "darwin":
		return path.Join(Homedir, "./Downloads"), nil
	}
	return "", errors.New("get download dir failed")
}

// GetDesktopDir get desktop dir
//  * The returned value depends on the operating system and is either a string,
//  * containing a value from the following table, or `null`.
//  *
//  * |Platform | Value                | Example                    |
//  * | ------- | -------------------- | -------------------------- |
//  * | Linux   | `XDG_DESKTOP_DIR`    | /home/justjavac/Desktop    |
//  * | macOS   | `$HOME`/Desktop      | /Users/justjavac/Desktop   |
//  * | Windows | `{FOLDERID_Desktop}` | C:\Users\justjavac\Desktop |
func GetDesktopDir() (string, error) {
	switch bos {
	case "linux":
		return os.Getenv(LinuxDesktopEnv), nil
	case "windows":
		return os.Getenv(WindowsDesktopEnv), nil
	case "darwin":
		return path.Join(Homedir, "./Desktop"), nil
	}
	return "", errors.New("get desktop dir failed")
}

// GetAudioDir get audio dir
//
//  * The returned value depends on the operating system and is either a string,
//  * containing a value from the following table, or `null`.
//  *
//  * |Platform | Value              | Example                  |
//  * | ------- | ------------------ | ------------------------ |
//  * | Linux   | `XDG_MUSIC_DIR`    | /home/justjavac/Music    |
//  * | macOS   | `$HOME`/Music      | /Users/justjavac/Music   |
//  * | Windows | `{FOLDERID_Music}` | C:\Users\justjavac\Music |
// https://github.com/justjavac/deno_audio_dir/blob/master/mod.ts
func GetAudioDir() (string, error) {
	switch bos {
	case "linux":
		return os.Getenv(LinuxAudioEnv), nil
	case "darwin":
		return path.Join(Homedir, "Music"), nil
	case "windows":
		return os.Getenv(WindowsAudioEnv), nil
	}
	return "", errors.New("get audio dir failed")
}

// GetTemplateDir Returns the path to the user's template directory.
//
// Only supports linux/windows
//
//  * The returned value depends on the operating system and is either a string,
//  * containing a value from the following table, or `null`.
//  *
//  * |Platform | Value                  | Example                                                        |
//  * | ------- | ---------------------- | -------------------------------------------------------------- |
//  * | Linux   | `XDG_TEMPLATES_DIR`    | /home/justjavac/Templates                                      |
//  * | macOS   | –                      | –                                                              |
//  * | Windows | `{FOLDERID_Templates}` | C:\Users\justjavac\AppData\Roaming\Microsoft\Windows\Templates |
// https://github.com/justjavac/deno_template_dir/blob/master/mod.ts
func GetTemplateDir() (string, error) {
	switch bos {
	case "windows":
		return os.Getenv(WindowsTemplateEnv), nil
	case "linux":
		return os.Getenv(LinuxTemplateEnv), nil
	}
	return "", errors.New("get template directory failed")
}

// GetVideoDir Returns the path to the user's video directory.
//  * The returned value depends on the operating system and is either a string,
//  * containing a value from the following table, or `null`.
//  *
//  * |Platform | Value               | Example                  |
//  * | ------- | ------------------- | ------------------------- |
//  * | Linux   | `XDG_VIDEOS_DIR`    | /home/justjavac/Videos    |
//  * | macOS   | `$HOME`/Movies      | /Users/justjavac/Movies   |
//  * | Windows | `{FOLDERID_Videos}` | C:\Users\justjavac\Videos |
// https://github.com/justjavac/deno_video_dir/blob/master/mod.ts
func GetVideoDir() (string, error) {
	switch bos {
	case "windows":
		return os.Getenv(WindowsVideoEnv), nil
	case "darwin":
		return path.Join(Homedir, "./Movies"), nil
	case "linux":
		return os.Getenv(LinuxVideosEnv), nil
	}
	return "", errors.New("get video dir failed")
}

func init() {
	dir, _ := os.UserHomeDir()
	Homedir = dir
}
