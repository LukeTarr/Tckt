import axios from 'axios';

const API = axios.create({baseURL : "http://localhost:4000"});

export const signin = (formData) => API.post("/auth/login", formData);
export const signup = (formData) => API.post("/auth/register", formData);