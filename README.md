# SW coupon code auto-claimer

From a functionnal viewpoint, this project provides a Dockerized service stack for automatic coupon fetching and claiming in Summoners war.
In reality this is a very over-engineered project I'm developping as an excuse to learn and practice go, database management as well as CI/CD.

## Running the project

Since each micro-service of this app is provided as a docker image, running it locally requires to simply deploy the stack using the docker-compose file
provided in the root of the project. 

First, create a .env file containing all the variables shown [in the provided example](.env.example).
Then, deploy the stack using docker compose

```bash
docker compose up -d --build
```

To ensure the backend is alive, you can run 
```bash
curl http://$(backend-host-ip):$(backend-port)/v1/health
```

## Functionnality

Currently, the discord bot implements a single command that allows to register users 
```
/register <hive_id> <server>
```

The backend service will ensure that the registered user indeed exists and then add it to the user database. Periodically, the service will claim all fetched 
coupons that were not fetched for all registered accounts.

## Project structure

![project structure](./docs/figures/project_architecture.png)