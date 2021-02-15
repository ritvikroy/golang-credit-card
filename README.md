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
$ git clone git@github.com:bitcyber/golang-credit-card.git
```
OR

```bash
$ go get github.com/bitcyber/golang-credit-card
```



## Downloads

Download all go-gettable dependencies for your project.

```bash
$ go get -d ./... 
```

## Quick Start

Run this command on Terminal after installation
```bash
$ go run main.go
```

Once the application is running, navigate to **http://localhost:3000** in your browser. If all went well, you should see level 1 header text **Welcome to the Credit Card Validation** displayed.


**REST API: Show Credit Card Validation**
----
  Returns json data about a single credit card.

* **URL**

  ```http
  http://localhost:3000/card-scheme/verify/:cardNumber
  ```

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `cardNumber=[integer]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ "success" : "true", "valid": "true", "scheme": "AMEX/VISA/MASTERCARD" }`

    OR

  * **Code:** 200  <br />
    **Content:** `{ "INVALID CARD NUMBER" }`
    
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ 404 page not found }`


* **Sample Request:**
    ```http
    http://localhost:3000/card-scheme/verify/373473833566829
    ```

* **Sample Response:**


  ```javascript
    {
    "success": "true",
    "valid": "true",
    "scheme": "AMEX"
    }
  ```


## Contributing ✨
> Any help and suggestions are very welcome and appreciated.
