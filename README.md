# go_fileserver
Simple file server in golang

### How To Run:
Build:  `docker build -t danielorf/gofileserver:0.1-scratch . -f scratch.Dockerfile`  
Run: `docker run -it -e FSDirectory='/' -p 8080:8080 danielorf/gofileserver:0.1-scratch`  
Visit localhost:8080 in your browser  

  
  
Multistage and scratch build details here:  https://www.cloudreach.com/blog/containerize-this-golang-dockerfiles/
