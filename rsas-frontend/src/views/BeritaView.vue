<script setup>
import { ref, onMounted } from 'vue';
import api from '../services/api';

const newsList = ref([]);
const isLoading = ref(true);
const searchQuery = ref('');
const selectedCategory = ref('Semua Kategori');

const fetchPublicNews = async () => {
  try {
    const response = await api.get('/public/news');
    newsList.value = response.data.data;
  } catch (error) {
    console.error('Failed to fetch news:', error);
  } finally {
    isLoading.value = false;
  }
};

const stripHtml = (html) => {
  const tmp = document.createElement("DIV");
  tmp.innerHTML = html;
  return tmp.textContent || tmp.innerText || "";
};

onMounted(fetchPublicNews);
</script>

<template>
  <div class="berita-page">
    <!-- Hero Section with Background Image -->
    <section class="hero-section">
      <div class="hero-overlay"></div>
      <div class="container">
        <div class="hero-content" data-aos="zoom-out">
          <span class="hero-badge">Pusat Media & Edukasi</span>
          <h1 class="hero-title">Berita & <br><span>Kabar Kesehatan</span></h1>
          <p class="hero-subtitle">
            Tetap terinformasi dengan berita terbaru, pengumuman resmi, dan tips kesehatan terpercaya dari para ahli kami.
          </p>
        </div>
      </div>

      <!-- Search Bar Overlapping (Positioning Wrapper) -->
      <div class="search-container">
        <!-- Animation Wrapper -->
        <div class="search-box-wrapper" data-aos="fade-up" data-aos-delay="200">
          <div class="search-box">
            <div class="search-input-group">
              <i class="fa-solid fa-magnifying-glass search-icon"></i>
              <input 
                v-model="searchQuery" 
                type="text" 
                placeholder="Cari berita atau pengumuman..." 
              />
            </div>
            <div class="search-filter">
              <select v-model="selectedCategory">
                <option>Semua Kategori</option>
                <option>Kesehatan</option>
                <option>Kegiatan</option>
                <option>Pengumuman</option>
              </select>
            </div>
            <button class="btn-search">Cari</button>
          </div>
        </div>
      </div>
    </section>

    <!-- News List Section -->
    <section class="news-section">
      <div class="container">
        <div v-if="isLoading" class="loading-state">
          <div class="spinner"></div>
          <p>Memuat kabar terbaru untuk Anda...</p>
        </div>

        <div v-else-if="newsList.length === 0" class="empty-state">
          <i class="fa-solid fa-newspaper"></i>
          <h3>Belum ada berita yang diterbitkan.</h3>
          <p>Nantikan informasi terbaru dari kami segera.</p>
        </div>

        <div v-else class="news-grid">
          <div 
            v-for="news in newsList" 
            :key="news.id" 
            class="news-card"
            data-aos="fade-up"
          >
            <div class="news-card__img">
              <img :src="news.thumbnail || 'https://images.unsplash.com/photo-1576091160550-2173dba999ef?auto=format&fit=crop&q=80&w=800'" alt="News Image" />
              <div class="news-card__category">{{ news.category || 'Umum' }}</div>
            </div>
            <div class="news-card__body">
              <span class="news-card__date">
                <i class="fa-regular fa-calendar-check"></i>
                {{ new Date(news.created_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' }) }}
              </span>
              <h3>{{ news.title }}</h3>
              <p>{{ stripHtml(news.content).substring(0, 120) }}...</p>
              <router-link :to="`/berita/${news.slug}`" class="news-card__link">
                Baca Selengkapnya <i class="fa-solid fa-arrow-right-long"></i>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<style lang="scss" scoped>
@use "../assets/sass/base/variables" as *;

.berita-page {
  background: #fff;
}

.hero-section {
  min-height: 80vh;
  position: relative;
  background-image: url('https://images.unsplash.com/photo-1464822759023-fed622ff2c3b?auto=format&fit=crop&q=80&w=2000');
  background-size: cover;
  background-position: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  color: #fff;
  padding: 15rem 2rem 12rem; /* Tambah padding top agar tidak tertutup navbar */

  .hero-overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(to bottom, rgba(0,0,0,0.7) 0%, rgba(0,0,0,0.3) 100%);
    z-index: 1;
  }

  .container {
    position: relative;
    z-index: 2;
    width: 100%;
    max-width: 120rem;
  }

  .hero-badge {
    display: inline-block;
    padding: 1rem 2.5rem;
    background: rgba(85, 197, 122, 0.2);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.3);
    border-radius: 50px;
    font-size: 1.3rem;
    font-weight: 800;
    letter-spacing: 3px;
    text-transform: uppercase;
    margin-bottom: 3.5rem;
    color: #fff;
  }

  .hero-title {
    font-family: 'Outfit', sans-serif;
    font-size: clamp(4rem, 8vw, 8.5rem);
    font-weight: 900;
    line-height: 1;
    margin-bottom: 3.5rem;
    letter-spacing: -3px;
    text-shadow: 0 10px 30px rgba(0,0,0,0.3);

    span {
      font-family: 'Playfair Display', serif;
      font-weight: 400;
      font-style: italic;
      opacity: 0.95;
      letter-spacing: 0;
    }
  }

  .hero-subtitle {
    font-size: 2.2rem;
    max-width: 85rem;
    margin: 0 auto;
    opacity: 0.95;
    line-height: 1.6;
    text-shadow: 0 5px 15px rgba(0,0,0,0.2);
  }
}

.search-container {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  width: 100%;
  display: flex;
  justify-content: center;
  z-index: 100;
}

.search-box-wrapper {
  width: 90%;
  max-width: 110rem;
  transform: translateY(50%); /* Half-overlap with the next section */
}

.search-box {
  background: #fff;
  padding: 1.5rem;
  border-radius: 25px;
  box-shadow: 0 2rem 5rem rgba(0,0,0,0.15);
  display: flex;
  align-items: center;
  gap: 1.5rem;

  @include respond("phone") {
    flex-direction: column;
    border-radius: 20px;
  }

  .search-input-group {
    flex: 1;
    display: flex;
    align-items: center;
    gap: 1.5rem;
    padding: 0 2rem;

    .search-icon {
      color: #55c57a;
      font-size: 2rem;
    }

    input {
      width: 100%;
      border: none;
      padding: 1.5rem 0;
      font-size: 1.6rem;
      color: #333;
      &:focus { outline: none; }
    }
  }

  .search-filter {
    border-left: 1px solid #eee;
    padding: 0 2rem;

    select {
      border: none;
      font-size: 1.6rem;
      font-weight: 600;
      color: #666;
      background: none;
      cursor: pointer;
      &:focus { outline: none; }
    }

    @include respond("phone") {
      border-left: none;
      border-top: 1px solid #eee;
      width: 100%;
      padding: 1.5rem 2rem;
    }
  }

  .btn-search {
    padding: 1.8rem 4rem;
    background: #55c57a;
    color: #fff;
    border: none;
    border-radius: 18px;
    font-size: 1.6rem;
    font-weight: 800;
    cursor: pointer;
    transition: all 0.3s;

    &:hover {
      background: #4ab46d;
      transform: scale(1.05);
      box-shadow: 0 1rem 2rem rgba(85, 197, 122, 0.3);
    }
  }
}

.news-section {
  padding: 18rem 0 15rem;

  .container {
    max-width: 120rem;
    margin: 0 auto;
    padding: 0 3rem;
  }
}

.news-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(35rem, 1fr));
  gap: 5rem;
}

.news-card {
  background: #fff;
  border-radius: 30px;
  overflow: hidden;
  box-shadow: 0 1rem 4rem rgba(0,0,0,0.05);
  transition: all 0.4s;
  display: flex;
  flex-direction: column;

  &:hover {
    transform: translateY(-1.5rem);
    box-shadow: 0 3rem 6rem rgba(85, 197, 122, 0.15);
  }

  &__img {
    height: 25rem;
    position: relative;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      transition: all 0.6s;
    }

    .news-card__category {
      position: absolute;
      bottom: 2rem;
      left: 2rem;
      padding: 0.8rem 1.8rem;
      background: #55c57a;
      color: #fff;
      border-radius: 12px;
      font-size: 1.2rem;
      font-weight: 800;
      text-transform: uppercase;
      letter-spacing: 1px;
    }
  }

  &:hover &__img img {
    transform: scale(1.1);
  }

  &__body {
    padding: 4rem;
    flex: 1;
    display: flex;
    flex-direction: column;

    .news-card__date {
      display: flex;
      align-items: center;
      gap: 1rem;
      font-size: 1.4rem;
      color: #999;
      font-weight: 600;
      margin-bottom: 2rem;

      i { color: #55c57a; }
    }

    h3 {
      font-family: 'Outfit', sans-serif;
      font-size: 2.4rem;
      font-weight: 800;
      line-height: 1.3;
      color: #333;
      margin-bottom: 2rem;
    }

    p {
      font-size: 1.6rem;
      color: #666;
      line-height: 1.7;
      margin-bottom: 3rem;
      flex: 1;
    }

    .news-card__link {
      display: flex;
      align-items: center;
      gap: 1.2rem;
      text-decoration: none;
      color: #55c57a;
      font-size: 1.6rem;
      font-weight: 800;
      transition: all 0.3s;

      &:hover {
        gap: 2rem;
        color: #333;
      }
    }
  }
}

.loading-state, .empty-state {
  text-align: center;
  padding: 10rem 0;
  
  i { font-size: 10rem; color: #eee; margin-bottom: 3rem; }
  h3 { font-size: 3rem; font-weight: 800; margin-bottom: 1.5rem; }
  p { font-size: 1.8rem; color: #999; }
}

.spinner {
  width: 7rem;
  height: 7rem;
  border: 6px solid #f3f3f3;
  border-top: 6px solid #55c57a;
  border-radius: 50%;
  margin: 0 auto 3rem;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
