import React, { Component, useEffect, useState } from 'react';
import './App.css';
import Header from "./components/Header";
import Posts from "./components/Posts"
import index from './index';
function App() {
  // Our useStates
  const [posts, setPosts] = useState([])

  //our useEffects
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
          console.log(`Error: ${err.message}`)
        }
      }
    }
    //calling our get
    fetchPosts();
  }, [])

  return (
    <div className="appDiv">
      <Header></Header>

      <Posts posts={posts}></Posts>
    </div >
  )
}

export default App;