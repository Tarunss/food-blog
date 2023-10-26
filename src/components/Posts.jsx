import React from 'react';
import './Posts.css';

const Posts = ({ post }) => {
    return (
        <article className="post">
            <h2>
                {post.body}
            </h2>
        </article>
    )
}

export default Posts