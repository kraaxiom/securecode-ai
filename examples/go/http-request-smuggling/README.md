# http-request-smuggling (CWE-444)

La version vulnerable transfere aveuglement les en-tetes `Content-Length` et `Transfer-Encoding` du client vers le backend, permettant une desynchronisation exploitable entre le proxy et le serveur. La version corrigee rejette toute requete presentant les deux en-tetes simultanement, conformement a la RFC 7230.
