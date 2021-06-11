package cmd

import (
	"flag"
	"fmt"

	"loadtest.github.com/lib"
)

var cmdLineOptions lib.Options = lib.NewOptions()

func init()  {

	flag.StringVar(&cmdLineOptions.Method, "method", "GET", "method to url")
	flag.StringVar(&cmdLineOptions.Method, "m", "GET", "method to url")

	flag.StringVar(&cmdLineOptions.Body, "body", "", "Send string as Request body")
	flag.StringVar(&cmdLineOptions.Body, "b", "", "Send string as Request body")
 
	flag.StringVar(&cmdLineOptions.File, "file", "", "Send the contents of the file as Request body")
	flag.StringVar(&cmdLineOptions.File, "f", "", "Send the contents of the file as Request body")

	flag.IntVar(&cmdLineOptions.MaxRequests, "maxRequests", -1, "Number of requests to perform")
	flag.IntVar(&cmdLineOptions.MaxRequests, "n", -1, "Number of requests to perform")

	flag.IntVar(&cmdLineOptions.Concurrency, "concurrency", 1, "Number of requests to make")
	flag.IntVar(&cmdLineOptions.Concurrency, "c", 1, "Number of requests to make")

	flag.IntVar(&cmdLineOptions.RequestPerSecond, "rps", 1, "Specify the requests per second for each client")

	flag.Var(cmdLineOptions.RequestHeaders, "headers", "Send a header as header:value (multiple)")
	flag.Var(cmdLineOptions.RequestHeaders, "H", "Send a header as header:value (multiple)")

	flag.Var(cmdLineOptions.Cookies, "cookie", "Send a cookie as name=value (multiple)")
	flag.Var(cmdLineOptions.Cookies, "C", "Send a cookie as name=value (multiple)")

}

func Parse()  {
	flag.Parse()
	extraArgs := flag.Args()
	if len(extraArgs) == 0 {
		lib.HelpAndExit("Missing URL to load-test")
	} else {
		cmdLineOptions.URL = extraArgs[0]
		flag.CommandLine.Parse(extraArgs[1:])
		extraArgs = append([]string{}, cmdLineOptions.URL)
		extraArgs = append(extraArgs, flag.CommandLine.Args()...)
		if len(extraArgs) > 1 {
			lib.HelpAndExit("Too many arguments: ", fmt.Sprint(extraArgs))
		}
	}
}

func Loadtest()  {
	Parse()
	cmdLineOptions.ExecuteTest()
}