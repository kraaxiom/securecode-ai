// vulnerable.go - CWE-1427: Improper Neutralization of Input Used for LLM Prompting
// L'agent planifie des actions via le LLM puis execute directement chaque
// outil retourne par le modele, sans validation metier ni confirmation
// humaine, meme pour des actions a fort impact (paiement, suppression).
package agent

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ToolCall represente une action que le modele souhaite declencher.
type ToolCall struct {
	Name string          `json:"name"`
	Args json.RawMessage `json:"args"`
}

// PlanResponse est la reponse brute renvoyee par l'API du LLM.
type PlanResponse struct {
	Steps []ToolCall `json:"steps"`
}

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// RunAgent envoie l'instruction utilisateur au LLM puis execute chaque
// etape du plan retourne, sans aucun controle intermediaire.
func RunAgent(userInput string) error {
	plan, err := requestPlan(userInput)
	if err != nil {
		return err
	}

	// VULNERABLE : chaque outil propose par le modele est execute directement,
	// y compris les outils a fort impact (paiement, suppression de compte).
	// Aucune couche de validation independante du LLM, aucune confirmation
	// humaine, aucun cloisonnement de privileges.
	for _, step := range plan.Steps {
		if err := executeTool(step); err != nil {
			return fmt.Errorf("echec execution outil %s: %w", step.Name, err)
		}
	}
	return nil
}

// requestPlan interroge l'API du LLM pour obtenir un plan d'actions.
func requestPlan(userInput string) (*PlanResponse, error) {
	client := &http.Client{}
	_ = client
	_ = apiKey
	// Appel HTTP vers l'API du fournisseur LLM (details omis).
	// Le modele recoit la liste des outils disponibles (send_payment,
	// delete_account, send_email, execute_shell) et renvoie un plan.
	return &PlanResponse{Steps: []ToolCall{}}, nil
}

// executeTool execute directement l'outil demande par le modele, sans
// distinguer les outils sensibles des outils benins.
func executeTool(step ToolCall) error {
	switch step.Name {
	case "send_payment":
		return sendPayment(step.Args)
	case "delete_account":
		return deleteAccount(step.Args)
	case "send_email":
		return sendEmail(step.Args)
	default:
		return fmt.Errorf("outil inconnu: %s", step.Name)
	}
}

func sendPayment(args json.RawMessage) error  { return nil }
func deleteAccount(args json.RawMessage) error { return nil }
func sendEmail(args json.RawMessage) error     { return nil }
