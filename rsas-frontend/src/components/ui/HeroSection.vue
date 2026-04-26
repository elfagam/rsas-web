<script setup>
import { onMounted } from 'vue';
import AOS from 'aos';
import { ASSETS } from '../../config/assets';

onMounted(() => {
  AOS.refresh();
});
</script>

<template>
  <header 
    class="hero" 
    :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.5), rgba(0,0,0,0.7)), url(${ASSETS.hero.home})` }"
  >
    <div class="hero__overlay"></div>
    
    <div class="hero__content">
      <h1 class="hero__title">
        <span class="hero__title-main" data-aos="fade-down" data-aos-duration="1200">
          RSUD Prof. Dr. H. <br>
          <span class="hero__title-name">Aloei Saboe</span>
        </span>
        <span class="hero__title-sub" data-aos="fade-up" data-aos-duration="1200" data-aos-delay="300">
          Menjadi Lebih Baik Untuk Semua
        </span>
      </h1>
      
      <p class="hero__description" data-aos="fade-up" data-aos-duration="1200" data-aos-delay="600">
        Rumah sakit rujukan utama di Gorontalo yang berkomitmen memberikan pelayanan kesehatan 
        terbaik dengan fasilitas modern dan tenaga medis profesional.
      </p>

      <div class="hero__buttons" data-aos="zoom-in" data-aos-duration="1000" data-aos-delay="900">
        <a href="#jadwal-dokter" class="btn btn--primary">Cek Jadwal Dokter</a>
        <router-link to="/layanan" class="btn btn--white-outline">Pusat Unggulan</router-link>
      </div>
    </div>

    <!-- Quick Action Box -->
    <div class="hero__actions-box" data-aos="fade-up" data-aos-offset="0" data-aos-delay="1200">
      <div class="action-item">
        <div class="action-item__icon"><i class="fa-solid fa-calendar-check"></i></div>
        <div class="action-item__text">
          <span>Pendaftaran</span>
          <strong>Rawat Jalan Online</strong>
        </div>
        <a href="#" class="action-item__link"></a>
      </div>
      
      <div class="action-item">
        <div class="action-item__icon"><i class="fa-solid fa-user-doctor"></i></div>
        <div class="action-item__text">
          <span>Cari Dokter</span>
          <strong>Spesialis & Jadwal</strong>
        </div>
        <a href="#jadwal-dokter" class="action-item__link"></a>
      </div>

      <div class="action-item action-item--emergency">
        <div class="action-item__icon"><i class="fa-solid fa-phone-volume"></i></div>
        <div class="action-item__text">
          <span>Gawat Darurat</span>
          <strong>IGD 24 Jam: (0435) 821010</strong>
        </div>
        <a href="tel:0435821010" class="action-item__link"></a>
      </div>
    </div>

    <div class="hero__scroll-indicator">
      <span></span>
    </div>
  </header>
</template>

<style scoped>
.hero {
  height: 100vh;
  position: relative;
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  color: #fff;
  width: 100%;
  margin: 0;
  padding: 0;
}

.hero__content {
  max-width: 100rem; /* Diperlebar sedikit agar lebih megah */
  z-index: 10;
}

.hero__title {
  display: flex;
  flex-direction: column;
  margin-bottom: 3.5rem;
}

.hero__title {
  display: flex;
  flex-direction: column;
  margin-bottom: 5rem; /* ~3.5rem * 1.4 */
}

.hero__title-main {
  font-size: 5rem; /* ~3.5rem * 1.4 */
  font-weight: 400;
  letter-spacing: 4px;
  text-transform: uppercase;
  line-height: 1.1;
  text-shadow: 0 1rem 3rem rgba(0, 0, 0, 0.4);
}

.hero__title-name {
  display: block;
  font-size: 8.4rem; /* 6rem * 1.4 */
  font-weight: 950;
  letter-spacing: 2px;
  margin-top: 1.4rem;
}

.hero__title-sub {
  font-size: 3.6rem; /* ~2.6rem * 1.4 */
  font-weight: 700;
  letter-spacing: 6px;
  margin-top: 2.8rem;
  color: #55c57a;
  text-transform: uppercase;
  text-shadow: 0 2px 15px rgba(0, 0, 0, 0.8);
  display: inline-block;
  padding-top: 1.4rem;
  border-top: 2px solid rgba(255, 255, 255, 0.2);
}

.hero__description {
  font-size: 2.4rem; /* ~1.7rem * 1.4 */
  line-height: 1.8;
  margin-bottom: 5.6rem;
  color: rgba(255, 255, 255, 0.9);
  padding: 0 7rem;
}

.hero__buttons {
  display: flex;
  gap: 2rem;
  justify-content: center;
}

.btn {
  padding: 1.5rem 4rem;
  border-radius: 50px;
  font-size: 1.4rem;
  font-weight: 700;
  text-transform: uppercase;
  text-decoration: none;
  transition: all 0.3s;
  cursor: pointer;
}

.btn--primary {
  background-color: #55c57a;
  color: white;
  border: none;
}

.btn--primary:hover {
  background-color: #40945c;
  transform: translateY(-3px);
  box-shadow: 0 1rem 2rem rgba(0, 0, 0, 0.2);
}

.btn--white-outline {
  background: transparent;
  color: white;
  border: 2px solid white;
}

.btn--white-outline:hover {
  background: white;
  color: #333;
  transform: translateY(-3px);
}

.hero__actions-box {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  margin: 0 auto;
  transform: translateY(30%); /* Diturunkan agar tidak menutupi tombol */
  width: 90%;
  max-width: 120rem;
  background: rgba(255, 255, 255, 0.98);
  backdrop-filter: blur(15px);
  display: flex;
  border-radius: 24px;
  box-shadow: 0 2.5rem 6rem rgba(0, 0, 0, 0.12);
  z-index: 100;
  overflow: visible;
  box-sizing: border-box;
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.action-item {
  flex: 1;
  flex-shrink: 1;
  min-width: 0;
  padding: 3.5rem 2rem; /* Lebar dikurangi sedikit untuk fleksibilitas */
  display: flex;
  align-items: center;
  gap: 1.5rem;
  position: relative;
  transition: all 0.4s cubic-bezier(0.165, 0.84, 0.44, 1);
  border-right: 1px solid #f0f0f0;
}

.action-item:last-child { border-right: none; }

.action-item:hover {
  background-color: #fcfdfc;
  transform: translateY(-2px);
}

.action-item__icon {
  width: 6rem;
  height: 6rem;
  background-color: #f0faf3;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2.8rem;
  color: #55c57a;
  transition: all 0.3s;
}

.action-item:hover .action-item__icon {
  background-color: #55c57a;
  color: #fff;
  transform: rotate(-10deg);
}

.action-item__text {
  display: flex;
  flex-direction: column;
}

.action-item__text span {
  font-size: 1.2rem;
  text-transform: uppercase;
  color: #888;
  letter-spacing: 1.5px;
  font-weight: 700;
  margin-bottom: 0.3rem;
}

.action-item__text strong {
  font-size: 1.9rem;
  color: #2d3436;
  font-weight: 900;
}

.action-item__link {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 5;
}

.action-item--emergency {
  background: linear-gradient(135deg, #ff5252 0%, #b33939 100%);
  color: #fff;
  border-right: none;
  animation: pulse-red 2s infinite;
  border-radius: 0 24px 24px 0; /* Mengikuti radius kontainer */
}

.action-item--emergency:hover {
  background: linear-gradient(135deg, #ff4141 0%, #9c2e2e 100%);
}

.action-item--emergency .action-item__icon {
  background-color: rgba(255, 255, 255, 0.2);
  color: #fff;
}

.action-item--emergency .action-item__text span,
.action-item--emergency .action-item__text strong {
  color: #fff;
}

@keyframes pulse-red {
  0% { box-shadow: 0 0 0 0 rgba(255, 82, 82, 0.4); }
  70% { box-shadow: 0 0 0 20px rgba(255, 82, 82, 0); }
  100% { box-shadow: 0 0 0 0 rgba(255, 82, 82, 0); }
}

@media (max-width: 1024px) {
  .action-item { padding: 2.5rem 3rem; }
  .action-item__text strong { font-size: 1.6rem; }
}

.hero__scroll-indicator {
  position: absolute;
  bottom: 3rem;
  left: 50%;
  transform: translateX(-50%);
  width: 3rem;
  height: 5rem;
  border: 2px solid rgba(255, 255, 255, 0.5);
  border-radius: 50px;
}

.hero__scroll-indicator span {
  position: absolute;
  top: 1rem;
  left: 50%;
  transform: translateX(-50%);
  width: 6px;
  height: 6px;
  background-color: #fff;
  border-radius: 50%;
  animation: scrollDown 2s infinite;
}

@keyframes scrollDown {
  0% { transform: translate(-50%, 0); opacity: 0; }
  30% { opacity: 1; }
  100% { transform: translate(-50%, 20px); opacity: 0; }
}

@media (max-width: 768px) {
  .hero {
    height: auto;
    min-height: 100vh;
    padding: 15rem 0 10rem; /* Memberikan ruang untuk Navbar dan Footer Box */
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
  }

  .hero__content {
    margin-bottom: 6rem; /* Jarak aman sebelum kotak aksi */
  }

  .hero__title-main { font-size: 3rem; }
  .hero__title-name { font-size: 5rem; }
  .hero__title-sub { font-size: 1.6rem; letter-spacing: 3px; }
  .hero__description { font-size: 1.5rem; padding: 0 2rem; margin-bottom: 4rem; }
  .hero__buttons { flex-direction: column; gap: 1.5rem; width: 80%; margin: 0 auto; }
  
  .hero__actions-box {
    position: relative; /* Berubah dari absolute ke relative agar tidak menumpuk */
    transform: none; /* Menghapus geseran vertikal */
    margin: 0 auto;
    flex-direction: column;
    width: 92%;
    background: transparent;
    box-shadow: none;
    border: none;
    gap: 1.5rem;
  }
  
  .action-item {
    background: rgba(255, 255, 255, 0.98);
    border-radius: 20px;
    border: 1px solid rgba(0,0,0,0.05);
    padding: 2.5rem 2rem;
    box-shadow: 0 1rem 3rem rgba(0,0,0,0.1);
  }

  .action-item--emergency {
    background: linear-gradient(135deg, #ff5252 0%, #b33939 100%) !important;
    color: #fff;
    border-radius: 20px;
    margin-top: 0.5rem;
    border: none;
  }
  
  .action-item__icon {
    width: 5rem;
    height: 5rem;
    font-size: 2.2rem;
  }

  .action-item__text strong {
    font-size: 1.7rem;
  }

  .hero__scroll-indicator {
    display: none; /* Sembunyikan indikator scroll di mobile agar tidak mengganggu flow */
  }
}
</style>
