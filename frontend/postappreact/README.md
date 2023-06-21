React Blog Application
======================

This is a simple blog application built with React. It allows users to add and delete posts.

Getting Started
---------------

### Prerequisites

*   Docker

### Installation

1.  Clone the repository:

    git clone https://github.com/your-username/react-blog-app.git
    

2.  Navigate to the project directory:

    cd react-blog-app
    

Usage
-----

1.  Build the Docker image:

    docker build -t react-blog-app .
    

2.  Run the Docker container:

    docker run -p 3000:3000 react-blog-app
    

3.  Open your browser and visit `http://localhost:3000` to access the application.

Features
--------

*   Add new posts: Enter the post details in the form and click "Add" to create a new post.
*   Delete posts: Click the "Delete" button next to a post to remove it from the list.

Technologies Used
-----------------

*   React
*   Axios

API
---

This application relies on a RESTful API to fetch and manipulate blog posts. Make sure you have the API server running on `http://localhost:8000/api/posts` or update the URL in the code accordingly.

License
-------

This project is licensed under the [MIT License](LICENSE).
