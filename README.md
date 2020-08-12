# Stably find lower prime number

This repo mean to host a service where can find nearest smaller prime number with the input

## Technologis

- Go as backend
  - Using mux for routing

- Algorithm
  - Using Segmented Sieve of Eratosthenes for finding smaller prime number
  - Fixed step for each finding at 100

- Frontend as Vuejs
  - Client side rendering

- CI/CD
  - Docker image for both code compiler and serving
  - Easier job with provided script

## How to use

1. Checkout project
2. Make sure you have `make` installed
3. Make sure you have `Docker` and `docker-compose` installed

Staying at root folder, trigger script:

- `./script/run_local.sh` for running solution at port `8080`
- `./script/deploy_heroku.sh` for deploying solution to heroku server, hosted at <http://stably-prime-number.herokuapp.com/>
- `./script/build.sh` for building solution on specific platform
