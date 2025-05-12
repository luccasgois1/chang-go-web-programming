# Chitchat app

App where you can create threads and dicuss interesting topics as you
would in a blog page.

This app was developed to run using Docker containers, the reason for that is
to reduce problems related to different types of OSs, which make it easy to deploy and
use.

## Requirements
* Docker version 27.5.1


## How to run
### Customizing your enviroment
```bash
# Create your .env file in the root dir of your application
cp .env.example .env
```

The following variables should be modified if you want to change the default configurations
```
POSTGRES_USER=<PUT THE USER THAT YOU WANT TO ACCESS YOUR DATABASE>
POSTGRES_PASSWORD=<PUT THE PASSWORD THAT YOU WANT TO ACCESS THE DATABASE>
POSTGRES_HOST=<USE localhost IF YOUR DATABASE IS NOT RUNNING IN DOCKER>
```

### Run your application using Docker
```bash
# To start the app server run:
docker compose up -d
# When you want to destroy the app, just run:
docker compose down
```
