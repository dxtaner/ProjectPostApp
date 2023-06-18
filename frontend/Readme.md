Docker Compose PostApp
======================

This is a basic configuration to start the PostApp application using Docker Compose. It consists of two services: a PostgreSQL database server and a React application.

Usage
-----

1.  Install Docker Compose on your computer: [Docker Compose Installation](https://docs.docker.com/compose/install/)
2.  Navigate to the project directory:
`cd postapp-docker-compose`4.  Start the application with Docker Compose:
`docker-compose up -d`6.  Access the React application by opening [http://localhost:3000](http://localhost:3000) in your browser.
7.  Use the following information for the database connection:

*   Host: localhost
*   Port: 5439
*   Database: postgres
*   Username: postgres
*   Password: depixen-pass

Services
--------

### db

*   Container name: postapp
*   A container that includes a PostgreSQL database server.
*   Configures POSTGRES\_PASSWORD, POSTGRES\_DB, and POSTGRES\_USER environment variables.
*   Port: 5439 (mapped to port 5432 on the local machine)
*   Uses a Docker volume to store the database data.

### react-app

*   Container name: react-postapp
*   A container that includes the React application.
*   Port: 3000
*   Depends on the db service, so it won't start until the db service is up.

Docker Data Volume
------------------

*   depixen-volume: A Docker volume used to store the data of the PostgreSQL database container.
