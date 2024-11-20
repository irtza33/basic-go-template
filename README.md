# basic-go-template
This repository contains a template for how a basic go project should look like. It is still a work in progress.

It is supposed to follow clean architecture for separation of concerns and act as a robust boiler plate for someone starting Go.

After the cloning the project, run go mod tidy and configure local variables in config/local.yml
You need postgresql set up to run the DB. Replace your username, password and database name in the file and execute run main.go from the root directory

Health check is at /health
