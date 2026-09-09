// CWE-349 : Acceptance of Extraneous Untrusted Data With Trust
// Les retours utilisateurs (feedback loop) sont réintégrés directement dans
// le jeu de données de fine-tuning, sans validation de provenance, sans
// détection d'anomalies et sans comparaison à un jeu de référence, ouvrant
// la voie à l'empoisonnement du modèle.

use std::fs::OpenOptions;
use std::io::Write;

struct ExempleEntrainement {
    prompt: String,
    completion_attendue: String,
}

fn integrer_feedback_utilisateur(
    chemin_dataset: &str,
    prompt: &str,
    correction_utilisateur: &str,
) -> std::io::Result<()> {
    // Vulnérable : toute correction soumise par un utilisateur est ajoutée
    // telle quelle au jeu d'entraînement, sans contrôle de provenance ni
    // détection d'anomalie statistique, ni échantillonnage humain.
    let exemple = ExempleEntrainement {
        prompt: prompt.to_string(),
        completion_attendue: correction_utilisateur.to_string(),
    };

    let mut fichier = OpenOptions::new().append(true).create(true).open(chemin_dataset)?;
    writeln!(fichier, "{{\"prompt\":\"{}\",\"completion\":\"{}\"}}", exemple.prompt, exemple.completion_attendue)?;

    Ok(())
}

fn lancer_cycle_finetuning(chemin_dataset: &str) {
    // Vulnérable : aucun jeu de référence (golden set) n'est utilisé pour
    // détecter une dérive de comportement après ce cycle d'entraînement.
    println!("Lancement du fine-tuning sur {chemin_dataset} sans validation préalable.");
}
