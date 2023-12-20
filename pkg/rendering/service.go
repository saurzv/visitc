package rendering

import (
	"bytes"
	"path"
	"runtime"
	"text/template"

	"github.com/saurzv/visitc/pkg/listing"
)

func SVG(site listing.Site) ([]byte, error) {
	_, file, _, _ := runtime.Caller(0)
	tmplPath := path.Dir(file)

	var out bytes.Buffer

	t, err := template.New("svg").ParseFiles(tmplPath + "/template/tmpl.svg")
	if err != nil {
		return out.Bytes(), err
	}

	err = t.ExecuteTemplate(&out, "svg", site)
	if err != nil {
		return out.Bytes(), err
	}
	return out.Bytes(), nil
}
