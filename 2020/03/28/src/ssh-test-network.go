package main

import (
  "flag"
  "fmt"
  "net"
  "os"
)

func main() {

  // Variable declaration
  var cli_cidr string
  var cli_debug bool
  var cli_host string

  // Flags declaration 
  flag.StringVar(&cli_cidr, "cidr", "", "CIDR network to test against")
  flag.StringVar(&cli_host, "host", "", "Hostname to test")
  flag.BoolVar(&cli_debug, "debug", false, "Enable debugging")

  // Customise usage message
  flag.Usage = Usage

  // Parse the provided CLI options
  flag.Parse()

  // Parse the cidr and resolve the hostname
  _,cidr,_ := net.ParseCIDR(cli_cidr)
  ips,_ := net.LookupIP(cli_host)

  // Iterate through all of the IP addresses resolved
  for _, ip := range ips {
    if(cidr.Contains(ip)) {

      if(cli_debug) { fmt.Println("Hostname resolves", ip) }

      // Matched, exit code 0
      os.Exit(0)
    }
  }

  // If we got here, we didn't match, exit code 1
  os.Exit(1)
}

var Usage = func() {
  fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
  flag.PrintDefaults()
  fmt.Println()
  fmt.Println("Use this command to route subnet ssh via a jumphost in your ssh config:")
  fmt.Println()
  fmt.Println("  Match exec \"~/bin/ssh-test-network --cidr=10.0.0.0/8 --host=%n")
  fmt.Println("    ProxyCommand ssh -q jumpbox.example.com -W %h:%p")
  fmt.Println("    ForwardAgent yes")
  fmt.Println()
}
