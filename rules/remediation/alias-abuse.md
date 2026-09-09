# Remédiation — Abus d'alias GraphQL

## Principe
Plafonner strictement le nombre d'alias/champs autorisés par requête et faire porter le rate limiting sur le coût réel de la requête plutôt que sur le nombre de requêtes HTTP.

## PHP (webonyx/graphql-php)
```php
// Avant — vulnérable : aucune limite sur le nombre d'alias, permet des centaines de tentatives de login en 1 requête
$schema = BuildSchema::build($sdl);
$result = GraphQL::executeQuery($schema, $query, null, null, $variables);

// Après — sécurisé
use GraphQL\Validator\Rules\QueryComplexity;
use GraphQL\Validator\Rules\DisableIntrospection;

DocumentValidator::addRule(new QueryComplexity(200));

$aliasCount = substr_count($query, ': ') + substr_count($query, ':');
if ($aliasCount > 20) {
    throw new UserError('Trop d\'alias dans la requête.');
}

$result = GraphQL::executeQuery($schema, $query, null, null, $variables);
```

## Node.js (graphql-armor / Apollo Server)
```js
// Avant — vulnérable : pas de limite d'alias, un attaquant envoie 500 alias de login en une requête
const server = new ApolloServer({ typeDefs, resolvers });

// Après — sécurisé
import { maxAliasesPlugin } from '@escape.tech/graphql-armor-max-aliases';

const server = new ApolloServer({
  typeDefs,
  resolvers,
  plugins: [maxAliasesPlugin({ n: 15 })],
});

// Limitation dédiée sur le resolver de login, indépendante du transport
const loginRateLimiter = new RateLimiterMemory({ points: 5, duration: 60 });
async function loginResolver(_, { email, password }, ctx) {
  await loginRateLimiter.consume(ctx.ip + ':' + email);
  return authService.login(email, password);
}
```

## Checklist de vérification post-patch
- [ ] Une limite explicite du nombre d'alias/champs par requête est configurée sur le serveur GraphQL.
- [ ] Le rate limiting sur les resolvers sensibles (login, OTP, reset password) est indépendant du nombre de requêtes HTTP.
- [ ] Un test confirme qu'une requête avec plus d'alias que la limite est rejetée avant exécution des resolvers.
- [ ] Les tentatives d'authentification via alias multiples sont journalisées et déclenchent une alerte au-delà d'un seuil.
