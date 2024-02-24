# food_out

**Food Out** is a food delivery application, written using Micro-service approach and primarily Go programming language.

## Content

<!--toc:start-->
- [food_out](#foodout)
  - [Content](#content)
  - [How to run](#how-to-run)
  - [Task](#task)
    - [System Requirements](#system-requirements)
  - [Architecture](#architecture)
    - [Overview](#overview)
    - [Accounts Service](#accounts-service)
    - [Auth Service](#auth-service)
    - [Client Service](#client-service)
    - [Restaurant Service](#restaurant-service)
  - [Courier Service](#courier-service)
    - [Menu Service](#menu-service)
    - [Order Service](#order-service)
    - [Media Manager](#media-manager)
    - [Infrastructure](#infrastructure)
<!--toc:end-->

## How to run

Coming soon...

## Task

The task is to implement whole food delivery application with micro-service architecture and Go programming language.

### System Requirements 

- [ ] Account management;
- [ ] Auth;
- [ ] Managing restaurant info and its menu.
- [ ] Managing working status for couriers.
- [ ] Serving shop.
- [ ] Creating orders.
- [ ] Functionality to accept or reject orders for restaurants and couriers. 

## Architecture

### Overview

(./docs/architecture.png)[System Architecture]

![architecture diagram](./docs/architecture.png "Architecture Diagram")

### Accounts Service

Account Service is responsible for managing user accounts.

It provides CRUD operations on the account resource.

### Auth Service

Auth service is responsible for user sessions, it uses Redis for storing these sessions.

### Client Service

Client Service is an entrance to the system for clients. It handles clients requests and generates events for the rest of the system.

It stores client-related information, f.e. addresses, in PostgreSQL.

### Restaurant Service

Restaurant Service handles restaurant's requests and generates events for the system.

It stores the restaurant specific information in PostgreSQL.

### Courier Service

Courier service handles the courier related requests.

It stores the courier related information in PostgreSQL.

### Menu Service

Menu service is responsible for managing restaurant's menus. It uses PostgreSQL for storing this information.

This service handles three domain fields, which are:

- Shop Viewing(makes a selection of categories and dishes for particular restaurant and (optional) category).
- Menu Management Service(handles CRUD operations for restaurant's dishes and categories).
- Order Content Validation(validates a set of dishes, that was required by user and generates an event of complete Order).

Shop Viewer is exposed, but the Menu Management Service and Order Content Validation Service are not, they are private for this system.

### Order Service

Order service is responsible for managing orders in the system.

Order service contains two domain fields:

- Order Viewer Service, which provides an interface to query orders by the participants(clients, couriers, restaurants).
- Order Management Service, which is responsible for perfoming Create,Update,Remove operations on the dishes and categories.

The OrderViewer is exposed, but the Order Management Service is private for the system and does not have public API.

### Media Manager 

Image Uploader is a utillity-service that handles media uploading/receiving,
it is primarily used as an image storage for dishes, categories, restaurants, user profiles.

### Infrastructure

Most of the services are using PostgreSQL for storing information.

Auth service is using Redis for storing sessions, due to it's great functionality of automatically deleting objects.

For storing media(images) MinIO is used.

Kafka is used as a message broker.

