DROP TABLE IF EXISTS itens;

CREATE TABLE itens (
    nome VARCHAR(100),
    descricao TEXT,
    tipo_item varchar(100),
    att_forca INT,
    att_destreza INT,
    att_inteligencia INT,
    peso_gramas INT,
    nivel_minimo INT
);

INSERT INTO itens (nome, descricao, tipo_item, att_forca, att_destreza, att_inteligencia, peso_gramas, nivel_minimo)
VALUES

    -- Cetros ou arinhas
    ('Cetro Iniciante', 'Um cetro básico para magos e clérigos iniciantes.', 'arma', 0, 0, 1, 300, 1),
    ('Cetro Arcano', 'Um cetro mágico esculpido em cristal, emanando poderes arcanos.', 'arma', -1, 1, 12, 600, 5),
    ('Varinha do Sábio', 'Uma varinha elegante feita de carvalho, concedendo grande sabedoria ao usuário.', 'arma', -2, 2, 15, 550, 8),
    ('Cetro do Fogo', 'Um cetro incandescente que permite ao usuário lançar chamas ardentes contra os inimigos.', 'arma', -1, 3, 10, 700, 4),

    -- Espadas
    ('Espada Curta', 'Uma para aprendizes de espadachins.', 'arma', 1, 0, 0, 1000, 1),
    ('Espada Longa', 'Uma espada equilibrada e confiável, perfeita para guerreiros de todos os estilos.', 'arma', 12, 0, 0, 1200, 12),
    ('Espada das Sombras', 'Uma espada envolta em escuridão, ampliando a força de ataques furtivos.', 'arma', 14, 2, -2, 1300, 15),
    ('Espada Élfica da Velocidade', 'Uma espada fina e leve, permitindo ataques rápidos e precisos.', 'arma', 8, 8, 1, 900, 10),

    -- Machado
    ('Machete básica', 'Machete simples para atques rápidos e curtos.', 'arma', 1, 0, 0, 1000, 1),

    -- Adaga
    ('Adaga da cozinha', 'Adaga simples para pequenos furtos.', 'arma', 0, 1, 0, 150, 1),

    -- Arcos
    ('Arco Simples', 'Arco leve para curtas distâncias.', 'arma', 0, 1, 0, 650, 1),
    ('Arco Longo', 'Um arco de longo alcance, perfeito para atiradores precisos.', 'arma', 3, 15, 2, 800, 7),
    ('Arco do Caçador', 'Um arco especialmente projetado para emboscar presas com ataques precisos.', 'arma', 2, 14, 3, 850, 9),
    ('Arco Místico', 'Um arco entalhado com runas místicas, conferindo poderes arcanos a cada flecha.', 'arma', 1, 12, 6, 1000, 11),

    -- Armaduras
    ('Armadura de Couro', 'Uma armadura de couro leve e flexível, ideal para movimentos ágeis.', 'armadura', 2, 6, 1, 2500, 2),
    ('Armadura Completa do Paladino', 'Uma armadura completa feita para paladinos, oferecendo grande proteção.', 'armadura', 8, -3, 2, 4000, 18),
    ('Armadura Élfica da Floresta', 'Uma armadura feita por elfos, misturando natureza e proteção mágica.', 'armadura', 4, 4, 5, 3200, 10),

    -- Elmos
    ('Elmo do Cavaleiro', 'Um elmo robusto e resistente, usado por cavaleiros em batalhas.', 'elmo', 4, -1, 0, 900, 6),
    ('Elmo da Visão Noturna', 'Um elmo especial que permite ao usuário enxergar na escuridão total.', 'elmo', 0, 1, 3, 800, 4),
    ('Elmo do Guardião', 'Um elmo forjado para proteger a cabeça do usuário em combate, resistente e imponente.', 'elmo', 3, 0, 0, 1200, 5),

    -- Botas
    ('Botas Velozes', 'Um par de botas encantadas que aumenta a velocidade do usuário, permitindo movimentos mais rápidos.', 'botas', 1, 8, 0, 500, 3),
    ('Botas de Saltos Altos', 'Botas com saltos altos que aumentam a capacidade de saltar grandes distâncias.', 'botas', 0, 10, 1, 550, 4),
    ('Botas do Viajante', 'Botas leves e confortáveis, ideais para longas jornadas e exploração.', 'botas', 1, 5, 3, 400, 2),

    -- Calças
    ('Calças Resistentes', 'Calças feitas de materiais resistentes, protegendo as pernas do usuário.', 'calcas', 2, 2, 0, 700, 2),
    ('Calças Leves de Lã', 'Calças leves e confortáveis feitas de lã, ideais para viajar longas distâncias.', 'calcas', 1, 5, 1, 600, 3),
    ('Calças de Combate', 'Calças projetadas para combate, permitindo movimentos ágeis e proteção adicional.', 'calcas', 3, 4, 0, 800, 4)

;