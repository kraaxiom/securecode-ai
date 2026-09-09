// vulnerable.go - CWE-79: Mutation-based Cross-Site Scripting (mXSS)
// Le HTML utilisateur est "sanitise" par une suppression naive de
// balises via une expression reguliere. Ce type d'approche ne tient
// pas compte des regles de parsing du navigateur (auto-fermeture de
// balises, normalisation d'attributs) : un payload structure de facon
// a "muter" une fois reparse par le navigateur (ex: dans une balise
// <noscript> ou <template> imbriquee) peut recomposer un vecteur actif
// apres coup, alors qu'il semblait neutralise avant stockage.
package richtext

import "regexp"

var scriptTagPattern = regexp.MustCompile(`(?i)<script[^>]*>.*?</script>`)

// SanitizeNaive tente de retirer les balises <script> par regex.
func SanitizeNaive(userHTML string) string {
	// VULNERABLE: une regex ne "comprend" pas la grammaire HTML et ne
	// couvre pas les vecteurs de mutation lies au reparsing navigateur
	return scriptTagPattern.ReplaceAllString(userHTML, "")
}
