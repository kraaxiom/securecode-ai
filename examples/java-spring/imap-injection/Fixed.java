// Corrigé : IMAP Injection (CWE-93)
// Le nom de dossier est validé par liste blanche et la recherche utilise
// exclusivement l'API structurée de Jakarta Mail (SearchTerm typé), qui
// encode nativement les littéraux au lieu d'assembler une commande IMAP en
// texte brut. Aucune entrée utilisateur n'est jamais concaténée dans une
// commande IMAP.
package com.example.security.imap;

import jakarta.mail.*;
import jakarta.mail.search.SearchTerm;
import jakarta.mail.search.SubjectTerm;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

import java.util.Properties;
import java.util.Set;

@RestController
public class ImapInjectionService {

    // Liste blanche des dossiers accessibles : jamais de nom de dossier
    // construit dynamiquement depuis l'entrée utilisateur.
    private static final Set<String> ALLOWED_FOLDERS = Set.of("INBOX", "Sent", "Drafts");

    @GetMapping("/api/mail/search")
    public Message[] search(@RequestParam String q, @RequestParam String folder) throws MessagingException {
        if (!ALLOWED_FOLDERS.contains(folder)) {
            throw new ResponseStatusException(HttpStatus.BAD_REQUEST, "Dossier non autorisé");
        }

        // Longueur bornée et suppression défensive des caractères de
        // contrôle, même si l'API structurée ci-dessous échappe déjà
        // correctement le critère de recherche.
        String criteria = q.replace("\r", "").replace("\n", "");
        if (criteria.length() > 200) {
            criteria = criteria.substring(0, 200);
        }

        Properties props = new Properties();
        props.put("mail.store.protocol", "imaps");
        Session session = Session.getInstance(props);
        Store store = session.getStore("imaps");
        store.connect("imap.example.com", "user", "password");

        Folder mailFolder = store.getFolder(folder);
        mailFolder.open(Folder.READ_ONLY);

        // Utilisation de l'API structurée SearchTerm : le critère est
        // transmis comme valeur typée, jamais interpolé dans une commande
        // texte brute, ce qui empêche toute injection de mots-clés IMAP.
        SearchTerm term = new SubjectTerm(criteria);
        return mailFolder.search(term);
    }
}
