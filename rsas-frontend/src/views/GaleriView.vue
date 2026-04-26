<script setup>
import { ref, onMounted } from 'vue';
import SectionHeader from '../components/ui/SectionHeader.vue';
import { ASSETS } from '../config/assets';

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

const activeTab = ref('fasilitas');

const galleryData = {
  fasilitas: [
    { id: 1, title: 'Gedung Utama RSAS', img: getImageUrl('nat-1-large.jpg'), category: 'Eksterior' },
    { id: 2, title: 'Ruang Operasi Hybrid', img: getImageUrl('nat-2-large.jpg'), category: 'Medis' },
    { id: 3, title: 'Lobby Utama & Resepsionis', img: getImageUrl('nat-3-large.jpg'), category: 'Interior' },
    { id: 4, title: 'Laboratorium Terpadu', img: getImageUrl('nat-1-large.jpg'), category: 'Medis' },
    { id: 5, title: 'Ruang Rawat Inap VIP', img: getImageUrl('nat-2-large.jpg'), category: 'Interior' },
    { id: 6, title: 'Unit Kateterisasi Jantung', img: getImageUrl('nat-3-large.jpg'), category: 'Medis' }
  ],
  aktivitas: [
    { id: 7, title: 'Seminar Kesehatan Nasional', img: getImageUrl('nat-1-large.jpg'), category: 'Edukasi' },
    { id: 8, title: 'Bakti Sosial & Pengabdian', img: getImageUrl('nat-2-large.jpg'), category: 'Sosial' },
    { id: 9, title: 'Diskusi Klinis Dokter Spesialis', img: getImageUrl('nat-3-large.jpg'), category: 'Edukasi' },
    { id: 10, title: 'Simulasi Penanganan Bencana', img: getImageUrl('nat-1-large.jpg'), category: 'Internal' },
    { id: 11, title: 'Pelatihan Keperawatan Intensif', img: getImageUrl('nat-2-large.jpg'), category: 'Edukasi' },
    { id: 12, title: 'Pemeriksaan Kesehatan Gratis', img: getImageUrl('nat-3-large.jpg'), category: 'Sosial' }
  ]
};

const selectedImage = ref(null);
const openLightbox = (img) => {
  selectedImage.value = img;
};

onMounted(() => {
  window.scrollTo(0, 0);
});
</script>

<template>
  <div class="gallery-page">
    <!-- Hero Section -->
    <header 
      class="gallery-hero" 
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.galeri})` }"
    >
      <div class="container">
        <div class="gallery-hero__content" data-aos="fade-up">
          <div class="gallery-hero__badge">Lensa RSAS</div>
          <h1 class="heading-academic">Galeri <span>Fasilitas & Aktivitas</span></h1>
          <p class="gallery-hero__sub">Dokumentasi visual fasilitas medis modern dan berbagai momen pelayanan tulus kami di RSUD Prof. Dr. H. Aloei Saboe.</p>
        </div>
      </div>
    </header>

    <!-- Tab Section -->
    <section class="section-tabs">
      <div class="container">
        <div class="tab-box" data-aos="fade-up">
          <button 
            class="tab-btn" 
            :class="{ 'tab-btn--active': activeTab === 'fasilitas' }"
            @click="activeTab = 'fasilitas'"
          >
            <i class="fa-solid fa-hospital"></i> Fasilitas & Sarana
          </button>
          <button 
            class="tab-btn" 
            :class="{ 'tab-btn--active': activeTab === 'aktivitas' }"
            @click="activeTab = 'aktivitas'"
          >
            <i class="fa-solid fa-users-viewfinder"></i> Aktivitas & Event
          </button>
        </div>
      </div>
    </section>

    <!-- Grid Section -->
    <section class="section-grid-gallery">
      <div class="container">
        <div class="gallery-grid">
          <div 
            v-for="(item, index) in galleryData[activeTab]" 
            :key="item.id" 
            class="gallery-item"
            data-aos="zoom-in"
            :data-aos-delay="index * 100"
            @click="openLightbox(item)"
          >
            <div class="gallery-item__img-box">
              <img :src="item.img" :alt="item.title" />
              <div class="gallery-item__overlay">
                <span class="gallery-item__cat">{{ item.category }}</span>
                <i class="fa-solid fa-expand"></i>
              </div>
            </div>
            <div class="gallery-item__info">
              <h3>{{ item.title }}</h3>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Lightbox -->
    <div v-if="selectedImage" class="lightbox" @click="selectedImage = null">
      <button class="lightbox__close" @click="selectedImage = null"><i class="fa-solid fa-xmark"></i></button>
      <div class="lightbox__content" @click.stop>
        <img :src="selectedImage.img" :alt="selectedImage.title" />
        <div class="lightbox__caption">
          <h3>{{ selectedImage.title }}</h3>
          <span>{{ selectedImage.category }}</span>
        </div>
      </div>
    </div>

    <!-- Contact CTA -->
    <section class="section-gallery-cta">
      <div class="container">
        <div class="cta-banner">
          <h2>Tertarik Bergabung atau Bekerja Sama?</h2>
          <p>RSAS selalu terbuka untuk kolaborasi pendidikan, penelitian, dan pelayanan kesehatan.</p>
          <router-link to="/pengaduan" class="btn btn--white">Hubungi Kami</router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.container {
  max-width: 130rem;
  margin: 0 auto;
  padding: 0 2rem;
}

.gallery-page {
  padding-top: 8rem;
}

.gallery-hero {
  height: 60vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
  text-align: center;
}

.gallery-hero__badge {
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

.gallery-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  margin: 0 auto;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
  line-height: 1.6;
}

/* Tabs */
.section-tabs {
  margin-top: -4rem;
  position: relative;
  z-index: 10;
}

.tab-box {
  background: #fff;
  padding: 1rem;
  border-radius: 50px;
  box-shadow: 0 1.5rem 4rem rgba(0,0,0,0.1);
  display: flex;
  justify-content: center;
  gap: 1rem;
  max-width: 60rem;
  margin: 0 auto;
}

.tab-btn {
  padding: 1.5rem 4rem;
  border-radius: 50px;
  border: none;
  background: transparent;
  font-size: 1.6rem;
  font-weight: 700;
  color: #666;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.tab-btn i { font-size: 2rem; color: #999; }
.tab-btn:hover { color: #55c57a; }
.tab-btn--active { background: #55c57a; color: #fff; }
.tab-btn--active i { color: #fff; }

/* Grid */
.section-grid-gallery { padding: 10rem 0; background-color: #f9fbf9; }

.gallery-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(35rem, 1fr));
  gap: 4rem;
}

.gallery-item {
  background: #fff;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
  cursor: pointer;
  transition: all 0.3s;
}

.gallery-item:hover { transform: translateY(-1rem); box-shadow: 0 2rem 5rem rgba(0,0,0,0.1); }

.gallery-item__img-box {
  height: 25rem;
  position: relative;
  overflow: hidden;
  background-color: #f0f0f0; /* Fallback color */
}

.gallery-item__img-box img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block; /* Avoid whitespace issue */
  transition: all 0.5s ease-in-out;
}

.gallery-item:hover .gallery-item__img-box img { transform: scale(1.1); }

.gallery-item__overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(85, 197, 122, 0.8);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: all 0.3s;
  color: #fff;
}

.gallery-item:hover .gallery-item__overlay { opacity: 1; }

.gallery-item__cat {
  font-size: 1.2rem;
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 2px;
  margin-bottom: 1.5rem;
}

.gallery-item__overlay i { font-size: 3rem; }

.gallery-item__info { padding: 2.5rem; text-align: center; }
.gallery-item__info h3 { font-size: 1.8rem; font-weight: 700; color: #333; }

/* Lightbox */
.lightbox {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,0.9);
  z-index: 9999;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 5rem;
}

.lightbox__close {
  position: absolute;
  top: 3rem;
  right: 3rem;
  background: transparent;
  border: none;
  color: #fff;
  font-size: 4rem;
  cursor: pointer;
}

.lightbox__content {
  max-width: 110rem;
  width: 100%;
  background: #fff;
  border-radius: 20px;
  overflow: hidden;
  position: relative;
}

.lightbox__content img { width: 100%; max-height: 80vh; object-fit: contain; background: #000; }

.lightbox__caption { padding: 3rem; text-align: center; }
.lightbox__caption h3 { font-size: 2.4rem; color: #333; margin-bottom: 0.5rem; }
.lightbox__caption span { font-size: 1.4rem; color: #55c57a; font-weight: 800; text-transform: uppercase; letter-spacing: 1px; }

/* CTA */
.section-gallery-cta { padding-bottom: 12rem; }
.cta-banner {
  background: #2d3436;
  padding: 8rem;
  border-radius: 40px;
  text-align: center;
  color: #fff;
}
.cta-banner h2 { font-size: 3.5rem; font-weight: 900; margin-bottom: 1.5rem; }
.cta-banner p { font-size: 2rem; opacity: 0.8; margin-bottom: 4rem; }

.btn--white {
  padding: 1.6rem 4rem;
  background: #fff;
  color: #333;
  text-decoration: none;
  font-size: 1.6rem;
  font-weight: 700;
  border-radius: 50px;
  transition: all 0.3s;
}
.btn--white:hover { background: #55c57a; color: #fff; transform: translateY(-3px); }

@media (max-width: 1000px) {
  .gallery-grid { grid-template-columns: 1fr; }
  .tab-box { flex-direction: column; border-radius: 20px; }
  .heading-primary { font-size: 3.5rem; }
  .lightbox { padding: 2rem; }
}
</style>
