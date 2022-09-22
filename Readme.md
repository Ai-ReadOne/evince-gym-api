# Author : Ai ReadOne
# Name: Evince-gym-api
# Description: This mini project is meant for self development in Golang and REST API.

1. This API is built for a sudo gym management application

#Endpoints:

#All Members 
- URL: localhost:8000/all-members //for all registered customers
- Method : GET
<hr>

#Get all Instructors 
- URL: localhost:8000/all-instructors //for all gym instructors 
- Method : GET
<hr>

#Get an Instructor record
- URL: localhost:8000/get-instructor/:instructorID //for a single gym instructor
- Method : GET
- Path parameter : instructorID
<hr>

#Get a members' record
- URL: localhost:8000/get-member/:gymID // for a single gym member
- Method : GET
- Path parameter : gymID
<hr>

#Create a member record
- URL: localhost:8000/create-member registers new customer to the gym
- Method : POST
<hr>

#Create an Instructor record
- URL: localhost:8000/create instructor registers new instructor to the gym
- Method : POST
