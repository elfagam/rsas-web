<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const isMenuOpen = ref(false);
const isScrolled = ref(false);

// Daftar halaman yang memiliki Hero gelap dan mendukung Navbar Transparan
const heroPages = ['/', '/berita', '/profil', '/layanan', '/dokter', '/pendidikan', '/zona-integritas', '/faq', '/pengaduan', '/galeri'];
const isHeroPage = computed(() => heroPages.includes(route.path));

// Navbar akan transparan HANYA jika di halaman Hero DAN belum di-scroll
const shouldShowBg = computed(() => !isHeroPage.value || isScrolled.value);

const handleScroll = () => {
  isScrolled.value = window.scrollY > 50;
};

// Fungsi helper resolusi gambar
const getImageUrl = (name) => {
  return new URL(`../../assets/img/${name}`, import.meta.url).href;
};

onMounted(() => {
  window.addEventListener('scroll', handleScroll);
});

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll);
});

const menuItems = [
  { name: 'Home', link: '/' },
  { name: 'Berita', link: '/berita' },
  { name: 'Profil', link: '/profil' },
  { name: 'Layanan', link: '/layanan' },
  { name: 'Jadwal Dokter', link: '/dokter' },
  { name: 'Pendidikan', link: '/pendidikan' },
  { name: 'Integritas', link: '/zona-integritas' },
  { name: 'FAQ', link: '/faq' },
];
</script>

<template>
  <nav class="navbar" :class="{ 'navbar--scrolled': shouldShowBg }">
    <div class="navbar__container">
      <router-link to="/" class="navbar__logo-box">
        <!-- Icon Medis Pengganti Logo Lama -->
        <div class="navbar__medical-icon">
          <i class="fa-solid fa-square-h"></i>
        </div>
        <div class="navbar__brand">
          <span class="navbar__brand-main" :class="{ 'navbar__brand-main--white': !shouldShowBg }">RSUD Prof. Dr. H. Aloei Saboe</span>
          <span class="navbar__brand-sub">Kota Gorontalo</span>
        </div>
      </router-link>

      <!-- Desktop Menu -->
      <ul class="navbar__list">
        <li v-for="item in menuItems" :key="item.name" class="navbar__item">
          <router-link :to="item.link" class="navbar__link">{{ item.name }}</router-link>
        </li>
      </ul>

      <!-- Mobile Button -->
      <div class="navbar__mobile-toggle" @click="isMenuOpen = !isMenuOpen">
        <div class="navbar__icon" :class="{ 'navbar__icon--open': isMenuOpen }">&nbsp;</div>
      </div>
    </div>

    <!-- Mobile Menu Overlay -->
    <transition name="slide-fade">
      <div v-if="isMenuOpen" class="navbar__mobile-overlay">
        <div class="navbar__mobile-header">
          <div class="navbar__logo-box">
             <div class="navbar__medical-icon"><i class="fa-solid fa-square-h"></i></div>
             <div class="navbar__brand">
               <span class="navbar__brand-main">RSUD Aloei Saboe</span>
               <span class="navbar__brand-sub">Kota Gorontalo</span>
             </div>
          </div>
          <div class="navbar__mobile-close" @click="isMenuOpen = false">
            <i class="fa-solid fa-xmark"></i>
          </div>
        </div>

        <div class="navbar__mobile-content">
          <!-- Quick Action Buttons -->
          <div class="navbar__mobile-actions">
            <router-link to="/pengaduan" class="mobile-action-btn mobile-action-btn--emergency" @click="isMenuOpen = false">
              <i class="fa-solid fa-phone-volume"></i>
              <span>IGD 24 JAM</span>
            </router-link>
            <a href="#" class="mobile-action-btn mobile-action-btn--primary" @click="isMenuOpen = false">
              <i class="fa-solid fa-calendar-check"></i>
              <span>DAFTAR ONLINE</span>
            </a>
          </div>

          <!-- Categorized Links -->
          <div class="navbar__mobile-groups">
            <div class="navbar__mobile-group">
              <h4 class="group-title">Layanan Pasien</h4>
              <ul class="group-list">
                <li><router-link to="/layanan" @click="isMenuOpen = false">Layanan Unggulan</router-link></li>
                <li><router-link to="/dokter" @click="isMenuOpen = false">Cari Dokter & Jadwal</router-link></li>
                <li><router-link to="/faq" @click="isMenuOpen = false">Pusat Bantuan (FAQ)</router-link></li>
                <li><router-link to="/galeri" @click="isMenuOpen = false">Galeri Fasilitas</router-link></li>
              </ul>
            </div>

            <div class="navbar__mobile-group">
              <h4 class="group-title">Rumah Sakit</h4>
              <ul class="group-list">
                <li><router-link to="/profil" @click="isMenuOpen = false">Profil & Sejarah</router-link></li>
                <li><router-link to="/berita" @click="isMenuOpen = false">Berita & Edukasi</router-link></li>
                <li><router-link to="/pendidikan" @click="isMenuOpen = false">Pendidikan & Riset</router-link></li>
                <li><router-link to="/zona-integritas" @click="isMenuOpen = false">Zona Integritas</router-link></li>
              </ul>
            </div>
          </div>
        </div>

        <div class="navbar__mobile-footer">
          <div class="contact-info">
            <p><i class="fa-solid fa-location-dot"></i> Jl. Prof. Dr. H. Aloei Saboe, Gorontalo</p>
            <p><i class="fa-solid fa-phone"></i> (0435) 821010</p>
          </div>
          <div class="social-links">
             <i class="fa-brands fa-facebook"></i>
             <i class="fa-brands fa-instagram"></i>
             <i class="fa-brands fa-youtube"></i>
          </div>
        </div>
      </div>
    </transition>
  </nav>
</template>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000;
  padding: 1.5rem 2rem;
  transition: all 0.4s ease;
  background: transparent;
}

.navbar--scrolled {
  padding: 1rem 2rem;
  background: rgba(255, 255, 255, 0.98);
  backdrop-filter: blur(10px);
  box-shadow: 0 1rem 3rem rgba(0, 0, 0, 0.08);
}

.navbar__container {
  max-width: 144rem;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.navbar__logo-box {
  display: flex;
  align-items: center;
  gap: 1rem;
  text-decoration: none;
}

.navbar__medical-icon {
  font-size: 4.2rem; /* 3rem * 1.4 */
  color: #55c57a;
  display: flex;
  align-items: center;
}

.navbar__brand {
  display: flex;
  flex-direction: column;
}

.navbar__brand-main {
  font-size: 1.7rem; /* 1.2rem * 1.4 */
  font-weight: 800;
  color: #333;
  line-height: 1.1;
  transition: color 0.3s;
}

.navbar__brand-sub {
  font-size: 1.3rem; /* 0.9rem * 1.4 */
  color: #55c57a;
  letter-spacing: 0.5px;
  text-transform: uppercase;
  font-weight: 700;
}

.navbar__brand-main--white {
  color: #fff;
}

.navbar__list {
  display: flex;
  list-style: none;
  gap: 1.1rem; /* 0.8rem * 1.4 */
}

@media (max-width: 1350px) {
  .navbar__list {
    display: none;
  }
}

.navbar__link {
  text-decoration: none;
  color: #fff;
  font-size: 1.3rem; /* 0.9rem * 1.4 */
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  transition: all 0.3s;
  padding: 1.1rem 0;
}

.navbar--scrolled .navbar__link {
  color: #444;
}

.navbar__link:hover,
.navbar__link.router-link-active {
  color: #55c57a;
}

.navbar--scrolled .navbar__link.router-link-active {
  color: #55c57a;
  border-bottom: 2px solid #55c57a;
}

/* Mobile styles */
.navbar__mobile-toggle {
  display: none;
  cursor: pointer;
  z-index: 2000;
}

@media (max-width: 1200px) {
  .navbar__mobile-toggle {
    display: block;
  }
}

.navbar__icon {
  position: relative;
  width: 3rem;
  height: 2px;
  background-color: #fff;
  transition: all 0.3s;
}

.navbar--scrolled .navbar__icon {
  background-color: #333;
}

.navbar__icon::before,
.navbar__icon::after {
  content: "";
  position: absolute;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: inherit;
  transition: all 0.3s;
}

.navbar__icon::before { top: -0.8rem; }
.navbar__icon::after { top: 0.8rem; }

.navbar__icon--open {
  background-color: transparent !important;
}

.navbar__icon--open::before {
  top: 0;
  transform: rotate(135deg);
}

.navbar__icon--open::after {
  top: 0;
  transform: rotate(-135deg);
}

/* New Mobile Overlay Styles */
.navbar__mobile-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  background-color: #fff;
  z-index: 3000;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
}

.navbar__mobile-header {
  padding: 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #eee;
}

.navbar__mobile-close {
  font-size: 3rem;
  color: #333;
  cursor: pointer;
}

.navbar__mobile-content {
  padding: 3rem 2rem;
  flex-grow: 1;
}

.navbar__mobile-actions {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
  margin-bottom: 4rem;
}

.mobile-action-btn {
  padding: 2rem 1.5rem;
  border-radius: 15px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  text-decoration: none;
  font-weight: 800;
  font-size: 1.2rem;
  text-align: center;
}

.mobile-action-btn i {
  font-size: 2.2rem;
}

.mobile-action-btn--emergency {
  background-color: #fff5f5;
  color: #ff5252;
  border: 2px solid #ff5252;
}

.mobile-action-btn--primary {
  background-color: #f0faf3;
  color: #55c57a;
  border: 2px solid #55c57a;
}

.navbar__mobile-groups {
  display: flex;
  flex-direction: column;
  gap: 4rem;
}

.group-title {
  font-size: 1.3rem;
  text-transform: uppercase;
  color: #999;
  letter-spacing: 2px;
  margin-bottom: 2rem;
  font-weight: 800;
}

.group-list {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.group-list a {
  text-decoration: none;
  font-size: 2.2rem;
  color: #333;
  font-weight: 400;
  display: block;
}

.group-list a.router-link-active {
  color: #55c57a;
  font-weight: 800;
}

.navbar__mobile-footer {
  padding: 4rem 2rem;
  background-color: #f9f9f9;
  border-top: 1px solid #eee;
}

.contact-info {
  margin-bottom: 2.5rem;
}


.contact-info p {
  font-size: 1.5rem;
  color: #666;
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  gap: 1rem;
}

.contact-info i { color: #55c57a; }

.social-links {
  display: flex;
  gap: 2rem;
  font-size: 2.2rem;
  color: #55c57a;
}

/* Transitions */
.slide-fade-enter-active, .slide-fade-leave-active {
  transition: all 0.4s ease;
}
.slide-fade-enter-from, .slide-fade-leave-to {
  transform: translateY(-20px);
  opacity: 0;
}
</style>
