<script setup>
import { ref, onMounted, onUnmounted, computed, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();
const isLoading = ref(true);
const news = ref(null);
const scrollProgress = ref(0);

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

// Simulasi Database (Idealnya diambil dari API berdasarkan route.params.id)
const allNews = [
  {
    id: 1,
    kategori: "PENGUMUMAN",
    tanggal: "25 April 2026",
    judul: "Perkuat Mutu Pelayanan, Kemenkes RI Verifikasi Standarisasi RSUD",
    ringkasan: "Kementerian Kesehatan RI melakukan verifikasi lapangan terhadap standar kualitas pelayanan di RSUD Prof. Dr. H. Aloei Saboe.",
    konten: `
      <p>GORONTALO - Kementerian Kesehatan Republik Indonesia melakukan kunjungan kerja dalam rangka verifikasi lapangan terhadap standar kualitas pelayanan di RSUD Prof. Dr. H. Aloei Saboe.</p>
      <p>Kunjungan ini bertujuan untuk memastikan bahwa setiap unit layanan telah memenuhi Standar Pelayanan Minimal (SPM) yang ditetapkan secara nasional. Tim verifikator meninjau langsung beberapa unit vital seperti IGD Terpadu, Pusat Jantung, dan Laboratorium.</p>
      <blockquote>"Kualitas pelayanan adalah prioritas utama kami. Verifikasi ini membantu kami mengidentifikasi area yang perlu ditingkatkan untuk kepuasan pasien yang lebih baik," ujar perwakilan manajemen RSAS.</blockquote>
      <p>Hasil dari verifikasi ini diharapkan dapat mempertahankan akreditasi paripurna yang telah diraih oleh RSUD Aloei Saboe selama beberapa tahun terakhir.</p>
    `,
    gambar: "nat-1-large.jpg",
    author: "Humas RSAS",
    tags: ["Kemenkes", "Pelayanan", "Verifikasi"]
  },
  {
    id: 2,
    kategori: "LAYANAN",
    tanggal: "24 April 2026",
    judul: "Pelayanan Tetap Optimal, RSUD Prof. Dr. H. Aloei Saboe Pastikan Kesiapan SDM",
    ringkasan: "Manajemen RSAS memastikan seluruh layanan kesehatan tetap berjalan 24 jam penuh selama libur panjang.",
    konten: `
      <p>Menjelang libur panjang, manajemen RSUD Prof. Dr. H. Aloei Saboe memastikan bahwa seluruh layanan kesehatan akan tetap berjalan selama 24 jam penuh.</p>
      <p>Pengaturan jadwal piket tenaga medis, termasuk dokter spesialis on-call, telah disiapkan untuk mengantisipasi lonjakan pasien di unit gawat darurat.</p>
      <p>Selain pelayanan medis, unit penunjang seperti farmasi dan kasir juga tetap beroperasi normal guna memudahkan masyarakat Gorontalo dalam mendapatkan akses kesehatan tanpa hambatan.</p>
    `,
    gambar: "nat-2-large.jpg",
    author: "Redaksi RSAS",
    tags: ["Libur Nasional", "SDM", "Layanan"]
  },
  {
    id: 3,
    kategori: "KESEHATAN",
    tanggal: "22 April 2026",
    judul: "Tips Menjaga Kesehatan Jantung di Usia Produktif",
    ringkasan: "Pusat Jantung RSAS berbagi tips sederhana untuk menjaga kesehatan kardiovaskular bagi pekerja muda.",
    konten: `
      <p>Penyakit jantung tidak lagi hanya menyerang usia lanjut. Pola hidup tidak sehat membuat usia produktif kini rentan terhadap masalah kardiovaskular.</p>
      <p>Pusat Jantung RSUD Prof. Dr. H. Aloei Saboe menyarankan pola makan seimbang, rutin berolahraga minimal 30 menit sehari, dan melakukan check-up rutin untuk mendeteksi dini faktor risiko.</p>
    `,
    gambar: "nat-3-large.jpg",
    author: "dr. Andi Specialist",
    tags: ["Jantung", "Tips Sehat", "Edukasi"]
  },
  {
    id: 4,
    kategori: "INOVASI",
    tanggal: "20 April 2026",
    judul: "RSUD Aloei Saboe Luncurkan Sistem Pendaftaran Online Terbaru",
    ringkasan: "Memudahkan pasien, RSAS kini menghadirkan aplikasi pendaftaran online dengan fitur antrean real-time.",
    konten: `
      <p>Transformasi digital terus dilakukan oleh RSUD Prof. Dr. H. Aloei Saboe. Kini, masyarakat bisa melakukan pendaftaran rawat jalan melalui aplikasi mobile.</p>
      <p>Sistem baru ini diharapkan dapat memangkas waktu tunggu di loket pendaftaran secara signifikan dan memberikan kenyamanan lebih bagi pasien.</p>
    `,
    gambar: "nat-4.jpg",
    author: "Tim IT RSAS",
    tags: ["Digital", "Inovasi", "Layanan"]
  }
];

// Menghitung Berita Terkait (3 berita lain)
const relatedNews = computed(() => {
  if (!news.value) return [];
  return allNews
    .filter(item => item.id !== news.value.id)
    .slice(0, 3);
});

// Update SEO (Title & Schema)
const updateSEO = (article) => {
  if (!article) return;
  
  // Title
  document.title = `${article.judul} | RSUD Prof. Dr. H. Aloei Saboe`;

  // JSON-LD Schema
  const schemaId = 'news-article-schema';
  let script = document.getElementById(schemaId);
  if (!script) {
    script = document.createElement('script');
    script.id = schemaId;
    script.type = 'application/ld+json';
    document.head.appendChild(script);
  }

  const schemaData = {
    "@context": "https://schema.org",
    "@type": "NewsArticle",
    "headline": article.judul,
    "image": [getImageUrl(article.gambar)],
    "datePublished": article.tanggal, // Idealnya format ISO
    "author": [{
      "@type": "Organization",
      "name": "RSUD Prof. Dr. H. Aloei Saboe",
      "url": window.location.origin
    }]
  };
  
  script.text = JSON.stringify(schemaData);
};

// Tracking Scroll Progress
const updateScrollProgress = () => {
  const winScroll = document.body.scrollTop || document.documentElement.scrollTop;
  const height = document.documentElement.scrollHeight - document.documentElement.clientHeight;
  const scrolled = (winScroll / height) * 100;
  scrollProgress.value = scrolled;
};

// Fungsi untuk Load Data Berita
const fetchNewsData = () => {
  isLoading.value = true;
  window.scrollTo(0, 0); // Scroll ke atas setiap ganti berita
  
  const newsId = parseInt(route.params.id);
  
  setTimeout(() => {
    const found = allNews.find(item => item.id === newsId);
    if (found) {
      news.value = found;
      updateSEO(found);
    } else {
      console.warn("Berita tidak ditemukan!");
      news.value = null;
    }
    isLoading.value = false;
  }, 600); // Sedikit lebih cepat untuk UX navigasi antar berita
};

// Pantau perubahan ID di URL (untuk navigasi berita terkait)
watch(() => route.params.id, (newId) => {
  if (newId) fetchNewsData();
});

onMounted(() => {
  window.addEventListener('scroll', updateScrollProgress);
  fetchNewsData();
});

onUnmounted(() => {
  window.removeEventListener('scroll', updateScrollProgress);
  // Cleanup Schema
  const script = document.getElementById('news-article-schema');
  if (script) script.remove();
});
</script>

<template>
  <div class="news-detail-page">
    <!-- Reading Progress Bar -->
    <div class="reading-progress-container">
      <div class="reading-progress-bar" :style="{ width: scrollProgress + '%' }"></div>
    </div>

    <div class="container" v-if="isLoading">
      <div class="skeleton-detail">
        <div class="skeleton-line skeleton-line--title"></div>
        <div class="skeleton-line skeleton-line--meta"></div>
        <div class="skeleton-img"></div>
        <div class="skeleton-line"></div>
        <div class="skeleton-line"></div>
        <div class="skeleton-line" style="width: 60%"></div>
      </div>
    </div>

    <template v-else-if="news">
      <!-- Breadcrumbs -->
      <nav class="breadcrumb">
        <div class="container">
          <router-link to="/">Beranda</router-link>
          <span class="separator">&raquo;</span>
          <router-link to="/berita">Berita</router-link>
          <span class="separator">&raquo;</span>
          <span class="current">{{ news.judul }}</span>
        </div>
      </nav>

      <article class="article">
        <div class="container container--narrow">
          <header class="article__header">
            <div class="article__meta">
              <span class="article__category">{{ news.kategori }}</span>
              <span class="article__date">{{ news.tanggal }}</span>
            </div>
            <h1 class="article__title">{{ news.judul }}</h1>
            <div class="article__author">
              Oleh: <strong>{{ news.author }}</strong>
            </div>
          </header>

          <figure class="article__img-box">
            <img :src="getImageUrl(news.gambar)" :alt="news.judul" class="article__img" />
          </figure>

          <div class="article__content" v-html="news.konten"></div>

          <footer class="article__footer">
            <div class="article__tags">
              <span v-for="tag in news.tags" :key="tag" class="tag">#{{ tag }}</span>
            </div>
            <div class="article__share">
              <span>Bagikan Berita:</span>
              <div class="share-links">
                <a href="#" title="Bagikan ke Facebook"><i class="fa-brands fa-facebook"></i></a>
                <a href="#" title="Bagikan ke WhatsApp"><i class="fa-brands fa-whatsapp"></i></a>
                <a href="#" title="Bagikan ke Twitter"><i class="fa-brands fa-twitter"></i></a>
              </div>
            </div>
          </footer>

          <!-- Related News Section -->
          <section class="related-news">
            <h3 class="related-news__title">Berita Terkait</h3>
            <div class="related-news__grid">
              <router-link 
                v-for="item in relatedNews" 
                :key="item.id" 
                :to="'/berita/' + item.id"
                class="related-card"
              >
                <div class="related-card__img-box">
                  <img :src="getImageUrl(item.gambar)" :alt="item.judul">
                </div>
                <div class="related-card__info">
                  <span class="related-card__category">{{ item.kategori }}</span>
                  <h4 class="related-card__judul">{{ item.judul }}</h4>
                </div>
              </router-link>
            </div>
          </section>

          <div class="article__navigation">
            <router-link to="/berita" class="btn-back">
              &larr; Kembali ke Daftar Berita
            </router-link>
          </div>
        </div>
      </article>
    </template>

    <div v-else class="container u-center-text">
       <h2>Maaf, Berita tidak ditemukan.</h2>
       <router-link to="/berita" class="btn-back">Kembali</router-link>
    </div>
  </div>
</template>

<style scoped>
/* Reading Progress Bar */
.reading-progress-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 4px;
  z-index: 1001;
  background: transparent;
}

.reading-progress-bar {
  height: 100%;
  background: linear-gradient(to right, #55c57a, #28b485);
  width: 0%;
  transition: width 0.1s ease-out;
}

.news-detail-page {
  padding-top: 8rem; 
  background-color: #fff;
  min-height: 100vh;
}

.container {
  max-width: 120rem;
  margin: 0 auto;
  padding: 0 2rem;
}

.container--narrow {
  max-width: 85rem; /* Sedikit lebih sempit untuk readability maksimal */
}

.breadcrumb {
  padding: 3rem 0;
  font-size: 1.4rem;
  color: #888;
}

.breadcrumb a {
  color: #55c57a;
  text-decoration: none;
  transition: color 0.3s;
}

.breadcrumb a:hover {
  color: #28b485;
}

.breadcrumb .separator {
  margin: 0 1.2rem;
}

.breadcrumb .current {
  color: #444;
  font-weight: 600;
}

.article {
  padding-bottom: 8rem;
}

.article__header {
  margin-bottom: 4rem;
}

.article__meta {
  display: flex;
  gap: 1.5rem;
  margin-bottom: 2rem;
  font-size: 1.4rem;
  font-weight: 700;
}

.article__category {
  color: #55c57a;
  background-color: rgba(85, 197, 122, 0.1);
  padding: 0.2rem 1.2rem;
  border-radius: 50px;
  text-transform: uppercase;
}

.article__date {
  color: #999;
  display: flex;
  align-items: center;
}

.article__title {
  font-size: 4.5rem;
  line-height: 1.1;
  color: #1a1a1a;
  margin-bottom: 2.5rem;
  font-weight: 900;
  letter-spacing: -1px;
}

.article__author {
  font-size: 1.5rem;
  color: #666;
  padding-bottom: 2rem;
  border-bottom: 1px solid #f1f1f1;
}

.article__img-box {
  margin: 4rem 0;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 3rem 6rem rgba(0, 0, 0, 0.1);
}

.article__img {
  width: 100%;
  height: auto;
  display: block;
  transition: transform 0.5s;
}

.article__img:hover {
  transform: scale(1.02);
}

.article__content {
  font-size: 1.9rem;
  line-height: 1.9;
  color: #333;
  font-family: 'Inter', sans-serif;
}

.article__content :deep(p) {
  margin-bottom: 3rem;
}

.article__content :deep(blockquote) {
  border-left: 6px solid #55c57a;
  padding: 3rem 4rem;
  font-style: italic;
  font-size: 2.2rem;
  background-color: #f9fdfb;
  margin: 5rem 0;
  color: #222;
  border-radius: 0 15px 15px 0;
}

.article__footer {
  margin-top: 8rem;
  padding: 4rem 0;
  border-top: 1px solid #eee;
  border-bottom: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 3rem;
}

.tag {
  display: inline-block;
  background-color: #f4f4f4;
  padding: 0.6rem 1.5rem;
  border-radius: 50px;
  font-size: 1.3rem;
  color: #555;
  margin-right: 1rem;
  transition: all 0.3s;
  cursor: pointer;
}

.tag:hover {
  background-color: #55c57a;
  color: #fff;
}

.article__share {
  display: flex;
  align-items: center;
  gap: 2rem;
}

.article__share span {
  font-size: 1.5rem;
  font-weight: 700;
  color: #333;
}

.share-links {
  display: flex;
  gap: 1.5rem;
}

.share-links a {
  font-size: 2.2rem;
  color: #55c57a;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 4.5rem;
  height: 4.5rem;
  border-radius: 50%;
  background-color: #f9fdfb;
}

.share-links a:hover {
  transform: translateY(-3px);
  background-color: #55c57a;
  color: #fff;
  box-shadow: 0 1rem 2rem rgba(85, 197, 122, 0.2);
}

/* Related News Section */
.related-news {
  margin-top: 8rem;
}

.related-news__title {
  font-size: 2.4rem;
  font-weight: 800;
  margin-bottom: 3rem;
  color: #1a1a1a;
  position: relative;
  display: inline-block;
}

.related-news__title::after {
  content: "";
  position: absolute;
  bottom: -10px;
  left: 0;
  width: 50px;
  height: 4px;
  background-color: #55c57a;
  border-radius: 2px;
}

.related-news__grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 3rem;
}

.related-card {
  text-decoration: none;
  color: inherit;
  transition: all 0.3s;
}

.related-card:hover .related-card__judul {
  color: #55c57a;
}

.related-card:hover .related-card__img-box img {
  transform: scale(1.1);
}

.related-card__img-box {
  aspect-ratio: 16/9;
  border-radius: 12px;
  overflow: hidden;
  margin-bottom: 1.5rem;
}

.related-card__img-box img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s;
}

.related-card__category {
  font-size: 1.2rem;
  font-weight: 700;
  color: #55c57a;
  text-transform: uppercase;
  margin-bottom: 0.5rem;
  display: block;
}

.related-card__judul {
  font-size: 1.6rem;
  line-height: 1.4;
  font-weight: 700;
  color: #333;
}

.article__navigation {
  margin-top: 8rem;
  text-align: center;
}

.btn-back {
  display: inline-block;
  padding: 1.5rem 3.5rem;
  background-color: #fff;
  border: 2px solid #55c57a;
  color: #55c57a;
  text-decoration: none;
  font-size: 1.5rem;
  font-weight: 700;
  border-radius: 50px;
  transition: all 0.3s;
}

.btn-back:hover {
  background-color: #55c57a;
  color: #fff;
  box-shadow: 0 1.5rem 3rem rgba(85, 197, 122, 0.2);
}

/* Skeleton */
.skeleton-line {
  height: 2rem;
  background: linear-gradient(90deg, #f0f0f0 25%, #e0e0e0 50%, #f0f0f0 75%);
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
  margin-bottom: 1.5rem;
  border-radius: 4px;
}

.skeleton-line--title { height: 4rem; width: 80%; margin-bottom: 3rem; }
.skeleton-line--meta { height: 1.5rem; width: 30%; margin-bottom: 4rem; }

.skeleton-img {
  height: 45rem;
  background: #eee;
  width: 100%;
  border-radius: 20px;
  margin-bottom: 5rem;
}

@keyframes loading {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

@media (max-width: 900px) {
  .related-news__grid { grid-template-columns: 1fr; }
  .article__title { font-size: 3.5rem; }
}

@media (max-width: 600px) {
  .article__title { font-size: 2.8rem; }
  .article__content { font-size: 1.7rem; }
  .article__footer { flex-direction: column; align-items: flex-start; }
}
</style>
