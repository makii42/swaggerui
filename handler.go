package swaggerui

//go:generate go-bindata-assetfs -nocompress -pkg ui -prefix ../static/ ../static/swagger/js/ ../static/swagger/css/ ../static/swagger/

import (
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

type swaggerUIHandler struct {
	urlPrefix  string
	fileServer http.Handler
}

// NewSwaggerUIHandler - Http Handler that serves swagger ui
func NewSwaggerUIHandler(urlPrefix string) http.Handler {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
		Prefix:   "swagger/", // filesystem prefix
	})

	return &swaggerUIHandler{
		urlPrefix:  urlPrefix,
		fileServer: fileServer,
	}
}

// Serves swagger-ui page from the index.html file located under static directory
func (h swaggerUIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix(h.urlPrefix, h.fileServer).ServeHTTP(w, r)
}
