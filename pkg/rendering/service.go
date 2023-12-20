package rendering

import (
	"bytes"
	"log"
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/saurzv/visitc/pkg/listing"
)

func SVG(site listing.Site) error {
	_, file, _, _ := runtime.Caller(0)
	p := path.Dir(file)
	t, err := template.New("svg").ParseFiles(p + "/template/tmpl.svg")
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	err = t.ExecuteTemplate(&out, "svg", site)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile("out.svg", out.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}
