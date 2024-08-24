import axios from 'axios';
//Get auth0 jwt token from auth0 dashboard and paste into this variable
const bearerToken = ''; 
axios.post("http://0.0.0.0:8080/api/generate-bearer-token", null, {
    headers: {
        'authorization': `Bearer ${bearerToken}`,
    }
}).then((response) => {
    console.log("here is the skyflow bearer token generated: " + response.data.token);
});
