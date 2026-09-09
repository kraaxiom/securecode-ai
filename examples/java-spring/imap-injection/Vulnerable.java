// Vulnérable : IMAP Injection (CWE-93)
// Le critère de recherche fourni par l'utilisateur est concaténé
// directement dans la commande IMAP SEARCH. Un attaquant peut injecter des
// guillemets et des mots-clés IMAP (ex: `" BODY "x" UID FETCH 1 (BODY[])`)
// pour modifier la commande exécutée sur le serveur de messagerie et
// accéder à des données ou des dossiers non autorisés.
package com.example.security.imap;

import jakarta.mail.*;
import jakarta.mail.search.SearchTerm;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.Properties;

@RestController
public class ImapInjectionService {

    @GetMapping("/api/mail/search")
    public Message[] search(@RequestParam String q, @RequestParam String folder) throws MessagingException {
        Properties props = new Properties();
        props.put("mail.store.protocol", "imaps");
        Session session = Session.getInstance(props);
        Store store = session.getStore("imaps");
        store.connect("imap.example.com", "user", "password");

        // Le nom de dossier utilisateur est utilisé tel quel : un attaquant
        // peut naviguer vers un dossier arbitraire.
        Folder mailFolder = store.getFolder(folder);
        mailFolder.open(Folder.READ_ONLY);

        // La commande de recherche brute est construite par concaténation :
        // les guillemets et retours chariot de "q" ne sont pas neutralisés.
        String rawCommand = "SUBJECT \"" + q + "\"";
        SearchTerm term = buildRawSearchTerm(rawCommand);
        return mailFolder.search(term);
    }

    private SearchTerm buildRawSearchTerm(String rawCommand) {
        // Représente la construction bas niveau d'une commande IMAP à
        // partir d'une chaîne non échappée (simplifié pour l'exemple).
        return null;
    }
}
