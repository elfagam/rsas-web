<script setup>
import { ref, computed, onMounted } from "vue";
import NewsCard from "../components/ui/NewsCard.vue";
import { ASSETS } from '../config/assets';

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

const isLoading = ref(true);
const searchQuery = ref('');
const activeCategory = ref('');

// Simulasi Data
const daftarBerita = ref([
  {
    id: 1,
    kategori: "PENGUMUMAN",
    tanggal: "25 April 2026",
    judul: "Perkuat Mutu Pelayanan, Kemenkes RI Verifikasi Standarisasi RSUD",
    ringkasan: "GORONTALO - Kementerian Kesehatan Republik Indonesia melakukan kunjungan kerja dalam rangka verifikasi lapangan terhadap standar kualitas pelayanan...",
    gambar: getImageUrl("nat-1.jpg"),
    tautan: "/berita/1",
  },
  {
    id: 2,
    kategori: "LAYANAN",
    tanggal: "24 April 2026",
    judul: "Pelayanan Tetap Optimal, RSUD Prof. Dr. H. Aloei Saboe Pastikan Kesiapan SDM",
    ringkasan: "GORONTALO - RSUD memastikan seluruh layanan kesehatan tetap berjalan optimal dengan melakukan rotasi jadwal tenaga medis selama periode libur panjang...",
    gambar: getImageUrl("nat-2.jpg"),
    tautan: "/berita/2",
  },
  {
    id: 3,
    kategori: "PENGUMUMAN",
    tanggal: "22 April 2026",
    judul: "Jadwal Operasional Klinik Spesialis Selama Bulan Ramadhan",
    ringkasan: "Diberitahukan kepada seluruh pasien dan pengunjung, terdapat penyesuaian jadwal jam operasional untuk beberapa poli klinik spesialis...",
    gambar: getImageUrl("nat-3.jpg"),
    tautan: "/berita/3",
  },
]);

const filteredBerita = computed(() => {
  return daftarBerita.value.filter(berita => {
    const matchesSearch = berita.judul.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
                         berita.ringkasan.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchesCategory = !activeCategory.value || berita.kategori.toLowerCase() === activeCategory.value.toLowerCase();
    return matchesSearch && matchesCategory;
  });
});

onMounted(() => {
  // Simulasi fetching data
  setTimeout(() => {
    isLoading.value = false;
  }, 1500);
});
</script>

<template>
  <div class="berita-view">
    <!-- Hero Section -->
    <header 
      class="berita-hero"
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.berita})` }"
    >
      <div class="container">
        <div class="berita-hero__content" data-aos="fade-up">
          <div class="berita-hero__badge">Pusat Media & Edukasi</div>
          <h1 class="heading-academic">Berita & <span>Kabar Kesehatan</span></h1>
          <p class="berita-hero__sub">Tetap terinformasi dengan berita terbaru, pengumuman resmi, dan tips kesehatan terpercaya dari para ahli kami.</p>
        </div>
      </div>
    </header>

    <div class="filter-wrapper">
      <div class="filter-bar" data-aos="fade-up">
        <span class="search-icon">
          <i class="fa-solid fa-magnifying-glass"></i>
        </span>

        <input
          v-model="searchQuery"
          type="text"
          placeholder="Cari berita atau pengumuman..."
          class="search-input"
        />

        <div class="divider"></div>

        <select v-model="activeCategory" class="category-select">
          <option value="">Semua Kategori</option>
          <option value="pengumuman">Pengumuman</option>
          <option value="layanan">Layanan</option>
        </select>

        <button class="btn-search">
          Cari
        </button>
      </div>
    </div>

    <section class="section-berita">
      <div class="container">
        <div class="news-grid">
          <template v-if="isLoading">
            <div class="skeleton-news" v-for="n in 3" :key="'skeleton-' + n">
              <div class="skeleton-img"></div>
              <div class="skeleton-content">
                <div class="skeleton-line sm"></div>
                <div class="skeleton-line lg"></div>
                <div class="skeleton-line md"></div>
              </div>
            </div>
          </template>

          <template v-else>
            <NewsCard
              v-for="berita in filteredBerita"
              :key="berita.id"
              v-bind="berita"
              data-aos="fade-up"
            />
          </template>
        </div>

        <!-- Empty State -->
        <div v-if="!isLoading && filteredBerita.length === 0" class="empty-state" data-aos="fade-up">
          <i class="fa-solid fa-newspaper"></i>
          <h3>Berita tidak ditemukan</h3>
          <p>Coba gunakan kata kunci lain atau pilih kategori yang berbeda.</p>
          <button @click="searchQuery = ''; activeCategory = ''" class="btn-reset">Lihat Semua Berita</button>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.berita-view {
  min-height: 100vh;
}

.container {
  max-width: 120rem;
  margin: 0 auto;
  padding: 0 2rem;
}

/* Hero */
.berita-hero {
  height: 60vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
  text-align: center;
}

.berita-hero__badge {
  display: inline-block;
  padding: 0.8rem 2.5rem;
  background: rgba(85, 197, 122, 0.2);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 50px;
  font-size: 1.3rem;
  font-weight: 700;
  color: #fff;
  margin-bottom: 3rem;
  text-transform: uppercase;
  letter-spacing: 1.5px;
}

.heading-academic {
  font-family: 'Playfair Display', serif;
  font-size: 6.5rem;
  font-weight: 900;
  line-height: 1.1;
  margin-bottom: 2.5rem;
}

.heading-academic span {
  font-style: italic;
  font-weight: 400;
  display: block;
  font-size: 5rem;
  margin-top: 1rem;
}

.berita-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  margin: 0 auto;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
  line-height: 1.6;
}

.filter-wrapper {
  margin-top: -6rem; /* Lebih dalam ke arah Hero untuk efek layering yang kuat */
  position: relative;
  z-index: 100;
  display: flex;
  justify-content: center;
  padding: 0 2rem;
}

.filter-bar {
  display: flex;
  gap: 2.5rem; /* Menambah jarak antar elemen input */
  background: #fff;
  padding: 1.8rem 3.5rem; /* Padding lebih luas untuk kesan lega */
  border-radius: 100px; /* Lebih membulat sempurna */
  box-shadow: 0 3rem 6rem rgba(0, 0, 0, 0.12);
  align-items: center;
  width: 100%;
  max-width: 100rem; /* Sedikit lebih lebar */
}

.search-icon {
  color: #55c57a; /* Warna hijau brand agar lebih menonjol */
  margin-left: 0.5rem;
  font-size: 1.8rem;
}

.search-input {
  flex: 1;
  border: none;
  outline: none;
  font-size: 1.7rem; /* Sedikit lebih besar untuk keterbacaan */
  color: #2d3436;
  background: transparent;
  min-width: 15rem;
  font-weight: 500;
}

.divider {
  width: 1px;
  height: 4rem; /* Lebih tinggi agar seimbang dengan padding baru */
  background-color: #f1f1f1;
}

.category-select {
  border: none;
  outline: none;
  font-size: 1.6rem;
  color: #636e72;
  background: transparent;
  cursor: pointer;
  padding: 0 1.5rem;
  font-weight: 600;
}

.btn-search {
  padding: 1.5rem 4.5rem; /* Padding tombol lebih mantap */
  border-radius: 50px;
  border: none;
  background: linear-gradient(135deg, #55c57a 0%, #33ab5a 100%);
  color: white;
  font-weight: 700;
  font-size: 1.5rem;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  box-shadow: 0 1rem 2rem rgba(85, 197, 122, 0.3);
}

.btn-search:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 1.5rem 3rem rgba(85, 197, 122, 0.4);
}

.section-berita {
  padding: 12rem 0 10rem; /* Jarak atas bawah lebih lega */
  max-width: 120rem;
  margin: 0 auto;
}

.news-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(32rem, 1fr));
  gap: 4rem;
}

/* Skeleton Loading styles */
.skeleton-news {
  background: #fff;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 1rem 3rem rgba(0, 0, 0, 0.05);
}

.skeleton-img {
  height: 22rem;
  background: linear-gradient(90deg, #f0f0f0 25%, #e0e0e0 50%, #f0f0f0 75%);
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
}

.skeleton-content {
  padding: 2.5rem;
}

.skeleton-line {
  height: 1.5rem;
  background: linear-gradient(90deg, #f0f0f0 25%, #e0e0e0 50%, #f0f0f0 75%);
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
  margin-bottom: 1.5rem;
  border-radius: 4px;
}

.skeleton-line.sm { width: 30%; }
.skeleton-line.md { width: 60%; }
.skeleton-line.lg { width: 90%; }

@keyframes loading {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 10rem 0;
  color: #999;
}

.empty-state i {
  font-size: 8rem;
  margin-bottom: 2rem;
  opacity: 0.3;
}

.empty-state h3 {
  font-size: 2.4rem;
  color: #333;
  margin-bottom: 1rem;
}

.btn-reset {
  margin-top: 3rem;
  padding: 1.2rem 3rem;
  background: #55c57a;
  color: #fff;
  border: none;
  border-radius: 50px;
  font-weight: 700;
  cursor: pointer;
}

/* Responsivitas */
@media (max-width: 768px) {
  .section-berita {
    padding: 8rem 2rem 5rem;
  }
  
  .filter-bar {
    flex-wrap: wrap;
    border-radius: 20px;
    padding: 2rem;
  }
  
  .divider {
    display: none;
  }
  
  .btn-search {
    width: 100%;
  }

  .heading-academic {
    font-size: 4rem;
  }
}
</style>
