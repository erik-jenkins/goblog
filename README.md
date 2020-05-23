# goBlog

This repo is complete overkill for a simple blog, but I'm using it as a learning experience for building more resilient, decoupled apps.

Run the `scripts/webstart.sh` script to start all of the containers. The `scripts/webstop.sh` script stops them. Make sure to generate self-signed certs (`scripts/generate_self_signed_cert.sh`) beforehand! 

## Development environment

In dev, the react app runs in its own NodeJS container. This allows for hot-reloading during development. The client makes a request of the nginx server (also containerized) running on `localhost:80` (which is probably a bad choice, but that can be changed later). The nginx server will forward all requests with the endpoint '/api' to the golang backend. All other requests will be forwarded to the Node dev server.

## Production environment

This hasn't been implemented yet, but in prod, the react app will be compiled down to static files and served directly from the nginx server.
