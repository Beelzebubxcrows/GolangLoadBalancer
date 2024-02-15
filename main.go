package main

import (
	"GolangLoadBalancer/pkg"
	"fmt"
	"net/http"
)

func main() {
	servers := []pkg.Server{
		pkg.CreateNewSimpleServer("https://leetcode.com/", true),
		pkg.CreateNewSimpleServer("https://www.facebook.com/", true),
		pkg.CreateNewSimpleServer("https://www.google.com/", true),
	}

	roundRobinLoadBalancer := pkg.CreateNewRoundRobinLoadBalancer("8080", servers)
	leastConnectionLoadBalancer := pkg.CreateNewLeastConnectionLoadBalancer("8080", servers)

	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("\n\nNew request received. Redirecting...")
		fmt.Print("Enter which load balancer algorithm to use: \n 1) For Round Robin - Enter \"0\" \n 2) For Least Connection - Enter \"1\" ")
		var loadBalancerToUse string
		fmt.Scan(&loadBalancerToUse)

		if loadBalancerToUse == "0" {
			fmt.Println("Using Round Robin algorithm for load balancing.")
			roundRobinLoadBalancer.ServeProxyRequest(rw, r)
		} else {
			fmt.Println("Using Least Connection for load balancing.")
			leastConnectionLoadBalancer.ServeProxyRequest(rw, r)
		}

	}
	http.HandleFunc("/", handleRedirect)
	http.ListenAndServe(":8080", nil)
}
