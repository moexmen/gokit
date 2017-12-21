package web

import (
	"net/http"
	"os"
)

// FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
// This implementation will always returns 404 Not Found if the request is a directory, and will not serve `index.html`.
// https://groups.google.com/d/msg/golang-nuts/bStLPdIVM6w/wSKqNoaSji8J
func FileServer(root http.Dir) http.Handler {
	return http.FileServer(justFilesFilesystem{root})
}

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)

	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}

	return f, nil
}
