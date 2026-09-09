// vulnerable.go - CWE-284: Improper Access Control
// Ce code de deploiement Go publie des regles de securite Firestore en "mode test"
// (allow read, write: if true) via l'API Firebase Rules, exposant toute la base
// de donnees en lecture/ecriture sans authentification.
package deploy

import (
	"context"
	"fmt"

	firebaserules "google.golang.org/api/firebaserules/v1"
)

const projectID = "REPLACE_WITH_YOUR_PROJECT_ID"

// DeployTestModeRules publie des regles Firestore en mode test, censees etre temporaires,
// mais laissees en production.
func DeployTestModeRules(ctx context.Context, svc *firebaserules.Service) error {
	// VULNERABLE: regle "if true" autorisant lecture ET ecriture a tout le monde,
	// authentifie ou non, sur l'integralite de la base de donnees
	rulesContent := `
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /{document=**} {
      allow read, write: if true;
    }
  }
}
`

	ruleset := &firebaserules.Ruleset{
		Source: &firebaserules.Source{
			Files: []*firebaserules.File{
				{Name: "firestore.rules", Content: rulesContent},
			},
		},
	}

	created, err := svc.Projects.Rulesets.Create(
		fmt.Sprintf("projects/%s", projectID), ruleset,
	).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creation ruleset: %w", err)
	}

	release := &firebaserules.Release{
		Name:      fmt.Sprintf("projects/%s/releases/cloud.firestore", projectID),
		RulesetName: created.Name,
	}
	_, err = svc.Projects.Releases.Update(
		release.Name, release,
	).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("publication release: %w", err)
	}

	return nil
}
