PostApp
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
`git clone https://github.com/your-username/PostApp.git`3.  Change to the project directory:
`cd PostApp`5.  Install the dependencies:
`go mod download`

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
