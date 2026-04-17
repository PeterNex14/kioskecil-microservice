module github.com/PeterNex14/kioskecil-microservice/user-service

go 1.24

replace github.com/PeterNex14/kioskecil-microservice/common => ../common

require (
	github.com/PeterNex14/kioskecil-microservice/common v0.0.0
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.12.3
)
