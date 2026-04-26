import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8000/api', // Alamat backend Go kita
  headers: {
    'Content-Type': 'application/json',
  },
});

// Interceptor untuk menyisipkan token JWT di setiap permintaan
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('rsas_token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Interceptor untuk menangani error (misal: token kedaluwarsa)
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && error.response.status === 401) {
      localStorage.removeItem('rsas_token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default api;
