# devops

**Build Commands:**

- go build -o ./build/currency-conversion ./currency-conversion/.
- go build -o ./build/exchange-rates ./exchange-rates/.

**Docker Commands:**

1. Build Docker Image:  
    docker build -t devops .
2. Start a network:  
    docker network create currency-network
3. Start redis connected to this network:  
    docker run --rm --net currency-network --name my_redis -p 6379:6379 -d redis
4. Start exchange-rates container:  
    docker run -it --net currency-network --name exchange_rate -p 9001:9001 devops
5. Start the other service:  
    docker run -it --net currency-network --name currency_conversion -p 9000:9000 devops ./currency-conversion
