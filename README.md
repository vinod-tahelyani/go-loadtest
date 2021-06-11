# loadtest
A simple command line tool, for running quick http loads to URLs.

### Usage
##### `-n, -maxRequests`

Number of requests to perform (default infinite)
##### `-rps int`

Specify the requests per second for each client (default 1)
##### `-b, -body`

Send string as Request body
##### `-c, -concurrency `

Number of client (default 1)
##### `-f, -file`

Send the contents of the file as Request body
##### `-m, -method string`

Method to url (default "GET")
##### `-C, -cookie `

Send a cookie as name=value (multiple)
##### `-H, -headers`

Send a header as header:value (multiple)
### Build
    $ git clone git@github.com:vinod-tahelyani/go-loadtest.git
    $ make make buildLoadtest
    $ ./bin/loadtest -h

**Note: This project is currently under developement.**