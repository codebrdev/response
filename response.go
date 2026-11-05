package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	// O status já foi enviado; um erro de encode só pode ser de escrita (cliente desconectado) e não há como corrigir a resposta
	_ = json.NewEncoder(w).Encode(data)
}
