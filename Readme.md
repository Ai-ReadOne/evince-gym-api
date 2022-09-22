# Author* : Ai ReadOne
# Name*: Evince-gym-api
# Description: This mini project is meant to test my proficiency in Golang and REST API.

1. This API is built for a sudo gym management application

its endpoints are
#All Members 
localhost:8000/all-members //for all registered customers
Method : GET
#Get all Instructors 
localhost:8000/all-instructors //for all gym instructors 
Method : GET
#Get an Instructor record:
localhost:8000/get-instructor/:instructorID //for a single gym instructor
Method : GET
Path parameter : instructorID
#Get a members' record:
localhost:8000/get-member/:gymID // for a single gym member
Method : GET
Path parameter : gymID
#Create a member record:
localhost:8000/create-member registers new customer to the gym
Method : POST
#Create an Instructor record:
calhost:8000/create instructor registers new instructor to the gym
Method : POST
