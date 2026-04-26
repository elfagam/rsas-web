<script setup>
import { ref, onMounted } from 'vue';
import api from '../../services/api';
import AdminBreadcrumbs from '../../components/admin/AdminBreadcrumbs.vue';

const breadcrumbItems = [{ label: 'Kelola Berita' }];
const newsList = ref([]);
const isLoading = ref(true);
const user = ref(JSON.parse(localStorage.getItem('rsas_user') || '{}'));

const fetchNews = async () => {
  try {
    const response = await api.get('/admin/news');
    newsList.value = response.data.data;
  } catch (error) {
    console.error('Failed to fetch news:', error);
  } finally {
    isLoading.value = false;
  }
};

const deleteNews = async (id) => {
  if (confirm('Apakah Anda yakin ingin menghapus berita ini?')) {
    try {
      await api.delete(`/admin/news/${id}`);
      fetchNews(); // Refresh daftar
    } catch (error) {
      alert('Gagal menghapus berita: ' + (error.response?.data?.error || 'Unknown error'));
    }
  }
};

const togglePublish = async (id) => {
  try {
    await api.patch(`/admin/news/${id}/publish`);
    fetchNews(); // Refresh daftar
  } catch (error) {
    alert('Gagal mengubah status publikasi: ' + (error.response?.data?.error || 'Unknown error'));
  }
};

onMounted(fetchNews);
</script>

<template>
  <div class="news-list-view">
    <div class="page-header">
      <div class="page-header__text">
        <AdminBreadcrumbs :items="breadcrumbItems" />
        <h1>Kelola <span>Berita</span></h1>
        <p>Manajemen publikasi artikel dan informasi RSUD Prof. Dr. H. Aloei Saboe</p>
      </div>
      <router-link to="/admin/berita/tambah" class="btn-add">
        <i class="fa-solid fa-plus"></i> Tulis Berita
      </router-link>
    </div>

    <div class="card-container">
      <div v-if="isLoading" class="loading-state">
        <i class="fa-solid fa-circle-notch fa-spin"></i>
        <p>Memuat data berita...</p>
      </div>

      <div v-else-if="newsList.length === 0" class="empty-state">
        <i class="fa-solid fa-newspaper"></i>
        <p>Belum ada berita yang diterbitkan.</p>
        <router-link to="/admin/berita/tambah" class="btn-link">Mulai Menulis</router-link>
      </div>

      <div v-else class="table-responsive">
        <table class="news-table">
          <thead>
            <tr>
              <th>Judul Berita</th>
              <th>Kategori</th>
              <th>Penulis</th>
              <th>Tanggal</th>
              <th>Status</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="news in newsList" :key="news.id">
              <td class="col-title">
                <strong>{{ news.title }}</strong>
                <span class="slug">{{ news.slug }}</span>
              </td>
              <td><span class="badge-category">{{ news.category || 'Umum' }}</span></td>
              <td>{{ news.author?.full_name || 'System' }}</td>
              <td>{{ new Date(news.created_at).toLocaleDateString('id-ID') }}</td>
              <td>
                <span :class="['status-pill', news.is_draft ? 'draft' : 'published']">
                  {{ news.is_draft ? 'Draft' : 'Published' }}
                </span>
              </td>
              <td class="col-actions">
                <!-- Tombol Lihat (Publik) -->
                <a 
                  :href="`/berita/${news.slug}`" 
                  target="_blank" 
                  class="action-btn view"
                  title="Lihat Berita di Web"
                >
                  <i class="fa-solid fa-eye"></i>
                </a>

                <!-- Editor & Admin bisa Publish/Unpublish -->
                <button 
                  v-if="user.role === 'admin' || user.role === 'editor'"
                  @click="togglePublish(news.id)"
                  :class="['action-btn publish', { 'active': !news.is_draft }]"
                  :title="news.is_draft ? 'Terbitkan Berita' : 'Kembalikan ke Draft'"
                >
                  <i :class="['fa-solid', news.is_draft ? 'fa-check' : 'fa-xmark']"></i>
                </button>

                <!-- Editor & Admin bisa edit -->
                <router-link 
                  v-if="user.role === 'admin' || user.role === 'editor'"
                  :to="`/admin/berita/edit/${news.id}`" 
                  class="action-btn edit"
                  title="Edit Berita"
                >
                  <i class="fa-solid fa-pen-to-square"></i>
                </router-link>

                <!-- Hanya Admin bisa hapus -->
                <button 
                  v-if="user.role === 'admin'"
                  @click="deleteNews(news.id)" 
                  class="action-btn delete"
                  title="Hapus Berita"
                >
                  <i class="fa-solid fa-trash"></i>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@use "../../assets/sass/base/variables" as *;

.news-list-view {
  padding: 5rem;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4rem;

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

.btn-add {
  padding: 1.5rem 3rem;
  background: #55c57a;
  color: #fff;
  text-decoration: none;
  border-radius: 15px;
  font-weight: 800;
  display: flex;
  align-items: center;
  gap: 1rem;
  box-shadow: 0 1rem 3rem rgba(85, 197, 122, 0.3);
  transition: all 0.3s;

  &:hover {
    transform: translateY(-3px);
    background: #4ab46d;
    box-shadow: 0 1.5rem 4rem rgba(85, 197, 122, 0.4);
  }
}

.card-container {
  background: #fff;
  border-radius: 30px;
  padding: 3rem;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.03);
}

.news-table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0 1rem;

  th {
    padding: 2rem;
    text-align: left;
    font-size: 1.4rem;
    color: #888;
    font-weight: 800;
    text-transform: uppercase;
    letter-spacing: 1px;
    border-bottom: 2px solid #f8f9fa;
  }

  td {
    padding: 2.5rem 2rem;
    background: #fff;
    font-size: 1.5rem;
    vertical-align: middle;
    border-bottom: 1px solid #f8f9fa;
  }

  tr:hover td {
    background: #fcfdfc;
  }
}

.col-title {
  max-width: 40rem;
  
  strong {
    display: block;
    font-size: 1.7rem;
    color: #333;
    margin-bottom: 0.5rem;
  }

  .slug {
    font-size: 1.2rem;
    color: #aaa;
    font-family: monospace;
  }
}

.badge-category {
  padding: 0.8rem 1.5rem;
  background: #f0faf3;
  color: #55c57a;
  border-radius: 10px;
  font-size: 1.3rem;
  font-weight: 700;
}

.status-pill {
  padding: 0.6rem 1.2rem;
  border-radius: 8px;
  font-size: 1.2rem;
  font-weight: 800;
  text-transform: uppercase;

  &.draft {
    background: #fff5f5;
    color: #e74c3c;
  }

  &.published {
    background: #f0faf3;
    color: #55c57a;
  }
}

.col-actions {
  display: flex;
  gap: 1rem;
}

.action-btn {
  width: 4.5rem;
  height: 4.5rem;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-decoration: none;
  border: 1px solid #eee;
  background: #fff;
  color: #666;
  font-size: 1.8rem;
  transition: all 0.3s;
  cursor: pointer;

  &:hover {
    transform: translateY(-2px);
  }

  &.edit:hover {
    color: #3498db;
    border-color: #3498db;
    background: #ebf5fb;
  }

  &.view:hover {
    color: #9b59b6;
    border-color: #9b59b6;
    background: #f4ecf7;
  }

  &.publish:hover {
    color: #f39c12;
    border-color: #f39c12;
    background: #fef5e7;
  }

  &.publish.active {
    background: #55c57a;
    color: #fff;
    border-color: #55c57a;
  }

  &.delete:hover {
    color: #e74c3c;
    border-color: #e74c3c;
    background: #fff5f5;
  }
}

.loading-state, .empty-state {
  text-align: center;
  padding: 8rem 0;

  i {
    font-size: 6rem;
    color: #eee;
    margin-bottom: 2rem;
  }

  p {
    font-size: 1.8rem;
    color: #999;
  }
}

.empty-state i { color: #55c57a; opacity: 0.2; }

.btn-link {
  display: inline-block;
  margin-top: 2rem;
  color: #55c57a;
  font-weight: 800;
  text-decoration: underline;
}
</style>
