package handlers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// NewFrontendHandler returns a new FrontendHandler.
func NewFrontendHandler(baseDir string) *FrontendHandler {
	return &FrontendHandler{
		BaseDir: baseDir,
	}
}

// FrontendHandler deals with serving up the thing needed for Vugu.
type FrontendHandler struct {
	BaseDir string // absolute path of project
	DevMode bool   // true if we should be automatically rebuilding the frontend (only on local)
}

// ServeHTTP implements http.Handler but only writes a response for files we serve.
func (h *FrontendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	distDir := http.Dir(filepath.Join(h.BaseDir, "dist"))

	buildFrontend := func() (ok bool) {

		cmd := exec.Command("sh", "build-wasm.sh")

		if runtime.GOOS == "windows" {
			cmd = exec.Command(`.\build-wasm.bat`)
		}

		cmd.Env = append(os.Environ(), "GO111MODULE=auto")

		cmd.Dir = h.BaseDir
		b, err := cmd.CombinedOutput()
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(500)
			_, _ = fmt.Fprintf(w, "build-frontend.go error: %v\n%s", err, b)
			return false
		}
		return true
	}

	ext := path.Ext(r.URL.Path)
	if ext != "" {
		f, err := distDir.Open(r.URL.Path)
		if os.IsNotExist(err) {
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("File error: %v", err), 500)
			return
		}
		defer f.Close()

		st, err := f.Stat()
		if err != nil {
			http.Error(w, fmt.Sprintf("File stat error: %v", err), 500)
			return
		}

		// manually handle some mime types
		switch ext {
		case ".wasm":
			w.Header().Set("Content-Type", "application/wasm")

		case ".js":
			w.Header().Set("Content-Type", "application/javascript")

		case ".css":
			w.Header().Set("Content-Type", "text/css")

		}
		http.ServeContent(w, r, r.URL.Path, st.ModTime(), f)
		return
	}

	if !buildFrontend() {
		return
	}
	h.serveHtml(w, r)
	return
}

func (h *FrontendHandler) serveHtml(w http.ResponseWriter, r *http.Request) {
	distDir := http.Dir(filepath.Join(h.BaseDir, "dist"))
	p := r.URL.Path
	if p == "/" {
		p = "/index"
	}
	f, err := distDir.Open(p + ".html")
	if os.IsNotExist(err) {
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("File error: %v", err), 500)
		return
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil {
		http.Error(w, fmt.Sprintf("File stat error: %v", err), 500)
		return
	}
	http.ServeContent(w, r, r.URL.Path, st.ModTime(), f)
	return
}

func (h *FrontendHandler) serveIndex(w http.ResponseWriter, r *http.Request) {
	in := strings.NewReader(indexHTML)
	http.ServeContent(w, r, "tacostore", startupTime, in)
}

var startupTime = time.Now()

var indexHTML = `<!doctype html>
<html>
<head>
<title vg-html='c.Title'></title>
<meta name="msapplication-TileColor" content="#da532c">
<meta name="theme-color" content="#ffffff">
<meta charset="utf-8"/>
<link rel="stylesheet" href="/css/main.css">
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.3.1/dist/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
<!-- styles -->
</head>
<body>
<div id="root">
<img style="position: absolute; top: 50%; left: 50%;" src="https://cdnjs.cloudflare.com/ajax/libs/galleriffic/2.0.1/css/loader.gif">
</div>
<script src="https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js"></script> <!-- MS Edge polyfill -->
<script src="/wasm_exec.js"></script>
<!-- scripts -->
<script>
	var wasmSupported = (typeof WebAssembly === "object");
	if (wasmSupported) {
		if (!WebAssembly.instantiateStreaming) { 
			WebAssembly.instantiateStreaming = async (resp, importObject) => {
				const source = await (await resp).arrayBuffer();
				return await WebAssembly.instantiate(source, importObject);
			};
		}
		const go = new Go();
		WebAssembly.instantiateStreaming(fetch("/main.wasm"), go.importObject).then((result) => {
			go.run(result.instance);
		});
	} else {
		console.log("Full functionality requires WebAssembly support.  Please upgrade your browser.")
	}
</script>
</body>
</html>`
