SELECT
    t2.*

FROM chars_itens as t1

LEFT JOIN itens AS t2
ON t1.nome_item = t2.nome

WHERE nome_char = '{nome}'