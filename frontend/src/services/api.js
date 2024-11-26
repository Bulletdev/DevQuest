import axios from 'axios';

const API_URL = 'http://localhost:8080/api';

export const getLessons = () => {
  return axios.get(`${API_URL}/lessons`);
};

export const getLessonById = (id) => {
  return axios.get(`${API_URL}/lessons/${id}`);
};
