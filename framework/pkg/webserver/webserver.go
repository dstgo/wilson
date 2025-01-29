package webserver

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/dstgo/wilson/framework/pkg/strs"
)

func ServeDir(dir string, addr string, data map[string]any) error {
	path := filepath.Join(dir, "index.html")
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if len(data) != 0 {
		var out = bytes.NewBuffer(nil)
		tpl := template.New("html")
		parser, err := tpl.Parse(strs.BytesToString(content))
		if err != nil {
			return err
		}

		if err = parser.Execute(out, data); err != nil {
			panic(err)
		}
		content = out.Bytes()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(dir, r.URL.Path)
		if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
			http.ServeFile(w, r, path)
			return
		}

		accept := r.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(content)
			return
		}
		http.NotFound(w, r)
	})

	log.Infof("static web server lisenting at %s\n", addr)

	// G114 (CWE-676): Use of net/http serve function that has no support for setting timeouts
	// #nosec
	if err := http.ListenAndServe(addr, nil); err != nil {
		return fmt.Errorf("failed to start web server: %s", err)
	}

	return err
}
