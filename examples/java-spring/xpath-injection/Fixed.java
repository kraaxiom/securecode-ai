// Corrigé : XPath Injection (CWE-643)
// L'expression XPath utilise des variables liées via un XPathVariableResolver
// au lieu de concaténer les identifiants dans la chaîne. Les valeurs
// utilisateur ne peuvent plus modifier la structure de l'expression.
// À noter : XPath reste déconseillé comme mécanisme d'authentification ;
// préférer une base de données avec hachage de mot de passe (ex: bcrypt/argon2).
package com.example.security.xpathinjection;

import org.springframework.stereotype.Service;
import org.w3c.dom.Document;

import javax.xml.namespace.QName;
import javax.xml.xpath.XPath;
import javax.xml.xpath.XPathExpression;
import javax.xml.xpath.XPathFactory;

@Service
public class XpathInjectionService {

    private final Document usersDocument;

    public XpathInjectionService(Document usersDocument) {
        this.usersDocument = usersDocument;
    }

    public boolean authenticate(String username, String password) throws Exception {
        XPath xpath = XPathFactory.newInstance().newXPath();

        // Résolveur de variables liées : les valeurs ne sont jamais insérées dans le texte de l'expression
        xpath.setXPathVariableResolver(variableName -> {
            if ("user".equals(variableName.getLocalPart())) {
                return username;
            }
            if ("pass".equals(variableName.getLocalPart())) {
                return password;
            }
            return null;
        });

        XPathExpression xpe = xpath.compile("//user[username=$user and password=$pass]");
        Object result = xpe.evaluate(usersDocument, javax.xml.xpath.XPathConstants.NODE);
        return result != null;
    }
}
