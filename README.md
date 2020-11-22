# miniature-giggle
simple http client and server programs written in golang

# Run Server

### run with default port
`go run main.go server`
starts server on default port 3000

### run with defined port
`go run main.go server <PORT_NUM>`
starts server on above definedd PORT_NUM

# Run Client
`go run client.go client`
starts client on default URL http://localhost and port 3000

### run with defined URL and PORT
`go run main.go client <URL> <PORT_NUM>`
starts client on defined URL and PORT_NUM
