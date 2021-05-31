# simple-crud-go-mongo
A simple crud using Golang, MongoDB and Docker

## Dev and Prod environment implemented
If you want to start a server in dev environment, just to run:
```bash 
  docker-compose -f docker-compose.dev.yml up -d --no-deps --build server
```
It will be available on
```bash
  localhost:3000/simple-crud
```

This Project is already setup to run in a Production Environment with Nginx and a SSL Certbot. To deploy in your host you will need a current ssl cert on host.
After that, just to run:
```bash 
  docker-compose up -d --no-deps --build server
```

## Auto Swagger doc implemented
![image](https://user-images.githubusercontent.com/53406077/120241927-246dcf00-c23a-11eb-8cf9-90f642e5f17b.png)

If you want to test this simple project, it is available on my personal page: [Simple Crud Swagger](https://ogustavobelo.com/simple-crud/swagger/index.html)
