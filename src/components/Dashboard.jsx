import React from 'react'
import './Dashboard.css'
import Posts from './Posts.jsx'
const Dashboard = ({ posts }) => (
    <div className='dashboardDiv'>
        <Posts posts={posts}></Posts>
        <div className='postList'>
            {posts.map((post) => (
                <li className="posts" key={post._id}>
                </li>
            ))} </div>
    </div>

);

export default Dashboard;
