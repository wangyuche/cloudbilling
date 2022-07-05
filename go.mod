module github.com/wangyuche/cloudbilling

go 1.18

require (
	github.com/wangyuche/cloudbilling/src/sql v0.0.0-20220704071454-d5afab0047ad
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/go-sql-driver/mysql v1.6.0 // indirect

replace github.com/wangyuche/cloudbilling/src/sql => ./src/sql
