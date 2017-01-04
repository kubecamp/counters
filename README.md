# Counters
Microservice to manage counters in redis

This little go utility creates a web server that interfaces with a Redis server.

To increase a counter

        curl -XPOST localhost:9090/mycounter

To get a counter

        curl localhost:9090/docker


We provide a Dockerfile in case you want to run it in a docker environment:

        docker network create simple-network
        docker run -d -P --name redis --network simple-network redis
        docker run -p 9090:9090  -e REDIS_URL=redis:6379  --network simple-network counters

We provide a Helm Chart to deploy the application into a kubernetes cluster.

        helm install ./charts/counter --name webcounter

The counters chart will install and configure `Redis` also.

