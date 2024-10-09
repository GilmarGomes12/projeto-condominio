package controllers

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "time"
    "condominio/models"
)

// FuncionariosHandler lida com as requisições HTTP para gerenciar funcionários do condomínio.
func FuncionariosHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            // Processa a requisição POST para criação de um novo funcionário.
            if err := r.ParseForm(); err != nil {
                http.Error(w, "Erro ao parsear o formulário", http.StatusBadRequest)
                return
            }

            // Obtém e parseia a data de admissão do formulário.
            admitidoEmString := r.FormValue("admitido_em")
            var admitidoEm time.Time
            var err error
            if admitidoEmString != "" {
                layout := "2006-01-02"
                admitidoEm, err = time.Parse(layout, admitidoEmString)
                if err != nil {
                    http.Error(w, "Erro ao parsear a data de admissão: "+err.Error(), http.StatusBadRequest)
                    return
                }
            } else {
                admitidoEm = time.Now() // Define a data atual como padrão se não fornecida.
            }

            // Cria uma nova instância de Funcionario com os dados do formulário.
            funcionario := models.Funcionarios{
                Nome:           r.FormValue("nome"),
                Endereco:       r.FormValue("endereco"),
                Bairro:         r.FormValue("bairro"),
                CEP:            r.FormValue("cep"),
                Cidade:         r.FormValue("cidade"),
                UF:             r.FormValue("uf"),
                Telefone:       r.FormValue("telefone"),
                Celular:        r.FormValue("celular"),
                Email:          r.FormValue("email"),
                FuncaoCargo:    r.FormValue("funcao_cargo"),
                HorarioTrabalho:r.FormValue("horario_trabalho"),
                AdmitidoEm:     admitidoEm,
                Observacoes:    r.FormValue("observacoes"),
            }

            log.Printf("Recebido funcionário: %+v", funcionario)

            // Insere o novo funcionário no banco de dados.
            if err := funcionario.Create(db); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Redireciona para a página de sucesso ou para a mesma página para limpar o formulário.
            http.Redirect(w, r, "/funcionarios", http.StatusSeeOther)

        case "GET":
            // Carrega o template para a página de listagem/cadastro dos funcionários.
            tmpl, err := template.ParseFiles("templates/funcionarios.html")
            if err != nil {
                log.Printf("Erro ao carregar o template: %v", err)
                http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
                return
            }
            // Executa o template e envia a resposta ao cliente.
            err = tmpl.Execute(w, nil)
            if err != nil {
                log.Printf("Erro ao executar o template: %v", err)
                http.Error(w, "Erro ao executar o template", http.StatusInternalServerError)
            }

        default:
            // Retorna erro 405 Method Not Allowed para métodos não suportados.
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    }
}