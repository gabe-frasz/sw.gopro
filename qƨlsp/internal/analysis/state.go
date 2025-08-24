package analysis

import (
	"fmt"
	"strings"

	"github.com/gabe-frasz/qslsp/internal/lsp"
)

type State struct {
	// map of file URIs to its contents
	Documents map[string]string
}

func NewState() *State {
	return &State{
		Documents: make(map[string]string),
	}
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return s.GetDiagnostics(uri)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return s.GetDiagnostics(uri)
}

func (s *State) Hover(uri string, position lsp.Position) string {
	return fmt.Sprintf(
		"File %s, line %d, column %d, total characters %d",
		uri,
		position.Line,
		position.Character,
		len(s.Documents[uri]),
	)
}

func (s *State) GoToDefinition(uri string, position lsp.Position) lsp.Location {
	definitionLine := position.Line - 1
	if position.Line == 0 {
		definitionLine = 1
	}

	return lsp.Location{
		URI: uri,
		Range: lsp.Range{
			Start: lsp.Position{
				Line:      definitionLine,
				Character: 0,
			},
			End: lsp.Position{
				Line:      definitionLine,
				Character: 0,
			},
		},
	}
}

func (s *State) GetCodeActions(uri string) []lsp.CodeAction {
	text := s.Documents[uri]

	codeActions := []lsp.CodeAction{}

	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx == -1 {
			continue
		}

		changeRange := lsp.Range{
			Start: lsp.Position{
				Line:      uint(row),
				Character: uint(idx),
			},
			End: lsp.Position{
				Line:      uint(row),
				Character: uint(idx + len("VS Code")),
			},
		}

		replaceChange := map[string][]lsp.TextEdit{}
		replaceChange[uri] = []lsp.TextEdit{
			{
				Range:   changeRange,
				NewText: "Neovim",
			},
		}

		censorChange := map[string][]lsp.TextEdit{}
		censorChange[uri] = []lsp.TextEdit{
			{
				Range:   changeRange,
				NewText: "VS C*de",
			},
		}

		codeActions = append(
			codeActions,
			lsp.CodeAction{
				Title: "Replace VS C*de editor to a superior one",
				Edit: &lsp.WorkspaceEdit{
					Changes: replaceChange,
				},
			},
			lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit: &lsp.WorkspaceEdit{
					Changes: censorChange,
				},
			},
		)
	}

	return codeActions
}

func (s *State) GetCompletions(uri string) []lsp.CompletionItem {
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim btw",
			Detail:        "Neovim is a free and open-source, community-driven text editor.",
			Documentation: "Visit https://neovim.io/",
		},
	}

	return items
}

func (s *State) GetDiagnostics(uri string) []lsp.Diagnostic {
	text := s.Documents[uri]

	diagnostics := []lsp.Diagnostic{}

	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx == -1 {
			continue
		}

		diagnostics = append(diagnostics,
			lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      uint(row),
						Character: uint(idx),
					},
					End: lsp.Position{
						Line:      uint(row),
						Character: uint(len("VS Code")),
					},
				},
				Severity: 1,
				Source:   "the-ghost-of-Bram-Moolenaar",
				Message:  "Detected forbidden incantation: 'VS C*de'. Consider replacing it with 'Neovim' (for extra IQ points) or deleting it (for inner peace).",
			},
		)
	}

	return diagnostics
}
