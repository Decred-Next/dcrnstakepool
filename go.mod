module github.com/Decred-Next/dcrnstakepool

go 1.12

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Decred-Next/dcrnd/blockchain/stake/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/blockchain/standalone/v8 v8.0.11
	github.com/Decred-Next/dcrnd/certgen/v8 v8.0.11
	github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8 v8.0.11
	github.com/Decred-Next/dcrnd/chaincfg/v8 v8.0.11
	github.com/Decred-Next/dcrnd/dcrec/v8 v8.0.11
	github.com/Decred-Next/dcrnd/dcrutil/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/hdkeychain/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/rpcclient/version4/v8 v8.0.10
	github.com/Decred-Next/dcrnd/wire/v8 v8.0.11
	github.com/Decred-Next/dcrndata/v8 v8.0.2
	github.com/Decred-Next/dcrnwallet/rpc/jsonrpc/types/version14/v8 v8.0.12
	github.com/Decred-Next/dcrnwallet/wallet/version3/v8 v8.0.12
	github.com/Decred-Next/slog/v8 v8.0.1
	github.com/apoydence/onpar v0.0.0-20190519213022-ee068f8ea4d1 // indirect
	github.com/dchest/captcha v0.0.0-20170622155422-6a29415a8364
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-gorp/gorp v2.0.1-0.20181104192722-f3677d4a0a88+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/csrf v1.6.2
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/jessevdk/go-flags v1.4.1-0.20200711081900-c17162fe8fd7
	github.com/jrick/logrotate v1.0.0
	github.com/poy/onpar v0.0.0-20190519213022-ee068f8ea4d1 // indirect
	github.com/prometheus/client_golang v1.12.2
	github.com/zenazn/goji v0.9.0
	github.com/ziutek/mymysql v1.5.4 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	google.golang.org/grpc v1.32.0
)


replace (
	github.com/Decred-Next/dcrnd/chaincfg/v8 => ../dcrnd/chaincfg
	github.com/Decred-Next/dcrnd/wire/v8 => ../dcrnd/wire

)