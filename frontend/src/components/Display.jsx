import Header from "./Header"
import Posts from "./Posts"
import { useState, useEffect } from "react"
import index from '.././index'

export default function Display() {
    // Our useStates
    const [posts, setPosts] = useState([]);
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
        <div className="displayDiv">
            <Header></Header>
            <Posts posts={posts}></Posts>
        </div>
    )
}