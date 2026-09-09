# Memory Poisoning (CWE-349)

La version vulnérable laisse le modèle extraire des « faits » de la conversation et les écrire directement dans la mémoire persistante de l'agent, réinjectée telle quelle comme contexte de confiance dans les sessions futures, sans confirmation de l'utilisateur ni cloisonnement entre utilisateurs. La correction exige une confirmation explicite de l'utilisateur avant toute écriture durable en mémoire, cloisonne strictement la mémoire par utilisateur, et retraite le contenu réinjecté comme une donnée à revalider plutôt que comme une instruction de confiance absolue.
