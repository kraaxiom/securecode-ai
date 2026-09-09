# Remédiation — Data Poisoning

## Principe
Valider la provenance et l'intégrité de chaque source de données utilisée pour l'entraînement ou le fine-tuning, détecter les anomalies statistiques avant intégration, et comparer les performances du modèle à un jeu de référence après chaque cycle d'entraînement.

## Python
```python
# Avant — vulnérable : ingestion directe de données non vérifiées dans le pipeline de fine-tuning
def build_training_set(sources):
    dataset = []
    for url in sources:
        dataset.extend(fetch_and_parse(url))  # aucune validation de provenance/intégrité
    return dataset

# Après — sécurisé : provenance vérifiée, détection d'anomalies, jeu de référence
TRUSTED_SOURCES = {"internal-labeled-set", "verified-partner-feed"}

def build_training_set(sources):
    dataset = []
    for source in sources:
        if source.id not in TRUSTED_SOURCES:
            audit_log.record("rejected_untrusted_source", source.id)
            continue
        verify_integrity(source, expected_hash=source.signed_hash)
        batch = fetch_and_parse(source.url)
        batch = filter_statistical_outliers(batch)  # détection d'anomalies
        dataset.extend(batch)
    return dataset

def validate_model_after_training(model, golden_set):
    score = evaluate(model, golden_set)
    if score < ACCEPTABLE_DRIFT_THRESHOLD:
        raise ModelDriftDetected(score)
```

## JavaScript / Node.js
```js
// Avant — vulnérable
async function buildTrainingSet(sources) {
  const dataset = [];
  for (const url of sources) {
    dataset.push(...await fetchAndParse(url)); // pas de vérification
  }
  return dataset;
}

// Après — sécurisé
const TRUSTED_SOURCES = new Set(["internal-labeled-set", "verified-partner-feed"]);

async function buildTrainingSet(sources) {
  const dataset = [];
  for (const source of sources) {
    if (!TRUSTED_SOURCES.has(source.id)) { auditLog.record("rejected_untrusted_source", source.id); continue; }
    verifyIntegrity(source, source.signedHash);
    let batch = await fetchAndParse(source.url);
    batch = filterStatisticalOutliers(batch);
    dataset.push(...batch);
  }
  return dataset;
}
```

## Checklist de vérification post-patch
- [ ] Chaque source de données d'entraînement est identifiée, signée/hashée, et vérifiée avant ingestion.
- [ ] Une étape de détection d'anomalies statistiques filtre les données aberrantes avant intégration au pipeline.
- [ ] Les données issues du feedback utilisateur passent par un échantillonnage humain avant réintégration à l'entraînement.
- [ ] Un jeu de référence (golden set) fixe est utilisé pour détecter une dérive de comportement après chaque cycle d'entraînement.
