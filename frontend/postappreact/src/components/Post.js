import { getDownloadURL, ref } from 'firebase/storage';
import React, { useEffect, useState } from 'react';
import { storage } from '../firebase';
import "./Post.css";

const Post = ({ post, deletePost }) => {
  const { title, description, imageuri } = post;
  const id = post.ID;
  const [imageURL, setImageURL] = useState(null);

  useEffect(() => {
    const imageRef = ref(storage, `images/${imageuri}`);
    getDownloadURL(imageRef).then(url => {
      setImageURL(url)
    }).catch((err) => {
      console.log(err)
    });
  },);

  const handleDelete = () => {
    deletePost(id);
  };

  return (
    <div className="post-container">
      <div className="post-description">
        <h3>{title}</h3>
        <p>{description}</p>
        {imageURL && <img src={imageURL} alt={title} className="post-image" />}
        <br></br>
        <button className="deletePost" onClick={handleDelete}>Delete Post</button>
      </div>
    </div>
  );
};




export default Post;
