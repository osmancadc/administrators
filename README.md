# Administrators

This project allows you to Create, Read and Update several administrators in the system

## Installation

Requires [golang](https://golang.org/dl/) v1.16+ to run.

Install the dependencies and devDependencies and start the server.

### Manual installation

```sh
cd admins
go mod init admins
go mod tidy
```
Run the server

```sh
go run .
```

### Docker installation

Install the dependencies and devDependencies and start the server.

```sh
cd admins
docker build --tag docker-admin .
```
Run the server

```sh
docker run -dp 3000:3000 docker-admin
```

## How to use

### See all the administrator's data

#### Request
```bash
    <host>:3000/administrators [GET]
```

#### Response

```javascript
    [
        {
            id: <Integer>,
            name: <String>,
            owner: <String>,
            criticality: <String>
        },
    ]
```

---
### Create Administrator

#### Request

```bash
    <host>:3000/administrator [POST]
```
```javascript
{
    name:<String>,
    owner:<String>,
    criticality:<String>
}
```

#### Response

```javascript
    200 - Administrator created successfully!
    400 - administrator's name is missing
    400 - the owner doesn't exists
    400 - the criticality doesn't exists
    500 - <Error>
```

---
### Create Administrator

#### Request

```bash
    <host>:3000/administrator/:id [PUT]
```
```javascript
{
    name:<String>,
    owner:<String>,
    criticality:<String>
}
```

#### Response

```javascript
    200 - Administrator updated successfully!
    400 - administrator's name is missing
    400 - the owner doesn't exists
    400 - the criticality doesn't exists
    500 - <Error>
```

---
---
### Update security

To improve the API's security is recommended to use authentication and JWT tokens to access the most critic endpoints like **_CREATE ADMIN_** and **_UPDATE ADMIN_**, depending of the displayed data of the admins the **_GET ADMIN_** endpoint also may be authorized via jwt 