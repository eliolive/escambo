package usuariorepo

type Usuario struct {
	Nome         string `json:"nome"`
	Email        string `json:"email"`
	Senha        string `json:"senha"`
	Telefone     string `json:"telefone"`
	WhatsappLink string `json:"whatsapp_link"`
}
