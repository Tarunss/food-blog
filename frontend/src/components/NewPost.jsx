import React, { useState } from 'react'
import './NewPost.css'
export default function newPost(props) {
    const [title, setTitle] = useState()
    const [post, setPost] = useState()
    const [date, setDate] = useState()

    function handleSubmit(e) {
        e.preventDefault;
        //code to update post goes here (POST to datatbase)
        props.toggle()
    }
    return (
        <div className="popup">
            <div className="popup-inner">
                <h2 className="titleHeader">New Post</h2>
                <form onSubmit={handleSubmit}>
                    <label>
                        Title:
                        <input type="text" value={title} onChange={e => setTitle(e.target.value)} />
                    </label>
                    <label>
                        Post:
                        <input type="textarea" value={post} onChange={e => setPost(e.target.value)} rows="4" cols="50" />
                    </label>
                    <label>
                        Date:
                        <input type="text" value={date} onChange={e => setDate(e.target.value)} />
                    </label>
                    <button type="submit" onClick={props.toggle}>Create Post</button>
                </form>
            </div>
        </div>
    );
};