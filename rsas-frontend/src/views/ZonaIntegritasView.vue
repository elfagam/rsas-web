<script setup>
import { onMounted } from 'vue';
import SectionHeader from '../components/ui/SectionHeader.vue';
import { ASSETS } from '../config/assets';

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

const areas = [
  { title: 'Manajemen Perubahan', icon: 'fa-solid fa-users-gear', desc: 'Membangun komitmen pimpinan dan integritas budaya kerja seluruh pegawai.' },
  { title: 'Penataan Tata Laksana', icon: 'fa-solid fa-diagram-project', desc: 'Meningkatkan efisiensi sistem dan prosedur kerja melalui digitalisasi.' },
  { title: 'Manajemen SDM', icon: 'fa-solid fa-user-check', desc: 'Transformasi pengelolaan SDM yang transparan, kompetitif, dan akuntabel.' },
  { title: 'Penguatan Akuntabilitas', icon: 'fa-solid fa-chart-line', desc: 'Meningkatkan kapasitas dan tanggung jawab kinerja rumah sakit.' },
  { title: 'Penguatan Pengawasan', icon: 'fa-solid fa-eye', desc: 'Mewujudkan penyelenggaraan RSAS yang bersih dan bebas dari KKN.' },
  { title: 'Kualitas Layanan Publik', icon: 'fa-solid fa-hand-holding-medical', desc: 'Inovasi pelayanan yang cepat, mudah, dan berorientasi pada kepuasan masyarakat.' }
];

const satisfactionStats = [
  { label: 'Indeks Kepuasan Masyarakat', value: '88.45', suffix: '/100', color: '#55c57a' },
  { label: 'Tingkat Kepatuhan SOP', value: '96', suffix: '%', color: '#3498db' },
  { label: 'Penanganan Keluhan', value: '< 24', suffix: ' Jam', color: '#f39c12' }
];

onMounted(() => {
  window.scrollTo(0, 0);
  if (window.AOS) {
    window.AOS.refresh();
  }
});
</script>

<template>
  <div class="zi-page">
    <!-- Hero Section -->
    <header 
      class="zi-hero" 
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.zonaIntegritas})` }"
    >
      <div class="container">
        <div class="zi-hero__content" data-aos="fade-up">
          <div class="zi-hero__badge">Pemerintah Kota Gorontalo</div>
          <h1 class="heading-primary">
            <span class="heading-primary__main">Zona Integritas</span>
            <span class="heading-primary__sub">Bersih & Melayani</span>
          </h1>
          <p class="zi-hero__sub">Mewujudkan RSUD Prof. Dr. H. Aloei Saboe sebagai Wilayah Bebas dari Korupsi (WBK) dan Wilayah Birokrasi Bersih dan Melayani (WBBM).</p>
          <div class="zi-hero__logos">
            <div class="zi-logo">WBK</div>
            <div class="zi-logo">WBBM</div>
          </div>
        </div>
      </div>
    </header>

    <!-- 6 Areas Section -->
    <section class="section-areas">
      <div class="container">
        <SectionHeader 
          title="6 Area Perubahan" 
          subtitle="Transformasi tata kelola rumah sakit menuju birokrasi yang bersih dan melayani."
          align="center"
        />

        <div class="areas-grid">
          <div v-for="(area, index) in areas" :key="area.title" class="area-card" data-aos="fade-up" :data-aos-delay="index * 100">
            <div class="area-card__icon"><i :class="area.icon"></i></div>
            <h3 class="area-card__title">{{ area.title }}</h3>
            <p class="area-card__desc">{{ area.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Satisfaction Stats -->
    <section class="section-zi-stats">
      <div class="container">
        <div class="zi-stats-box">
          <div class="zi-stats-header">
            <h2 class="heading-secondary">Transparansi <span>Kinerja & Kepuasan</span></h2>
            <p class="heading-tertiary">Data capaian indikator pelayanan publik RSAS periode 2024</p>
          </div>
          <div class="zi-stats-grid">
            <div v-for="stat in satisfactionStats" :key="stat.label" class="zi-stat-item">
              <span class="zi-stat-value" :style="{ color: stat.color }">{{ stat.value }}<small>{{ stat.suffix }}</small></span>
              <span class="zi-stat-label">{{ stat.label }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Anti-Corruption Commitment -->
    <section class="section-anti-corruption">
      <div class="container">
        <div class="ac-banner" data-aos="zoom-in">
          <div class="ac-banner__icon"><i class="fa-solid fa-handshake-slash"></i></div>
          <div class="ac-banner__text">
            <h2 class="heading-secondary">Katakan TIDAK <span>pada Gratifikasi!</span></h2>
            <p class="heading-tertiary">Seluruh layanan di RSUD Prof. Dr. H. Aloei Saboe bebas dari pungutan liar. Jangan memberikan tips, hadiah, atau uang dalam bentuk apapun kepada petugas kami.</p>
          </div>
          <a href="#" class="btn btn--white">Lapor Jika Ada Pungli</a>
        </div>
      </div>
    </section>

    <!-- Whistleblowing System -->
    <section class="section-wbs">
      <div class="container">
        <div class="wbs-layout">
          <div class="wbs-content" data-aos="fade-right">
            <h2 class="heading-secondary">Whistleblowing <span>System (WBS)</span></h2>
            <p class="heading-tertiary">Mekanisme penyampaian pengaduan dugaan tindak pidana korupsi atau pelanggaran kode etik yang terjadi di lingkungan RSUD Prof. Dr. H. Aloei Saboe. Kami menjamin kerahasiaan identitas Anda.</p>
            <div class="wbs-steps">
              <div class="wbs-step">
                <span class="wbs-step__num">1</span>
                <p>Identifikasi pelanggaran yang terjadi.</p>
              </div>
              <div class="wbs-step">
                <span class="wbs-step__num">2</span>
                <p>Siapkan bukti pendukung (Foto/Dokumen).</p>
              </div>
              <div class="wbs-step">
                <span class="wbs-step__num">3</span>
                <p>Lapor melalui portal WBS resmi kami.</p>
              </div>
            </div>
            <button class="btn btn--green">Buka Portal Pengaduan (WBS)</button>
          </div>
          <div class="wbs-image" data-aos="fade-left">
            <img :src="getImageUrl('nat-1-large.jpg')" alt="Legal and Transparency" />
          </div>
        </div>
      </div>
    </section>

    <!-- Transparency Documents Section -->
    <section class="section-zi-docs">
      <div class="container">
        <SectionHeader 
          title="Dokumen Transparansi" 
          subtitle="Akses publik terhadap dokumen pakta integritas dan laporan zona integritas RSAS."
          align="center"
        />

        <div class="zi-docs-grid">
          <div class="zi-doc-card" data-aos="fade-up">
            <div class="zi-doc-card__info">
              <i class="fa-solid fa-file-contract"></i>
              <div>
                <h4>Pakta Integritas Pegawai</h4>
                <span>Komitmen kolektif seluruh staf RSAS</span>
              </div>
            </div>
            <a href="#" class="btn-icon-download"><i class="fa-solid fa-download"></i></a>
          </div>

          <div class="zi-doc-card" data-aos="fade-up" data-aos-delay="100">
            <div class="zi-doc-card__info">
              <i class="fa-solid fa-file-invoice"></i>
              <div>
                <h4>Laporan Kinerja (LKjIP)</h4>
                <span>Laporan akuntabilitas instansi pemerintah</span>
              </div>
            </div>
            <a href="#" class="btn-icon-download"><i class="fa-solid fa-download"></i></a>
          </div>

          <div class="zi-doc-card" data-aos="fade-up" data-aos-delay="200">
            <div class="zi-doc-card__info">
              <i class="fa-solid fa-hand-holding-heart"></i>
              <div>
                <h4>Laporan IKM 2024</h4>
                <span>Hasil survei kepuasan masyarakat tahunan</span>
              </div>
            </div>
            <a href="#" class="btn-icon-download"><i class="fa-solid fa-download"></i></a>
          </div>
        </div>

        <div class="ikm-cta" data-aos="zoom-in">
          <div class="ikm-cta__content">
            <h3>Bantu Kami Meningkatkan Layanan</h3>
            <p>Pendapat Anda sangat berharga. Isi Survei Kepuasan Masyarakat secara online untuk membantu kami mengevaluasi kualitas pelayanan kami.</p>
          </div>
          <button class="btn btn--white">Isi Survei IKM Online <i class="fa-solid fa-arrow-right-long u-margin-left-small"></i></button>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.container {
  max-width: 120rem;
  margin: 0 auto;
  padding: 0 2rem;
}

.zi-page {
  /* Padding top dihapus agar Hero menyentuh bagian paling atas layar */
}

/* Hero */
.zi-hero {
  height: 65vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
  text-align: center;
}

.zi-hero__content {
  max-width: 90rem;
  margin: 0 auto;
}

.zi-hero__badge {
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
  font-size: 5.5rem;
  margin-top: 1rem;
}

.zi-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  line-height: 1.6;
  margin-bottom: 5rem;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
}

.zi-hero__logos {
  display: flex;
  justify-content: center;
  gap: 3rem;
}

.zi-logo {
  width: 10rem;
  height: 10rem;
  border: 3px solid #55c57a;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.8rem;
  font-weight: 900;
  color: #55c57a;
  background: rgba(255,255,255,0.05);
}

/* 6 Areas */
.section-areas {
  padding: 12rem 0;
  background-color: #f9fbf9;
}

.areas-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 4rem;
  margin-top: 6rem;
}

.area-card {
  background: #fff;
  padding: 4rem;
  border-radius: 25px;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
  transition: all 0.3s;
  border: 1px solid transparent;
}

.area-card:hover {
  transform: translateY(-1rem);
  border-color: #55c57a;
  box-shadow: 0 2rem 5rem rgba(0,0,0,0.1);
}

.area-card__icon {
  font-size: 4rem;
  color: #55c57a;
  margin-bottom: 2.5rem;
}

.area-card__title {
  font-size: 2rem;
  font-weight: 800;
  color: #333;
  margin-bottom: 1.5rem;
}

.area-card__desc {
  font-size: 1.5rem;
  color: #777;
  line-height: 1.7;
}

/* Stats */
.section-zi-stats {
  padding: 10rem 0;
}

.zi-stats-box {
  background: #fff;
  padding: 6rem;
  border-radius: 30px;
  box-shadow: 0 3rem 6rem rgba(0,0,0,0.05);
  text-align: center;
}

.zi-stats-header {
  margin-bottom: 6rem;
}

.zi-stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 4rem;
}

.zi-stat-value {
  display: block;
  font-size: 5rem;
  font-weight: 900;
  margin-bottom: 1rem;
}

.zi-stat-value small {
  font-size: 2rem;
  color: #999;
}

.zi-stat-label {
  font-size: 1.6rem;
  font-weight: 700;
  color: #555;
  text-transform: uppercase;
  letter-spacing: 1px;
}

/* Anti-Corruption Banner */
.section-anti-corruption {
  padding-bottom: 12rem;
}

.ac-banner {
  background: linear-gradient(to right, #eb4d4b, #c0392b);
  padding: 6rem;
  border-radius: 30px;
  display: flex;
  align-items: center;
  gap: 4rem;
  color: #fff;
}

.ac-banner__icon {
  font-size: 6rem;
  opacity: 0.8;
}

.ac-banner__text {
  flex: 1;
}

.ac-banner__text h2 {
  font-size: 3.2rem;
  margin-bottom: 1rem;
}

.ac-banner__text p {
  font-size: 1.8rem;
  opacity: 0.9;
  line-height: 1.6;
}

/* WBS Section */
.section-wbs {
  padding-bottom: 12rem;
}

.wbs-layout {
  display: flex;
  gap: 8rem;
  align-items: center;
}

.wbs-content {
  flex: 1;
}

.wbs-content p {
  font-size: 1.8rem;
  color: #666;
  line-height: 1.8;
  margin: 2rem 0 4rem;
}

.wbs-steps {
  display: grid;
  gap: 2rem;
  margin-bottom: 4rem;
}

.wbs-step {
  display: flex;
  align-items: center;
  gap: 2rem;
  padding: 2rem;
  background: #f9fbf9;
  border-radius: 15px;
}

.wbs-step__num {
  width: 4rem;
  height: 4rem;
  background: #55c57a;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.8rem;
  font-weight: 800;
  flex-shrink: 0;
}

.wbs-step p {
  margin: 0;
  font-size: 1.6rem;
  font-weight: 700;
  color: #333;
}

.wbs-image {
  flex: 1;
  border-radius: 30px;
  overflow: hidden;
  box-shadow: 0 3rem 6rem rgba(0,0,0,0.1);
}

.wbs-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Transparency Docs */
.section-zi-docs {
  padding: 10rem 0;
  background-color: #f9fbf9;
}

.zi-docs-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(35rem, 1fr));
  gap: 3rem;
  margin-top: 5rem;
}

.zi-doc-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 3rem;
  background: #fff;
  border-radius: 20px;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
  transition: all 0.3s;
}

.zi-doc-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 1.5rem 4rem rgba(0,0,0,0.08);
}

.zi-doc-card__info {
  display: flex;
  align-items: center;
  gap: 2rem;
}

.zi-doc-card__info i {
  font-size: 3rem;
  color: #55c57a;
}

.zi-doc-card__info h4 {
  font-size: 1.8rem;
  font-weight: 800;
  color: #333;
}

.zi-doc-card__info span {
  display: block;
  font-size: 1.3rem;
  color: #888;
  margin-top: 0.3rem;
}

.btn-icon-download {
  width: 5rem;
  height: 5rem;
  background: #f9fbf9;
  color: #55c57a;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.8rem;
  text-decoration: none;
  transition: all 0.3s;
}

.btn-icon-download:hover {
  background: #55c57a;
  color: #fff;
}

/* IKM CTA */
.ikm-cta {
  margin-top: 8rem;
  background: linear-gradient(to right, #3498db, #2980b9);
  padding: 5rem;
  border-radius: 30px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 4rem;
  color: #fff;
}

.ikm-cta__content h3 {
  font-size: 2.8rem;
  font-weight: 800;
  margin-bottom: 1rem;
}

.ikm-cta__content p {
  font-size: 1.6rem;
  opacity: 0.9;
  max-width: 60rem;
}

/* Buttons */
.btn {
  padding: 1.6rem 4rem;
  border-radius: 50px;
  text-decoration: none;
  font-size: 1.6rem;
  font-weight: 700;
  transition: all 0.3s;
  display: inline-block;
  cursor: pointer;
  border: none;
}

.btn--white { background: #fff; color: #3498db; }
.btn--white:hover { transform: translateY(-3px); box-shadow: 0 1rem 2rem rgba(0,0,0,0.2); }

.btn--green { background: #55c57a; color: #fff; }

@media (max-width: 1000px) {
  .areas-grid { grid-template-columns: repeat(2, 1fr); }
  .zi-stats-grid { grid-template-columns: 1fr; }
  .ac-banner { flex-direction: column; text-align: center; }
  .wbs-layout { flex-direction: column; }
  .ikm-cta { flex-direction: column; text-align: center; }
  .ikm-cta__content p { margin: 0 auto; }
  .heading-primary { font-size: 4rem; }
}

@media (max-width: 600px) {
  .areas-grid { grid-template-columns: 1fr; }
}
</style>
