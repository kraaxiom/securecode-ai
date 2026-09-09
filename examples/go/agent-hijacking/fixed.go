// fixed.go - Correctif CWE-1427
// La couche de decision du LLM est separee de la couche d'execution.
// Toute action a fort impact exige une validation metier independante
// du modele et une confirmation humaine explicite avant execution.
// Chaque decision et resultat est journalise pour l'audit.
package agent

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ToolCall struct {
	Name string          `json:"name"`
	Args json.RawMessage `json:"args"`
}

type PlanResponse struct {
	Steps []ToolCall `json:"steps"`
}

const apiKey = "REPLACE_WITH_YOUR_API_KEY"

// highImpactTools liste les outils qui exigent une confirmation humaine
// et une validation metier avant toute execution.
var highImpactTools = map[string]bool{
	"send_payment":   true,
	"delete_account": true,
}

// AuditLogger journalise chaque decision et action de l'agent pour
// permettre l'audit et la detection d'anomalies post-incident.
type AuditLogger interface {
	Record(event string, step ToolCall)
}

// HumanConfirmer demande une confirmation humaine explicite avant
// l'execution d'une action irreversible ou a fort impact.
type HumanConfirmer interface {
	Confirm(step ToolCall) bool
}

// BusinessValidator applique des controles metier independants du modele
// (limites, coherence des arguments, permissions) sur chaque action.
type BusinessValidator interface {
	Validate(step ToolCall) error
}

type Agent struct {
	Audit     AuditLogger
	Confirmer HumanConfirmer
	Validator BusinessValidator
	// UntrustedContext indique si l'agent traite actuellement du contenu
	// externe non fiable (page web, document) : dans ce cas ses privileges
	// sont reduits automatiquement.
	UntrustedContext bool
}

// RunAgent envoie l'instruction utilisateur au LLM puis controle chaque
// etape avant execution, avec confirmation humaine pour les actions
// irreversibles et journalisation systematique.
func (a *Agent) RunAgent(userInput string) error {
	plan, err := requestPlan(userInput)
	if err != nil {
		return err
	}

	for _, step := range plan.Steps {
		// Cloisonnement : privileges reduits lors du traitement de contenu
		// externe non fiable, les outils a fort impact sont bloques d'office.
		if a.UntrustedContext && highImpactTools[step.Name] {
			a.Audit.Record("blocked_high_impact_untrusted_context", step)
			continue
		}

		if highImpactTools[step.Name] {
			if !a.Confirmer.Confirm(step) {
				a.Audit.Record("blocked_high_impact_no_confirmation", step)
				continue
			}
		}

		// Validation metier independante de ce que le modele a produit,
		// quel que soit l'outil.
		if err := a.Validator.Validate(step); err != nil {
			a.Audit.Record("blocked_business_rule_violation", step)
			continue
		}

		if err := executeTool(step); err != nil {
			a.Audit.Record("execution_failed", step)
			return fmt.Errorf("echec execution outil %s: %w", step.Name, err)
		}
		a.Audit.Record("action_executed", step)
	}
	return nil
}

func requestPlan(userInput string) (*PlanResponse, error) {
	client := &http.Client{}
	_ = client
	_ = apiKey
	return &PlanResponse{Steps: []ToolCall{}}, nil
}

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
