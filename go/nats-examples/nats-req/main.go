package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"time"
)

func usage() {
	log.Printf("Usage: nats-req [-s server] [-creds file] -nKey file <subject> <msg> \n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

/**
* @link https://github.dev/nats-io/nats.go/tree/main/examples/nats-rply
 */
func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var userCreds = flag.String("creds", "", "User Credentials File")
	var nkeyFile = flag.String("nkey", "", "NKey Seed File")
	var showHelp = flag.Bool("h", false, "Show help message")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	args := flag.Args()
	fmt.Println(args, len(args))
	if len(args) < 2 {
		showUsageAndExit(1)
	}

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Requestor")}

	if *userCreds != "" && *nkeyFile != "" {
		log.Fatal("specify -seed or -creds")
	}

	// Use UserCredentials
	if *userCreds != "" {
		opts = append(opts, nats.UserCredentials(*userCreds))
	}

	// Use Nkey authentication.
	if *nkeyFile != "" {
		opt, err := nats.NkeyOptionFromSeed(*nkeyFile)
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, opt)
	}

	// Connect to NATS
	fmt.Println("connect")
	nc, err := nats.Connect(*urls, opts...)
	if err != nil {
		fmt.Println("CONNECT ERROR!!")
		log.Fatal(err)
	}
	defer nc.Close()
	subj, payload := args[0], []byte(args[1])

	msg, err := nc.Request(subj, payload, 2*time.Second)
	if err != nil {

		fmt.Println("NC ERROR!!")

		if nc.LastError() != nil {
			log.Fatalf("%v for request", nc.LastError())
		}
		log.Fatalf("%v for request", err)
	}

	log.Printf("Published [%s] : '%s'", subj, payload)
	log.Printf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))
}
