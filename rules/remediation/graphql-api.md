# Remédiation — Sécurité des API GraphQL

## Principe
Désactiver l'introspection en production, limiter la profondeur/complexité des requêtes, appliquer l'autorisation au niveau de chaque résolveur, et ne jamais renvoyer d'erreurs détaillées au client.

## Node.js (Apollo Server)
```js
// Avant — vulnérable
const server = new ApolloServer({
  typeDefs,
  resolvers,
  introspection: true, // actif en production
});

// Après — sécurisé
const server = new ApolloServer({
  typeDefs,
  resolvers,
  introspection: process.env.NODE_ENV !== 'production',
  validationRules: [depthLimit(5), createComplexityLimitRule(1000)],
  formatError: (err) => {
    // Ne jamais exposer la stack trace / détails internes
    return { message: 'Une erreur est survenue.', code: err.extensions?.code };
  },
});
```

## Résolveur avec autorisation au niveau champ (Node.js)
```js
// Avant — vulnérable
const resolvers = {
  User: {
    ssn: (parent) => parent.ssn, // aucun contrôle d'accès
  },
};

// Après — sécurisé
const resolvers = {
  User: {
    ssn: (parent, args, context) => {
      if (!context.user || context.user.id !== parent.id) {
        throw new ForbiddenError('Accès non autorisé');
      }
      return parent.ssn;
    },
  },
};
```

## Python (Graphene / Django)
```python
# Avant — vulnérable
schema = graphene.Schema(query=Query)  # introspection active par défaut

# Après — sécurisé
from graphql import validate, parse
from graphql.validation import NoSchemaIntrospectionCustomRule

# désactiver l'introspection en production via une règle de validation personnalisée
validation_rules = [NoSchemaIntrospectionCustomRule] if not settings.DEBUG else []
```

## Checklist de vérification post-patch
- [ ] L'introspection GraphQL est désactivée (ou restreinte à un usage interne authentifié) en production.
- [ ] Une limite de profondeur et de complexité de requête est configurée.
- [ ] Chaque résolveur de champ sensible applique son propre contrôle d'autorisation, pas seulement la requête racine.
- [ ] Le formatteur d'erreurs personnalisé ne renvoie jamais de stack trace ni de détails d'implémentation.
- [ ] Un test confirme qu'une requête profondément imbriquée est rejetée avant exécution.
