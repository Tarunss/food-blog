import React, { Component, useEffect, useState } from 'react';
import { Navigate, BrowserRouter, Routes, Route, useNavigate } from 'react-router-dom';
import './App.css';
import Header from "./components/Header";
import Posts from "./components/Posts"
import Dashboard from './components/Dashboard';
import Preferences from './components/Preferences';
import Login from './components/Login';
import index from './index';
import useToken from './components/useToken';



function App() {
  // custom useAuthorizatino
  const { token, setToken } = useToken();
  // Our useStates
  const [posts, setPosts] = useState([]);

  //our useEffect for fetchPosts
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
    <Routes>
      <Route
        path="/"
        element={
          <div className="appDiv">
            <Header></Header>
            <Posts posts={posts}></Posts>
          </div>}>
      </Route>
      <Route
        path="/adminlogin"
        element={token ? <Navigate to="/dashboard"></Navigate> : <Login setToken={setToken}></Login>}>
      </Route>
      <Route path="/dashboard" element={!token ? <Navigate to="/"></Navigate> : <Dashboard posts={posts}></Dashboard>}>
      </Route>
      <Route path="/preferences" element={<Preferences></Preferences>}>
      </Route>
    </Routes>
  )
}

export default App;