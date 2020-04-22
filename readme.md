# Move db client creation in the repository module instead of injecting it.

# It is an extra step that is not needed



# Configuration with env variables
## Database config
- DB_NAME
- DB_PASSWORD
- DB_PORT
- DB_HOST
- DB_USER



# Start locust
Run the locust.sh bash script that will run locust at localhost:9089
Specify 
- Users
- Hatch rate
- Host (dont forget the http)



# Experiment
50 Users
0.05 Hatch Rate
15mins
http://35.241.184.139:30353