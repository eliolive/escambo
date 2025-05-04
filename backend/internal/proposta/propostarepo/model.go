package propostarepo

type Proposta struct {
	ID             string `json:"id"`
	PostagemID     string `json:"postagem_id"`
	RemetenteID    string `json:"remetente_id"`
	DestinatarioID string `json:"destinatario_id"`
	Status         string `json:"status"`
	Excluida       bool   `json:"excluida"`
}
