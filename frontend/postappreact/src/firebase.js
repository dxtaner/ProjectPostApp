import { initializeApp } from "firebase/app";
import { getStorage } from "firebase/storage"

const firebaseConfig = {
  apiKey: "AIzaSyApLaD8rmZjp0BXgXzJk4r6d4hQUYAGvAc",
  authDomain: "imageuri-62ab5.firebaseapp.com",
  databaseURL: "https://imageuri-62ab5-default-rtdb.firebaseio.com",
  projectId: "imageuri-62ab5",
  storageBucket: "imageuri-62ab5.appspot.com",
  messagingSenderId: "98237005624",
  appId: "1:98237005624:web:0b665fb40f6bc89dd00bcd",
  measurementId: "G-XEDXFR9VWF"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const storage = getStorage(app);