// CWE-798 : correction — aucun secret en dur dans le code. Les identifiants
// sont résolus au démarrage depuis un gestionnaire de secrets (ici illustré
// via variables d'environnement injectées par Vault/Secrets Manager), et ne
// sont jamais journalisés en clair.

use aws_sdk_s3::Client;
use aws_sdk_s3::config::Region;
use std::env;

struct ConfigurationSensible {
    stripe_api_key: String,
    database_url: String,
}

impl ConfigurationSensible {
    // Sécurisé : les secrets sont chargés depuis l'environnement d'exécution
    // (lui-même alimenté par un gestionnaire de secrets, pas committé),
    // avec échec explicite si absent plutôt qu'une valeur par défaut.
    fn charger() -> Result<Self, env::VarError> {
        Ok(Self {
            stripe_api_key: env::var("STRIPE_API_KEY")?,
            database_url: env::var("DATABASE_URL")?,
        })
    }
}

impl std::fmt::Debug for ConfigurationSensible {
    // Sécurisé : implémentation Debug personnalisée qui masque les secrets,
    // pour éviter toute fuite accidentelle via `{:?}` dans un log.
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("ConfigurationSensible")
            .field("stripe_api_key", &"***REDACTED***")
            .field("database_url", &"***REDACTED***")
            .finish()
    }
}

async fn construire_client_s3() -> Client {
    // Sécurisé : aucune credential codée en dur. Le SDK résout la chaîne de
    // credentials standard (rôle IAM d'instance/tâche, SSO, variables
    // d'environnement injectées par le gestionnaire de secrets).
    let config = aws_config::defaults(aws_config::BehaviorVersion::latest())
        .region(Region::new("eu-west-3"))
        .load()
        .await;

    Client::new(&config)
}

fn journaliser_configuration(config: &ConfigurationSensible) {
    // Sécurisé : on journalise la représentation masquée, jamais les valeurs
    // réelles des secrets.
    println!("Configuration chargée : {config:?}");
}
