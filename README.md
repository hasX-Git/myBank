# Bank API
## How to use
### Pull the repository
Pull the repository on your local machine with

    git clone https://github.com/hasX-Git/myBank.git

### Working with .example files
There are 3 example files: .env.example, Dockerfile.Example, and docker-compose_example.yml.

These are templates for .env, Dockerfile and docker-compose.yml respectively.

Wherever there is <ins>#change</ins> comment, you can change those to any preferred name.

Wherever there is <ins>#AP</ins>, which stands for App Port, it must be the same for all ports with #AP comment. For example, if in .env file the port for #AP is 5432, it must be 5432 for #AP in Dockerfile and docker-compose. Same with <ins>#DP</ins>, Database Port

Before running programs, delete .gitkeep in folders "files" and "db-data"

### Running program
Run the following line

    docker compose up --build