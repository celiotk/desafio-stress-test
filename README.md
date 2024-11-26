# Desafio Stress Test
Um sistema CLI em Go para realizar testes de carga em serviços web, permitindo especificar a URL, o número total de requisições e a quantidade de chamadas simultâneas.

## Gerando imagem do Docker
Para criar a imagem Docker, execute o comando abaixo, substituindo `<nome-imagem-docker>` pelo nome desejado para a imagem:
```bash
docker build -t <nome-imagem-docker> .
```

## Executando a Aplicação
Execute o container Docker gerado, fornecendo os seguintes parâmetros:
* `--url`: URL do serviço a ser testado.  
* `--requests`: Número total de requests.  
* `--concurrency`: Número de chamadas simultâneas.

Exemplo do comando:
```bash
docker run <nome imagem docker> --url=http://google.com --requests=1000 --concurrency=10
```

## Resultado do teste
Ao término do teste de carga, o sistema exibe um relatório com o seguinte formato:
```json
{
  "TotalTime": "3m0.668107199s",
  "TotalReq": 961,
  "Status200": 906,
  "OtherStatus": {
    "429": 55
  },
  "Errors": 39
}
```
### Campos do Relatório 
* `TotalTime`: Tempo total gasto na execução.  
* `TotalReq`: Quantidade total de requests realizados.  
* `Status200`: Quantidade de requests com status HTTP 200.  
* `OtherStatus`: Distribuição de outros códigos de status HTTP.  
* `Errors`: Quantidade requests com erro.
