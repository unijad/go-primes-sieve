# go primes
this is an experamintal project for golang, serving a wrapper project over primes sieve algorithm.

### requirements

Create a web application that takes in a number and returns to the user the highest prime number lower than the input number. For example, an input of 55 would return 53. Treat this like a real production application that you would be proud of.

### project architecture

*technologies to be used:*

- golang/gogin *base server*
- typescript/react-jsx *client side*
- sass *styles*
- docker *one config, one container*
- workbox *offline caching, PWA*

*project files*

- **build** *binary build folder*
- **src**
    - **assets** *tsx/sass and other client side assets*
    - **public** *public exposed folder*
    - **views** *html templates*
    - main.go
    - home.controller.go
    - primes.controller.go
- .gitignore
- .babelrc
- .sassrc
- go.mod
- Dockerfile
- package.json
- tsconfig.json
- workbox-config.js

## prime function algorithm

Sieve of Eratosthenes allows us to generate a list of primes. with preformance of
*O(n x n)*

reference:

[khan academy](https://duckduckgo.com "***sieve of eratosthenes***") ***sieve of eratosthenes***


go modules we will be using for this project

[gin-gonic](github.com/gin-gonic/gin)
