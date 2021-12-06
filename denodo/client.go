package denodo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Database     string
	Host         string
	MaxOpenConns int
	MaxIdleConns int
	Port         int
	SslMode      string
}

type Client struct {
	Connection *sql.DB
}

func (c *Config) NewClient(password, username *string) (*Client, error) {
	var client Client
	var err error

	if (password != nil) && (username != nil) {
		denodoConnUrl := fmt.Sprintf(
			"user=%s password=%s host=%s port=%d sslmode=%s database=%s",
			*username,
			*password,
			c.Host,
			c.Port,
			c.SslMode,
			c.Database,
		)

		client.Connection, err = sql.Open("postgres", denodoConnUrl)
		client.Connection.SetMaxOpenConns(c.MaxOpenConns)
		client.Connection.SetMaxIdleConns(c.MaxIdleConns)

		if err != nil {
			return nil, err
		}
	}

	return &client, nil
}

func (c *Client) ResultSet(sqlStmt *string) ([][]string, error) {
	var err error
	var results *sql.Rows

	results, err = c.Connection.Query(*sqlStmt)
	if err != nil {
		return nil, err
	}

	columns, err := results.Columns()
	if err != nil {
		return nil, err
	}

	count := len(columns)
	tableData := [][]string{}
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for results.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		results.Scan(valuePtrs...)
		record := make([]string, count)
		for i := range columns {
			switch values[i].(type) {
			case nil:
				var s string
				record[i] = s
			case []byte:
				record[i] = string(values[i].([]byte))
			default:
				record[i] = fmt.Sprintf("%s", values[i])
			}
		}
		tableData = append(tableData, record)
	}

	results.Close()

	return tableData, nil
}

func (c *Client) ExecuteSQL(sqlStmt *string) error {
	var err error

	_, err = c.Connection.Exec(*sqlStmt)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) CloseSession() error {
	var err error

	err = c.Connection.Close()

	return err
}
