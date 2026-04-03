package repositories_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbDatabase,
	)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func execSqlFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	cmd := exec.Command(
		"docker",
		"exec",
		"-i", // 標準入力を受け付けるために必須
		"db-for-go",
		"mysql",
		"-u",
		"docker",
		"--password=docker",
		"sampledb",
	)
	cmd.Stdin = file

	// エラーメッセージを受け取る
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"実行エラー: %v\n【MySQLからの詳細メッセージ】:\n%s",
			err,
			stderr.String(),
		)
	}
	return nil
}

func setupTestData() error {
	return execSqlFile("./testdata/setupDB.sql")
}

func cleanupDB() error {
	return execSqlFile("./testdata/cleanupDB.sql")
}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanupDB", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setupDB", err)
		return err
	}
	return nil
}

func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1) // 1: 異常終了
	}

	m.Run()

	teardown()
}
