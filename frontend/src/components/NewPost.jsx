import React, { useState } from 'react'
import index from '.././index.jsx'
import ReactQuill from 'react-quill'
import 'react-quill/dist/quill.snow.css';
import './NewPost.css'

const formats = ['header', 'bold', 'italic', 'underline', 'strike', 'blockquote', 'list', 'bullet', 'indent', 'link', 'image']
const modules = {
    toolbar: [
        [{ 'header': [1, 2, false] }],

    ]
}
export default function newPost(props) {
    const [title, setTitle] = useState("")
    const [post, setPost] = useState("")
    const [date, setDate] = useState("")
    const [content, setContent] = useState("")
    function handleSubmit(e) {
        e.preventDefault();
        //code to update post goes here (POST to database)
        //create formdata to post to our API
        const postPosts = async () => {
            try {
                //use our index and axios to make a get requests (/posts is our endpoint)
                const response = await index.post('/api/post', { title: title, body: post, date: date });
                console.log(response);
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
        //calling our post
        postPosts();
        props.handleClose();
        console.log(props);
    }
    return (
        <div className="popup">
            <div className="popup-inner">
                <h2 className="titleHeader" flex="5">New Post</h2>
                <form onSubmit={handleSubmit}>
                    <label className="newPostTitle">
                        Title:
                        <input type="text" value={title} onChange={e => setTitle(e.target.value)} />
                    </label>
                    <label>
                        Post:
                        <textarea className="newPostArea" rows="10" value={post} onChange={e => setPost(e.target.value)} />
                    </label>
                    <label className="newPostDate">
                        Date:
                        <input type="text" value={date} onChange={e => setDate(e.target.value)} />
                    </label>
                    <label className="postBody">
                        Body:
                        <ReactQuill value={content} modules={modules} format={formats}></ReactQuill>
                    </label>
                    <button type="submit">Create Post</button>
                </form>
            </div>
        </div>
    );
};