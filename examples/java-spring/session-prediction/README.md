# Session Prediction — CWE-330

La version vulnérable génère les identifiants de session à l'aide d'un simple compteur séquentiel (`session-100001`, `session-100002`, ...). Un attaquant disposant d'un seul identifiant de session valide — le sien — peut en déduire la plage de valeurs voisines et énumérer des sessions actives d'autres utilisateurs, usurpant ainsi leur identité sans jamais connaître leur mot de passe.

La version corrigée génère chaque identifiant via un générateur aléatoire cryptographiquement sûr (`SecureRandom`), avec au moins 128 bits d'entropie, rendant toute prédiction ou énumération computationnellement infaisable. En production, il est préférable de déléguer entièrement cette génération à un mécanisme de session éprouvé (Spring Session, conteneur d'application), qui applique déjà ces garanties.

**Référence** : CWE-330 (Use of Insufficiently Random Values).
