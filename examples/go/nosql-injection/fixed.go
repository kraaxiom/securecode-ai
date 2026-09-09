// fixed.go - Correction CWE-943
// Les identifiants sont decodes dans une struct typee (string, string) :
// tout champ JSON envoye comme objet ({"$ne": null}) provoque une erreur
// de decodage au lieu d'etre transmis comme operateur MongoDB. Le
// filtre est ensuite construit explicitement avec des valeurs scalaires
// uniquement.
package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(coll *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds loginRequest
		// FIXED: struct typee -> un objet/operateur JSON provoque une erreur de decodage
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "requete invalide", http.StatusBadRequest)
			return
		}
		if creds.Username == "" || creds.Password == "" {
			http.Error(w, "requete invalide", http.StatusBadRequest)
			return
		}

		// FIXED: le filtre est construit avec des valeurs scalaires explicites
		filter := bson.M{
			"username": creds.Username,
			"password": creds.Password, // en pratique: comparer un hash, jamais le mot de passe en clair
		}

		var user bson.M
		err := coll.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			http.Error(w, "identifiants invalides", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
