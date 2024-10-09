package utils

import (
    "fmt"
    "net/smtp"
)

func EnviarEmailRecuperacaoSenha(email, token string) {
    from := "seu_email@example.com"
    password := "sua_senha"
    to := []string{email}
    smtpHost := "smtp.example.com"
    smtpPort := "587"

    mensagem := []byte(fmt.Sprintf("Para redefinir sua senha, clique no link: http://seu_dominio.com/redefinir-senha?token=%s", token))

    auth := smtp.PlainAuth("", from, password, smtpHost)
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, mensagem)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("E-mail enviado com sucesso")
}