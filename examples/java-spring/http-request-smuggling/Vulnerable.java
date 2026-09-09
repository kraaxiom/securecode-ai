// Vulnérable : HTTP Request Smuggling (CWE-444)
// Ce filtre applicatif fait confiance à la requête telle que transmise par
// la chaîne de proxys sans vérifier la cohérence des en-têtes de
// délimitation du corps (Content-Length / Transfer-Encoding). Un frontal et
// un backend interprétant différemment une requête ambiguë permettent de
// "contrebander" une seconde requête cachée dans le corps de la première.
package com.example.security.smuggling;

import jakarta.servlet.*;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.stereotype.Component;

import java.io.IOException;

@Component
public class HttpRequestSmugglingConfig implements Filter {

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
            throws IOException, ServletException {
        HttpServletRequest httpRequest = (HttpServletRequest) request;

        // Aucune vérification : la requête est transmise telle quelle au
        // reste de la chaîne, même si elle contient à la fois
        // Content-Length et Transfer-Encoding (ambiguïté classique de
        // smuggling CL.TE / TE.CL).
        chain.doFilter(httpRequest, response);
    }
}
