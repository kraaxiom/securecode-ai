<?php

// Correction : le service vérifie le jeton d'identité Firebase de
// l'utilisateur côté serveur avant toute écriture, et stocke le document
// sous le chemin `users/{userId}/documents/{docId}` attendu par des règles
// Firestore restrictives (`request.auth.uid == userId`), garantissant que
// même en cas d'appel direct au SDK, seules les données du propriétaire
// authentifié sont modifiées.

namespace App\Services;

use Kreait\Firebase\Auth as FirebaseAuth;
use Kreait\Firebase\Factory;
use RuntimeException;

class UserNoteService
{
    public function saveNote(string $idToken, string $noteId, array $data): void
    {
        $firebase = (new Factory())->withServiceAccount(config('firebase.credentials'));
        $auth = $firebase->createAuth();

        $verifiedIdToken = $auth->verifyIdToken($idToken);
        $userId = $verifiedIdToken->claims()->get('sub');

        if (empty($userId)) {
            throw new RuntimeException('Utilisateur non authentifié.');
        }

        $firestore = $firebase->createFirestore();

        $firestore->database()
            ->collection('users')->document($userId)
            ->collection('documents')->document($noteId)
            ->set($data);
    }
}
