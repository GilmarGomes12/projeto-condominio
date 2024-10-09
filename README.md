# Projeto Condomínio Varandas do Praia

## Descrição
Este projeto é um sistema de gerenciamento para o condomínio Varandas do Praia. Ele inclui funcionalidades para registrar e gerenciar visitantes, prestadores de serviço, ocorrências e veículos. O sistema é desenvolvido utilizando HTML, CSS, JavaScript, Go e várias bibliotecas de terceiros.

## Funcionalidades

- **Registro de Visitantes**:
  - **Formulário Intuitivo**: O sistema oferece um formulário simples e intuitivo para o registro de visitantes, solicitando informações como nome, documento, apartamento a ser visitado, placa do veículo e motivo da visita.
  - **Validação de Dados**: Todos os dados inseridos são validados para garantir a integridade das informações.
  - **Histórico de Visitas**: O sistema mantém um histórico completo das visitas, permitindo que seja realizado consulta as informações a qualquer momento.

- **Registro de Prestadores de Serviço**:
  - **Formulário Completo**: Registro de prestadores de serviço com informações detalhadas sobre a empresa, tipo de serviço, horário de entrada e saída, e detalhes do contratante.
  - **Verificação de Antecedentes**: Validação dos dados e verificação de antecedentes dos prestadores de serviço.

- **Registro de Ocorrências**:
  - **Formulário Detalhado**: Registro de ocorrências com detalhes como data, hora, descrição da ocorrência e autor.
  - **Notificação Automática**: Notificação automática para o síndico e administradores sobre novas ocorrências.

- **Gerenciamento de Veículos**:
  - **Controle de Acesso**: Sistema de controle de acesso para veículos, com registro de entradas e saídas.

## Fluxos de Usuário

- **Registro de Ocorrência**:
  1. O usuário (porteiro ou morador) acessa a página de registro de ocorrências.
  2. Preenche o formulário com os detalhes da ocorrência.
  3. Submete o formulário, que é validado e salvo no sistema.

- **Consulta de Visitante**:
  1. O porteiro acessa a página de consulta de visitantes.
  2. Insere os critérios de busca (nome, documento, apartamento visitado).
  3. O sistema exibe a lista de visitantes que correspondem aos critérios.
  4. O porteiro pode visualizar os detalhes do visitante selecionado.

## Usuários e Permissões

- **Porteiros**:
  - Registrar e consultar visitantes.
  - Registrar prestadores de serviço.
  - Registrar ocorrências.

- **Síndico**:
  - Consultar todas as informações registradas no sistema.
  - Gerenciar ocorrências.
  - Receber notificações sobre novas ocorrências e registros importantes.

- **Administradores**:
  - Acesso total ao sistema.
  - Gerenciar usuários e permissões.
  - Consultar e gerenciar todas as informações do sistema.

## Tecnologias Utilizadas

- **HTML5**: Estrutura das páginas web.
- **CSS3**: Estilização das páginas web.
- **JavaScript**: Funcionalidades dinâmicas e manipulação do DOM.
- **Go**: Backend do sistema.
- **Bootstrap**: Framework CSS para design responsivo.
- **jQuery**: Biblioteca JavaScript para simplificar a manipulação do DOM.
- **AOS (Animate On Scroll)**: Biblioteca para animações ao rolar a página.
- **GLightbox**: Biblioteca para exibição de imagens e vídeos em lightbox.
- **Swiper**: Biblioteca para criação de sliders/carrosséis.

## Estrutura do Projeto

📦condominio
 ┣ 📂config
 ┃ ┣ 📜.env
 ┃ ┣ 📜db.go
 ┃ ┗ 📜db_test.go
 ┣ 📂controllers
 ┃ ┣ 📜agendamentos.go
 ┃ ┣ 📜configuracao.go
 ┃ ┣ 📜domesticos.go
 ┃ ┣ 📜encomenda.go
 ┃ ┣ 📜funcionarios.go
 ┃ ┣ 📜index.go
 ┃ ┣ 📜login.go
 ┃ ┣ 📜menu.go
 ┃ ┣ 📜morador.go
 ┃ ┣ 📜mudanca.go
 ┃ ┣ 📜ocorrencia.go
 ┃ ┣ 📜pesquisa.go
 ┃ ┣ 📜pesquisa_agendamento.go
 ┃ ┣ 📜pesquisa_encomenda.go
 ┃ ┣ 📜pesquisa_morador.go
 ┃ ┣ 📜prestador.go
 ┃ ┣ 📜register.go
 ┃ ┣ 📜usuario.go
 ┃ ┣ 📜veiculos.go
 ┃ ┗ 📜visitantes.go
 ┣ 📂middleware
 ┃ ┗ 📜auth.go
 ┣ 📂models
 ┃ ┣ 📜agendamentos.go
 ┃ ┣ 📜configuracao.go
 ┃ ┣ 📜domesticos.go
 ┃ ┣ 📜encomenda.go
 ┃ ┣ 📜funcionarios.go
 ┃ ┣ 📜morador.go
 ┃ ┣ 📜mudanca.go
 ┃ ┣ 📜ocorrencia.go
 ┃ ┣ 📜prestador.go
 ┃ ┣ 📜usuario.go
 ┃ ┗ 📜visitantes.go
 ┣ 📂routes
 ┃ ┗ 📜routes.go
 ┣ 📂templates
 ┃ ┣ 📂assets
 ┃ ┃ ┣ 📂css
 ┃ ┃ ┃ ┣ 📜custom.css
 ┃ ┃ ┃ ┣ 📜customs.css
 ┃ ┃ ┃ ┗ 📜main.css
 ┃ ┃ ┣ 📂img
 ┃ ┃ ┃ ┣ 📂menu
 ┃ ┃ ┃ ┃ ┣ 📜agenda.png
 ┃ ┃ ┃ ┃ ┣ 📜automovel.png
 ┃ ┃ ┃ ┃ ┣ 📜cadastro.png
 ┃ ┃ ┃ ┃ ┣ 📜configuracoes.png
 ┃ ┃ ┃ ┃ ┣ 📜domesticas.png
 ┃ ┃ ┃ ┃ ┣ 📜encomenda.png
 ┃ ┃ ┃ ┃ ┣ 📜funcionarios.png
 ┃ ┃ ┃ ┃ ┣ 📜incidente.png
 ┃ ┃ ┃ ┃ ┣ 📜mudanca.png
 ┃ ┃ ┃ ┃ ┣ 📜pesquisar.png
 ┃ ┃ ┃ ┃ ┣ 📜suporte-tecnico.png
 ┃ ┃ ┃ ┃ ┗ 📜visitantes.png
 ┃ ┃ ┃ ┣ 📜apple-touch-icon.png
 ┃ ┃ ┃ ┣ 📜favicon.png
 ┃ ┃ ┃ ┣ 📜hero-img.png
 ┃ ┃ ┃ ┣ 📜swiper-bundle.min.js
 ┃ ┃ ┃ ┗ 📂waypoints
 ┃ ┃ ┃ ┃ ┗ 📜noframework.waypoints.js
 ┃ ┣ 📜agendamentos.html
 ┃ ┣ 📜configuracoes.html
 ┃ ┣ 📜domesticos.html
 ┃ ┣ 📜encomendas.html
 ┃ ┣ 📜funcionarios.html
 ┃ ┣ 📜login.html
 ┃ ┣ 📜menu.html
 ┃ ┣ 📜moradores.html
 ┃ ┣ 📜mudanca.html
 ┃ ┣ 📜ocorrencias.html
 ┃ ┣ 📜pesquisar.html
 ┃ ┣ 📜prestadores.html
 ┃ ┣ 📜register.html
 ┃ ┣ 📜veiculos.html
 ┃ ┗ 📜visitantes.html
 ┣ 📜.gitignore
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜LICENSE
 ┣ 📜main.go
 ┗ 📜README.md

 ## Banco de Dados

O projeto utiliza PostgreSQL como banco de dados. Abaixo estão os passos para configurar o banco de dados:

## Diagrama UML

Abaixo está o diagrama UML que representa a arquitetura do sistema:

![Diagrama UML](docs/Diagrama_UML.png)

1. **Criação do Banco de Dados**:
    - Crie um banco de dados PostgreSQL com o nome desejado.
    - Exemplo de comando para criar o banco de dados:
      ```sql
      CREATE DATABASE condominio;
      ```

2. **Configuração das Variáveis de Ambiente**:
    - Configure as variáveis de ambiente necessárias no arquivo `.env`. Exemplo:
      ```env
      DB_HOST=localhost
      DB_PORT=5432
      DB_USER=seu_usuario
      DB_PASSWORD=sua_senha
      DB_NAME=condominio
      ```

3. **Migrações do Banco de Dados**:
    - Execute as migrações para criar as tabelas necessárias no banco de dados.
    - Exemplo de comando para executar as migrações:
      ```sh
      go run main.go migrate
      ```

4. **Estrutura das Tabelas**:
    - Abaixo está o modelo da estrutura das tabelas principais:
      
    ```sql
    -- Tabela de usuário para criação dos usuários do sistema com 3 níveis de permissão
    CREATE TABLE usuario (
        id SERIAL PRIMARY KEY,
        username VARCHAR(150) UNIQUE NOT NULL,
        password VARCHAR(128) NOT NULL,
        permissao VARCHAR(50) NOT NULL, -- administrador, sindico, porteiro
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    -- Tabela principal de moradores
    CREATE TABLE moradores (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(100) NOT NULL,
        apartamento VARCHAR(10) NOT NULL,
        bloco VARCHAR(10) NOT NULL,
        telefone1 VARCHAR(15) NOT NULL,
        telefone2 VARCHAR(15),
        email VARCHAR(255) NOT NULL,
        email2 VARCHAR(255),
        observacao TEXT,
        UNIQUE (apartamento, bloco)
    );

    -- Tabela para informações adicionais dos moradores
    CREATE TABLE morador_info (
        id SERIAL PRIMARY KEY,
        morador_id INT NOT NULL REFERENCES moradores(id) ON DELETE CASCADE,
        nome VARCHAR(100) NOT NULL,
        data_nascimento DATE NOT NULL,
        UNIQUE (morador_id, nome)
    );

    -- Tabela para veículos
    CREATE TABLE veiculo (
        id SERIAL PRIMARY KEY,
        placa VARCHAR(10) NOT NULL,
        cor VARCHAR(20) NOT NULL,
        marca VARCHAR(50) NOT NULL,
        modelo VARCHAR(50) NOT NULL,
        UNIQUE (placa)
    );

    -- Tabela para associar moradores aos veículos
    CREATE TABLE morador_veiculo (
        morador_id INT REFERENCES moradores(id) ON DELETE CASCADE,
        veiculo_id INT REFERENCES veiculo(id) ON DELETE CASCADE,
        PRIMARY KEY (morador_id, veiculo_id)
    );

    -- Tabela para agendamento de mudanças
    CREATE TABLE agendamento_mudanca (
        id SERIAL PRIMARY KEY,
        data_mudanca DATE NOT NULL,
        responsavel_nome VARCHAR(100) NOT NULL,
        responsavel_apto VARCHAR(10) NOT NULL,
        responsavel_bloco VARCHAR(10) NOT NULL,
        horario VARCHAR(20) NOT NULL,
        nome_empresa VARCHAR(100) NOT NULL,
        iscar_entrando BOOLEAN DEFAULT FALSE,
        iscar_saindo BOOLEAN DEFAULT FALSE,
        uso_elevador BOOLEAN DEFAULT FALSE,
        uso_escada BOOLEAN DEFAULT FALSE,
        iscar BOOLEAN DEFAULT FALSE
    );

    -- Tabela para agendamento de espaços comuns
    CREATE TABLE agendamento (
        id SERIAL PRIMARY KEY,
        nome_morador VARCHAR(100) NOT NULL,
        apartamento VARCHAR(10) NOT NULL,
        bloco VARCHAR(10) NOT NULL,
        local VARCHAR(100) NOT NULL,
        dia INTEGER DEFAULT 1,
        mes INTEGER DEFAULT 1,
        ano INTEGER DEFAULT 2023,
        periodo VARCHAR(20) NOT NULL,
        funcionario VARCHAR(100) NOT NULL,
        observacoes TEXT,
        convidados TEXT
    );

    -- Tabela para ocorrências
    CREATE TABLE ocorrencia (
        id SERIAL PRIMARY KEY,
        numero_ocorrencia VARCHAR(100) UNIQUE NOT NULL,
        data_ocorrencia DATE NOT NULL,
        nome_funcionario VARCHAR(100) NOT NULL,
        funcao_funcionario VARCHAR(100) NOT NULL,
        hora_registro TIME NOT NULL,
        data_registro DATE NOT NULL,
        unidade_envolvida VARCHAR(100) NOT NULL,
        bloco VARCHAR(100) DEFAULT 'A',
        autor_ocorrencia VARCHAR(3000) NOT NULL,
        descricao_ocorrencia TEXT NOT NULL
    );

    -- Tabela para prestadores de serviço
    CREATE TABLE prestador (
        id SERIAL PRIMARY KEY,
        nome_empresa VARCHAR(100) NOT NULL,
        tipo_servico VARCHAR(100) NOT NULL,
        data DATE NOT NULL,
        hora_entrada TIME NOT NULL,
        hora_saida TIME NOT NULL,
        nome_prestador VARCHAR(100) NOT NULL,
        rg_cpf VARCHAR(20) NOT NULL,
        telefone VARCHAR(15) NOT NULL,
        contratante_nome VARCHAR(100) NOT NULL,
        contratante_tipo VARCHAR(20) NOT NULL,
        contratante_apto VARCHAR(10),
        contratante_bloco VARCHAR(10),
        autorizou VARCHAR(100) NOT NULL
    );

    -- Tabela para visitantes
    CREATE TABLE visitante (
        id SERIAL PRIMARY KEY,
        data DATE NOT NULL,
        nome_visitante VARCHAR(100) NOT NULL,
        rg_cpf VARCHAR(20) NOT NULL,
        visitando VARCHAR(100) NOT NULL,
        apartamento VARCHAR(10) NOT NULL,
        bloco VARCHAR(10) NOT NULL,
        hora_entrada TIME NOT NULL,
        hora_saida TIME,
        autorizou VARCHAR(100) NOT NULL,
        vaga VARCHAR(20) NOT NULL,
        placa VARCHAR(10),
        marca VARCHAR(50),
        modelo VARCHAR(50),
        cor VARCHAR(20)
    );

    -- Tabela para encomendas
    CREATE TABLE encomenda (
        id SERIAL PRIMARY KEY,
        numero_protocolo VARCHAR(100) UNIQUE NOT NULL,
        data_entrega TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        data_hora_recebimento TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        nome_destinatario VARCHAR(100) NOT NULL,
        apartamento VARCHAR(10) NOT NULL,
        bloco VARCHAR(10) NOT NULL,
        numero_rastreamento VARCHAR(100),
        tipo_encomenda VARCHAR(50) NOT NULL,
        descricao TEXT,
        empresa_entrega VARCHAR(100),
        observacoes TEXT,
        nome_entregador VARCHAR(100),
        cpf_rg_entregador VARCHAR(20),
        nome_porteiro VARCHAR(100) NOT NULL,
        nome_retirou VARCHAR(100)
    );

    -- Tabela para funcionários domésticos
    CREATE TABLE funcionario_domestico (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(100) NOT NULL,
        apartamento VARCHAR(10) NOT NULL,
        bloco VARCHAR(10) NOT NULL,
        funcao VARCHAR(100) NOT NULL,
        horario VARCHAR(100) NOT NULL,
        telefone VARCHAR(15) NOT NULL
    );

    -- Tabela para funcionários do condomínio
    CREATE TABLE funcionario_condominio (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(100) NOT NULL,
        endereco VARCHAR(255) NOT NULL,
        bairro VARCHAR(100) NOT NULL,
        cep VARCHAR(10) NOT NULL,
        cidade VARCHAR(100) NOT NULL,
        uf CHAR(2) NOT NULL,
        telefone VARCHAR(15) NOT NULL,
        celular VARCHAR(15) NOT NULL,
        email VARCHAR(254) NOT NULL,
        funcao_cargo VARCHAR(100) NOT NULL,
        horario_trabalho VARCHAR(100) NOT NULL,
        admitido_em DATE NOT NULL,
        observacoes TEXT
    );
    ```

## Instalação

1. Clone o repositório:
    ```sh
    git clone https://github.com/GilmarGomes12/Projeto-Condominio.git
    cd condominio-varandas-do-praia
    ```

2. Instale as dependências do backend:
    ```sh
    go mod download
    ```

3. Configure o banco de dados:
    - Crie um banco de dados PostgreSQL.
    - Configure as variáveis de ambiente necessárias no arquivo `.env`.

## Execução

1. Inicie o servidor backend:
    ```sh
    go run main.go
    ```

2. Abra o navegador e acesse `http://localhost:8080`.

## Contribuição

1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Commit suas mudanças (`git commit -m 'Adiciona nova feature'`).
4. Faça o push para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
