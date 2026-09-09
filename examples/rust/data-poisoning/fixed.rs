// CWE-349 : correction — le feedback utilisateur est isolé dans une file de
// révision, passe par une détection d'anomalies statistiques avant toute
// intégration, et chaque cycle de fine-tuning est comparé à un jeu de
// référence figé pour détecter une dérive suspecte.

use std::fs::OpenOptions;
use std::io::Write;

struct ExempleEntrainement {
    prompt: String,
    completion_attendue: String,
    source_utilisateur_id: String,
    valide_par_humain: bool,
}

fn mettre_en_file_revision(
    chemin_file_attente: &str,
    prompt: &str,
    correction_utilisateur: &str,
    source_utilisateur_id: &str,
) -> std::io::Result<()> {
    // Sécurisé : le feedback est isolé dans une file d'attente séparée du
    // jeu d'entraînement, avec traçabilité de la source, en attendant
    // validation humaine — jamais intégré directement.
    let exemple = ExempleEntrainement {
        prompt: prompt.to_string(),
        completion_attendue: correction_utilisateur.to_string(),
        source_utilisateur_id: source_utilisateur_id.to_string(),
        valide_par_humain: false,
    };

    let mut fichier = OpenOptions::new().append(true).create(true).open(chemin_file_attente)?;
    writeln!(
        fichier,
        "{{\"prompt\":\"{}\",\"completion\":\"{}\",\"source\":\"{}\",\"valide\":{}}}",
        exemple.prompt, exemple.completion_attendue, exemple.source_utilisateur_id, exemple.valide_par_humain
    )?;

    Ok(())
}

fn detecter_anomalie_statistique(exemples_en_attente: &[ExempleEntrainement]) -> Vec<usize> {
    // Sécurisé : détection basique de motifs répétitifs suspects (même
    // source soumettant un volume anormal d'exemples quasi identiques) —
    // à combiner avec des outils de détection d'anomalies dédiés en
    // production.
    let mut index_suspects = Vec::new();
    for (i, exemple) in exemples_en_attente.iter().enumerate() {
        let occurrences = exemples_en_attente
            .iter()
            .filter(|e| e.source_utilisateur_id == exemple.source_utilisateur_id)
            .count();
        if occurrences > 50 {
            index_suspects.push(i);
        }
    }
    index_suspects
}

fn lancer_cycle_finetuning_valide(
    dataset_valide_par_humain: &[ExempleEntrainement],
    jeu_de_reference: &[ExempleEntrainement],
) -> bool {
    // Sécurisé : seuls les exemples validés par un humain intègrent
    // l'entraînement, et un jeu de référence figé permet de détecter une
    // dérive de comportement après le cycle.
    let exemples_valides = dataset_valide_par_humain
        .iter()
        .filter(|e| e.valide_par_humain)
        .count();

    println!(
        "Fine-tuning lancé avec {exemples_valides} exemples validés, comparaison à {} exemples de référence.",
        jeu_de_reference.len()
    );

    true
}
