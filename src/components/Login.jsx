import { React, useState } from 'react'
import './Login.css'
import PropTypes from 'prop-types'
//POST request function
async function loginUser(credentials) {
    return fetch('http://localhost:8081/admin', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
    })
        .then(data => data.json).then(console.log(data))
}
const handleSubmit = async e => {
    e.preventDefault();
    console.log(e)
    const token = await loginUser({
        username,
        password
    });
    setToken(token);
}

export default function Login({ setToken }) {
    //useStates
    const [username, setUserName] = useState([]);
    const [password, setPassWord] = useState([]);
    return (
        <div className="login-wrapper:">
            <h1>Please log in</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    <p>Username:</p>
                    <input type="text" onChange={e => setUserName(e.target.value)} />

                </label>
                <label>
                    <p>Password:</p>
                    <input type="password" onChange={e => setPassWord(e.target.value)} />

                </label>
                <div>
                    <button type="submit">Submit</button>
                </div>
            </form>
        </div>
    )
};

Login.propTypes = {
    setToken: PropTypes.func.isRequired
};