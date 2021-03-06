# LifeTrack

## Project to practice full stack development

### Operations:
* CI / CD (Travis CI)
* Docker
* Kubernetes
* AWS 

### Stack:
* Golang
* React
* Postgres

### Security / Infrastructure (Maybe)
* Microsoft Certificate Authority (MSCA) 
* Veeam Backup and Replication

### Logging
* ELK Stack
*    - Elasticsearch 
*    - Logstash
*    - Kibana

## Dashboard to track the following: 

* Sleep - Using Fitbit API,
* Steps - Using Fitbit API,
* Heart Rate - Using Fitbit API,
* Weight - Using Fitbit API,
* Bodyfat Percentage - Using Fitbit API,
* Water - Using Fitbit API,
* Workouts - Manual,
* Reading - Manual,
* Language - Using Duolingo API?,
* Vitamins - Manual,
* Asset Management - Manual,
* Budgeting - Manual,
* Diet Record - Manual / Fitbit APi,
* Password Storage - Manual

#### Currently Setup: 

* Travis: Travis.yml file set up for Golang, currently building new pushes. 
* Docker: Currently sets up an instance with Golang and copies over the Main.go in binary form.
* Amazon EC2 Instance: Setup with Docker running - this is being treated as my production environment. 
* Docker Hub: Setup to allow for easy production updates.
