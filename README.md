Primeira implementação de API em Go

Este projeto desenvolve uma API simples em Go que executa operações matemáticas básicas (adição, subtração, multiplicação, divisão e raiz quadrada).
Endpoints disponíveis:

- /adicao: Adiciona um número à lista de números registrados.
- /subtracao: Subtrai um número da lista de números registrados.
- /multiplicacao: Multiplica todos os números registrados por um número especificado.
- /divisao: Divide todos os números registrados por um número especificado.
- /raizquadrada: Calcula a raiz quadrada de todos os números registrados.
- /seetotal: Retorna o total atual dos números registrados.
- /delete: Remove todos os números da lista de registros.

Funcionamento:

A API espera um corpo de requisição no formato JSON contendo um número para realizar operações de adição, subtração, multiplicação, divisão ou calcular a raiz quadrada. Os números são armazenados até serem removidos pelo endpoint /delete.
Utilização:

Para utilizar a API, envie requisições HTTP aos endpoints mencionados acima, especificando o número desejado no corpo da requisição em formato JSON.
