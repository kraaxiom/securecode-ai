# Remédiation — Requêtes GraphQL profondément imbriquées

## Principe
Limiter la profondeur maximale des requêtes acceptées, combiner avec une analyse de coût de requête, et imposer une pagination obligatoire sur toute relation potentiellement récursive.

## PHP (webonyx/graphql-php)
```php
// Avant — vulnérable : aucune limite de profondeur, requête récursive (user->friends->friends->...) acceptée
$result = GraphQL::executeQuery($schema, $query, null, null, $variables);

// Après — sécurisé
use GraphQL\Validator\Rules\QueryDepth;
use GraphQL\Validator\DocumentValidator;

DocumentValidator::addRule(new QueryDepth(8));

$result = GraphQL::executeQuery($schema, $query, null, null, $variables);
```

## Node.js (graphql-depth-limit / Apollo Server)
```js
// Avant — vulnérable
const server = new ApolloServer({ typeDefs, resolvers });

// Après — sécurisé
import depthLimit from 'graphql-depth-limit';

const server = new ApolloServer({
  typeDefs,
  resolvers,
  validationRules: [depthLimit(8)],
});

// Pagination obligatoire sur une relation récursive
const typeDefs = gql`
  type User {
    id: ID!
    friends(first: Int = 10, after: String): FriendConnection!
  }
`;
```

## Checklist de vérification post-patch
- [ ] Une limite de profondeur maximale (ex: 8 niveaux) est appliquée à toutes les requêtes GraphQL avant exécution.
- [ ] Une analyse de complexité/coût de requête complète la limite de profondeur.
- [ ] Toute relation récursive du schéma (amis, catégories, commentaires) impose une pagination obligatoire.
- [ ] Un test confirme qu'une requête dépassant la profondeur maximale est rejetée avant d'atteindre les resolvers.
