# Remédiation — Mauvaise configuration Firebase (Firestore/Realtime Database/Storage)

## Principe
Remplacer les règles de sécurité en mode test (`if true`) par des règles explicites basées sur l'authentification ET la propriété des données, avec validation de schéma sur les écritures.

## Firestore (firestore.rules)
```
// Avant — vulnérable : mode test, tout le monde peut lire/écrire
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /{document=**} {
      allow read, write: if true;
    }
  }
}

// Après — sécurisé : lecture/écriture limitées au propriétaire du document
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /users/{userId}/documents/{docId} {
      allow read, write: if request.auth != null
                          && request.auth.uid == userId;
      allow create: if request.auth != null
                     && request.auth.uid == userId
                     && request.resource.data.keys().hasOnly(['title', 'content', 'createdAt']);
    }
  }
}
```

## Realtime Database (database.rules.json)
```json
// Avant — vulnérable
{
  "rules": {
    ".read": true,
    ".write": true
  }
}

// Après — sécurisé
{
  "rules": {
    "users": {
      "$uid": {
        ".read": "auth != null && auth.uid === $uid",
        ".write": "auth != null && auth.uid === $uid"
      }
    }
  }
}
```

## Cloud Storage (storage.rules)
```
// Avant — vulnérable
rules_version = '2';
service firebase.storage {
  match /b/{bucket}/o {
    match /{allPaths=**} {
      allow read, write: if true;
    }
  }
}

// Après — sécurisé : upload limité au propriétaire, type et taille contrôlés
rules_version = '2';
service firebase.storage {
  match /b/{bucket}/o {
    match /users/{userId}/uploads/{fileName} {
      allow read: if request.auth != null && request.auth.uid == userId;
      allow write: if request.auth != null && request.auth.uid == userId
                    && request.resource.size < 5 * 1024 * 1024
                    && request.resource.contentType.matches('image/.*');
    }
  }
}
```

## Checklist de vérification post-patch
- [ ] Aucune règle `allow read, write: if true;` ne subsiste en production.
- [ ] Chaque règle de lecture/écriture vérifie `request.auth != null` ET la propriété de la ressource.
- [ ] Les écritures valident le schéma des champs autorisés (`hasOnly`, types, tailles).
- [ ] Les règles Storage limitent le type MIME et la taille des fichiers uploadés.
- [ ] Les règles ont été testées avec l'émulateur Firebase (cas autorisé + cas refusé).
- [ ] Le déploiement des règles (`firebase deploy --only firestore:rules,storage:rules`) est confirmé en production.
