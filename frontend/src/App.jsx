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
import NewPost from './components/NewPost';
import SinglePost from './components/SinglePost';
import Display from './components/Display';



function App() {
  // custom useAuthorizatino
  const { token, setToken } = useToken();

  //our useEffect for fetchPosts
  return (
    <Routes>
      <Route
        path="/"
        element={<Display></Display>}>
      </Route>
      <Route
        path=":id"
        element={
          <SinglePost></SinglePost>}>
      </Route>
      <Route
        path="/adminlogin"
        element={token ? <Navigate to="/dashboard"></Navigate> : <Login setToken={setToken}></Login>}>
      </Route>
      <Route path="/dashboard" element={!token ? <Navigate to="/"></Navigate> : <Dashboard></Dashboard>}>
      </Route>
      <Route path="/newpost" element={!token ? <Navigate to="/"></Navigate> : <NewPost></NewPost>}></Route>
      <Route path="/preferences" element={<Preferences></Preferences>}>
      </Route>
    </Routes>
  )
}

export default App;