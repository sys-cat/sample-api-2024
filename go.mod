module local/sample-api-2024

go 1.23.1

require github.com/julienschmidt/httprouter v1.3.0

replace local/sample-api-20024/handler => ./handler
