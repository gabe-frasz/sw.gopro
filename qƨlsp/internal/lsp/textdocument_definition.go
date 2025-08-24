package lsp

type TextDocumentDefinitionRequest struct {
	Request
	Params TextDocumentDefinitionParams `json:"params"`
}

type TextDocumentDefinitionParams struct {
	TextDocumentPositionParams
}

type TextDocumentDefinitionResponse struct {
	Response
	Result Location `json:"result"`
}

func NewTextDocumentDefinitionResponse(id int, location Location) *TextDocumentDefinitionResponse {
	return &TextDocumentDefinitionResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: location,
	}
}
