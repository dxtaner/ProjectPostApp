PostApp
======================


Docker Compose PostApp
======================

This is a basic configuration to start the PostApp application using Docker Compose. It consists of two services: a PostgreSQL database server and a React application.

Usage
-----

1.  Install Docker Compose on your computer: [Docker Compose Installation](https://docs.docker.com/compose/install/)
2.  Navigate to the project directory:

    cd postapp-docker-compose
    

3.  Start the application with Docker Compose:

    docker-compose up -d
    

4.  Access the React application by opening [http://localhost:3000](http://localhost:3000) in your browser.
5.  Use the following information for the database connection:

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

PostApp GoLang
=======

PostApp is a simple RESTful API application that allows users to manage posts. It provides endpoints to retrieve, create, update, and delete posts.

Prerequisites
-------------

Before running the application, make sure you have the following installed:

*   Go programming language
*   gorilla/mux package
*   jinzhu/gorm package

Installation
------------

1.  Clone the repository:

    git clone https://github.com/your-username/PostApp.git
    

2.  Change to the project directory:

    cd PostApp
    

3.  Install the dependencies:

    go mod download
    

Usage
-----

To run the application, execute the following command:

    go run main.go
    

The application will start listening on `localhost` at port `8000` by default.

API Endpoints
-------------

The following API endpoints are available:

*   `GET /api/posts` - Retrieve all posts
*   `GET /api/posts/{id}` - Retrieve a specific post by ID
*   `POST /api/posts` - Create a new post
*   `DELETE /api/posts/{id}` - Delete a post by ID
*   `PATCH /api/posts/{id}` - Update a post by ID

Example Usage
-------------

Assuming the application is running on `localhost:8000`, you can interact with the API endpoints using tools like cURL or Postman.

To retrieve all posts:

    GET http://localhost:8000/api/posts
    

To retrieve a specific post:

    GET http://localhost:8000/api/posts/{id}
    

To create a new post:

    POST http://localhost:8000/api/posts
    
    Request Body:
    {
      "title": "New Post",
      "description": "This is a new post",
      "imageuri": "https://example.com/image.jpg"
    }
    

To delete a post:

    DELETE http://localhost:8000/api/posts/{id}
    

To update a post:

    PATCH http://localhost:8000/api/posts/{id}
    
    Request Body:
    {
      "title": "Updated Post",
      "description": "This post has been updated",
      "imageuri": "https://example.com/new-image.jpg"
    }
    

Contributing
------------

Contributions to the project are welcome. If you find any issues or want to add new features, please open an issue or submit a pull request.

License
-------

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
