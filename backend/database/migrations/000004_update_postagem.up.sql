ALTER TABLE postagens
ADD CONSTRAINT unique_titulo_user UNIQUE (titulo, user_id);
