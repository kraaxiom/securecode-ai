// Corrigé : LDAP Injection (CWE-90)
// Le filtre est construit via l'API Spring LDAP (EqualsFilter), qui échappe
// automatiquement les caractères spéciaux du filtre selon RFC 4515
// (*, (, ), \, NUL), empêchant toute modification de la structure du filtre.
package com.example.security.ldap;

import org.springframework.ldap.core.LdapTemplate;
import org.springframework.ldap.filter.EqualsFilter;
import org.springframework.ldap.filter.Filter;
import org.springframework.ldap.query.LdapQuery;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import static org.springframework.ldap.query.LdapQueryBuilder.query;

@RestController
public class LdapInjectionService {

    private final LdapTemplate ldapTemplate;

    public LdapInjectionService(LdapTemplate ldapTemplate) {
        this.ldapTemplate = ldapTemplate;
    }

    @GetMapping("/api/ldap/search")
    public boolean findUser(@RequestParam String uid) {
        // Validation de format en amont : identifiant alphanumérique attendu.
        if (!uid.matches("^[a-zA-Z0-9._-]{1,64}$")) {
            throw new IllegalArgumentException("Identifiant LDAP invalide");
        }

        // EqualsFilter échappe automatiquement les caractères spéciaux du filtre.
        Filter filter = new EqualsFilter("uid", uid);

        LdapQuery ldapQuery = query()
                .base("ou=users,dc=example,dc=com")
                .filter(filter.encode());

        return !ldapTemplate.search(ldapQuery, ctx -> ctx).isEmpty();
    }
}
