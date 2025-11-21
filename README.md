Conversor de Moedas (CLI)

Este projeto é um conversor de moedas em linha de comando (CLI) escrito em Go.
Ele utiliza um JSON interno contendo taxas de câmbio fixas e converte valores de BRL (Real) para a moeda desejada.

📌 Como funciona

O usuário executa o binário passando dois parâmetros:

O valor em BRL que deseja converter

A moeda de destino, como USD, EUR, JPY etc.

Exemplo:

./convert 10 USD


Isso converte R$ 10,00 para Dólar americano (USD) usando a taxa definida no JSON.

🛠️ Instalação

Clone o repositório e acesse a pasta:

git clone https://github.com/SEU_USUARIO/seu-repo.git
cd seu-repo/cmd


Compile o binário:

go build -o convert


Agora execute:

./convert 10 EUR

📤 Uso
./convert [valor_em_brl] [moeda_destino]


Exemplos:

./convert 50 USD
./convert 120 eur
./convert 200 jpy


A moeda não diferencia maiúsculas de minúsculas (usd, UsD, USD — tudo funciona).

🧩 Estrutura do JSON

O programa utiliza internamente um JSON no formato:

{
  "base": "BRL",
  "date": "2025-04-14",
  "rates": {
    "USD": 0.151,
    "EUR": 0.137
  }
}


A conversão é feita multiplicando o valor em BRL pela taxa correspondente.

✔️ Exemplos de saída
./convert 10 USD
1.51

./convert 200 EUR
27.40

📎 Observação

As taxas de câmbio são fixas e apenas para fins de estudo.
