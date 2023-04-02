package config

import (
	"database/sql"
	"fmt"
)

func ConnectToDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	// dsn = "mysql", "gyun@tcp(localhost:3306)/balansnack" (사용자 이름:비밀번호@tcp(호스트 or IP:포트 번호)/데이터베이스 이름)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %s", err)
	}

	return db, nil
}