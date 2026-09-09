# Remédiation — RAG Poisoning

## Principe
Restreindre et authentifier les sources autorisées à alimenter la base documentaire, scanner les documents entrants pour détecter des instructions cachées avant indexation, et conserver la traçabilité de la provenance jusqu'à la réponse générée.

## Python
```python
# Avant — vulnérable : indexation directe de documents crawlés sans validation
def ingest_documents(urls):
    for url in urls:
        doc = crawl(url)
        chunks = split_into_chunks(doc)
        for chunk in chunks:
            vector_store.upsert(embed(chunk), metadata={"text": chunk})  # pas de provenance ni de scan

# Après — sécurisé : allowlist de sources + scan de contenu suspect + traçabilité
TRUSTED_SOURCES = {"docs.internal.acme.com", "verified-partner-wiki.acme.com"}

def ingest_documents(urls):
    for url in urls:
        domain = extract_domain(url)
        if domain not in TRUSTED_SOURCES:
            audit_log.record("rejected_untrusted_source", url)
            continue

        doc = crawl(url)
        if content_scanner.contains_suspicious_instructions(doc):  # motifs impératifs cachés
            audit_log.record("flagged_suspicious_document", url)
            continue

        chunks = split_into_chunks(doc)
        for chunk in chunks:
            vector_store.upsert(
                embed(chunk),
                metadata={"text": chunk, "source_url": url, "ingested_at": now(), "trust_level": "verified"},
            )

def answer_query(query):
    results = vector_store.search(embed(query))
    response = llm.chat(messages=build_prompt(query, results))
    return response, [r.metadata["source_url"] for r in results]  # traçabilité jusqu'à la réponse
```

## JavaScript / Node.js
```js
// Avant — vulnérable
async function ingestDocuments(urls) {
  for (const url of urls) {
    const doc = await crawl(url);
    const chunks = splitIntoChunks(doc);
    for (const chunk of chunks) {
      await vectorStore.upsert(await embed(chunk), { text: chunk });
    }
  }
}

// Après — sécurisé
const TRUSTED_SOURCES = new Set(["docs.internal.acme.com", "verified-partner-wiki.acme.com"]);

async function ingestDocuments(urls) {
  for (const url of urls) {
    const domain = extractDomain(url);
    if (!TRUSTED_SOURCES.has(domain)) { auditLog.record("rejected_untrusted_source", url); continue; }

    const doc = await crawl(url);
    if (await contentScanner.containsSuspiciousInstructions(doc)) {
      auditLog.record("flagged_suspicious_document", url);
      continue;
    }

    const chunks = splitIntoChunks(doc);
    for (const chunk of chunks) {
      await vectorStore.upsert(await embed(chunk), { text: chunk, sourceUrl: url, ingestedAt: Date.now(), trustLevel: "verified" });
    }
  }
}
```

## Checklist de vérification post-patch
- [ ] Seules des sources authentifiées et listées en allowlist alimentent la base documentaire.
- [ ] Chaque document entrant est scanné pour détecter des motifs d'instructions impératives cachées avant indexation.
- [ ] La provenance de chaque chunk indexé est conservée et exposée jusqu'à la réponse générée (traçabilité/audit).
- [ ] Une procédure de retrait rapide existe pour désindexer un document identifié comme malveillant.
