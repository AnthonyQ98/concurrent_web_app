# Concurrent Go Web Application Docker Image

This Docker image hosts a concurrent Go web application. It's available on the public Docker Hub registry.

## How to use it?

### Run the image
```docker run -d -p 8080:8080 --restart always --name <container_name> -t devanthony/concurrent_golang_web_app:1```

### Access the web app
```Go to the URL: http://localhost:8080/home```
