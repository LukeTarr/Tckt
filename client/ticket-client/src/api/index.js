import axios from 'axios';

const API = axios.create({baseURL : "http://tckt-demo.herokuapp.com"});

export const signin = (formData) => API.post("/auth/login", formData);
export const signup = (formData) => API.post("/auth/register", formData);