# Command Injection (CWE-78)

La version vulnérable exécute `cmd.exe /c ping -n 3 {host}` en concaténant directement l'entrée utilisateur, ce qui permet d'injecter des métacaractères shell (`&`, `|`, `;`) pour exécuter des commandes arbitraires. La correction valide d'abord que `host` est une adresse IP valide via `IPAddress.TryParse`, puis invoque directement l'exécutable `ping` avec des arguments passés séparément dans `ArgumentList`, sans jamais passer par un shell.
