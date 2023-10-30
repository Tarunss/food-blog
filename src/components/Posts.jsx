import React from 'react';
import './Posts.css';

const Posts = ({ posts }) => {
    return (
        <ul>
            {posts.map((post) => (
                <li className="posts" key={post._id}>
                    <ul className="post">
                        <li className="postTitle">{post.title}</li>
                        <li className="postBody">{post.body}</li>
                        <li className="postDate">{post.date}</li>
                    </ul>
                </li>
            ))}
        </ul>
    )
}

export default Posts