This project uses a jwt acces token taken by auth0 in order to generate a skyflow bearer token. It can be helpful for potential customers new to the skyflow API, and needing to see how it works while using their previous authentication methods. 

TECH USED: Javascript, Go

HOW TO RUN:

1. Auth0 Setup: 
- Go to manage.auth0.com and create an account.
- In the auth0 dashboard, go to Authentication -> API's -> Auth0 managment API -> API explorer.
- Go to the Management API Explorer
- Click the Set API Token button at the top left.
- Set the API token by pasting the API Token that you copied in the first step.
- Click the Set Token button.
2. Skyflow setup: 
- Go to skyflow studio and download your credentials file
- paste that path in main.go
3. Run the server:
- cd backend
- go run main.go
4. Run the frontend:
- cd frontend
- npm init
- npm install
- node client.js

