To build docker image:
- docker build -t "image_name"

To build docker container
- docker run --name "container_name" -p 4000:4000 "image_name"

To run docker container with volumes attached.
- docker run --name node_c2_nodemon -p 4000:4000 --rm -v /Users/sumit/personal-projects/devops-rnd/docker/api:/app -v /app/node_modules  node_app:nodemon