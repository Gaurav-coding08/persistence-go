# persistence-go
Repository consumes from kafka producer and persists data in database (PostgrSQL).

Steps to Setup & produce updates locally

# Setup env variables in .env file 
cp .env-example .env

# install the dependecies 
make setup 

# start PostgreSQL and Kakfka 
make db-up       # Start PostgreSQL

# make the migrations up
make migrate-up 

# test the application 
make test 

# run the application 
make run 

# now consumer is ready and will start consuming from the brocker and persists message in db. 

