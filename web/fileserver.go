package web

import (
	"net/http"
	"os"
	"strings"
)

type fileHandler struct {
	Root     string
	NotFound http.HandlerFunc
}

// FileServer is a FileServerHandler with the NotFound handler set to http.NotFound.
func FileServer(root string) http.Handler {
	return fileHandler{
		Root:     root,
		NotFound: http.NotFound,
	}
}

// FileServerHandler creates a file server that serves files from from a "Root" folder.
// It will call "NotFound" HandlerFunc if the file cannot be found on the system.
// In the event where the path is a directory, it will not redirect and instead call "NotFound" HandlerFunc.
func FileServerHandler(root string, notFound http.HandlerFunc) http.Handler {
	if notFound == nil {
		return FileServer(root)
	}
	return fileHandler{
		Root:     root,
		NotFound: notFound,
	}
}

func (h fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if containsDotDot(r.URL.Path) {
		h.NotFound(w, r)
		return
	}

	name := h.Root + r.URL.Path
	info, err := os.Stat(name)
	if os.IsNotExist(err) || info.IsDir() {
		h.NotFound(w, r)
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
