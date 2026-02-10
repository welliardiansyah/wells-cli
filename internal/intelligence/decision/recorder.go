package decision

import "github.com/welliardiansyah/wells-cli/internal/intelligence/storage"

const decisionFile = "decisions.json"

func Record(action, entity string) {
	var decisions []Decision
	_ = storage.ReadJSON(decisionFile, &decisions)

	decisions = append(decisions, New(action, entity))
	_ = storage.WriteJSON(decisionFile, decisions)
}
