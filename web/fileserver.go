package web

import (
	"net/http"
	"os"
	"strings"
)

// FileServer creates a file server that serves files from from a "Root" folder.
// It will call "NotFound" HandlerFunc if the path contains '..' or if the file cannot be found on the system.
type FileServer struct {
	Root     string
	NotFound http.HandlerFunc
}

func (f FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if containsDotDot(r.URL.Path) {
		f.NotFound(w, r)
		return
	}
	name := f.Root + r.URL.Path
	if _, err := os.Stat(name); os.IsNotExist(err) {
		f.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, name)
}

// This is copied from https://github.com/golang/go/blob/master/src/net/http/fs.go#L676
func containsDotDot(v string) bool {
	if !strings.Contains(v, "..") {
		return false
	}
	for _, ent := range strings.FieldsFunc(v, isSlashRune) {
		if ent == ".." {
			return true
		}
	}
	return false
}

// This is copied from https://github.com/golang/go/blob/master/src/net/http/fs.go#L688
func isSlashRune(r rune) bool { return r == '/' || r == '\\' }
