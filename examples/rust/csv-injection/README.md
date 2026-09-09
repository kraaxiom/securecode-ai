# CSV Injection (CWE-1236)

Le code vulnérable écrit les champs utilisateur bruts dans le CSV exporté sans vérifier s'ils commencent par un caractère déclencheur de formule (`=`, `+`, `-`, `@`), ce qui peut provoquer l'exécution d'une formule à l'ouverture dans un tableur. La correction ajoute une fonction `sanitize_csv_cell` qui préfixe ces cellules d'une apostrophe neutralisante avant écriture, conformément à la remédiation CWE-1236.
