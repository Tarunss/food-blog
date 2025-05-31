import { React, useState } from 'react'
import './Login.css'
import PropTypes from 'prop-types'
import index from '.././index'

import bcrypt from 'bcryptjs'
//POST request function
async function loginUser(credentials) {
    return fetch('http://localhost:8080/api/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
    })
        .then(data => {
            console.log(data);
            return data.json()
        })
}


export default function Login({ setToken }) {
    //useStates
    const [username, setUserName] = useState([]);
    const [password, setPassWord] = useState([]);
    async function handleSubmit(e) {
        e.preventDefault();
        //console.log(e)
        const token = await loginUser({
            username,
            password
        });
        //console.log(token);
        setToken(token);
    }
    return (
        <div className="login-wrapper">
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
                <div className="submitButton">
                    <button type="submit">Submit</button>
                </div>
            </form>
        </div>
    )
};

Login.propTypes = {
    setToken: PropTypes.func.isRequired
};