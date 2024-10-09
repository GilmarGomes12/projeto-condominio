package config

import (
    "database/sql"
    _ "github.com/jackc/pgx/v4/stdlib"
    "github.com/joho/godotenv"
    "go.uber.org/zap"
    "log"
    "os"
    "fmt"
)

var (
    DB *sql.DB
    logger *zap.Logger
)

// InitConfig carrega variáveis de ambiente e inicializa o logger.
func InitConfig() error {
    if err := godotenv.Load("config/.env"); err != nil {
        log.Println("Aviso: Erro ao carregar o arquivo .env. Usando variáveis de ambiente do sistema.")
    }

    var err error
    logger, err = zap.NewProduction()
    if err != nil {
        return err
    }
    return nil
}

// ConnectDB estabelece uma conexão com o banco de dados PostgreSQL.
func ConnectDB() error {
    dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = sql.Open("pgx", dbURL)
    if err != nil {
        logger.Error("Erro ao abrir a conexão com o banco de dados", zap.Error(err))
        return err
    }

    // Verifica a conexão de imediato.
    if err = DB.Ping(); err != nil {
        logger.Error("Erro ao pingar o banco de dados", zap.Error(err))
        return err
    }

    return nil
}

// CloseDB fecha a conexão com o banco de dados.
func CloseDB() {
    if err := DB.Close(); err != nil {
        logger.Error("Erro ao fechar a conexão com o banco de dados", zap.Error(err))
    }
}

// GetDB retorna a instância do banco de dados.
func GetDB() *sql.DB {
    return DB
}

// Logger retorna a instância do logger.
func Logger() *zap.Logger {
    return logger
}