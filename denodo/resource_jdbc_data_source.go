package denodo

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceJDBCDataSource() *schema.Resource {
	return &schema.Resource{
		CreateContext: createJDBCDataSource,
		DeleteContext: deleteJDBCDataSource,
		ReadContext:   readJDBCDataSource,
		UpdateContext: updateJDBCDataSource,
		Schema: map[string]*schema.Schema{
			"batch_insert_size": &schema.Schema{
				Description: "When the data source has to insert several rows into the database of this data source, it can insert them in batches. This number sets the number of queries per batch.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"bcp_executable_location": &schema.Schema{
				Description: "When the data source has to insert several rows into the database of this data source, it can insert them in batches. This number sets the number of queries per batch.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"class_path": &schema.Schema{
				Description: "Path to the JAR file containing the JDBC driver for the specified source (optional).",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"denodo_database": &schema.Schema{
				Description: "Database where the data source will be created.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"data_source_database_type": &schema.Schema{
				Description: "Data source database type.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"data_source_database_version": &schema.Schema{
				Description: "Data source database version.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"data_source_description": &schema.Schema{
				Default:     "",
				Description: "The description of the data source.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"database_uri": &schema.Schema{
				Description: "The connection URL to the database.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"driver_class_name": &schema.Schema{
				Description: "The driver class to be used for connection to the data source.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"exhausted_action": &schema.Schema{
				Default:     "1",
				Description: "Specifies the behavior of the pool when the pool is empty (all the connections are running queries). (Default value: 1)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"fetch_size": &schema.Schema{
				Default:     "1000",
				Description: "Gives the JDBC driver a hint as to the number of rows that should be fetched from the database when more rows are needed.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"folder": &schema.Schema{
				Default:     "/",
				Description: "Name of the folder where the data source will be stored.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"ignore_trailing_spaces": &schema.Schema{
				Default:     "TRUE",
				Description: "If true, the Server removes the space characters at the end of text type values of the results returned by these data source’s views.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"initial_size": &schema.Schema{
				Default:     "4",
				Description: "Number of connections with which the pool is initialized. (default value: 4)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"isolation_level": &schema.Schema{
				Default:     "TRANSACTION_READ_COMMITTED",
				Description: "Sets the desired isolation level for the queries and transactions executed in the database. If not present, the data source uses the default isolation level of the database.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"jdbc_driver_properties": &schema.Schema{
				Description: "List of name/value pairs that will be passed to the JDBC driver when creating connection with this database.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"kerberos_properties": &schema.Schema{
				Description: "List of name/value pairs that will be passed to the JDBC driver when creating connection with this database. The properties on this list are meant to configure the Kerberos authentication mechanism between the Virtual DataPort server and the database. ",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"max_active": &schema.Schema{
				Default:     "20",
				Description: "Maximum number of connections that can be opened at a given time. These are the connections currently used to run queries plus the connections idle in the pool. (default value: 20)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"max_idle": &schema.Schema{
				Default:     "-1",
				Description: "Maximum number of connections that can sit idle in the pool at any time. (default value: -1)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"max_open_prepared_statements": &schema.Schema{
				Default:     "-1",
				Description: "Maximum number of opened prepared statements. Only taken into account if POOLPREPAREDSTATEMENTS is true.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"min_evictable_time": &schema.Schema{
				Default:     "1800000",
				Description: "Minimum amount of time in milliseconds that a connection may sit idle in the pool before it is eligible for eviction. (default value: 1800000 - 30 minutes)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"name": &schema.Schema{
				Description: "Name of the data source to be created.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"number_test_per_eviction": &schema.Schema{
				Default:     "3",
				Description: "Number of connections examined in each run of the connection eviction thread. The reason for not examining all the connections of the pool in each run of the thread is because while a connection is being examined, it cannot be used by a query. (default value: 3)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"on_move_read": &schema.Schema{
				Description: "if true, when the Execution Engine reads data from the Netezza database to perform a data movement, it will do so using its “External tables” feature. Setting this to true is equivalent to selecting the check box “Use external tables for data movement” of the “Read settings”, on the “Read & Write” tab of the data source",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"on_move_write": &schema.Schema{
				Description: "if true, when the Execution Engine writes data to this database to perform a data movement, it does so using its proprietary API. Setting this to yes is equivalent to selecting the check box “Use Bulk Data Load APIs” of the “Write settings”, on the “Read & Write” tab of the data source.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"pool_prepared_statements": &schema.Schema{
				Default:     "false",
				Description: "if true, the pool of prepared statements is enabled.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"sql_ldr_executable_location": &schema.Schema{
				Description: "SQLDR executable location to use bulk data load APIs. (Oracle)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"target_catalog": &schema.Schema{
				Description: "Target catalog to use for bulk load api work.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"target_schema": &schema.Schema{
				Description: "Target schema to use for bulk load api work.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"time_between_eviction": &schema.Schema{
				Default:     "-1",
				Description: "How long in milliseconds the eviction thread should sleep before “runs” of examining idle connections. If negative, the eviction thread is not launched. (default value: -1)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"test_on_borrow": &schema.Schema{
				Default:     "true",
				Description: "if true and the parameter VALIDATIONQUERY is not empty, the pool will execute the VALIDATIONQUERY on the selected connection before returning it to the execution engine. If the VALIDATIONQUERY fails, the pool will discard the connection and select another one. If there are no more idle connections, it will create one. (default value: true)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"test_on_return": &schema.Schema{
				Default:     "false",
				Description: " if true and the parameter VALIDATIONQUERY is not empty, the pool will execute the VALIDATIONQUERY on the connections returned to the pool. If the VALIDATIONQUERY fails, the pool will discard the connection. (default value: false)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"test_while_idle": &schema.Schema{
				Default:     "false",
				Description: "If true and the parameter VALIDATIONQUERY is not empty, the connections will be validated by the connection eviction thread. To validate a connection, the thread runs the validation query on the connection. If it fails, the pool drops the connection from the pool. (default value: false)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"user_for_query_optimization": &schema.Schema{
				Description: "Password to connect to the data source.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"username": &schema.Schema{
				Description: "Username to connect to the data source.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"user_password": &schema.Schema{
				Description: "Password to connect to the data source.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"validation_query": &schema.Schema{
				Default:     "",
				Description: "SQL query executed by the connection pool to check if a connection is still valid; also known as “ping query”. It is only used when at least one of TESTONBORROW, TESTONRETURN or TESTWHILEIDLE are true. (default value: depends on the adapter)",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"work_dir": &schema.Schema{
				Description: "Work directory used by bulk load configuration. (Oracle only)",
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func createJDBCDataSource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var batchInsertSize string
	var bcpExecutableLocation string
	var classPath string
	var client *Client
	var dataSourceDatabaseType string
	var dataSourceDatabaseVersion string
	var dataSourceDescription string
	var databaseURI string
	var denodoDatabase string
	var diags diag.Diagnostics
	var driverClassName string
	var err error
	var exhaustedAction string
	var fetchSize string
	var folder string
	var ignoreTrailingSpaces string
	var initialSize string
	var isolationLevel string
	var jdbcDriverProperties string
	var kerberosProperties string
	var maxActive string
	var maxIdle string
	var maxOpenPreparedStatements string
	var minEvictableTime string
	var minIdle string
	var name string
	var numberTestPerEviction string
	var onMoveRead string
	var onMoveWrite string
	var poolPreparedStatements string
	var sqlldrExecutableLocation string
	var sqlStmt string
	var targetCatalog string
	var targetSchema string
	var timeBetweenEviction string
	var testOnBorrow string
	var testOnReturn string
	var testWhileIdle string
	var username string
	var userPassword string
	var useForQueryOptimization string
	var validationQuery string
	var workDir string

	batchInsertSize = d.Get("batch_insert_size").(string)
	bcpExecutableLocation = d.Get("bcp_executable_location").(string)
	classPath = d.Get("class_path").(string)
	denodoDatabase = d.Get("denodo_database").(string)
	dataSourceDatabaseType = d.Get("data_source_database_type").(string)
	dataSourceDatabaseVersion = d.Get("data_source_database_version").(string)
	dataSourceDescription = d.Get("data_source_description").(string)
	databaseURI = d.Get("database_uri").(string)
	driverClassName = d.Get("driver_class_name").(string)
	exhaustedAction = d.Get("exhausted_action").(string)
	fetchSize = d.Get("fetch_size").(string)
	folder = d.Get("folder").(string)
	ignoreTrailingSpaces = d.Get("ignore_trailing_spaces").(string)
	initialSize = d.Get("initial_size").(string)
	isolationLevel = d.Get("isolation_level").(string)
	jdbcDriverProperties = d.Get("jdbc_driver_properties").(string)
	kerberosProperties = d.Get("kerberos_properties").(string)
	maxActive = d.Get("max_active").(string)
	maxIdle = d.Get("max_idle").(string)
	maxOpenPreparedStatements = d.Get("max_open_prepared_statements").(string)
	minEvictableTime = d.Get("min_evictable_time").(string)
	minIdle = d.Get("max_idle").(string)
	name = d.Get("name").(string)
	numberTestPerEviction = d.Get("number_test_per_eviction").(string)
	onMoveRead = d.Get("on_move_read").(string)
	onMoveWrite = d.Get("on_move_write").(string)
	poolPreparedStatements = d.Get("pool_prepared_statements").(string)
	sqlldrExecutableLocation = d.Get("sql_ldr_executable_location").(string)
	targetCatalog = d.Get("target_catalog").(string)
	targetSchema = d.Get("target_schema").(string)
	timeBetweenEviction = d.Get("time_between_eviction").(string)
	testOnBorrow = d.Get("test_on_borrow").(string)
	testOnReturn = d.Get("test_on_return").(string)
	testWhileIdle = d.Get("test_while_idle").(string)
	username = d.Get("username").(string)
	userPassword = d.Get("user_password").(string)
	useForQueryOptimization = d.Get("user_for_query_optimization").(string)
	validationQuery = d.Get("validation_query").(string)
	workDir = d.Get("work_dir").(string)

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
CREATE DATASOURCE JDBC %s
FOLDER = '%s'
DRIVERCLASSNAME = '%s'
DATABASEURI = '%s'
USERNAME = '%s'
USERPASSWORD = '%s' ENCRYPTED
CLASSPATH = '%s'
DATABASENAME = '%s'
DATABASEVERSION = '%s'
ISOLATIONLEVEL = %s
IGNORETRAILINGSPACES = %s
FETCHSIZE = %s
VALIDATIONQUERY = %s
INITIALSIZE = %s
MAXIDLE = %s
MINIDLE = %s
MAXACTIVE = %s
EXHAUSTEDACTION = %s
TESTONBORROW = %s
TESTONRETURN = %s
TESTWHILEIDLE = %s
TIMEBETWEENEVICTION = %s
NUMTESTPEREVICTION = %s
MINEVICTABLETIME = %s
POOLPREPAREDSTATEMENTS = %s
MAXOPENPREPAREDSTATEMENTS = %s`,
		denodoDatabase,
		name,
		folder,
		driverClassName,
		databaseURI,
		username,
		userPassword,
		classPath,
		dataSourceDatabaseType,
		dataSourceDatabaseVersion,
		isolationLevel,
		ignoreTrailingSpaces,
		fetchSize,
		validationQuery,
		initialSize,
		maxIdle,
		minIdle,
		maxActive,
		exhaustedAction,
		testOnBorrow,
		testOnReturn,
		testWhileIdle,
		timeBetweenEviction,
		numberTestPerEviction,
		minEvictableTime,
		poolPreparedStatements,
		maxOpenPreparedStatements,
	)

	if jdbcDriverProperties != "" {
		sqlStmt += fmt.Sprintf(
			`
PROPERTIES (
	%s
)`,
			jdbcDriverProperties,
		)
	}

	if kerberosProperties != "" {
		sqlStmt += fmt.Sprintf(
			`
KERBEROSPROPERTIES (
	%s
)`,
			kerberosProperties,
		)
	}

	sqlStmt += `
DATA_LOAD_CONFIGURATION (`

	if useForQueryOptimization != "" {
		sqlStmt += fmt.Sprintf(
			`
	USE_FOR_QUERY_OPTIMIZATION = %s`,
			useForQueryOptimization,
		)
	}

	if batchInsertSize != "" {
		sqlStmt += fmt.Sprintf(
			`
	BATCHINSERTSIZE = %s`,
			batchInsertSize,
		)
	}

	if workDir != "" || sqlldrExecutableLocation != "" || bcpExecutableLocation != "" {
		sqlStmt += `
	BULK_LOAD_CONFIGURATION (`

		if workDir != "" {
			sqlStmt += fmt.Sprintf(
				`
		WORK_DIR = '%s'`,
				workDir,
			)
		}

		if sqlldrExecutableLocation != "" {
			sqlStmt += fmt.Sprintf(
				`
		SQLLDR_EXECUTABLE_LOCATION = '%s'`,
				sqlldrExecutableLocation,
			)
		}

		if bcpExecutableLocation != "" {
			sqlStmt += fmt.Sprintf(
				`
		BCP_EXECUTABLE_LOCATION = '%s'`,
				bcpExecutableLocation,
			)
		}
	}

	if targetCatalog != "" {
		sqlStmt += fmt.Sprintf(
			`
	TARGET_CATALOG = %s`,
			targetCatalog,
		)
	}

	if targetSchema != "" {
		sqlStmt += fmt.Sprintf(
			`
	TARGET_SCHEMA = %s`,
			targetSchema,
		)
	}

	sqlStmt += fmt.Sprintf(
		`
	USEEXTERNALTABLES (
		ONMOVEREAD = %s,
		ONMOVEWRITE = %s
	)`,
		onMoveRead,
		onMoveWrite,
	)

	sqlStmt += `
)`

	sqlStmt += fmt.Sprintf(
		`
DESCRIPTION = '%s';`,
		dataSourceDescription,
	)
	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("name").(string))

	diags = readJDBCDataSource(ctx, d, meta)

	return diags
}

func deleteJDBCDataSource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var denodoDatabase string
	var diags diag.Diagnostics
	var err error
	var name string
	var sqlStmt string

	denodoDatabase = d.Get("denodo_database").(string)
	name = d.Id()
	sqlStmt = fmt.Sprintf(
		"CONNECT DATABASE %; DROP DATASOURCE JDBC %s;",
		denodoDatabase,
		name,
	)
	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId("")
	}

	return diags
}

func readJDBCDataSource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var client *Client
	var denodoDatabase string
	var diags diag.Diagnostics
	var err error
	var name string
	var resultSet [][]string
	var sqlStmt string

	denodoDatabase = d.Get("denodo_database").(string)
	name = d.Id()
	sqlStmt = fmt.Sprintf(
		"CONNECT DATABASE %s; DESC DATASOURCE JDBC %s;",
		denodoDatabase,
		name,
	)
	client = meta.(*Client)

	resultSet, err = client.ResultSet(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", resultSet[0][0])
	d.Set("denodo_database", resultSet[0][1])
	d.Set("database_uri", resultSet[0][2])
	d.Set("driver_class_name", resultSet[0][3])
	d.Set("username", resultSet[0][4])
	d.Set("data_source_database_type", resultSet[0][5])
	d.Set("data_source_database_version", resultSet[0][6])
	d.Set("validation_query", resultSet[0][7])
	d.Set("initial_size", resultSet[0][8])
	d.Set("max_active", resultSet[0][9])

	return diags
}

func updateJDBCDataSource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var batchInsertSize string
	var bcpExecutableLocation string
	var classPath string
	var client *Client
	var dataSourceDatabaseType string
	var dataSourceDatabaseVersion string
	var dataSourceDescription string
	var databaseURI string
	var denodoDatabase string
	var diags diag.Diagnostics
	var driverClassName string
	var err error
	var exhaustedAction string
	var fetchSize string
	var ignoreTrailingSpaces string
	var initialSize string
	var isolationLevel string
	var jdbcDriverProperties string
	var kerberosProperties string
	var maxActive string
	var maxIdle string
	var maxOpenPreparedStatements string
	var minEvictableTime string
	var minIdle string
	var name string
	var numberTestPerEviction string
	var onMoveRead string
	var onMoveWrite string
	var poolPreparedStatements string
	var sqlldrExecutableLocation string
	var sqlStmt string
	var targetCatalog string
	var targetSchema string
	var timeBetweenEviction string
	var testOnBorrow string
	var testOnReturn string
	var testWhileIdle string
	var username string
	var userPassword string
	var useForQueryOptimization string
	var validationQuery string
	var workDir string

	batchInsertSize = d.Get("batch_insert_size").(string)
	bcpExecutableLocation = d.Get("bcp_executable_location").(string)
	classPath = d.Get("class_path").(string)
	denodoDatabase = d.Get("denodo_database").(string)
	dataSourceDatabaseType = d.Get("data_source_database_type").(string)
	dataSourceDatabaseVersion = d.Get("data_source_database_version").(string)
	dataSourceDescription = d.Get("data_source_description").(string)
	databaseURI = d.Get("database_uri").(string)
	driverClassName = d.Get("driver_class_name").(string)
	exhaustedAction = d.Get("exhausted_action").(string)
	fetchSize = d.Get("fetch_size").(string)
	ignoreTrailingSpaces = d.Get("ignore_trailing_spaces").(string)
	initialSize = d.Get("initial_size").(string)
	isolationLevel = d.Get("isolation_level").(string)
	jdbcDriverProperties = d.Get("jdbc_driver_properties").(string)
	kerberosProperties = d.Get("kerberos_properties").(string)
	maxActive = d.Get("max_active").(string)
	maxIdle = d.Get("max_idle").(string)
	maxOpenPreparedStatements = d.Get("max_open_prepared_statements").(string)
	minEvictableTime = d.Get("min_evictable_time").(string)
	minIdle = d.Get("max_idle").(string)
	name = d.Get("name").(string)
	numberTestPerEviction = d.Get("number_test_per_eviction").(string)
	onMoveRead = d.Get("on_move_read").(string)
	onMoveWrite = d.Get("on_move_write").(string)
	poolPreparedStatements = d.Get("pool_prepared_statements").(string)
	sqlldrExecutableLocation = d.Get("sql_ldr_executable_location").(string)
	targetCatalog = d.Get("target_catalog").(string)
	targetSchema = d.Get("target_schema").(string)
	timeBetweenEviction = d.Get("time_between_eviction").(string)
	testOnBorrow = d.Get("test_on_borrow").(string)
	testOnReturn = d.Get("test_on_return").(string)
	testWhileIdle = d.Get("test_while_idle").(string)
	username = d.Get("username").(string)
	userPassword = d.Get("user_password").(string)
	useForQueryOptimization = d.Get("user_for_query_optimization").(string)
	validationQuery = d.Get("validation_query").(string)
	workDir = d.Get("work_dir").(string)

	sqlStmt = fmt.Sprintf(
		`
CONNECT DATABASE %s;
ALTER DATASOURCE JDBC %s
DRIVERCLASSNAME = '%s'
DATABASEURI = '%s'
USERNAME = '%s'
USERPASSWORD = '%s' ENCRYPTED
CLASSPATH = '%s'
DATABASENAME = '%s'
DATABASEVERSION = '%s'
ISOLATIONLEVEL = %s
IGNORETRAILINGSPACES = %s
FETCHSIZE = %s
VALIDATIONQUERY = %s
INITIALSIZE = %s
MAXIDLE = %s
MINIDLE = %s
MAXACTIVE = %s
EXHAUSTEDACTION = %s
TESTONBORROW = %s
TESTONRETURN = %s
TESTWHILEIDLE = %s
TIMEBETWEENEVICTION = %s
NUMTESTPEREVICTION = %s
MINEVICTABLETIME = %s
POOLPREPAREDSTATEMENTS = %s
MAXOPENPREPAREDSTATEMENTS = %s`,
		denodoDatabase,
		name,
		driverClassName,
		databaseURI,
		username,
		userPassword,
		classPath,
		dataSourceDatabaseType,
		dataSourceDatabaseVersion,
		isolationLevel,
		ignoreTrailingSpaces,
		fetchSize,
		validationQuery,
		initialSize,
		maxIdle,
		minIdle,
		maxActive,
		exhaustedAction,
		testOnBorrow,
		testOnReturn,
		testWhileIdle,
		timeBetweenEviction,
		numberTestPerEviction,
		minEvictableTime,
		poolPreparedStatements,
		maxOpenPreparedStatements,
	)

	if jdbcDriverProperties != "" {
		sqlStmt += fmt.Sprintf(
			`
PROPERTIES (
	%s
)`,
			jdbcDriverProperties,
		)
	}

	if kerberosProperties != "" {
		sqlStmt += fmt.Sprintf(
			`
KERBEROSPROPERTIES (
	%s
)`,
			kerberosProperties,
		)
	}

	sqlStmt += `
DATA_LOAD_CONFIGURATION (`

	if useForQueryOptimization != "" {
		sqlStmt += fmt.Sprintf(
			`
	USE_FOR_QUERY_OPTIMIZATION = %s`,
			useForQueryOptimization,
		)
	}

	if batchInsertSize != "" {
		sqlStmt += fmt.Sprintf(
			`
	BATCHINSERTSIZE = %s`,
			batchInsertSize,
		)
	}

	if workDir != "" || sqlldrExecutableLocation != "" || bcpExecutableLocation != "" {
		sqlStmt += `
	BULK_LOAD_CONFIGURATION (`

		if workDir != "" {
			sqlStmt += fmt.Sprintf(
				`
		WORK_DIR = '%s'`,
				workDir,
			)
		}

		if sqlldrExecutableLocation != "" {
			sqlStmt += fmt.Sprintf(
				`
		SQLLDR_EXECUTABLE_LOCATION = '%s'`,
				sqlldrExecutableLocation,
			)
		}

		if bcpExecutableLocation != "" {
			sqlStmt += fmt.Sprintf(
				`
		BCP_EXECUTABLE_LOCATION = '%s'`,
				bcpExecutableLocation,
			)
		}
	}

	if targetCatalog != "" {
		sqlStmt += fmt.Sprintf(
			`
	TARGET_CATALOG = %s`,
			targetCatalog,
		)
	}

	if targetSchema != "" {
		sqlStmt += fmt.Sprintf(
			`
	TARGET_SCHEMA = %s`,
			targetSchema,
		)
	}

	sqlStmt += fmt.Sprintf(
		`
	USEEXTERNALTABLES (
		ONMOVEREAD = %s,
		ONMOVEWRITE = %s
	)`,
		onMoveRead,
		onMoveWrite,
	)

	sqlStmt += `
)`

	sqlStmt += fmt.Sprintf(
		`
DESCRIPTION = '%s';`,
		dataSourceDescription,
	)

	client = meta.(*Client)

	err = client.ExecuteSQL(&sqlStmt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("name").(string))

	diags = readJDBCDataSource(ctx, d, meta)

	return diags
}
