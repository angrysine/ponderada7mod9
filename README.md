# ponderada7mod9

## Introdução

Este repositório contém a setima ponderada do módulo 9 do curso de engenharia da computação do Inteli. O objetivo é integrar o sistema com o confluent e um cluster do mongodb. Foi utilizado o metabase para a visualização dos dados.

## Como rodar

Somente é necessário rodar o seguinte comando:

```bash
go run *.go
```

Pois o sistema usa multithreading para rodar o publisher e o metabase simultaneamente


## Vídeo

O vídeo da apresentação da ponderada pode ser encontrado no seguinte link: [Vídeo](https://youtu.be/eaTM6IqzyNU)

## Testes

Os testes foram feitos utilizando o framework de testes do go. Para rodar os testes, basta rodar o seguinte comando:

```bash
go test
```

Existe somente um teste que publica uma mensagem e verifica se ela foi inserida no banco de dados.
