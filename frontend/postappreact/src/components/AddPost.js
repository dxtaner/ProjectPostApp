import React, { useState } from "react";
import "./AddPost.css";
import { storage } from "../firebase";
import { ref, uploadBytes } from "firebase/storage";
import { v4 as uuidv4 } from 'uuid';

const AddPost = ({ addPost }) => {
  const [title, setTitle] = useState("New Title");
  const [description, setDescription] = useState("New Description");
  const [image, setImage] = useState(null);
  const MAX_TITLE_LENGTH = 50;

  const handleTitleChange = (e) => {
    setTitle(e.target.value);
  };

  const handleDescriptionChange = (e) => {
    setDescription(e.target.value);
  };

  const onSubmit = (e) => {
    e.preventDefault();
    if (!title || !description || !image) {
      alert("Boş alan bırakmayın!!");
      return;
    }

    if (title.length > MAX_TITLE_LENGTH) {
      alert(`Başlık en fazla ${MAX_TITLE_LENGTH} karakter olabilir`);
      return;
    }

    const newPost = {
      ID: parseInt(uuidv4().slice(0, 7), 16),
      title: title,
      description: description,
      imageuri: image.name,
      createddate: new Date()
    };

    addPost(newPost);

    const imageName = image.name;
    const imageRef = ref(storage, `images/${imageName}`);
    uploadBytes(imageRef, image).then(() => {
      alert(`${imageName} adlı resim fireabase storage aktarıldı..`)
    });

    setTitle("New Title");
    setDescription("New Description");
    setImage(null);
  };

  const handleImageChange = (e) => {
    const selectedImage = e.target.files[0];
    if (!selectedImage) {
      return;
    }
    if (selectedImage.type === "image/png" || selectedImage.type === "image/jpeg") {
      setImage(selectedImage);
    } else {
      alert("Sadece PNG veya JPEG formatında resimler yüklenebilir.");
    }
  };

  const handleImageRemove = () => {
    setImage(null);
  };
  
  return (
    <div className="add-post-container">
      <h2>Post Area</h2>
      <form onSubmit={onSubmit}>
        <div className="title">
          <input type="text"  value={title} onChange={handleTitleChange} style={{ fontWeight: 'bold' }} />
        </div>
        <div className="description">
          <textarea value={description} onChange={handleDescriptionChange}></textarea>
        </div>
        <div className="file">
          {!image && (
            <label className="imageLabel" htmlFor="image">
              + Image
            </label>
          )}
          {!image && <input type="file" id="image" name="image" accept="image/*" onChange={handleImageChange} />}
          {image && (
            <div>
              <button onClick={handleImageRemove}>
                <img src={URL.createObjectURL(image)} className="image" alt={title} />
              </button>
            </div>
          )}
        </div>
        <button type="submit">Add Post</button>
      </form>
    </div>

  );
};

export default AddPost;




