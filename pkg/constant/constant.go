package constant

// App
const AppName = "Go-Microservice-Template"

const (
	AppEnvProd = "prod"
	AppEnvDev  = "dev"
	AppEnvTest = "test"
)

// Http + Grpc
const (
	HttpHost      = "localhost"
	HttpPort      = 4000
	EchoGzipLevel = 5

	GrpcHost = "localhost"
	GrpcPort = 3000
)

// Postgres
const (
	PgMaxConn         = 1
	PgMaxIdleConn     = 1
	PgMaxLifeTimeConn = 1
	PgSslMode         = "disable"
)

const (
	AccessTokenHint  = "access_token"
	RefreshTokenHint = "refresh_token"
)

const (
	StorageSessionName = "go_oauth2_server_session"
	UserSessionKey     = "go_oauth2_server_user"
)
