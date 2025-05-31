import axios from 'axios'
//define some functions to get our data
export default axios.create({
    baseURL: 'http://localhost:8080'
})