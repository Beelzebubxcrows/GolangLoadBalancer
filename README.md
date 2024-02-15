# GolangLoadBalancer

This project is a simple implementation of a load balancer in GoLang. It includes two load balancing algorithms: Round Robin and Least Connection. The Round Robin algorithm selects servers in a circular order, while the Least Connection algorithm selects the server with the fewest active connections.

## Key Features

- Implements two load balancing algorithms: Round Robin and Least Connection.
- Provides a simple interface for adding and removing servers.
- Includes a basic HTTP server for testing the load balancer.
- Allows users to choose the load balancing algorithm at runtime.

## Info

This project currently includes three endpoints, where request are redirected to imitating potential servers:

- [https://leetcode.com/](https://leetcode.com/)
- [https://www.facebook.com/](https://www.facebook.com/)
- [https://www.google.com/](https://www.google.com/)

  
## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-repo.git

2. Navigate to the project directory: `cd GolangLoadBalancer`

3. Run the project: `go run main.go`

4. Open a web browser and navigate to `http://localhost:8080/` to access the load balancer.


## Usage

When you hit `http://localhost:8080/` in your web browser, the server will ask you to choose a load balancing algorithm:

- Enter `0` for Round Robin.
- Enter `1` for Least Connection.

After choosing an algorithm, the server will use the selected algorithm to serve your request.

If you choose `Round Robin`, the server will forward your request to one of the following dummy servers.

Similarly, if you choose `Least Connection`, the server will forward your request to the dummy server with the fewest active connections.

