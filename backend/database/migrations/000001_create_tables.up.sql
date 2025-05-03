CREATE TABLE usuarios (
  id UUID NOT NULL PRIMARY KEY,
  nome VARCHAR(150) NOT NULL,
  email VARCHAR(150) NOT NULL,
  senha VARCHAR(100) NOT NULL,
  telefone VARCHAR(20) NOT NULL,
  whatsapp_link VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE categorias(
    id UUID NOT NULL PRIMARY KEY,
    titulo VARCHAR(100) not null
);

CREATE TABLE postagens (
  id UUID NOT NULL PRIMARY KEY,
  titulo VARCHAR(150) NOT NULL,
  descricao TEXT NOT NULL,
  imagem_url VARCHAR(255) NOT NULL,
  user_id UUID NOT NULL,
  categoria_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_postagens_usuarios
    FOREIGN KEY (user_id) REFERENCES usuarios(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_postagens_categorias
    FOREIGN KEY (categoria_id) REFERENCES categorias(id) ON DELETE SET NULL
);

CREATE TABLE endereco (
   id UUID NOT NULL PRIMARY KEY,
   cep VARCHAR(15) NOT NULL,
   rua VARCHAR(255) NOT NULL,
   numero INT NOT NULL,
   complemento TEXT,
   bairro VARCHAR(100),
   cidade VARCHAR(100),
   estado VARCHAR(2),
   user_id UUID NOT NULL,
   CONSTRAINT fk_endereco_usuarios
     FOREIGN KEY (user_id) REFERENCES usuarios(id)
     ON DELETE CASCADE
);
CREATE TABLE propostas (
    id UUID PRIMARY KEY,
    postagem_id UUID NOT NULL,
    remetente_id UUID NOT NULL,         -- quem faz a proposta
    destinatario_id UUID NOT NULL,      -- quem recebe a proposta
    status VARCHAR(20) NOT NULL DEFAULT 'pendente',  -- pendente, aceita, recusada
    excluida BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_proposta_postagem
        FOREIGN KEY (postagem_id) REFERENCES postagens(id),
    CONSTRAINT fk_proposta_remetente
        FOREIGN KEY (remetente_id) REFERENCES usuarios(id),
    CONSTRAINT fk_proposta_destinatario
        FOREIGN KEY (destinatario_id) REFERENCES usuarios(id)
);
