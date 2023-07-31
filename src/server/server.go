package server

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"text/template"

	"github.com/ericovis/resume-cli/src/resume"
)

var (
	//go:embed templates
	templates embed.FS
	//go:embed static
	static embed.FS
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func Start(resume resume.Resume, port int) {
	fSys, _ := fs.Sub(static, "static")
	http.Handle("/static/", http.FileServer(http.FS(fSys)))

	tmpl := template.Must(template.ParseFS(templates, "templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, resume)
	})

	fmt.Printf("Listening on port %v\n", port)
	openBrowser(fmt.Sprintf("http://localhost:%v", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
