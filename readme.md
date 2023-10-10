## How to install app
- clone the repository
- create .env file in codes folder from .env.example
- in docker/.envs create app.env file from app.env.example
- in docker folder create docker-compose.override.yml from docker-compose.override.example.yml
- exec docker folder from terminal and run `docker-compose build` after successful build run `docker-compose up golang-app`