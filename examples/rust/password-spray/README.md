# Password Spraying

Le handler `login` ne limitait les échecs que par compte, ce qui ne détecte pas un attaquant testant un mot de passe courant contre de nombreux comptes distincts (CWE-307), technique contournant les verrouillages classiques. La version corrigée ajoute une agrégation globale des échecs par IP (`per_ip`), déclenche un blocage et une alerte de supervision lorsqu'un volume anormal est atteint tous comptes confondus, en complément de la limite par compte.
