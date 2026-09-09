# RAG Poisoning (CWE-349)

La version vulnérable indexe automatiquement dans la base vectorielle tout document reçu, sans vérifier sa provenance, sans liste blanche de sources de confiance ni contrôle du contenu, ce qui permet à un contenu malveillant d'être ensuite récupéré et traité comme une source fiable par le modèle. La correction valide la source contre une liste blanche autorisée, applique une analyse de contenu et un score de confiance avant indexation, et conserve la traçabilité de la provenance de chaque document.
