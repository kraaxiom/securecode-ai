// Vulnérable : XPath Injection (CWE-643)
// L'expression XPath utilisée pour l'authentification est construite par
// concaténation directe des identifiants fournis par l'utilisateur. Un
// attaquant peut injecter une syntaxe XPath (ex: "' or '1'='1") pour
// contourner l'authentification.
package com.example.security.xpathinjection;

import org.springframework.stereotype.Service;
import org.w3c.dom.Document;

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

        // Concaténation directe des identifiants dans l'expression XPath
        String expr = "//user[username='" + username + "' and password='" + password + "']";
        XPathExpression xpe = xpath.compile(expr);

        Object result = xpe.evaluate(usersDocument, javax.xml.xpath.XPathConstants.NODE);
        return result != null;
    }
}
