// https://github.com/gookit/goutil/blob/master/envutil/info.go
// https://github.com/gookit/goutil/blob/master/sysutil/sysutil.go

package info

import (
	"io"
	"os"
	"runtime"
	"strings"
	"syscall"
)

// IsWin system. linux windows darwin
func IsWin() bool {
	return runtime.GOOS == "windows"
}

// IsMac system
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux system
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsConsole check out is console env. alias of the sysutil.IsConsole()
func IsConsole(out io.Writer) bool {
	o, ok := out.(*os.File)
	if !ok {
		return false
	}

	fd := o.Fd()

	// fix: cannot use 'o == os.Stdout' to compare
	return fd == uintptr(syscall.Stdout) || fd == uintptr(syscall.Stdin) || fd == uintptr(syscall.Stderr)
}

// IsMSys msys(MINGW64) env. alias of the sysutil.IsMSys()
func IsMSys() bool {
	// "MSYSTEM=MINGW64"
	if len(os.Getenv("MSYSTEM")) > 0 {
		return true
	}

	return false
}

// Support color:
// 	"TERM=xterm"
// 	"TERM=xterm-vt220"
// 	"TERM=xterm-256color"
// 	"TERM=screen-256color"
// 	"TERM=tmux-256color"
// 	"TERM=rxvt-unicode-256color"
// Don't support color:
// 	"TERM=cygwin"
var specialColorTerms = map[string]bool{
	"alacritty":             true,
	"screen-256color":       true,
	"tmux-256color":         true,
	"rxvt-unicode-256color": true,
}

// IsSupportColor check current console is support color.
//
// Supported:
// 	linux, mac, or windows's ConEmu, Cmder, putty, git-bash.exe
// Not support:
// 	windows cmd.exe, powerShell.exe
func IsSupportColor() bool {
	envTerm := os.Getenv("TERM")
	if strings.Contains(envTerm, "xterm") {
		return true
	}

	// it's special color term
	if _, ok := specialColorTerms[envTerm]; ok {
		return true
	}

	// like on ConEmu software, e.g "ConEmuANSI=ON"
	if os.Getenv("ConEmuANSI") == "ON" {
		return true
	}

	// like on ConEmu software, e.g "ANSICON=189x2000 (189x43)"
	if os.Getenv("ANSICON") != "" {
		return true
	}

	// up: if support 256-color, can also support basic color.
	return IsSupport256Color()
}

// IsSupport256Color render
func IsSupport256Color() bool {
	// "TERM=xterm-256color"
	// "TERM=screen-256color"
	// "TERM=tmux-256color"
	// "TERM=rxvt-unicode-256color"
	supported := strings.Contains(os.Getenv("TERM"), "256color")
	if !supported {
		// up: if support true-color, can also support 256-color.
		supported = IsSupportTrueColor()
	}

	return supported
}

// IsSupportTrueColor render. IsSupportRGBColor
func IsSupportTrueColor() bool {
	// "COLORTERM=truecolor"
	return strings.Contains(os.Getenv("COLORTERM"), "truecolor")
}
