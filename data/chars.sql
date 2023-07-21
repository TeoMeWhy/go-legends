DROP TABLE IF EXISTS chars;

CREATE TABLE IF NOT EXISTS chars (
    nome varchar(100),
    raca varchar(100),
    classe varchar(100),
    exp int,
    nivel int,
    pontos int,
    vitalidade int,
    att_forca int,
    att_destreza int,
    att_inteligencia int,
    mod_forca int,
    mod_destreza int,
    mod_inteligencia int
);