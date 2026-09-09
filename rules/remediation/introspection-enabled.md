# Remédiation — Introspection GraphQL activée en production

## Principe
Désactiver explicitement l'introspection et les interfaces exploratoires (GraphiQL, Playground) en production, avec une configuration différenciée par environnement.

## PHP (webonyx/graphql-php)
```php
// Avant — vulnérable : introspection toujours active, même en prod
$result = GraphQL::executeQuery($schema, $query, null, null, $variables);

// Après — sécurisé
use GraphQL\Validator\Rules\DisableIntrospection;
use GraphQL\Validator\DocumentValidator;

if (getenv('APP_ENV') === 'production') {
    DocumentValidator::addRule(new DisableIntrospection());
}

$result = GraphQL::executeQuery($schema, $query, null, null, $variables);
```

## Node.js (Apollo Server)
```js
// Avant — vulnérable : introspection et Playground actifs par défaut en prod
const server = new ApolloServer({ typeDefs, resolvers });

// Après — sécurisé
const isProd = process.env.NODE_ENV === 'production';

const server = new ApolloServer({
  typeDefs,
  resolvers,
  introspection: !isProd,
  plugins: isProd ? [ApolloServerPluginLandingPageDisabled()] : [],
});
```

## Checklist de vérification post-patch
- [ ] L'introspection est explicitement désactivée (`introspection: false` ou équivalent) en environnement de production.
- [ ] GraphiQL/Playground/Voyager ne sont pas accessibles publiquement en production.
- [ ] La configuration distingue clairement dev/staging/production, sans réglage de développement propagé par défaut.
- [ ] Un test confirme qu'une requête d'introspection (`__schema`) échoue en production.
