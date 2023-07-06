# golang_api

## Setup 
This project uses a Makefile for boilerplate. This includes a dockerized postgres instance for persistence. 
To get setup, ensure you have Docker installed on your system and run the following command:
```shell
make run_db
```
This will run a docker container that has a postgres DB in it.

Then, run:
```shell
make run
```