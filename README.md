# Multi Container App

Here I have developed a multi container app. project <a href="https://roadmap.sh/projects/multi-container-service">link</a> <br>

    1 - todos app (Golang server)
    2 - Mongo DB (database)

### Architecture

    1 - both apps are running inside container.
    2 - there is a private network where apps can communicate
    3 - there is a exposed network where Golang server can be accessed.

**Note:** sure there are many mode ways to improve this. but this is a practice to deploy an operational application.

### Tools

    1 - Docker: to handle containerization
    2 - Docker Compose: to handle multiple services of my application (+ volumes and networks)
    3 - Ansible: to handle configuration and deployment safely
    4 - Github Action: to implement CI/CD pipeline
