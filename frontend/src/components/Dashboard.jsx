import React, { useEffect, useState } from 'react'
import './Dashboard.css'
import index from '.././index'
import Posts from './Posts.jsx'
import NewPost from './NewPost.jsx'
import { Button, Modal } from 'react-bootstrap';
export default function Dashboard() {
    const [seen, setSeen] = useState(false);
    const [posts, setPosts] = useState([]);
    function handleDelete(id) {
        const deletePost = async () => {
            try {
                //use our index and axios to make a get requests (/posts is our endpoint)
                console.log(id);
                const response = await index.delete('/api/post/' + id);
                if (response.status === 200) {
                    // Filter out the deleted post from the current list of posts
                    const updatedPosts = posts.filter(post => post._id !== id);
                    // Update the state with the updated list of posts
                    setPosts(updatedPosts);
                }
                //setPosts(response.data);
                //print(posts.data)
            } catch (err) {
                //not in 200 response range
                if (err.response) {
                    console.log(err.response.data);
                    console.log(err.response.status);
                    console.log(err.response.headers);
                } else {
                    console.log(`Error: ${err.message}`);
                }
            }
        }
        //calling our get
        deletePost();
    }
    function togglePop() {
        setSeen(!seen);
    }
    function handleClose() {
        setSeen(false);
    }
    function handleShow() {
        setSeen(true);
    }
    useEffect(() => {
        const fetchPosts = async () => {
            try {
                //use our index and axios to make a get requests (/posts is our endpoint)
                const response = await index.get('/api/posts');
                setPosts(response.data);
                //print(posts.data)
            } catch (err) {
                //not in 200 response range
                if (err.response) {
                    console.log(err.response.data);
                    console.log(err.response.status);
                    console.log(err.response.headers);
                } else {
                    console.log(`Error: ${err.message}`);
                }
            }
        }
        //calling our get
        fetchPosts();
    }, [])
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
                    <NewPost handleClose={handleClose}></NewPost>
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
                                    <button type="submit" onClick={() => handleDelete(post._id)}>Delete</button>
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
