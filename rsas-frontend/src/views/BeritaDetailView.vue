<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import api from '../services/api';

const route = useRoute();
const router = useRouter();
const isLoading = ref(true);
const news = ref(null);
const scrollProgress = ref(0);
const relatedNews = ref([]);

const fetchNewsDetail = async () => {
  isLoading.value = true;
  window.scrollTo(0, 0);
  try {
    const slug = route.params.slug;
    const response = await api.get(`/public/news/${slug}`);
    news.value = response.data.data;
    document.title = `${news.value.title} | RSUD Prof. Dr. H. Aloei Saboe`;
    fetchRelatedNews();
  } catch (error) {
    console.error('Failed to fetch news detail:', error);
    router.push('/berita');
  } finally {
    isLoading.value = false;
  }
};

const fetchRelatedNews = async () => {
  try {
    const res = await api.get('/public/news');
    relatedNews.value = res.data.data
      .filter(item => item.id !== news.value?.id)
      .slice(0, 3);
  } catch (error) {
    console.error('Failed to fetch related news:', error);
  }
};

const updateScrollProgress = () => {
  const h = document.documentElement;
  const b = document.body;
  const st = 'scrollTop';
  const sh = 'scrollHeight';
  scrollProgress.value = (h[st] || b[st]) / ((h[sh] || b[sh]) - h.clientHeight) * 100;
};

const downloadMainImage = async () => {
  if (!news.value?.thumbnail) return;
  const imageUrl = news.value.thumbnail;
  const fileName = `RSAS-${news.value.slug}.jpg`;
  try {
    const response = await fetch(imageUrl);
    const blob = await response.blob();
    const url = window.URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = fileName;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    window.open(imageUrl, '_blank');
  }
};

const scrollToContent = () => {
  const contentSection = document.querySelector('.article-body');
  if (contentSection) {
    contentSection.scrollIntoView({ behavior: 'smooth' });
  }
};

onMounted(() => {
  fetchNewsDetail();
  window.addEventListener('scroll', updateScrollProgress);
});

onUnmounted(() => {
  window.removeEventListener('scroll', updateScrollProgress);
});

watch(() => route.params.slug, () => {
  fetchNewsDetail();
});
</script>

<template>
  <div class="news-detail-page">
    <div class="reading-progress-container">
      <div class="reading-progress-bar" :style="{ width: scrollProgress + '%' }"></div>
    </div>

    <template v-if="news">
      <!-- Hybrid Hero (Image Only) -->
      <section class="hybrid-hero">
        <div class="hybrid-hero__bg" :style="{ backgroundImage: `url(${news.thumbnail || 'https://images.unsplash.com/photo-1576091160550-2173dba999ef?auto=format&fit=crop&q=80&w=2000'})` }"></div>
      </section>

      <!-- Floating Title Card -->
      <header class="hybrid-header">
        <div class="container container--narrow">
          <div class="hybrid-card" data-aos="fade-up">
            <nav class="breadcrumb-hybrid">
              <router-link to="/">Home</router-link>
              <span class="sep">/</span>
              <router-link to="/berita">Berita</router-link>
              <span class="sep">/</span>
              <span class="cur">Detail</span>
            </nav>

            <div class="article-meta">
              <span class="badge">{{ news.category || 'Kesehatan' }}</span>
              <span class="date">
                <i class="fa-regular fa-calendar-check"></i>
                {{ new Date(news.created_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' }) }}
              </span>
            </div>
            
            <h1 class="article-title">{{ news.title }}</h1>
            
            <div class="article-author">
              <div class="author-info">
                <span class="by">Publikasi Oleh:</span>
                <span class="name">Humas RSAS Gorontalo</span>
              </div>
            </div>

            <!-- Action Button Corner -->
            <button @click="scrollToContent" class="card-action-btn" title="Mulai Membaca">
              <i class="fa-solid fa-chevron-down"></i>
            </button>
          </div>
        </div>
      </header>

      <!-- Main Content Area -->
      <main class="article-body">
        <div class="container container--narrow">
          <article class="article-content" v-html="news.content" data-aos="fade-up"></article>

          <!-- Article Actions & Footer -->
          <footer class="article-footer" data-aos="fade-up">
            <div class="article-actions">
              <button @click="downloadMainImage" class="btn-download">
                <i class="fa-solid fa-cloud-arrow-down"></i> Unduh Infografis
              </button>
            </div>
            
            <div class="share-box">
              <span class="share-label">Bagikan:</span>
              <div class="share-buttons">
                <a href="#" class="share-btn fb"><i class="fa-brands fa-facebook-f"></i></a>
                <a href="#" class="share-btn wa"><i class="fa-brands fa-whatsapp"></i></a>
                <a href="#" class="share-btn tw"><i class="fa-brands fa-x-twitter"></i></a>
              </div>
            </div>
          </footer>

          <!-- Related News Section -->
          <section v-if="relatedNews.length > 0" class="related-section" data-aos="fade-up">
            <h3 class="section-title">Berita <span>Terkait</span></h3>
            <div class="related-grid">
              <router-link 
                v-for="item in relatedNews" 
                :key="item.id" 
                :to="'/berita/' + item.slug"
                class="related-card"
              >
                <div class="related-card__img">
                  <img :src="item.thumbnail || 'https://images.unsplash.com/photo-1576091160550-2173dba999ef?auto=format&fit=crop&q=80&w=400'" alt="Related">
                </div>
                <div class="related-card__info">
                  <span class="cat">{{ item.category }}</span>
                  <h4>{{ item.title }}</h4>
                </div>
              </router-link>
            </div>
          </section>
        </div>
      </main>
    </template>

    <div v-else-if="isLoading" class="loading-full">
      <div class="spinner"></div>
      <p>Memuat konten premium...</p>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@use "../assets/sass/base/variables" as *;

.news-detail-page {
  background: #f8f9fa; /* Background sedikit abu-abu agar kartu putih menonjol */
  min-height: 100vh;
}

/* Containers */
.container--narrow {
  max-width: 90rem;
  margin: 0 auto;
  padding: 0 4rem;
}

/* Reading Progress */
.reading-progress-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 4px;
  z-index: 9999;
}
.reading-progress-bar {
  height: 100%;
  background: linear-gradient(to right, #55c57a, #2ecc71);
}

/* Hybrid Hero */
.hybrid-hero {
  height: 60vh;
  width: 100%;
  overflow: hidden;
  position: relative;

  &__bg {
    width: 100%;
    height: 100%;
    background-size: cover;
    background-position: center;
    background-attachment: fixed; /* Parallax effect */
  }
}

/* Floating Card Header */
.hybrid-header {
  margin-top: -15rem; /* Menindih gambar di atas */
  position: relative;
  z-index: 10;
}

.hybrid-card {
  background: rgba(255, 255, 255, 0.45); /* Jauh lebih transparan */
  backdrop-filter: blur(40px) saturate(200%); /* Blur lebih dalam & saturasi lebih kuat */
  -webkit-backdrop-filter: blur(40px) saturate(200%);
  padding: 6rem;
  border-radius: 40px;
  box-shadow: 0 4rem 10rem rgba(0,0,0,0.2);
  border: 1px solid rgba(255, 255, 255, 0.2); /* Border lebih tipis & halus */
  animation: pulse-fade 4s infinite ease-in-out; /* Efek Fade pada Bayangan */
  position: relative;
  
  .breadcrumb-hybrid {
    display: flex;
    gap: 1.2rem;
    font-size: 1.3rem;
    font-weight: 800; /* Lebih tebal */
    margin-bottom: 3rem;
    color: #888;
    a { 
      color: #2d8a4d; /* Hijau yang lebih kuat/gelap untuk teks */
      text-decoration: none; 
      text-shadow: 0 1px 2px rgba(255,255,255,0.8); /* Glow putih agar tidak tenggelam */
      transition: 0.3s; 
      &:hover { color: #55c57a; text-decoration: underline; } 
    }
    .sep { margin: 0 0.5rem; opacity: 0.5; }
    .cur { color: #222; }
  }

  .article-meta {
    display: flex;
    align-items: center;
    gap: 2rem;
    margin-bottom: 2.5rem;

    .badge {
      padding: 0.6rem 2rem;
      background: #55c57a;
      color: #fff;
      border-radius: 12px;
      font-size: 1.2rem;
      font-weight: 800;
      text-transform: uppercase;
      letter-spacing: 1px;
    }

    .date {
      font-size: 1.5rem;
      color: #555; /* Sedikit lebih gelap agar teks abu-abu terbaca */
      font-weight: 700;
      i { 
        margin-right: 0.8rem; 
        color: #2d8a4d; /* Hijau yang lebih berani */
        text-shadow: 0 1px 2px rgba(255,255,255,0.8);
      }
    }
  }

  .article-title {
    font-family: 'Outfit', sans-serif;
    font-size: clamp(3rem, 5vw, 5.5rem);
    font-weight: 900;
    line-height: 1.1;
    color: #1a1a1a;
    margin-bottom: 3.5rem;
    letter-spacing: -2px;
  }

  .article-author {
    display: flex;
    align-items: center;
    gap: 1.5rem;
    padding-top: 2.5rem;
    border-top: 1px solid #f0f0f0;
    
    .author-info {
      font-size: 1.5rem;
      .by { color: #999; margin-right: 0.8rem; }
      .name { font-weight: 800; color: #333; }
    }
  }

  /* Action Button Corner */
  .card-action-btn {
    position: absolute;
    bottom: 4rem;
    right: 4rem;
    width: 6rem;
    height: 6rem;
    background: #55c57a;
    color: #fff;
    border: none;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 2.2rem;
    cursor: pointer;
    box-shadow: 0 1rem 3rem rgba(85, 197, 122, 0.4);
    transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
    z-index: 15;

    &:hover {
      transform: translateY(-5px) scale(1.1);
      background: #2d8a4d;
      box-shadow: 0 1.5rem 4rem rgba(45, 138, 77, 0.5);
    }

    i {
      animation: bounce 2s infinite;
    }
  }
}

@keyframes pulse-fade {
  0% { box-shadow: 0 4rem 10rem rgba(0,0,0,0.2); }
  50% { box-shadow: 0 4rem 12rem rgba(85, 197, 122, 0.25); }
  100% { box-shadow: 0 4rem 10rem rgba(0,0,0,0.2); }
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {transform: translateY(0);}
  40% {transform: translateY(-5px);}
  60% {transform: translateY(-3px);}
}

/* Article Body */
.article-body {
  padding: 6rem 0 12rem;
  background: #fff; /* Konten utama tetap di latar putih */
}

.article-content {
  font-family: 'Lato', sans-serif;
  font-size: 2.1rem;
  line-height: 1.75; /* Lebih tegas dan padat */
  color: #222; /* Charcoal Deep untuk kesan mewah */

  /* Drop Cap: Huruf Pertama Mewah */
  & > p:first-of-type::first-letter {
    font-family: 'Playfair Display', serif;
    font-size: 8.5rem;
    font-weight: 900;
    float: left;
    line-height: 1;
    margin-right: 1.5rem;
    margin-top: 0.5rem;
    color: #55c57a;
    text-shadow: 2px 2px 0px rgba(85, 197, 122, 0.1);
  }

  /* Headings yang Tegas */
  :deep(h2), :deep(h3) {
    font-family: 'Outfit', sans-serif;
    color: #111;
    font-weight: 800;
    line-height: 1.2;
    margin: 6rem 0 3rem;
    letter-spacing: -1px;
  }

  :deep(h2) { font-size: 3.8rem; }
  :deep(h3) { font-size: 3rem; }

  :deep(p) { margin-bottom: 4rem; text-align: justify; }
  
  :deep(img) {
    width: 100%;
    border-radius: 25px;
    margin: 6rem 0;
    box-shadow: 0 4rem 8rem rgba(0,0,0,0.12);
  }

  :deep(blockquote) {
    font-family: 'Playfair Display', serif;
    font-style: italic;
    font-size: 2.8rem;
    line-height: 1.6;
    padding: 4rem 5rem;
    border-left: 10px solid #55c57a;
    background: #f9fdfb;
    margin: 7rem 0;
    color: #111;
    border-radius: 0 25px 25px 0;
    position: relative;

    &::before {
      content: '"';
      position: absolute;
      top: -2rem;
      left: 2rem;
      font-size: 10rem;
      color: rgba(85, 197, 122, 0.1);
      font-family: serif;
    }
  }

  /* Filter Pembersih Spasi: Mencegah <br> atau <p> kosong merusak layout */
  :deep(br + br) {
    display: none; 
  }

  /* Sembunyikan paragraf yang sengaja dikosongkan atau hanya berisi <br> */
  :deep(p:empty),
  :deep(p:has(> br:only-child)) {
    display: none;
    margin: 0;
    padding: 0;
  }

  /* List Styling */
  :deep(ul), :deep(ol) {
    margin-bottom: 4rem;
    padding-left: 3rem;
    li { margin-bottom: 1.5rem; }
  }
}

.article-footer {
  margin-top: 8rem;
  padding: 5rem 0;
  border-top: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 4rem;

  .btn-download {
    padding: 1.6rem 3.5rem;
    background: #f8f9fa;
    border: 2px solid #55c57a;
    color: #55c57a;
    border-radius: 50px;
    font-size: 1.5rem;
    font-weight: 800;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 1.5rem;
    transition: all 0.3s;
    &:hover { background: #55c57a; color: #fff; transform: translateY(-5px); box-shadow: 0 1rem 3rem rgba(85, 197, 122, 0.2); }
  }

  .share-box {
    display: flex;
    align-items: center;
    gap: 2.5rem;
    .share-label { font-weight: 800; font-size: 1.6rem; color: #333; }
    .share-buttons {
      display: flex;
      gap: 1.2rem;
      .share-btn {
        width: 4.8rem;
        height: 4.8rem;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        background: #f9fdfb;
        color: #55c57a;
        font-size: 1.8rem;
        transition: all 0.3s;
        &:hover { background: #55c57a; color: #fff; transform: translateY(-5px); }
      }
    }
  }
}

/* Related News */
.related-section {
  margin-top: 10rem;
  .section-title { font-family: 'Outfit', sans-serif; font-size: 3rem; font-weight: 900; margin-bottom: 5rem; span { color: #55c57a; } }
  .related-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(26rem, 1fr)); gap: 4rem; }
  .related-card {
    text-decoration: none; color: inherit;
    &__img { height: 18rem; border-radius: 20px; overflow: hidden; margin-bottom: 2rem; img { width: 100%; height: 100%; object-fit: cover; transition: 0.5s; } }
    &__info {
      .cat { font-size: 1.2rem; font-weight: 800; color: #55c57a; text-transform: uppercase; margin-bottom: 0.8rem; display: block; }
      h4 { font-size: 1.8rem; font-weight: 800; line-height: 1.4; color: #333; transition: 0.3s; }
    }
    &:hover { .related-card__img img { transform: scale(1.1); } h4 { color: #55c57a; } }
  }
}

/* Loading */
.loading-full { height: 100vh; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 2rem; .spinner { width: 5rem; height: 5rem; border: 4px solid #f3f3f3; border-top: 4px solid #55c57a; border-radius: 50%; animation: spin 1s linear infinite; } }

@media (max-width: 900px) {
  .hybrid-card { padding: 3rem; border-radius: 20px; margin-top: -10rem; }
}

@media (max-width: 600px) {
  .article-title { font-size: 2.8rem; }
  .article-content { font-size: 1.8rem; }
  .article-footer { flex-direction: column; align-items: flex-start; }
}
</style>
