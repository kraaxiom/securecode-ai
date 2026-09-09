# OS Command Injection (CWE-78)

La version vulnérable exécute `cmd.exe /c 7z a archive.zip {fileName}` en concaténant directement l'entrée utilisateur dans les arguments passés au shell, ce qui permet d'injecter des méta-caractères (`&`, `|`, `;`) pour exécuter des commandes arbitraires. La correction valide `fileName` avec une liste blanche stricte (caractères alphanumériques, `_`, `-`, `.` uniquement) puis invoque directement l'exécutable `7z` avec des arguments passés séparément via `ArgumentList`, sans interprétation shell.
