<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../../services/api';
import AdminBreadcrumbs from '../../components/admin/AdminBreadcrumbs.vue';

const router = useRouter();
const isLoading = ref(false);
const errorMessage = ref('');

const breadcrumbItems = [
  { label: 'Kelola Berita', link: '/admin/berita' },
  { label: 'Tulis Berita Baru' }
];

const form = ref({
  title: '',
  category: 'Umum',
  content: '',
  thumbnail: '',
});

const handleSubmit = async () => {
  isLoading.value = true;
  errorMessage.value = '';

  try {
    // Menghapus trailing slash agar sesuai dengan backend fix
    await api.post('/admin/news', form.value);
    alert('Berita berhasil disimpan!');
    router.push('/admin/berita');
  } catch (error) {
    errorMessage.value = error.response?.data?.error || 'Gagal menyimpan berita.';
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="news-form-view">
    <div class="page-header">
      <div class="page-header__text">
        <AdminBreadcrumbs :items="breadcrumbItems" />
        <router-link to="/admin/berita" class="btn-back">
          <i class="fa-solid fa-arrow-left"></i> Kembali ke Daftar
        </router-link>
        <h1>Tulis <span>Berita Baru</span></h1>
        <p>Gunakan bahasa yang informatif dan edukatif untuk masyarakat</p>
      </div>
    </div>

    <div class="form-container" data-aos="fade-up">
      <form @submit.prevent="handleSubmit" class="main-form">
        <div v-if="errorMessage" class="error-box">
          <i class="fa-solid fa-circle-exclamation"></i>
          {{ errorMessage }}
        </div>

        <div class="form-grid">
          <div class="form-main">
            <div class="form-group">
              <label>Judul Berita</label>
              <input 
                v-model="form.title" 
                type="text" 
                placeholder="Masukkan judul berita yang menarik..." 
                required
              />
            </div>

            <div class="form-group">
              <label>Konten Berita</label>
              <QuillEditor 
                v-model:content="form.content" 
                contentType="html"
                theme="snow"
                placeholder="Tuliskan isi berita selengkap mungkin di sini..."
                :toolbar="['bold', 'italic', 'underline', { 'list': 'ordered'}, { 'list': 'bullet' }, 'link', 'clean']"
              />
            </div>
          </div>

          <aside class="form-sidebar">
            <div class="form-group">
              <label>Kategori</label>
              <select v-model="form.category">
                <option value="Umum">Umum</option>
                <option value="Kesehatan">Kesehatan</option>
                <option value="Event">Event / Kegiatan</option>
                <option value="Pengumuman">Pengumuman</option>
                <option value="Pendidikan">Pendidikan</option>
              </select>
            </div>

            <div class="form-group">
              <label>Thumbnail URL</label>
              <input 
                v-model="form.thumbnail" 
                type="text" 
                placeholder="https://..." 
              />
              <p class="input-help">Gunakan link gambar (Google Cloud integration segera hadir)</p>
            </div>

            <div class="publish-card">
              <h3>Publikasi</h3>
              <p>Status saat ini: <strong>Draft</strong></p>
              <div class="publish-actions">
                <button type="submit" class="btn-save" :disabled="isLoading">
                  <span v-if="!isLoading"><i class="fa-solid fa-floppy-disk"></i> Simpan Berita</span>
                  <span v-else><i class="fa-solid fa-circle-notch fa-spin"></i> Menyimpan...</span>
                </button>
              </div>
            </div>
          </aside>
        </div>
      </form>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@use "../../assets/sass/base/variables" as *;

.news-form-view {
  padding: 5rem;
}

.btn-back {
  display: inline-flex;
  align-items: center;
  gap: 1rem;
  text-decoration: none;
  color: #888;
  font-size: 1.4rem;
  font-weight: 700;
  margin-bottom: 2rem;
  transition: all 0.3s;

  &:hover {
    color: #55c57a;
  }
}

.page-header {
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

.form-container {
  background: #fff;
  border-radius: 30px;
  padding: 4rem;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.03);
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 35rem;
  gap: 5rem;

  @include respond("tablet") {
    grid-template-columns: 1fr;
  }
}

.form-group {
  margin-bottom: 3rem;

  label {
    display: block;
    font-size: 1.4rem;
    font-weight: 800;
    color: #444;
    margin-bottom: 1.2rem;
    text-transform: uppercase;
    letter-spacing: 1px;
  }

  input, select, textarea {
    width: 100%;
    padding: 1.8rem 2rem;
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
    }
  }

  .ql-container {
    min-height: 40rem;
    font-size: 1.6rem;
    font-family: 'Inter', sans-serif;
    border-bottom-left-radius: 15px;
    border-bottom-right-radius: 15px;
  }

  .ql-toolbar {
    border-top-left-radius: 15px;
    border-top-right-radius: 15px;
    background: #fdfdfd;
  }

  .input-help {
    font-size: 1.2rem;
    color: #aaa;
    margin-top: 1rem;
    font-style: italic;
  }
}

.publish-card {
  background: #f8f9fa;
  padding: 3rem;
  border-radius: 20px;
  border: 1px solid #eee;

  h3 {
    font-size: 1.8rem;
    font-weight: 800;
    margin-bottom: 1.5rem;
  }

  p {
    font-size: 1.4rem;
    color: #666;
    margin-bottom: 2.5rem;

    strong {
      color: #e74c3c;
    }
  }
}

.btn-save {
  width: 100%;
  padding: 1.8rem;
  background: #55c57a;
  color: #fff;
  border: none;
  border-radius: 12px;
  font-size: 1.6rem;
  font-weight: 800;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;

  &:hover {
    background: #4ab46d;
    transform: translateY(-3px);
    box-shadow: 0 1rem 3rem rgba(85, 197, 122, 0.3);
  }

  &:disabled {
    background: #ccc;
    cursor: not-allowed;
  }
}

.error-box {
  background: #fff5f5;
  color: #e74c3c;
  padding: 2rem;
  border-radius: 15px;
  margin-bottom: 3rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  font-size: 1.5rem;
  font-weight: 700;
  border-left: 5px solid #e74c3c;
}
</style>
