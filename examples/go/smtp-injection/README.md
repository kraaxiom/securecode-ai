# smtp-injection (CWE-93)

La version vulnerable concatene le nom et le sujet fournis par l'utilisateur directement dans les en-tetes du message brut envoye via `smtp.SendMail` : une entree contenant CRLF permet d'injecter des en-tetes supplementaires (Bcc, Cc) ou de terminer les en-tetes pour injecter un corps de message arbitraire. La version corrigee rejette explicitement toute entree contenant CR ou LF avant de l'inserer dans les en-tetes.
