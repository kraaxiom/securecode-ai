// vulnerable.go - CWE-943: Injection NoSQL (MongoDB)
// Le corps JSON de la requete HTTP est decode directement en bson.M et
// transmis tel quel comme filtre de requete. Un attaquant peut envoyer
// {"username": "admin", "password": {"$ne": null}} pour contourner
// l'authentification en injectant un operateur MongoDB au lieu d'une
// valeur scalaire.
package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginHandler(coll *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds bson.M
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "requete invalide", http.StatusBadRequest)
			return
		}

		// VULNERABLE: creds peut contenir des operateurs Mongo ($ne, $gt...)
		var user bson.M
		err := coll.FindOne(context.Background(), creds).Decode(&user)
		if err != nil {
			http.Error(w, "identifiants invalides", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
