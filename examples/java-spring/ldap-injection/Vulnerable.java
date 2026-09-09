// Vulnérable : LDAP Injection (CWE-90)
// Le filtre de recherche LDAP est construit par concaténation directe d'une
// entrée utilisateur non échappée. Un attaquant peut injecter des méta-caractères
// LDAP (*, (, ), \, NUL) pour modifier la logique du filtre et contourner
// l'authentification ou extraire des entrées non autorisées de l'annuaire.
package com.example.security.ldap;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.naming.NamingEnumeration;
import javax.naming.directory.DirContext;
import javax.naming.directory.SearchControls;
import javax.naming.directory.SearchResult;

@RestController
public class LdapInjectionService {

    private final DirContext ldapContext;
    private static final String BASE_DN = "ou=users,dc=example,dc=com";

    public LdapInjectionService(DirContext ldapContext) {
        this.ldapContext = ldapContext;
    }

    @GetMapping("/api/ldap/search")
    public boolean findUser(@RequestParam String uid) throws Exception {
        // Concaténation directe : uid = "*)(uid=*))(|(uid=*" permet de contourner le filtre.
        String filter = "(uid=" + uid + ")";

        SearchControls controls = new SearchControls();
        controls.setSearchScope(SearchControls.SUBTREE_SCOPE);

        NamingEnumeration<SearchResult> results = ldapContext.search(BASE_DN, filter, controls);
        return results.hasMore();
    }
}
