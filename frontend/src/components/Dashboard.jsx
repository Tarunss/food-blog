import React, { useState } from 'react'
import './Dashboard.css'
import Posts from './Posts.jsx'
import NewPost from './NewPost.jsx'
import { Button, Modal } from 'react-bootstrap';
export default function Dashboard({ posts }) {
    const [seen, setSeen] = useState(false);
    // function togglePop() {
    //     setSeen(!seen);
    // }
    function handleClose() {
        setSeen(false);
    }
    function handleShow() {
        setSeen(true);
    }
    return (
        <div className='dashboardDiv'>
            <Button variant="primary" onClick={handleShow} className="newButton">New Post</Button>
            <Modal
                show={seen}
                onHide={handleClose}
                centered
                backdrop='static'
                keyboard='false'
                className="modal">
                <Modal.Body>
                    <NewPost></NewPost>
                </Modal.Body>
            </Modal>
            <div className='postList'>
                {posts.map((post) => (
                    <li className="buttons" key={post._id}>
                        <div className="postDiv">
                            <article className="post">
                                <section className="titleSection">
                                    <p className="postDate">{post.date}</p>
                                    <h1 className="postTitle">{post.title} </h1>
                                    <p className="postBody">{post.body}</p>
                                    <button type="submit">Delete</button>
                                    <button type="submit">Update</button>
                                </section>
                            </article>
                        </div>
                    </li>
                ))} </div>

        </div>
        //if seen is true, then we want to change our styles for this dashboard component
    );
};
