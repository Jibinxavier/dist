# Project Architecture  v0.0.1
[Work in Progress]
## Outline

- Overview 
- API

## Overview 

A simple project to learn about distributed systems and patterns and to learn Golang while doing so.

Map Reduce pattern will be attempted, to start off with.

## API

The APIs will be REST-like :-) 

### MapR


- **/health** : health status of entity (leader or worker)
- **/status** : Overall job status 

#### Server 

- POST **/register** : to register with leader 
- POST **/job**   : to add job
- GET **/job/:id/**  : job status

#### Worker 
- **/assign**   : assign jobs 




## Core components 
- leader 
- worker 
- job
### Leader 

Distributes work 

### Worker 

Does some type of job 


### Job 
A job has following attributes:

- job id 
- action type: map or reduce
- type: sort or grep 
- status: complete or incomplete

 