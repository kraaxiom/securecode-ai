// fixed.go - Correction CWE-79
// On remplace la sanitisation par regex par un vrai parseur HTML
// (golang.org/x/net/html via bluemonday), qui construit un arbre DOM
// conforme aux regles de parsing du navigateur avant d'appliquer une
// politique d'allowlist stricte. Cette approche est resistante aux
// vecteurs de mutation qui exploitent les divergences entre une
// analyse textuelle naive et le parsing reel du navigateur.
package richtext

import "github.com/microcosm-cc/bluemonday"

var strictPolicy = bluemonday.StrictPolicy() // aucune balise HTML autorisee

// SanitizeStrict s'appuie sur un vrai parseur HTML + allowlist.
func SanitizeStrict(userHTML string) string {
	// FIXED: sanitisation basee sur un arbre DOM, pas sur une regex textuelle
	return strictPolicy.Sanitize(userHTML)
}

// SanitizeRichText autorise un sous-ensemble limite de mise en forme,
// toujours via le meme parseur DOM conforme.
func SanitizeRichText(userHTML string) string {
	policy := bluemonday.UGCPolicy()
	return policy.Sanitize(userHTML)
}
