<?php

// Faille : le service écrit dans Firestore en supposant que les règles de
// sécurité sont en mode test (`allow read, write: if true;`). Le code ne
// transmet aucun jeton d'identité d'utilisateur et ne vérifie jamais la
// propriété du document côté serveur, en s'appuyant entièrement sur des
// règles ouvertes qui laissent n'importe quel client lire/écrire toute la
// base.

namespace App\Services;

use Kreait\Firebase\Factory;

class UserNoteService
{
    public function saveNote(string $userId, string $noteId, array $data): void
    {
        $firestore = (new Factory())
            ->withServiceAccount(config('firebase.credentials'))
            ->createFirestore();

        // Écriture directe sans vérifier que $userId correspond à
        // l'utilisateur authentifié : repose sur des règles Firestore
        // ouvertes (mode test) pour "fonctionner".
        $firestore->database()
            ->collection('documents')
            ->document($noteId)
            ->set($data);
    }
}
