<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const user = ref(JSON.parse(localStorage.getItem('rsas_user') || '{}'));

const handleLogout = () => {
  localStorage.removeItem('rsas_token');
  localStorage.removeItem('rsas_user');
  router.push('/login');
};

const stats = ref([
  { label: 'Total Berita', value: '12', icon: 'fa-newspaper', color: '#3498db' },
  { label: 'Dokter Aktif', value: '45', icon: 'fa-user-doctor', color: '#55c57a' },
  { label: 'Bed Tersedia', value: '28', icon: 'fa-bed-pulse', color: '#e67e22' },
  { label: 'Pengunjung Hari Ini', value: '1,240', icon: 'fa-chart-line', color: '#9b59b6' },
]);

onMounted(() => {
  // Verifikasi token (sederhana)
  if (!localStorage.getItem('rsas_token')) {
    router.push('/login');
  }
});
</script>

<template>
  <div class="admin-dashboard-page">
    <header class="top-bar">
      <div class="welcome-text">
        <h1>Dashboard <span>Utama</span></h1>
        <p>Selamat datang, <strong>{{ user.full_name }}</strong> ({{ user.role }})</p>
      </div>
      <div class="user-profile">
        <img src="https://ui-avatars.com/api/?name=Admin+RSAS&background=55c57a&color=fff" alt="User Avatar" />
      </div>
    </header>

    <section class="stats-grid">
      <div v-for="item in stats" :key="item.label" class="stat-card" data-aos="fade-up">
        <div class="stat-card__icon" :style="{ backgroundColor: item.color + '20', color: item.color }">
          <i :class="['fa-solid', item.icon]"></i>
        </div>
        <div class="stat-card__info">
          <h3>{{ item.value }}</h3>
          <p>{{ item.label }}</p>
        </div>
      </div>
    </section>

    <section class="quick-actions">
      <h2>Aksi <span>Cepat</span></h2>
      <div class="actions-grid">
        <router-link to="/admin/berita/tambah" class="action-btn">
          <i class="fa-solid fa-plus"></i>
          Tulis Berita Baru
        </router-link>
        <button class="action-btn">
          <i class="fa-solid fa-calendar-plus"></i>
          Update Jadwal Dokter
        </button>
        <button class="action-btn">
          <i class="fa-solid fa-wrench"></i>
          Pengaturan Sistem
        </button>
      </div>
    </section>
  </div>
</template>

<style lang="scss" scoped>
@use "../../assets/sass/base/variables" as *;

.admin-dashboard-page {
  padding: 5rem;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 5rem;

  .welcome-text {
    h1 {
      font-family: 'Outfit', sans-serif;
      font-size: 3.5rem;
      font-weight: 800;
      margin-bottom: 0.5rem;
      
      span {
        color: #55c57a;
      }
    }

    p {
      font-size: 1.6rem;
      color: #666;
    }
  }

  .user-profile img {
    width: 6rem;
    height: 6rem;
    border-radius: 50%;
    border: 3px solid #fff;
    box-shadow: 0 1rem 3rem rgba(0,0,0,0.1);
  }
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(24rem, 1fr));
  gap: 3rem;
  margin-bottom: 5rem;
}

.stat-card {
  background: #fff;
  padding: 3rem;
  border-radius: 30px;
  display: flex;
  align-items: center;
  gap: 2.5rem;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.03);
  transition: all 0.3s;

  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 2rem 5rem rgba(0,0,0,0.08);
  }

  &__icon {
    width: 7rem;
    height: 7rem;
    border-radius: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 2.5rem;
  }

  &__info {
    h3 {
      font-size: 2.8rem;
      font-weight: 800;
      color: #333;
    }

    p {
      font-size: 1.4rem;
      color: #888;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 1px;
    }
  }
}

.quick-actions {
  h2 {
    font-family: 'Outfit', sans-serif;
    font-size: 2.4rem;
    font-weight: 800;
    margin-bottom: 3rem;

    span {
      color: #55c57a;
    }
  }
}

.actions-grid {
  display: flex;
  gap: 2rem;
  flex-wrap: wrap;

  .action-btn {
    padding: 2rem 3rem;
    background: #fff;
    border: 1px solid #eee;
    border-radius: 20px;
    font-size: 1.6rem;
    font-weight: 700;
    color: #444;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 1.5rem;
    transition: all 0.3s;
    text-decoration: none;

    i {
      color: #55c57a;
      font-size: 2rem;
    }

    &:hover {
      background: #55c57a;
      color: #fff;
      border-color: #55c57a;

      i {
        color: #fff;
      }
    }
  }
}
</style>
