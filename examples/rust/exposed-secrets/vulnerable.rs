// CWE-798 : Use of Hard-coded Credentials
// La clé API et la chaîne de connexion base de données sont codées en dur
// dans le source, ce qui les expose dans l'historique Git dès le premier
// commit et à quiconque a accès au code (public ou non).

use aws_sdk_s3::Client;
use aws_sdk_s3::config::Credentials;

// Vulnérable : secrets en clair dans le code source.
const STRIPE_API_KEY: &str = "sk_live_EXAMPLE_NOT_A_REAL_KEY";
const DATABASE_URL: &str = "postgres://admin:SuperSecret123!@db.prod.internal:5432/app";

fn construire_client_s3() -> Client {
    // Vulnérable : identifiants AWS codés en dur au lieu d'être résolus via
    // un provider de credentials (rôle IAM, gestionnaire de secrets).
    let credentials = Credentials::new(
        "AKIAIOSFODNN7EXAMPLE",
        "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
        None,
        None,
        "hardcoded",
    );

    aws_sdk_s3::Config::builder()
        .credentials_provider(credentials)
        .region(aws_sdk_s3::config::Region::new("eu-west-3"))
        .build()
        .into()
}

fn journaliser_configuration() {
    // Vulnérable : le secret apparaît en clair dans les logs applicatifs.
    println!("Connexion DB avec : {DATABASE_URL}");
    println!("Clé Stripe active : {STRIPE_API_KEY}");
}
