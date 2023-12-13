package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/gallery3-backend/errors"
)

func WriteResponse(w http.ResponseWriter, v any) {
	if err, ok := v.(error); ok {
		w.WriteHeader(http.StatusInternalServerError)
		if _, ok := err.(errors.Error); !ok {
			v = errors.ErrUnknown.Wrap(err)
		}
		log.Error().AnErr("error", v.(error)).Msg("Handler response fails.")
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, _ := json.Marshal(v)
	w.Write(b)

	log.Info().RawJSON("response", b).Msg("Response")
}
