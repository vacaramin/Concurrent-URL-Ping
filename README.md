# Concurrent-URL-Ping

Concurrent-URL-Ping is a simple Go application designed to demonstrate how to make concurrent HTTP GET requests to a specified URL. The application hits the given URL 1000 times using goroutines, showcasing basic concurrency patterns in Go.

## Introduction

This project serves as an educational tool for understanding the basics of concurrent programming in Go, including the use of goroutines and synchronization mechanisms like WaitGroups. It targets a single URL with 1000 concurrent GET requests and logs the completion of each request.

## Getting Started

### Prerequisites

To run this project, you will need:
- Go

### Customization
You can modify the main.go file to change the target URL or the number of requests made by adjusting the url variable and the loop condition in the main function, respectively.

### Disclaimer
This tool is for educational purposes only. Please ensure you have permission to hit the target URL with a high number of requests to avoid unintended denial-of-service implications.