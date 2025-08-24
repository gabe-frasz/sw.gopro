package lsp

type TextDocumentCompletionRequest struct {
	Request
	Params TextDocumentCompletionParams `json:"params"`
}

type TextDocumentCompletionParams struct {
	TextDocumentPositionParams
}

type TextDocumentCompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label         string `json:"label"`
	Detail        string `json:"detail"`
	Documentation string `json:"documentation"`
}

func NewTextDocumentCompletionResponse(id int, items []CompletionItem) *TextDocumentCompletionResponse {
	return &TextDocumentCompletionResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
}
