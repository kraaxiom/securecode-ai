# expression-language-injection (CWE-917)

La version vulnerable concatene une condition fournie par l'utilisateur dans une expression evaluee dynamiquement par un moteur de regles, permettant l'injection de code d'expression. La version corrigee remplace l'evaluation dynamique par une allowlist de regles predefinies selectionnees par nom, supprimant tout interpreteur expose.
