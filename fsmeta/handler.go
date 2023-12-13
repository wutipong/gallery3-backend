package fsmeta

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/gallery3-backend/handler"
)

var rootDir = ""

func Init(dir string) error {
	r, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	rootDir = r

	return nil
}

type request struct {
	Path string `json:"path"`
}

type response struct {
	Directories []string `json:"directories"`
	Files       []string `json:"files"`
}

// @Param 		path path string false "path"
// @Success      200  {object}  fsmeta.response
// @Failure      500  {object}  errors.Error
// @Router /dir/{path} [get]
func Handler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	p := params.ByName("path")

	filesystem := os.DirFS(rootDir)
	entries, err := fs.ReadDir(filesystem, filepath.Join("./", p))
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	resp := response{
		Directories: []string{},
		Files:       []string{},
	}

	for _, entry := range entries {
		if entry.IsDir() {
			resp.Directories = append(resp.Directories, entry.Name())
		} else {
			resp.Files = append(resp.Files, entry.Name())
		}
	}

	log.Info().
		Interface("path", p).
		Msg("get")

	handler.WriteResponse(w, resp)
}
