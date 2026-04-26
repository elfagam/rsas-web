<script setup>
import { onMounted } from 'vue';
import SectionHeader from '../components/ui/SectionHeader.vue';
import { ASSETS } from '../config/assets';

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

const educationalStats = [
  { label: 'Mahasiswa Didik', value: '500+', icon: 'fa-solid fa-user-graduate' },
  { label: 'Dosen Klinik', value: '120+', icon: 'fa-solid fa-chalkboard-user' },
  { label: 'Program Studi', value: '15+', icon: 'fa-solid fa-book-medical' },
  { label: 'Penelitian Per Tahun', value: '80+', icon: 'fa-solid fa-microscope' }
];

const categories = [
  {
    title: 'Komkordik',
    desc: 'Komite Koordinasi Pendidikan yang mengatur seluruh alur pendidikan mahasiswa di RSAS.',
    icon: 'fa-solid fa-sitemap',
    links: ['Struktur Organisasi', 'Visi & Misi Pendidikan', 'Regulasi Pendidikan']
  },
  {
    title: 'Stase Klinik',
    desc: 'Informasi rotasi klinik bagi mahasiswa kedokteran di berbagai departemen spesialis.',
    icon: 'fa-solid fa-stethoscope',
    links: ['Jadwal Stase', 'Buku Panduan Mahasiswa', 'Sistem Penilaian']
  },
  {
    title: 'Penelitian & Etik',
    desc: 'Layanan administrasi penelitian dan pengajuan izin etik penelitian kesehatan (KEPK).',
    icon: 'fa-solid fa-vial-circle-check',
    links: ['Panduan Penelitian', 'Pengajuan Etik Online', 'Log Riset Mahasiswa']
  },
  {
    title: 'E-Learning & Library',
    desc: 'Akses ke perpustakaan medis digital dan platform pembelajaran online RSAS.',
    icon: 'fa-solid fa-laptop-code',
    links: ['Portal LMS', 'Perpustakaan Digital', 'Video Pembelajaran']
  }
];

onMounted(() => {
  window.scrollTo(0, 0);
  if (window.AOS) {
    window.AOS.refresh();
  }
});
</script>

<template>
  <div class="edu-page">
    <!-- Hero Section -->
    <header 
      class="edu-hero" 
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.pendidikan})` }"
    >
      <div class="container">
        <div class="edu-hero__content" data-aos="fade-up">
          <div class="edu-hero__badge">
            <i class="fa-solid fa-award"></i> Rumah Sakit Pendidikan Utama
          </div>
          <h1 class="heading-primary">
            <span class="heading-primary__main">Mencetak Masa Depan</span>
            <span class="heading-primary__sub">Tenaga Medis Indonesia</span>
          </h1>
          <p class="edu-hero__sub">Pusat integrasi keunggulan klinis, riset inovatif, dan pengembangan karakter tenaga kesehatan yang berintegritas tinggi.</p>
          <div class="edu-hero__actions">
            <a href="#portal" class="btn-academic">Jelajahi Portal Akademik</a>
            <a href="#penelitian" class="btn-academic-outline">Penelitian & Etik</a>
          </div>
        </div>
      </div>
    </header>

    <!-- Stats Bar (Floating) -->
    <section class="section-edu-stats">
      <div class="container">
        <div class="edu-stats-bar" data-aos="fade-up">
          <div v-for="stat in educationalStats" :key="stat.label" class="edu-stat-item">
            <div class="edu-stat-item__icon"><i :class="stat.icon"></i></div>
            <div class="edu-stat-item__content">
              <span class="edu-stat-item__value">{{ stat.value }}</span>
              <span class="edu-stat-item__label">{{ stat.label }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Main Categories (Bento Grid) -->
    <section class="section-categories" id="portal">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big" data-aos="fade-up">
          <h2 class="heading-secondary">Portal <span>Akademik RSAS</span></h2>
          <p class="heading-tertiary">Pusat koordinasi dan sumber daya bagi seluruh sivitas akademika.</p>
        </div>

        <div class="edu-bento-grid">
          <!-- Featured: Komkordik -->
          <div class="edu-bento-item edu-bento-item--featured" data-aos="fade-right">
            <div class="edu-bento-item__content">
              <div class="edu-bento-item__icon"><i class="fa-solid fa-sitemap"></i></div>
              <h3 class="edu-bento-item__title">Komkordik RSAS</h3>
              <p class="edu-bento-item__desc">Komite Koordinasi Pendidikan yang menjamin mutu dan sinkronisasi antara pelayanan medis, pendidikan, dan penelitian.</p>
              <ul class="edu-bento-item__links">
                <li><a href="#"><i class="fa-solid fa-id-badge"></i> Struktur Organisasi</a></li>
                <li><a href="#"><i class="fa-solid fa-bullseye"></i> Visi & Misi Pendidikan</a></li>
                <li><a href="#"><i class="fa-solid fa-scale-balanced"></i> Regulasi & Pedoman</a></li>
              </ul>
            </div>
            <div class="edu-bento-item__visual">
              <i class="fa-solid fa-building-columns"></i>
            </div>
          </div>

          <!-- Small Cards -->
          <div class="edu-bento-item" data-aos="fade-up" data-aos-delay="100">
            <div class="edu-bento-item__icon"><i class="fa-solid fa-stethoscope"></i></div>
            <h3 class="edu-bento-item__title">Stase Klinik</h3>
            <p class="edu-bento-item__desc">Panduan rotasi, jadwal, dan sistem penilaian mahasiswa kedokteran.</p>
            <a href="#" class="edu-bento-item__link">Buka Portal Stase &rarr;</a>
          </div>

          <div class="edu-bento-item" data-aos="fade-up" data-aos-delay="200">
            <div class="edu-bento-item__icon"><i class="fa-solid fa-vial-circle-check"></i></div>
            <h3 class="edu-bento-item__title">Penelitian & Etik</h3>
            <p class="edu-bento-item__desc">Layanan izin etik (KEPK) dan administrasi riset terpadu.</p>
            <a href="#" class="edu-bento-item__link">Info Etik Online &rarr;</a>
          </div>

          <div class="edu-bento-item edu-bento-item--dark" data-aos="fade-up" data-aos-delay="300">
            <div class="edu-bento-item__icon"><i class="fa-solid fa-laptop-code"></i></div>
            <h3 class="edu-bento-item__title">E-Learning</h3>
            <p class="edu-bento-item__desc">Akses perpustakaan digital dan platform LMS pendidikan RSAS.</p>
            <a href="#" class="edu-bento-item__link">Masuk Library &rarr;</a>
          </div>
        </div>
      </div>
    </section>

    <!-- Accreditation Section -->
    <section class="section-accreditation">
      <div class="container">
        <div class="accreditation-wrapper" data-aos="zoom-in">
          <div class="accreditation-info">
            <h3 class="accreditation-title">Kredibilitas Akademik Terakreditasi</h3>
            <p class="accreditation-text">RSUD Prof. Dr. H. Aloei Saboe telah memenuhi standar tertinggi sebagai institusi pendidikan medis nasional.</p>
          </div>
          <div class="accreditation-logos">
            <div class="acc-logo-item">
              <img :src="getImageUrl('logo-kars.png')" alt="KARS" />
              <span>Paripurna</span>
            </div>
            <div class="acc-logo-item">
              <i class="fa-solid fa-university"></i>
              <span>RS Pendidikan Utama</span>
            </div>
            <div class="acc-logo-item">
              <i class="fa-solid fa-shield-halved"></i>
              <span>Komite Etik Independen</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Research Focus -->
    <section class="section-research" id="penelitian">
      <div class="container">
        <div class="research-layout">
          <div class="research-content" data-aos="fade-right">
            <h2 class="heading-secondary">Mendukung Inovasi <span>Melalui Riset</span></h2>
            <p class="heading-tertiary">RSUD Prof. Dr. H. Aloei Saboe berkomitmen penuh dalam mendukung setiap upaya penelitian kesehatan bagi mahasiswa maupun praktisi medis melalui Komite Etik yang independen dan fasilitas riset yang mumpuni.</p>
            
            <div class="research-steps">
              <div class="research-step">
                <span class="research-step__num">1</span>
                <p>Konsultasi & Pengajuan Protokol Penelitian</p>
              </div>
              <div class="research-step">
                <span class="research-step__num">2</span>
                <p>Kajian Etik & Terbit Surat Izin Penelitian</p>
              </div>
              <div class="research-step">
                <span class="research-step__num">3</span>
                <p>Akses Fasilitas Riset & Pengumpulan Data</p>
              </div>
            </div>

            <button class="btn btn--green">Ajukan Izin Penelitian</button>
          </div>
          <div class="research-image" data-aos="fade-left">
            <img :src="ASSETS.galeri.lab" alt="Penelitian RSAS" />
          </div>
        </div>
      </div>
    </section>

    <!-- Downloadable Forms Section -->
    <section class="section-downloads">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big" data-aos="fade-up">
          <h2 class="heading-secondary">Arsip <span>& Formulir</span></h2>
          <p class="heading-tertiary">Unduh dokumen pendukung untuk keperluan pendidikan dan penelitian Anda.</p>
        </div>

        <div class="downloads-grid">
          <div class="download-card" data-aos="fade-up">
            <div class="download-card__icon"><i class="fa-solid fa-file-pdf"></i></div>
            <div class="download-card__info">
              <h4>Formulir Izin Penelitian</h4>
              <p>Format PDF - 1.2 MB</p>
            </div>
            <a href="#" class="btn-download-small"><i class="fa-solid fa-download"></i></a>
          </div>

          <div class="download-card" data-aos="fade-up" data-aos-delay="100">
            <div class="download-card__icon"><i class="fa-solid fa-file-word"></i></div>
            <div class="download-card__info">
              <h4>Protokol Etik Kesehatan</h4>
              <p>Format DOCX - 850 KB</p>
            </div>
            <a href="#" class="btn-download-small"><i class="fa-solid fa-download"></i></a>
          </div>

          <div class="download-card" data-aos="fade-up" data-aos-delay="200">
            <div class="download-card__icon"><i class="fa-solid fa-file-pdf"></i></div>
            <div class="download-card__info">
              <h4>Logbook Stase Klinik</h4>
              <p>Format PDF - 2.5 MB</p>
            </div>
            <a href="#" class="btn-download-small"><i class="fa-solid fa-download"></i></a>
          </div>

          <div class="download-card" data-aos="fade-up" data-aos-delay="300">
            <div class="download-card__icon"><i class="fa-solid fa-file-pdf"></i></div>
            <div class="download-card__info">
              <h4>Panduan Mahasiswa Didik</h4>
              <p>Format PDF - 5.1 MB</p>
            </div>
            <a href="#" class="btn-download-small"><i class="fa-solid fa-download"></i></a>
          </div>

          <div class="download-card" data-aos="fade-up" data-aos-delay="400">
            <div class="download-card__icon download-card__icon--blue"><i class="fa-solid fa-calendar-check"></i></div>
            <div class="download-card__info">
              <h4>Form Pendaftaran Seminar</h4>
              <p>Format PDF - 1.1 MB</p>
            </div>
            <a href="#" class="btn-download-small"><i class="fa-solid fa-download"></i></a>
          </div>

          <div class="download-card" data-aos="fade-up" data-aos-delay="500">
            <div class="download-card__icon download-card__icon--orange"><i class="fa-solid fa-poll-h"></i></div>
            <div class="download-card__info">
              <h4>Survei Kepuasan Mahasiswa</h4>
              <p>Format Link - Online Form</p>
            </div>
            <a href="#" class="btn-download-small"><i class="fa-solid fa-external-link"></i></a>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="section-edu-cta">
      <div class="container">
        <div class="edu-cta" data-aos="flip-up">
          <h2>Butuh Bantuan Akademik?</h2>
          <p>Hubungi sekretariat Komkordik RSAS untuk informasi lebih lanjut mengenai rotasi klinik dan administrasi pendidikan.</p>
          <a href="#" class="btn btn--white">Hubungi Sekretariat</a>
        </div>
      </div>
    </section>
  </div>
</template>

<style lang="scss" scoped>
@use "../assets/sass/base/variables" as *;

.container {
  max-width: 144rem;
  margin: 0 auto;
  padding: 0 4rem; /* Memberikan jarak aman di sisi kiri & kanan */

  @include respond("tablet") {
    padding: 0 2rem;
  }
}

/* Hero Section Styles */

/* Hero */
.edu-hero {
  height: 70vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  justify-content: center; /* Menjamin container berada di tengah secara horizontal */
  color: #fff;
  position: relative;
  text-align: center;
}

.edu-hero__content {
  width: 100%;
}

.edu-hero__badge {
  display: inline-flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem 2.5rem;
  background: rgba(85, 197, 122, 0.2);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 50px;
  font-size: 1.4rem;
  font-weight: 700;
  color: #fff;
  margin-bottom: 3rem;
  text-transform: uppercase;
  letter-spacing: 1.5px;
}

.edu-hero__badge i {
  color: #f1c40f;
  font-size: 1.8rem;
}

.edu-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  line-height: 1.6;
  margin: 0 auto 5rem; /* Menambahkan margin auto kiri-kanan untuk menempatkannya di tengah */
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
}

.edu-hero__actions {
  display: flex;
  gap: 2.5rem;
  justify-content: center; /* Memastikan tombol aksi berada di tengah */
}

.btn-academic {
  background: #55c57a;
  color: #fff;
  padding: 2rem 4rem;
  border-radius: 50px;
  font-size: 1.7rem;
  font-weight: 800;
  text-decoration: none;
  transition: all 0.3s;
  box-shadow: 0 1rem 3rem rgba(85, 197, 122, 0.3);
}

.btn-academic:hover {
  transform: translateY(-3px);
  background: #28b485;
  box-shadow: 0 1.5rem 4rem rgba(85, 197, 122, 0.4);
}

.btn-academic-outline {
  background: transparent;
  color: #fff;
  padding: 2rem 4rem;
  border: 2px solid rgba(255, 255, 255, 0.5);
  border-radius: 50px;
  font-size: 1.7rem;
  font-weight: 700;
  text-decoration: none;
  transition: all 0.3s;
  backdrop-filter: blur(5px);
}

.btn-academic-outline:hover {
  background: #fff;
  color: #55c57a;
}

.edu-hero__actions {
  display: flex;
  gap: 2.5rem;
}

.btn-academic {
  background: #55c57a;
  color: #fff;
  padding: 2rem 4rem;
  border-radius: 50px;
  font-size: 1.7rem;
  font-weight: 800;
  text-decoration: none;
  transition: all 0.3s;
  box-shadow: 0 1rem 3rem rgba(85, 197, 122, 0.3);
}

.btn-academic:hover {
  transform: translateY(-3px);
  background: #28b485;
  box-shadow: 0 1.5rem 4rem rgba(85, 197, 122, 0.4);
}

.btn-academic-outline {
  background: transparent;
  color: #fff;
  padding: 2rem 4rem;
  border: 2px solid rgba(255, 255, 255, 0.5);
  border-radius: 50px;
  font-size: 1.7rem;
  font-weight: 700;
  text-decoration: none;
  transition: all 0.3s;
  backdrop-filter: blur(5px);
}

.btn-academic-outline:hover {
  background: #fff;
  color: #55c57a;
}

/* Stats Bar (Floating) */
.section-edu-stats {
  margin-top: -7rem;
  position: relative;
  z-index: 100;
}

.edu-stats-bar {
  background: #fff;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  padding: 4rem;
  border-radius: 30px;
  box-shadow: 0 3rem 8rem rgba(0,0,0,0.1);
  border: 1px solid rgba(255, 255, 255, 0.5);
}

.edu-stat-item {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 2.5rem;
}

.edu-stat-item:not(:last-child) {
  border-right: 1px solid #eee;
}

@media (max-width: 900px) {
  .edu-stat-item:not(:last-child) {
    border-right: none;
    border-bottom: 1px solid #eee;
    padding-bottom: 2rem;
  }
}

.edu-stat-item__icon {
  font-size: 3.5rem;
  color: #55c57a;
  background: rgba(85, 197, 122, 0.1);
  width: 7rem;
  height: 7rem;
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.edu-stat-item__value {
  display: block;
  font-size: 3.2rem;
  font-weight: 900;
  color: #333;
  line-height: 1;
}

.edu-stat-item__label {
  font-size: 1.3rem;
  font-weight: 700;
  color: #888;
  text-transform: uppercase;
  letter-spacing: 1px;
}

/* Bento Grid */
.section-categories {
  padding: 10rem 0;
  background-color: #fcfdfe;
}

.edu-bento-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 3rem;
  margin-top: 5rem;
  align-items: stretch;
}

@media (max-width: 1300px) {
  .edu-bento-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

.edu-bento-item {
  background: #fff;
  padding: 4rem;
  border-radius: 35px;
  box-shadow: 0 1rem 4rem rgba(0,0,0,0.03);
  border: 1px solid #f0f0f0;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.edu-bento-item--featured {
  grid-column: 1 / -1; /* Ubah menjadi full width agar tidak merusak baris grid */
  background: linear-gradient(135deg, #fff 0%, #f9fbf9 100%);
  flex-direction: row;
  overflow: hidden;
  position: relative;
  min-height: 35rem;
}

@media (max-width: 800px) {
  .edu-bento-item--featured {
    flex-direction: column;
    min-height: auto;
  }
}

.edu-bento-item--dark {
  background: #1a1a1a;
  color: #fff;
}

.edu-bento-item--dark .edu-bento-item__title { color: #fff; }
.edu-bento-item--dark .edu-bento-item__desc { color: rgba(255, 255, 255, 0.7); }
.edu-bento-item--dark .edu-bento-item__icon { background: rgba(85, 197, 122, 0.2); }

.edu-bento-item:hover {
  transform: translateY(-1rem);
  box-shadow: 0 2rem 6rem rgba(85, 197, 122, 0.1);
  border-color: #55c57a;
}

.edu-bento-item__icon {
  width: 6rem;
  height: 6rem;
  background: rgba(85, 197, 122, 0.1);
  border-radius: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2.5rem;
  color: #55c57a;
  margin-bottom: 2.5rem;
}

.edu-bento-item__title {
  font-size: 2.4rem;
  font-weight: 800;
  color: #333;
  margin-bottom: 1.5rem;
}

.edu-bento-item__desc {
  font-size: 1.6rem;
  color: #666;
  line-height: 1.7;
  margin-bottom: 3rem;
}

.edu-bento-item__link {
  font-size: 1.5rem;
  font-weight: 700;
  color: #55c57a;
  text-decoration: none;
}

.edu-bento-item--featured .edu-bento-item__content {
  flex: 1;
  position: relative;
  z-index: 2;
}

.edu-bento-item--featured .edu-bento-item__visual {
  flex: 0 0 30%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 15rem;
  color: rgba(85, 197, 122, 0.05);
  position: absolute;
  right: -2rem;
  bottom: -2rem;
  transform: rotate(-15deg);
}

.edu-bento-item__links {
  list-style: none;
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1.5rem;
  margin-top: 2rem;
}

.edu-bento-item__links a {
  display: flex;
  align-items: center;
  gap: 1.2rem;
  padding: 1.5rem 2rem;
  background: #fff;
  border: 1px solid #eee;
  border-radius: 15px;
  text-decoration: none;
  color: #333;
  font-size: 1.4rem;
  font-weight: 700;
  transition: all 0.3s;
}

.edu-bento-item__links a:hover {
  background: #55c57a;
  color: #fff;
  border-color: #55c57a;
}

/* Accreditation */
.section-accreditation {
  padding-bottom: 10rem;
  background-color: #fcfdfe;
}

.accreditation-wrapper {
  background: #fff;
  padding: 6rem;
  border-radius: 40px;
  display: flex;
  align-items: center;
  gap: 8rem;
  box-shadow: 0 2rem 5rem rgba(0,0,0,0.02);
  border: 1px solid #f0f0f0;
}

.accreditation-info {
  flex: 1;
}

.accreditation-title {
  font-size: 3rem;
  font-weight: 900;
  color: #333;
  margin-bottom: 1.5rem;
}

.accreditation-text {
  font-size: 1.8rem;
  color: #777;
  line-height: 1.6;
  margin-bottom: 0;
}

.accreditation-logos {
  display: flex;
  gap: 5rem;
}

.acc-logo-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
}

.acc-logo-item img {
  height: 7rem;
  filter: grayscale(1);
  opacity: 0.6;
  transition: all 0.3s;
}

.acc-logo-item i {
  font-size: 5rem;
  color: #55c57a;
  opacity: 0.6;
}

.acc-logo-item:hover img,
.acc-logo-item:hover i {
  filter: grayscale(0);
  opacity: 1;
}

.acc-logo-item span {
  font-size: 1.2rem;
  font-weight: 800;
  text-transform: uppercase;
  color: #999;
  letter-spacing: 1px;
}

/* Research Box */
.section-research {
  padding-bottom: 8rem;
  background-color: #f9fbf9;
}

/* Downloads Section */
.section-downloads {
  padding: 10rem 0;
  background-color: #fff;
}

.downloads-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(35rem, 1fr));
  gap: 2.5rem;
  margin-top: 5rem;
}

.download-card {
  display: flex;
  align-items: center;
  gap: 2rem;
  padding: 2.5rem;
  background: #f9fbf9;
  border-radius: 18px;
  border: 1px solid #eee;
  transition: all 0.3s;
}

.download-card:hover {
  background: #fff;
  border-color: #55c57a;
  transform: translateY(-5px);
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
}

.download-card__icon {
  width: 5.5rem;
  height: 5.5rem;
  background: #fff;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2.2rem;
  color: #e74c3c; /* Red for PDF/Docs */
  box-shadow: 0 0.5rem 1.5rem rgba(0,0,0,0.05);
}

.download-card__icon--blue { color: #3498db; }
.download-card__icon--orange { color: #f39c12; }

.download-card__info {
  flex-grow: 1;
}

.download-card__info h4 {
  font-size: 1.6rem;
  font-weight: 800;
  color: #333;
  margin-bottom: 0.3rem;
}

.download-card__info p {
  font-size: 1.2rem;
  color: #888;
  font-weight: 600;
}

.btn-download-small {
  width: 4.5rem;
  height: 4.5rem;
  background: #55c57a;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.6rem;
  text-decoration: none;
  transition: all 0.3s;
}

.btn-download-small:hover {
  background: #33ab5a;
  transform: scale(1.1);
}

/* CTA */
/* Research Section - WBS Style */
.section-research {
  padding: 12rem 0;
  background-color: #fff;
}

.research-layout {
  display: flex;
  gap: 8rem;
  align-items: center;

  @include respond("tablet") {
    flex-direction: column;
    gap: 4rem;
  }
}

.research-content {
  flex: 1;
}

.research-content p.heading-tertiary {
  margin: 2rem 0 4rem;
}

.research-steps {
  display: grid;
  gap: 2rem;
  margin-bottom: 4rem;
}

.research-step {
  display: flex;
  align-items: center;
  gap: 2rem;
  padding: 2.2rem;
  background: #f9fbf9;
  border-radius: 15px;
  transition: all 0.3s;
  border: 1px solid transparent;

  &:hover {
    background: #fff;
    border-color: #55c57a;
    box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
    transform: translateX(1rem);
  }
}

.research-step__num {
  width: 4.5rem;
  height: 4.5rem;
  background: #55c57a;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  font-weight: 800;
  flex-shrink: 0;
}

.research-step p {
  margin: 0;
  font-size: 1.7rem;
  font-weight: 700;
  color: $warna-text;
}

.research-image {
  flex: 1;
  border-radius: 30px;
  overflow: hidden;
  box-shadow: 0 3rem 6rem rgba(0,0,0,0.1);
}

.research-image img {
  width: 100%;
  height: 50rem;
  object-fit: cover;
}

.research-step {
  display: flex;
  align-items: center;
  gap: 2rem;
  padding: 2.2rem;
  background: #f9fbf9;
  border-radius: 15px;
  transition: all 0.3s;
  border: 1px solid transparent;

  &:hover {
    background: #fff;
    border-color: #55c57a;
    box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
    transform: translateX(1rem);
  }
}

.research-step__num {
  width: 4.5rem;
  height: 4.5rem;
  background: #55c57a;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  font-weight: 800;
  flex-shrink: 0;
}

.research-step p {
  margin: 0;
  font-size: 1.7rem;
  font-weight: 700;
  color: $warna-text;
}

.research-image {
  flex: 1;
  border-radius: 30px;
  overflow: hidden;
  box-shadow: 0 3rem 6rem rgba(0,0,0,0.1);
}

.research-image img {
  width: 100%;
  height: 50rem;
  object-fit: cover;
}

/* CTA */
.section-edu-cta {
  padding: 10rem 0;
}

.edu-cta {
  background: linear-gradient(to right, #55c57a, #33ab5a);
  padding: 8rem;
  border-radius: 30px;
  text-align: center;
  color: #fff;
}

.edu-cta h2 { font-size: 4rem; margin-bottom: 2rem; }
.edu-cta p { font-size: 1.8rem; margin-bottom: 4rem; opacity: 0.9; }

/* Buttons */
.btn {
  padding: 1.5rem 3.5rem;
  border-radius: 50px;
  text-decoration: none;
  font-size: 1.6rem;
  font-weight: 700;
  transition: all 0.3s;
  display: inline-block;
  cursor: pointer;
  border: none;
}

.btn--white { background: #fff; color: #55c57a; }
.btn--white:hover { transform: translateY(-3px); box-shadow: 0 1rem 2rem rgba(0,0,0,0.2); }

.btn--outline-white { background: transparent; border: 2px solid #fff; color: #fff; }
.btn--outline-white:hover { background: #fff; color: #55c57a; }

.btn--green { background: #55c57a; color: #fff; }

@media (max-width: 1000px) {
  .edu-stats-bar { 
    grid-template-columns: repeat(2, 1fr); 
    gap: 3rem;
  }
  .research-box { flex-direction: column; }
  .heading-primary { font-size: 4rem; }
}

@media (max-width: 600px) {
  .edu-stats-bar { 
    grid-template-columns: 1fr; 
    padding: 3rem;
  }
}
</style>
