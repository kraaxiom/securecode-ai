# Remédiation — IMAP Injection

## Principe
Ne jamais construire une commande IMAP par concaténation de chaînes à partir d'une entrée utilisateur. Utiliser exclusivement les fonctions/paramètres fournis par le client IMAP de la bibliothèque standard ou d'un framework, qui encodent correctement les littéraux (préfixage par longueur ou guillemets échappés) au lieu d'assembler la commande à la main.

## PHP (ext-imap)
```php
// Avant — vulnérable
$folder = $_GET['folder'];
$criteria = $_GET['q'];
$search = "SUBJECT \"$criteria\"";
$emails = imap_search($mbox, $search);
imap_reopen($mbox, "{imap.example.com:993/imap/ssl}" . $folder);

// Après — sécurisé
// 1) Le critère de recherche passe par imap_utf8() / échappement des guillemets
//    et n'accepte que du texte libre, jamais des mots-clés de commande.
$criteria = str_replace(['"', "\r", "\n"], '', $_GET['q']);
$emails = imap_search($mbox, 'SUBJECT "' . $criteria . '"');

// 2) Le nom de dossier est validé contre une liste blanche connue
//    (jamais construit dynamiquement depuis l'entrée utilisateur).
$allowedFolders = ['INBOX', 'Sent', 'Drafts'];
$folder = in_array($_GET['folder'], $allowedFolders, true) ? $_GET['folder'] : 'INBOX';
imap_reopen($mbox, "{imap.example.com:993/imap/ssl}" . $folder);
```

## Node.js (imapflow)
```js
// Avant — vulnérable
const term = req.query.q;
await client.mailboxOpen('INBOX');
const results = await client.search({ subject: term }); // OK seulement si la lib échappe déjà

// mais danger typique : construction manuelle d'une commande brute
connection.exec(`SEARCH SUBJECT "${term}"`); // vulnérable si connection bas niveau

// Après — sécurisé
// Utiliser exclusivement l'API structurée de la bibliothèque (imapflow, node-imap)
// qui encode les littéraux ; ne jamais construire de commande IMAP en texte brut.
const term = String(req.query.q ?? '').slice(0, 200);
await client.mailboxOpen('INBOX');
const results = await client.search({ subject: term });
```

## Python (imaplib)
```python
# Avant — vulnérable
criteria = request.args.get("q")
typ, data = mail.search(None, f'SUBJECT "{criteria}"')

# Après — sécurisé
# imaplib échappe automatiquement les arguments passés séparément à search();
# ne jamais interpoler l'entrée utilisateur dans une chaîne de commande.
criteria = request.args.get("q", "")
criteria = criteria.replace("\r", "").replace("\n", "")[:200]
typ, data = mail.search(None, "SUBJECT", criteria)
```

## Checklist de vérification post-patch
- [ ] Aucune commande IMAP n'est construite par concaténation directe d'une chaîne contenant une entrée utilisateur.
- [ ] Les noms de dossier IMAP proviennent d'une liste blanche, jamais directement de l'entrée utilisateur.
- [ ] Les caractères `"`, `\r` et `\n` sont neutralisés ou impossibles à injecter dans tout critère de recherche transmis.
- [ ] La bibliothèque cliente IMAP utilisée encode nativement les littéraux (arguments passés séparément, pas en texte concaténé).
- [ ] Un test confirme qu'une entrée contenant des guillemets ou des mots-clés IMAP (`SEARCH`, `LOGIN`, `\r\n`) ne modifie pas la commande exécutée.
