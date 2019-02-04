# ETL Project in GoLang

Um serviço de ETL consiste basicamente de 3 etapas principais: A extração dos dados; A transformação ou higienização destes dados; O carregamento dos dados manipulados.

Neste projeto de demonstração, este processo é feito utilizando GO:

**A extração:**

A extração dos dados pode se dar de várias formas diferentes, seja por APIs, serviços de mensageria, bancos de dados, arquivos, buscas periódicas através de _schedulers_, etc. Neste projeto, foi utilizado um arquivo fixo (`/source/base_teste.txt`).

**A Transformação:**

A tranformação ou higienização dos dados pode ser para filtrar dados não necessários para o objetivo final do processo ou para que fiquem em um formato que possa ser analisado. Neste projeto, a higienização divide os dados em 3 tipos, sendo eles: clientes sem compras (não efetuaram compras e talvez não sejam necessários para o objetivo da análise), clientes com inconsistências (possuem erros cadastrais) e clientes sanitizados (aqueles cujo as informações estão de acordo com as validações). O processo de transformação também pode agrupar dados para serem armazenados em um único local, como um banco não relacional.

**O Carregamento:**

Os dados transformados são persistidos, geralmente em um banco de dados, mas também pode criar arquivos ou enviar os dados para outros sistemas. Neste projeto, os dados préviamente separados são persistidos em 3 tabelas distintas do Postgres.


**Versões:**

Go 1.11.5<br>
Postgres 10.6


**Rodando o projeto:**

Como as tabelas necessárias são criadas em tempo de execução, para rodar o projeto, apenas dois passos são necessários:

1 - Criar o database (apenas caso ainda não exista):

`CREATE DATABASE postgres
     WITH 
     OWNER = postgres
     ENCODING = 'UTF8'
     CONNECTION LIMIT = -1;`

Obs.: O usuário para conexão com o banco deve ser o padrão do Postgres:
usuário: **postgres**
senha: **postgres**

2 - Rodar o projeto:

`go run main.go`

Ao rodar o projeto, todas as tabelas utilizadas são criadas/recriadas.

O resultado da execução pode ser vista pelo log do console e nas tabelas que foram criadas. Todo o processo leva entre 10 e 15 segundos para ser executado. São extraídos, analisados, sanitizados e persistidos quase 50 mil itens.