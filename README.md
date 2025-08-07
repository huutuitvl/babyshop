# File: .readme
# This file contains instructions for setting up and running the BabyShop application using Docker Compose.

```bash
docker-compose up -d
```
<!-- This command starts the services defined in Docker Compose in detached mode. -->

```bash
docker exec -it babyshop-mysql mysql -uroot -proot
```
<!-- This command connects to the MySQL service inside the container using the root user and password 'root'. -->

```bash
docker ps
```
<!-- Use this command to check if the MySQL service is running. -->

```bash
docker-compose down -v
docker-compose up -d
```
<!-- The 'down -v' command removes old MySQL data and restarts the services with the correct password. -->

```bash
docker-compose restart mysql
```
<!-- This command restarts the MySQL service in Docker Compose. -->

```bash
docker exec -it babyshop-mysql mysql -uroot -proot
```
<!-- This command connects to MySQL in the container with root user and password 'root'. -->

```bash
docker-compose down
docker-compose up -d
```
<!-- These commands remove the old MySQL data and restart the service. -->
