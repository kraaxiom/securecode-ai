# Remédiation — Déni de service par requête GraphQL coûteuse

## Principe
Plafonner strictement toute valeur de pagination fournie par le client, mettre en place une analyse de complexité de requête avec un budget de coût maximal, et ajouter des timeouts d'exécution.

## PHP (webonyx/graphql-php)
```php
// Avant — vulnérable : le client contrôle "first" sans plafond
function resolveItems($root, array $args) {
    return Item::query()->limit($args['first'])->get(); // first peut valoir 1000000
}

// Après — sécurisé
use GraphQL\Validator\Rules\QueryComplexity;

DocumentValidator::addRule(new QueryComplexity(300));

const MAX_PAGE_SIZE = 100;

function resolveItems($root, array $args) {
    $first = min((int) ($args['first'] ?? 20), MAX_PAGE_SIZE);
    return Item::query()->limit($first)->get();
}
```

## Node.js (graphql-query-complexity / Apollo Server)
```js
// Avant — vulnérable
const resolvers = {
  Query: {
    items: (_, { first }) => db.item.findMany({ take: first }), // first non plafonné
  },
};

// Après — sécurisé
import { createComplexityLimitRule } from 'graphql-validation-complexity';

const MAX_PAGE_SIZE = 100;

const resolvers = {
  Query: {
    items: (_, { first = 20 }) => {
      const take = Math.min(first, MAX_PAGE_SIZE);
      return db.item.findMany({ take });
    },
  },
};

const server = new ApolloServer({
  typeDefs,
  resolvers,
  validationRules: [createComplexityLimitRule(1000)],
  plugins: [{
    requestDidStart: () => ({
      async executionDidStart() {
        return { willResolveField: () => { /* timeout global géré en amont */ } };
      },
    }),
  }],
});
```

## Checklist de vérification post-patch
- [ ] Tout argument de pagination (`first`, `limit`, `take`) est plafonné côté serveur indépendamment de la valeur envoyée par le client.
- [ ] Une bibliothèque d'analyse de coût de requête est configurée avec un budget maximal par requête.
- [ ] Un timeout d'exécution global est appliqué à chaque requête GraphQL.
- [ ] Un test confirme qu'une requête demandant une pagination excessive est plafonnée sans erreur serveur.
