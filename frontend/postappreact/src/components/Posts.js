import React from 'react';
import Post from './Post.js';
import "./Posts.css";

const Posts = ({ posts, deletePost }) => {
  return (
    <div className="card">
      <div className="card-body">
        {posts.map((post, index) => (
          <Post key={index} post={post} deletePost={deletePost} />
        ))}
      </div>
    </div>
  );
};

export default Posts;