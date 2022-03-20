<div align="center">
   <a href="https://go.dev/"><img src = https://img.shields.io/badge/GO-v1.17.6-blue></a>
   <a href= "https://aws.amazon.com/id/rds/?p=ft&c=db&z=3"><img src = https://img.shields.io/badge/AWS_RDS-MySQL-orange></a>
   <a href="https://echo.labstack.com/"><img src = https://img.shields.io/badge/Echo-v4.7.0-blue></a>
   <a href="https://aws.amazon.com/id/s3/?did=ft_card&trk=ft_card"><img src = https://img.shields.io/badge/AWS-S3%20Bucket-green></a>
   <a href="https://www.edamam.com/"><img src = https://img.shields.io/badge/ExtAPI-%20EDAMAM-green></a>
   <a href="https://hub.docker.com/"><img src = https://img.shields.io/badge/Deploy-%20Docker-blue></a>
   <a href="https://www.okteto.com/"><img src = https://img.shields.io/badge/Deploy-%20Okteto-purple></a>
</div>
<br />

# Healthy Fit

<!-- PROJECT LOGO -->
<div align="center">
  <a href="https://raw.githubusercontent.com/FINAL-PROJECT-ALTA/FE/development/image/logo-white.png">
    <img src="https://raw.githubusercontent.com/FINAL-PROJECT-ALTA/FE/development/image/logo-white.png" alt="Logo" width="250" height="180">
  </a>

  <h3 align="center">Healthy Fit</h3>
  <p align="center">
   An application that aims its users to monitor eating behavior on a daily basis
    <br />
    <div id = "other-software-design"></div>
    ·
     <a href="https://app.swaggerhub.com/apis/aaryadewangga/Final_Project/1.0#/">Open API</a>
    ·
    <a href="https://github.com/FINAL-PROJECT-ALTA/FE">Front-End</a>
    ·
    <a href="https://github.com/FINAL-PROJECT-ALTA/QE">Quality-Engineer</a>
  </p>
</div>
<br />

<!-- TABLE OF CONTENTS -->
## Table of Contents
1. [About the Project](#about-the-project)
2. [Feature](#feture)
3. [Tech Stack](#tech-stack)
4. [Hight Level Architecture](#high-level-architecture)
5. [ERD](#erd)
6. [Unit Testing](#unit-testing)
6. [How to Contribute](#contribute)
7. [Authors](#authors)

<!-- ABOUT THE PROJECT -->
## About The Project
An application to get recommendations for health-supporting foods to lose weight or gain weight


<p align="right">(<a href="#top">back to top</a>)</p>

## Feature
-  Login and Logout
-  Create account registration

As Users
-  See the content of the food or drink
-  Set a goal to get a recommendation menu calculation
-  Choose the recommend menu
-  View the history of the selected menu

As Admin
-  CRUD types of food, drinks, fruit, snacks
-  CRUD types of recommend menu


<p align="right">(<a href="#top">back to top</a>)</p>

## Tech Stack
### Framework
- [Echo (Go Web Framework)](https://echo.labstack.com/)

### Build With
- [Golang (Language)](https://go.dev/) 
- [Gorm (ORM Library)](https://aws.amazon.com/id/?nc2=h_lg)
- [Testify (Unit Test)](https://github.com/stretchr/testify)

### Third Party
- [AWS S3 Bucket](https://aws.amazon.com/id/?nc2=h_lg)
- [Edamam Food Databases](https://www.edamam.com/)

### Database
- [AWS RDS](https://aws.amazon.com/id/?nc2=h_lg)

### Deployment
- [Docker (Container - image)](https://hub.docker.com/)
- [Okteto (Kubernetes Platform)](https://www.okteto.com/)

### Collaboration 
- [Trello](https://trello.com/) - Manage Project
- [Github](https://github.com/) - Versioning Project

<p align="right">(<a href="#top">back to top</a>)</p>

## Structure
``` bash
Healthy Fit
  ├── configs                
  │     └──config.go                  # Configs files
  ├── delivery                        # Endpoints handlers or controllers
  │     ├──controllers
  │     │   └── users
  │     │     ├── formatter.go        # Default response format for spesific controllers
  │     │     ├── users_test.go       # Unit tests for spesific controllers
  │     │     └── users.go            # Spesific controller
  │     ├──middlewares
  │     │   └── jwtMiddleware.go      # Middlewares Function
  │     └──routes  
  │         └── routes.go             # Endpoints list
  ├── deployment               
  │     └── app.yaml                  # Deploymen installer
  ├── entities                
  │     └── users.go                  # Database model
  ├── repository              
  │     ├── interface.go              # Repository Interface for controllers
  │     ├── users_test.go             # Unit test for spesific repository
  │     └── users.go                  # Spesific Repository
  ├── utils                 
  │     ├── mysql
  │     │    └── driver.go            # Database driver
  │     └── thrid_party
  │          └── thirdParty.go        # Third Party driver
  ├── .gitignore  
  ├── dockerfile                      # Which files to ignore when committing
  ├── go.mod                  
  ├── go.sum                  
  ├── main.go                         # Main Program
  └── README.md    
```
<p align="right">(<a href="#top">back to top</a>)</p>

<!-- HLA -->
## High Level Architecture

<img src="./doc/HLA.png" alt="display-preview">

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ERD -->
## ERD

<img src="./doc/ERD.png" alt="display-preview">

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- Testing -->
## Unit Testing

<img src="" alt="display-preview">

<p align="right">(<a href="#top">back to top</a>)</p>

## Contribute

- Fork this repository

    ```sh
    $ git clone https://github.com/YOUR_USERNAME/FINAL-PROJECT-ALTA/BE.git
    > Cloning into `healthy-fit`...
    > remote: Counting objects: 10, done.
    > remote: Compressing objects: 100% (8/8), done.
    > remove: Total 10 (delta 1), reused 10 (delta 1)
    > Unpacking objects: 100% (10/10), done.
    ```
<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
## Authors
* Arya Nur Dewangga Putra - [Github](https://github.com/aaryadewangga) · [LinkedIn](https://www.linkedin.com/in/aryadewangga/)
* Ade Mawan - [Github](https://github.com/ademawan) · [LinkedIn](https://www.linkedin.com/in/ade-mawan-527657177/)

<p align="right">(<a href="#top">back to top</a>)</p>
