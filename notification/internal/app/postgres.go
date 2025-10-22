package app

import "fmt"

var (
	//PostgresHost     = "localhost"
	//PostgresPort     = "5435"
	PostgresHost     = "postgresql-notification"
	PostgresPort     = "5432"
	PostgresDB       = "postgres-notification"
	PostgresUser     = "postgres-notification-user"
	PostgresPassword = "postgres-notification-psw"
)

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDB,
	)
}
