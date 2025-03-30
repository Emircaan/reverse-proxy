This project implements a gRPC-based reverse proxy server with load balancing and middleware functionalities like rate limiting, logging, and IP-based access control.

Overview
Reverse Proxy: Routes incoming gRPC requests to multiple backend servers.

Load Balancer: Distributes requests evenly across available backend servers.

Middleware: Processes requests through a chain of functions such as logging and rate limiting.

