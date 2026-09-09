# Remédiation — Attaque par batching GraphQL

## Principe
Désactiver le batching s'il n'est pas nécessaire, ou plafonner strictement le nombre d'opérations par requête batchée. Faire compter le rate limiter au niveau des opérations GraphQL exécutées, pas des requêtes HTTP.

## PHP (webonyx/graphql-php, endpoint custom)
```php
// Avant — vulnérable : tableau de requêtes traité sans limite de taille
$body = json_decode(file_get_contents('php://input'), true);
if (is_array($body) && array_is_list($body)) {
    foreach ($body as $op) {
        $results[] = GraphQL::executeQuery($schema, $op['query'], null, null, $op['variables']);
    }
}

// Après — sécurisé
const MAX_BATCH_SIZE = 5;

$body = json_decode(file_get_contents('php://input'), true);
if (is_array($body) && array_is_list($body)) {
    if (count($body) > MAX_BATCH_SIZE) {
        http_response_code(400);
        exit(json_encode(['error' => 'Batch trop volumineux']));
    }
    $rateLimiter->consume($clientId, count($body)); // compte les opérations, pas la requête HTTP
    foreach ($body as $op) {
        $results[] = GraphQL::executeQuery($schema, $op['query'], null, null, $op['variables']);
    }
}
```

## Node.js (Apollo Server / graphql-http)
```js
// Avant — vulnérable : batching activé sans plafond
const server = new ApolloServer({ typeDefs, resolvers, allowBatchedHttpRequests: true });

// Après — sécurisé
const MAX_BATCH_SIZE = 5;

app.use('/graphql', (req, res, next) => {
  if (Array.isArray(req.body) && req.body.length > MAX_BATCH_SIZE) {
    return res.status(400).json({ error: 'Batch trop volumineux' });
  }
  next();
});

const server = new ApolloServer({ typeDefs, resolvers, allowBatchedHttpRequests: true });

// Rate limiter conscient du nombre d'opérations réelles
async function batchAwareLimiter(req, res, next) {
  const opCount = Array.isArray(req.body) ? req.body.length : 1;
  try {
    await limiter.consume(req.ip, opCount);
    next();
  } catch {
    res.status(429).json({ error: 'Trop de requêtes' });
  }
}
```

## Checklist de vérification post-patch
- [ ] Le batching GraphQL est désactivé, ou une limite stricte (`maxBatchSize`) est appliquée et testée.
- [ ] Le rate limiter compte le nombre d'opérations GraphQL réellement exécutées, pas le nombre de requêtes HTTP.
- [ ] Les mutations sensibles (login, reset password) ont une limitation dédiée indépendante du batching.
- [ ] Un test confirme qu'un batch dépassant la limite est rejeté avec un code 400 avant exécution.
