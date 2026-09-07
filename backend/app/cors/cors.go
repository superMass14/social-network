package cors

import "net/http"

func SetCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "*, Authorisation")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func SetCorsy(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "*, Autorisation")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

	// Vérifier si la requête est de type OPTIONS (pour les pré-vols CORS)
	if r.Method == "OPTIONS" {
		// Répondre avec succès et terminer la requête
		(*w).WriteHeader(http.StatusOK)
		return
	}
}
