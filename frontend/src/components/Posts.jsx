import React from 'react';
import './Posts.css';

const Posts = ({ posts }) => (
    <div className="postListDiv">
        <ul className="postList">
            {posts.map((post) => (
                <li className="posts" key={post._id}>
                    <div className="postDiv">
                        <article className="post">
                            <section className="titleSection">
                                <p className="postDate">{post.date}</p>
                                <h1 className="postTitle">{post.title} </h1>
                                <p className="postSummary">{post.summary}</p>
                            </section>
                        </article>
                    </div>
                </li>
            ))}
        </ul>
    </div>
)


export default Posts