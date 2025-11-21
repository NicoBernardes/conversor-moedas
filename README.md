Conversor de Moedas (CLI)

Este é um conversor simples de moedas feito em Go.
O programa lê um JSON interno com taxas de câmbio e converte valores de BRL (Real) para a moeda desejada.

🚀 Como usar

Compile o projeto:

go build -o convert


Execute passando o valor em BRL e a moeda de destino:

./convert [valor] [moeda]

Exemplos:
./convert 10 USD
./convert 50 EUR
./convert 200 JPY


A moeda pode ser digitada em qualquer formato (usd, USD, UsD etc.).

📌 O que o programa faz

Carrega um JSON com as cotações.

Verifica se a moeda informada existe.

Converte o valor em BRL multiplicando pela taxa correspondente.

Exibe o valor convertido no terminal.
