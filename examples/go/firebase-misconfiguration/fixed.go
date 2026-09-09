// fixed.go - Correctif CWE-284: Improper Access Control
// Les regles de securite Firestore deployees restreignent explicitement la lecture
// et l'ecriture au proprietaire authentifie de chaque document, avec validation de schema.
package deploy

import (
	"context"
	"fmt"

	firebaserules "google.golang.org/api/firebaserules/v1"
)

const projectID = "REPLACE_WITH_YOUR_PROJECT_ID"

// DeployOwnerScopedRules publie des regles Firestore restreintes au proprietaire
// authentifie de chaque document, avec validation des champs autorises a l'ecriture.
func DeployOwnerScopedRules(ctx context.Context, svc *firebaserules.Service) error {
	// SECURISE: lecture/ecriture limitees au proprietaire authentifie du document,
	// et validation du schema sur les creations
	rulesContent := `
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /users/{userId}/documents/{docId} {
      allow read, write: if request.auth != null
                          && request.auth.uid == userId;
      allow create: if request.auth != null
                     && request.auth.uid == userId
                     && request.resource.data.keys().hasOnly(['title', 'content', 'createdAt']);
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
		Name:        fmt.Sprintf("projects/%s/releases/cloud.firestore", projectID),
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
