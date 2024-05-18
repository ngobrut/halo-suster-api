# Halo-suster-api
This repository is a backend application for nurses or other users to record patient medical records in hospitals.

### Features 
- Authentication
- Nurse Management
- Medical Record
- Image Upload

### API Documentations
[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/26331093-fc9f055f-fe58-4858-9a05-73db6c54f888?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D26331093-fc9f055f-fe58-4858-9a05-73db6c54f888%26entityType%3Dcollection%26workspaceId%3Dbe40ac12-1b46-4bfd-840a-d967465b3fbd)

## Getting Started

These instructions will give you a copy of the project up and running on
your local machine for development and testing purposes. See deployment
for notes on deploying the project on a live system.

### Prerequisites

List of prerequisites for development
- Golang
- Postgresql
- AWS S3

### Create your .env

```
DB_HOST=
DB_PORT=
DB_NAME=
DB_USERNAME=
DB_PASSWORD=
DB_PARAMS="sslmode=disable"
JWT_SECRET=
BCRYPT_SALT=8 

AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_S3_BUCKET_NAME=
AWS_REGION=
```

### Migrate database
- Execute all migration
```
$ run this
make migrate-up
```
- Rollback migration (one migration)
```
$ run this
make migrate-down
```
- Rollback migration (one migration)
```
$ run this
make migrate-drop
```

### Install all dependencies
```
$ run this
go download
```
### Run your application
```
$ run this
go run main.go
```

## Deployment
Using Docker to build docker image and push to registry
```
$ run this
docker build . -t <image-name>::latest
```
