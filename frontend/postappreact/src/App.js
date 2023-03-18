import React, { useState, useEffect } from 'react';
import AddPost from './components/AddPost.js';
import Posts from './components/Posts.js';
import axios from "axios";
import "./App.css";

const App = () => {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await axios.get("http://localhost:8000/api/posts");
        setPosts(res.data.data);
      } catch (error) {
        console.log(error);
      }
    };
    fetchData();
  }, []);

  const addPost = async (newPost) => {
    try {
      const res = await axios.post("http://localhost:8000/api/posts", newPost);
      setPosts([res.data.data, ...posts]);
    } catch (error) {
      console.log(error);
    }
  };
  
  const deletePost = async (id) => {
    try {
      await axios.delete(`http://localhost:8000/api/posts/${id}`);
      setPosts(posts.filter(post => post.ID !== id));
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <div className='container'>
      <AddPost addPost={addPost} />
      <hr />
      <Posts posts={posts} deletePost={deletePost} />
    </div>
  );
};

export default App;
