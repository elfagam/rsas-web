<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../services/api';
import { ASSETS } from '../config/assets';

const router = useRouter();
const username = ref('');
const password = ref('');
const showPassword = ref(false);
const isLoading = ref(false);
const errorMessage = ref('');

const handleLogin = async () => {
  isLoading.value = true;
  errorMessage.value = '';

  try {
    const response = await api.post('/auth/login', {
      username: username.value,
      password: password.value,
    });

    const { token, user } = response.data;

    // Simpan token dan data user ke localStorage
    localStorage.setItem('rsas_token', token);
    localStorage.setItem('rsas_user', JSON.stringify(user));

    // Redirect ke Dashboard (akan dibuat nanti)
    router.push('/admin/dashboard');
  } catch (error) {
    errorMessage.value = error.response?.data?.error || 'Login gagal. Periksa kembali akun Anda.';
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="login-view" :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.6)), url(${ASSETS.hero.profil})` }">
    <div class="container">
      <div class="login-card" data-aos="zoom-in">
        <div class="login-card__header">
          <div class="login-card__logo">
            <i class="fa-solid fa-hospital-user"></i>
          </div>
          <h1>Portal <span>Administrasi</span></h1>
          <p>Selamat datang kembali di sistem manajemen RSUD Prof. Dr. H. Aloei Saboe</p>
        </div>

        <form @submit.prevent="handleLogin" class="login-form">
          <div v-if="errorMessage" class="login-error">
            <i class="fa-solid fa-circle-exclamation"></i>
            {{ errorMessage }}
          </div>

          <div class="form-group">
            <label>Username</label>
            <div class="input-wrapper">
              <i class="fa-solid fa-user"></i>
              <input 
                v-model="username" 
                type="text" 
                placeholder="Masukkan username Anda" 
                required
              />
            </div>
          </div>

          <div class="form-group">
            <label>Password</label>
            <div class="input-wrapper">
              <i class="fa-solid fa-lock"></i>
              <input 
                v-model="password" 
                :type="showPassword ? 'text' : 'password'" 
                placeholder="Masukkan password Anda" 
                required
              />
              <button 
                type="button" 
                class="password-toggle" 
                @click="showPassword = !showPassword"
              >
                <i :class="['fa-solid', showPassword ? 'fa-eye-slash' : 'fa-eye']"></i>
              </button>
            </div>
          </div>

          <button type="submit" class="btn-login" :disabled="isLoading">
            <span v-if="!isLoading">Masuk ke Sistem</span>
            <span v-else><i class="fa-solid fa-circle-notch fa-spin"></i> Memverifikasi...</span>
          </button>
        </form>

        <div class="login-card__footer">
          <p>&copy; 2024 RSUD Prof. Dr. H. Aloei Saboe. All Rights Reserved.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@use "../assets/sass/base/variables" as *;

.login-view {
  min-height: 100vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 4rem 2rem;
}

.login-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  max-width: 50rem;
  width: 100%;
  margin: 0 auto;
  border-radius: 40px;
  padding: 6rem 5rem;
  box-shadow: 0 4rem 10rem rgba(0,0,0,0.3);
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.2);

  @include respond("tablet") {
    padding: 4rem 3rem;
  }
}

.login-card__logo {
  width: 9rem;
  height: 9rem;
  background: #55c57a;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 4rem;
  margin: 0 auto 3rem;
  box-shadow: 0 1.5rem 4rem rgba(85, 197, 122, 0.4);
}

.login-card__header h1 {
  font-family: 'Outfit', sans-serif;
  font-size: 3.2rem;
  font-weight: 800;
  color: #333;
  margin-bottom: 1.5rem;
  letter-spacing: -1px;

  span {
    color: #55c57a;
  }
}

.login-card__header p {
  font-size: 1.5rem;
  color: #666;
  line-height: 1.6;
  margin-bottom: 4rem;
}

.login-form {
  text-align: left;
}

.login-error {
  background: #fff5f5;
  color: #e74c3c;
  padding: 1.5rem 2rem;
  border-radius: 15px;
  font-size: 1.4rem;
  font-weight: 700;
  margin-bottom: 2.5rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  border-left: 4px solid #e74c3c;
}

.form-group {
  margin-bottom: 2.5rem;

  label {
    display: block;
    font-size: 1.4rem;
    font-weight: 800;
    color: #444;
    margin-bottom: 1rem;
    text-transform: uppercase;
    letter-spacing: 1px;
  }
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;

  i {
    position: absolute;
    left: 2rem;
    font-size: 1.8rem;
    color: #aaa;
    transition: all 0.3s;
  }

  input {
    width: 100%;
    padding: 1.8rem 2rem 1.8rem 5.5rem;
    border: 2px solid #eee;
    border-radius: 15px;
    font-size: 1.6rem;
    transition: all 0.3s;
    background: #fdfdfd;

    &:focus {
      outline: none;
      border-color: #55c57a;
      background: #fff;
      box-shadow: 0 1rem 3rem rgba(85, 197, 122, 0.1);

      & + .password-toggle i {
        color: #55c57a;
      }
    }
  }
}

.password-toggle {
  position: absolute;
  right: 2rem;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 10;

  i {
    font-size: 1.8rem;
    color: #aaa;
    transition: all 0.3s;
  }

  &:hover i {
    color: #55c57a;
  }
}

.btn-login {
  width: 100%;
  padding: 2rem;
  background: #55c57a;
  color: #fff;
  border: none;
  border-radius: 15px;
  font-size: 1.7rem;
  font-weight: 800;
  cursor: pointer;
  transition: all 0.3s;
  margin-top: 2rem;
  box-shadow: 0 1.5rem 4rem rgba(85, 197, 122, 0.3);

  &:hover {
    background: #4ab46d;
    transform: translateY(-3px);
    box-shadow: 0 2rem 5rem rgba(85, 197, 122, 0.4);
  }

  &:disabled {
    background: #ccc;
    cursor: not-allowed;
    transform: none;
    box-shadow: none;
  }
}

.login-card__footer {
  margin-top: 5rem;
  padding-top: 3rem;
  border-top: 1px solid #eee;

  p {
    font-size: 1.2rem;
    color: #aaa;
    text-transform: uppercase;
    letter-spacing: 1px;
  }
}
</style>
