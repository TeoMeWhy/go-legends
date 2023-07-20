DROP TABLE IF EXISTS racas;

CREATE TABLE IF NOT EXISTS racas (
    id int,
    nome varchar(100),
    mod_destreza int,
    mod_forca int,
    mod_inteligencia int
);

INSERT INTO racas
VALUES (1, 'humano', 2, 1, 1),
       (2, 'elfo', 2, 0, 2),
       (3, 'anao', 1, 3, 0),
       (4, 'halfling', 4, 0, 0),
       (5, 'orc', 0, 1, 0)
;
