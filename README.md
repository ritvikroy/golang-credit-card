# Golang-credit-card

It validates credit card details, display credit card types and reduces human
error using the Luhn algorithm using REST APIs.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Doc](https://img.shields.io/badge/go%20doc-read-blue.svg)](https://godoc.org/github.com/bitcyber/golang-credit-card)

Basic credit card validation using the [Luhn algorithm](http://en.wikipedia.org/wiki/Luhn_algorithm)

Currently identifies the following credit card companies:

- American Express
- MasterCard
- Visa
- Visa Electron

## Installation

```bash
go get github.com/bitcyber/golang-credit-card
```

## Quick Start

Run this command on Terminal after installation
```bash
go run main.go
```

```go
// Call GET method and checks weather GET is working or not:

We can test with cURL or postman , and then send a GET request to 
http://localhost:3000/card-scheme/ 
to check whether GET method is working or not, increment the likes of a joke.


// Pass credit card number to validate/verify
http://localhost:3000/card-scheme/verify/{cardNumber}

For example:
http://localhost:3000/card-scheme/verify/373473833566829

// Display credit card validation output



# Contibuting
> Any help and suggestions are very welcome and appreciated.
```