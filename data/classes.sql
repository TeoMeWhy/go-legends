DROP TABLE IF EXISTS classes;

CREATE TABLE IF NOT EXISTS classes (
    id int,
    nome varchar(100),
    atributo_primario varchar(100),
    atributo_secundario varchar(100),
    atributo_terciario varchar(100)
);

INSERT INTO classes
VALUES (1, 'guerreiro', 'forca', 'destreza', 'inteligencia'),
       (2, 'arqueiro', 'destreza', 'forca', 'inteligencia'),
       (3, 'ladino', 'destreza', 'inteligencia', 'forca'),
       (4, 'mago', 'inteligencia', 'destreza', 'forca'),
       (5, 'clerigo', 'inteligencia', 'forca', 'destreza')
;
