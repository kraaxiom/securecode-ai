# Remédiation — Session Replay

## Principe
Inclure et vérifier une expiration courte (`exp`) sur tout jeton d'authentification, marquer les jetons à usage unique comme consommés dès leur première utilisation valide côté serveur, et utiliser un nonce pour les mécanismes sensibles.

## PHP
```php
// Avant — vulnérable
function verifierTokenReset($token) {
    $data = decoderToken($token);
    return $data['user_id']; // jamais marqué comme utilisé, réutilisable indéfiniment
}

// Après — sécurisé
function verifierTokenReset($token) {
    $data = decoderToken($token);
    if ($data['exp'] < time()) {
        throw new Exception('Jeton expiré');
    }
    if (tokenDejaUtilise($data['jti'])) { // jti = identifiant unique du jeton
        throw new Exception('Jeton déjà utilisé');
    }
    marquerTokenCommeUtilise($data['jti']);
    return $data['user_id'];
}
```

## Node.js (Express)
```js
// Avant — vulnérable
function verifyToken(token) {
  const payload = jwt.decode(token);
  return payload.userId; // pas de vérification exp/jti, rejouable
}

// Après — sécurisé
async function verifyToken(token) {
  const payload = jwt.verify(token, process.env.JWT_SECRET); // vérifie exp/nbf automatiquement
  if (await isTokenUsed(payload.jti)) {
    throw new Error('Jeton déjà utilisé');
  }
  await markTokenUsed(payload.jti);
  return payload.userId;
}
```

## Python (Flask/Django)
```python
# Avant — vulnérable
def verify_reset_token(token):
    data = decode_token(token)
    return data['user_id']  # aucune vérification d'expiration ni d'unicité

# Après — sécurisé
def verify_reset_token(token):
    data = decode_token(token)  # ex: itsdangerous.URLSafeTimedSerializer avec max_age
    if is_token_used(data['jti']):
        raise ValueError('Jeton déjà utilisé')
    mark_token_used(data['jti'])
    return data['user_id']
```

## Checklist de vérification post-patch
- [ ] Tout jeton d'authentification porte une expiration courte vérifiée côté serveur.
- [ ] Les jetons à usage unique (reset password, magic link) sont marqués consommés après la première utilisation.
- [ ] Un stockage serveur (nonce/jti) permet d'invalider un jeton déjà présenté.
- [ ] Un test confirme qu'une seconde présentation du même jeton est rejetée.
