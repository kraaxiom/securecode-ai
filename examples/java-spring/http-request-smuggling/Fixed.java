// Corrigé : HTTP Request Smuggling (CWE-444)
// Ce filtre applicatif rejette explicitement toute requête présentant à la
// fois les en-têtes Content-Length et Transfer-Encoding, éliminant
// l'ambiguïté de délimitation du corps qui permet le smuggling. Ce contrôle
// est une défense en profondeur : la configuration du proxy/frontal en
// amont doit également normaliser ou rejeter ces requêtes.
package com.example.security.smuggling;

import jakarta.servlet.*;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.stereotype.Component;

import java.io.IOException;

@Component
public class HttpRequestSmugglingConfig implements Filter {

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
            throws IOException, ServletException {
        HttpServletRequest httpRequest = (HttpServletRequest) request;
        HttpServletResponse httpResponse = (HttpServletResponse) response;

        boolean hasContentLength = httpRequest.getHeader("Content-Length") != null;
        boolean hasTransferEncoding = httpRequest.getHeader("Transfer-Encoding") != null;

        // Rejet explicite de toute requête ambiguë : la présence conjointe
        // de ces deux en-têtes est le signal classique d'une tentative de
        // request smuggling.
        if (hasContentLength && hasTransferEncoding) {
            httpResponse.sendError(HttpServletResponse.SC_BAD_REQUEST,
                    "Requête ambiguë refusée (Content-Length et Transfer-Encoding conjoints)");
            return;
        }

        chain.doFilter(httpRequest, response);
    }
}
