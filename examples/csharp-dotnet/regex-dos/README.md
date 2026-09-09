# ReDoS - Regular Expression Denial of Service (CWE-1333)

La version vulnérable utilise la regex `^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$`, dont les quantificateurs imbriqués `(x+)+` provoquent un backtracking exponentiel sur certaines entrées, appliquée sans limite de taille ni timeout. La correction réécrit la regex sans ambiguïté de correspondance, impose une limite de longueur (254 caractères) avant tout test, et configure un `TimeSpan` de timeout sur le `Regex` .NET afin qu'une entrée pathologique déclenche une `RegexMatchTimeoutException` gérée plutôt que de bloquer le thread.
