package denodo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
)

var err error
var database string
var username string
var password string
var host string
var port string
var connOpts string

denodoConnUrl := fmt.Sprintf(
"postgres://%s:%s@%s:%s/%s?%s", // <- remove the '?sslmode=require' and have the value from conn-opts set this
username,
password,
host,
port,
database,
connOpts,
)