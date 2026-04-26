<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const handleLogout = () => {
  localStorage.removeItem('rsas_token');
  localStorage.removeItem('rsas_user');
  router.push('/login');
};
</script>

<template>
  <aside class="admin-sidebar">
    <div class="sidebar__header">
      <i class="fa-solid fa-hospital-user"></i>
      <span>RSAS Admin</span>
    </div>
    
    <nav class="sidebar__nav">
      <!-- Dashboard -->
      <router-link to="/admin/dashboard" class="nav-item" active-class="active">
        <i class="fa-solid fa-gauge"></i> Dashboard
      </router-link>

      <!-- Berita -->
      <router-link to="/admin/berita" class="nav-item" :class="{ 'active': $route.path.startsWith('/admin/berita') }">
        <i class="fa-solid fa-newspaper"></i> Kelola Berita
      </router-link>

      <!-- Dokter (Placeholder) -->
      <router-link to="/admin/dokter" class="nav-item" active-class="active">
        <i class="fa-solid fa-user-doctor"></i> Data Dokter
      </router-link>

      <!-- Bed (Placeholder) -->
      <router-link to="/admin/bed" class="nav-item" active-class="active">
        <i class="fa-solid fa-bed-pulse"></i> Ketersediaan Bed
      </router-link>
    </nav>

    <div class="sidebar__footer">
      <button @click="handleLogout" class="btn-logout">
        <i class="fa-solid fa-right-from-bracket"></i> Keluar
      </button>
    </div>
  </aside>
</template>

<style lang="scss" scoped>
.admin-sidebar {
  width: 28rem;
  background: #fff;
  border-right: 1px solid #eee;
  display: flex;
  flex-direction: column;
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  z-index: 1000;

  .sidebar__header {
    padding: 4rem 3rem;
    display: flex;
    align-items: center;
    gap: 1.5rem;
    font-size: 2.2rem;
    font-weight: 800;
    color: #333;
    font-family: 'Outfit', sans-serif;

    i {
      color: #55c57a;
      font-size: 3rem;
    }
  }

  .sidebar__nav {
    flex: 1;
    padding: 0 1.5rem;

    .nav-item {
      display: flex;
      align-items: center;
      gap: 1.5rem;
      padding: 1.5rem 2rem;
      text-decoration: none;
      color: #666;
      font-size: 1.5rem;
      font-weight: 600;
      border-radius: 15px;
      transition: all 0.3s;
      margin-bottom: 0.5rem;

      i {
        font-size: 1.8rem;
        width: 2.5rem;
        text-align: center;
      }

      &:hover, &.active {
        background: #f0faf3;
        color: #55c57a;
      }

      &.active {
        font-weight: 800;
      }
    }
  }

  .sidebar__footer {
    padding: 3rem;
    border-top: 1px solid #eee;

    .btn-logout {
      width: 100%;
      padding: 1.5rem;
      background: #fff5f5;
      color: #e74c3c;
      border: 1px solid #fed7d7;
      border-radius: 12px;
      cursor: pointer;
      font-weight: 700;
      display: flex;
      align-items: center;
      justify-content: center;
      gap: 1rem;
      transition: all 0.3s;

      &:hover {
        background: #e74c3c;
        color: #fff;
      }
    }
  }
}
</style>
