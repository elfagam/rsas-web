<script setup>
import { ref, onMounted } from 'vue';
import HistoryTimeline from '../components/ui/HistoryTimeline.vue';
import RenstraCard from '../components/ui/RenstraCard.vue';
import ManagementCard from '../components/ui/ManagementCard.vue';
import SectionHeader from '../components/ui/SectionHeader.vue';
import { ASSETS } from '../config/assets';

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

const orgData = {
  dewanPengawas: [
    { name: 'Ketua Dewan Pengawas', position: 'Ketua', image: null, motto: 'Mengawasi dengan Integritas' },
    { name: 'Anggota Dewas 1', position: 'Anggota', image: null },
    { name: 'Anggota Dewas 2', position: 'Anggota', image: null },
    { name: 'Anggota Dewas 3', position: 'Anggota', image: null },
    { name: 'Anggota Dewas 4', position: 'Anggota', image: null }
  ],
  direktur: { 
    name: 'dr. H. Ahmad Lihu, M.Si', 
    position: 'Direktur Utama', 
    nip: '19700101 199503 1 001',
    image: getImageUrl('direktur-aloeisaboe.jpg'),
    motto: 'Kepemimpinan adalah teladan dalam melayani masyarakat.'
  },
  wadir: [
    { name: 'Wadir Pelayanan', position: 'Wadir Pelayanan', image: null },
    { name: 'Wadir Umum & Keuangan', position: 'Wadir Umum & Keuangan', image: null },
    { name: 'Wadir Diklit & Pengembangan', position: 'Wadir Diklit', image: null }
  ],
  spi: [
    { name: 'Ketua SPI', position: 'Ketua SPI', image: null },
    { name: 'Anggota SPI 1', position: 'Anggota SPI', image: null },
    { name: 'Anggota SPI 2', position: 'Anggota SPI', image: null }
  ],
  bidang: [
    { name: 'Kabid Pelayanan Medik', position: 'Kabid', image: null },
    { name: 'Kabid Keperawatan', position: 'Kabid', image: null },
    { name: 'Kabid Penunjang', position: 'Kabid', image: null },
    { name: 'Kabag Tata Usaha', position: 'Kabag', image: null },
    { name: 'Kabag Perencanaan', position: 'Kabag', image: null },
    { name: 'Kabag Keuangan', position: 'Kabag', image: null },
    { name: 'Kabid Diklit', position: 'Kabid', image: null },
    { name: 'Kabid Humas & Pemasaran', position: 'Kabid', image: null },
    { name: 'Kabid SDM', position: 'Kabid', image: null }
  ]
};

const history = [
  { 
    year: '1926', 
    title: 'Awal Pembangunan',
    desc: 'Pembangunan RSU Kotamadya Gorontalo dimulai sebagai langkah awal pelayanan kesehatan modern di wilayah ini.',
    icon: 'fa-solid fa-hammer'
  },
  { 
    year: '1929', 
    title: 'Pemanfaatan Resmi',
    desc: 'Pemanfaatan secara resmi dengan nama RSU Kotamadya Gorontalo untuk melayani masyarakat umum.',
    icon: 'fa-solid fa-hospital'
  },
  { 
    year: '1979', 
    title: 'Perubahan Nama',
    desc: 'Berubah nama menjadi RSU Prof. Dr. H. Aloei Saboe Gorontalo sebagai penghormatan terhadap tokoh kesehatan daerah.',
    icon: 'fa-solid fa-signature'
  },
  { 
    year: '2005', 
    title: 'Rumah Sakit Kelas B',
    desc: 'Ditetapkan sebagai Rumah Sakit Kelas B dan menjadi Badan Layanan Umum Daerah (BLUD).',
    icon: 'fa-solid fa-award'
  },
  { 
    year: '2023', 
    title: 'Akreditasi Paripurna',
    desc: 'Meraih tingkat Akreditasi Paripurna oleh KARS sebagai bukti kualitas pelayanan standar nasional.',
    icon: 'fa-solid fa-star'
  }
];

const values = [
  { title: 'Profesional', desc: 'Melayani dengan kompetensi dan etika tinggi.', icon: 'fa-solid fa-user-tie' },
  { title: 'Inovatif', desc: 'Terus mengembangkan teknologi dan sistem layanan.', icon: 'fa-solid fa-lightbulb' },
  { title: 'Integritas', desc: 'Jujur, transparan, dan bertanggung jawab.', icon: 'fa-solid fa-shield-halved' },
  { title: 'Empati', desc: 'Melayani pasien dengan kasih sayang dan kepedulian.', icon: 'fa-solid fa-heart' }
];

const renstra = {
  periode: '2024 - 2029',
  pillars: [
    { 
      title: 'Transformasi Layanan digital', 
      desc: 'Implementasi penuh Rekam Medis Elektronik (RME) dan sistem antrean terintegrasi mobile.',
      icon: 'fa-solid fa-microchip',
      target: '2025'
    },
    { 
      title: 'Pusat Rujukan Regional', 
      desc: 'Menjadi pusat rujukan jantung dan onkologi terbaik di wilayah Sulawesi Utara & Gorontalo.',
      icon: 'fa-solid fa-hospital-user',
      target: '2027'
    },
    { 
      title: 'Green Hospital', 
      desc: 'Penerapan konsep rumah sakit ramah lingkungan dengan efisiensi energi dan pengelolaan limbah modern.',
      icon: 'fa-solid fa-leaf',
      target: '2028'
    },
    { 
      title: 'World Class Education', 
      desc: 'Menyelenggarakan pendidikan kedokteran dan riset kesehatan berskala internasional.',
      icon: 'fa-solid fa-earth-asia',
      target: '2029'
    }
  ]
};

onMounted(() => {
  window.scrollTo(0, 0);
  if (window.AOS) {
    window.AOS.refresh();
  }
});
</script>

<template>
  <div class="profil-page">
    <header 
      class="profil-hero" 
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.profil})` }"
    >
      <div class="container">
        <div class="profil-hero__content" data-aos="fade-up">
          <div class="profil-hero__badge">Legacy of Excellence</div>
          <h1 class="heading-primary">
            <span class="heading-primary__main">Jejak Bakti &</span>
            <span class="heading-primary__sub">Dedikasi Kami</span>
          </h1>
          <p class="profil-hero__sub">Menjadi mercusuar kesehatan yang mandiri dan terkemuka, memadukan inovasi medis dengan ketulusan hati dalam melayani.</p>
        </div>
      </div>
    </header>

    <!-- Vision & Mission -->
    <section class="section-vision">
      <div class="container">
        <div class="vision-wrapper" data-aos="fade-up">
          <div class="vision-content">
            <div class="vision-main">
              <span class="vision-label">Visi Kami</span>
              <h2 class="vision-title">"Terwujudnya Pelayanan Kesehatan yang Bermutu, Mandiri dan Berkeadilan bagi Seluruh Masyarakat."</h2>
            </div>
            
            <div class="mission-side">
              <span class="vision-label">Misi Utama</span>
              <ul class="mission-list">
                <li><i class="fa-solid fa-circle-check"></i> Meningkatkan kualitas pelayanan kesehatan paripurna.</li>
                <li><i class="fa-solid fa-circle-check"></i> Mengembangkan sarana prasarana medis modern.</li>
                <li><i class="fa-solid fa-circle-check"></i> Manajemen transparan & akuntabel.</li>
                <li><i class="fa-solid fa-circle-check"></i> Pendidikan & penelitian berkualitas.</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Cultural Values -->
    <section class="section-values">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big" data-aos="fade-up">
          <h2 class="heading-secondary">Pilar <span>Budaya Kami</span></h2>
          <p class="heading-tertiary">Nilai-nilai luhur yang menjadi kompas dalam setiap tindakan medis dan pelayanan kami.</p>
        </div>
        <div class="values-grid">
          <div v-for="val in values" :key="val.title" class="value-item" data-aos="zoom-in">
            <div class="value-item__icon"><i :class="val.icon"></i></div>
            <h3 class="value-item__title">{{ val.title }}</h3>
            <p class="value-item__desc">{{ val.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Strategic Plan (Renstra) -->
    <section class="section-renstra">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big">
          <h2 class="heading-secondary">Rencana <span>Strategis RSAS</span></h2>
          <p class="heading-tertiary">Arah kebijakan dan target utama untuk periode {{ renstra.periode }}</p>
        </div>

        <div class="renstra-grid">
          <RenstraCard 
            v-for="(pilar, index) in renstra.pillars" 
            :key="index"
            v-bind="pilar"
            :delay="index * 100"
          />
        </div>

        <div class="u-center-text u-margin-top-big">
          <a href="#" class="btn-download">
            <i class="fa-solid fa-file-pdf"></i> Unduh Dokumen Renstra Lengkap (PDF)
          </a>
        </div>
      </div>
    </section>

    <!-- Hospital Management Section -->
    <section class="section-management">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big" data-aos="fade-up">
          <h2 class="heading-secondary">Kepemimpinan <span>& Manajemen</span></h2>
          <p class="heading-tertiary">Sinergi para ahli dalam mewujudkan tata kelola rumah sakit yang unggul.</p>
        </div>

        <div class="mgmt-layout">
          <!-- Dewan Pengawas -->
          <div class="mgmt-group">
            <h3 class="mgmt-group-title">Dewan Pengawas</h3>
            <div class="mgmt-row mgmt-row--center">
              <ManagementCard 
                v-for="(dewas, index) in orgData.dewanPengawas" 
                :key="index"
                v-bind="dewas" 
                class="mgmt-card--small"
                :delay="index * 50"
              />
            </div>
          </div>

          <!-- Direktur Utama -->
          <div class="mgmt-group">
            <div class="mgmt-row mgmt-row--center">
              <ManagementCard 
                v-bind="orgData.direktur" 
                class="mgmt-card--large"
              />
            </div>
          </div>

          <!-- Wakil Direktur -->
          <div class="mgmt-group">
            <div class="mgmt-row">
              <ManagementCard 
                v-for="(wadir, index) in orgData.wadir" 
                :key="index"
                v-bind="wadir" 
                :delay="index * 100"
              />
            </div>
          </div>

          <!-- SPI -->
          <div class="mgmt-group">
            <h3 class="mgmt-group-title">Satuan Pengawas Intern (SPI)</h3>
            <div class="mgmt-row mgmt-row--center">
              <ManagementCard 
                v-for="(spi, index) in orgData.spi" 
                :key="index"
                v-bind="spi" 
                class="mgmt-card--small"
                :delay="index * 50"
              />
            </div>
          </div>

          <!-- Kepala Bidang / Bagian -->
          <div class="mgmt-group">
            <h3 class="mgmt-group-title">Kepala Bidang & Bagian</h3>
            <div class="mgmt-grid">
              <ManagementCard 
                v-for="(bidang, index) in orgData.bidang" 
                :key="index"
                v-bind="bidang" 
                class="mgmt-card--mini"
                :delay="index * 50"
              />
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- History Timeline -->
    <section class="section-history-detail">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big">
          <h2 class="heading-secondary">Jejak Langkah Kami</h2>
          <p class="section-subtitle">Perjalanan panjang dalam melayani masyarakat Gorontalo.</p>
        </div>
        
        <HistoryTimeline :historyData="history" />
      </div>
    </section>

    <!-- Accreditation -->
    <section class="section-accreditation">
      <div class="container">
        <div class="accreditation-box" data-aos="zoom-out">
          <div class="accreditation-box__img">
            <img 
              :src="getImageUrl('logo-kars.png')" 
              alt="KARS Logo"
              @error="(e) => e.target.src = 'https://ui-avatars.com/api/?name=KARS&background=55c57a&color=fff'"
            >
          </div>
          <div class="accreditation-box__content">
            <h2 class="heading-secondary">Terakreditasi Paripurna</h2>
            <p>RSUD Prof. Dr. H. Aloei Saboe telah meraih predikat **PARIPURNA (Bintang Lima)** dari Komisi Akreditasi Rumah Sakit (KARS) Republik Indonesia. Ini adalah jaminan komitmen kami terhadap keselamatan pasien dan mutu pelayanan.</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
/* Profil Section Styles */

.container {
  max-width: 120rem;
  margin: 0 auto;
  padding: 0 2rem;
}

.u-center-text { text-align: center; }
.u-margin-bottom-big { margin-bottom: 8rem; }
.u-margin-bottom-small { margin-bottom: 2rem; }

.section-subtitle {
  font-size: 1.8rem;
  color: #777;
  margin-top: 1rem;
}

/* Hero Profil */
.profil-hero {
  height: 60vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
  text-align: center;
}

.profil-hero__badge {
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

.profil-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  margin: 0 auto;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
  line-height: 1.6;
}

/* Vision Mission */
.section-vision { 
  padding: 12rem 0; 
  background: #fff;
}

.vision-wrapper {
  background: #f9fbf9;
  border-radius: 40px;
  padding: 8rem;
  position: relative;
  overflow: hidden;
  border: 1px solid #f1f1f1;
}

.vision-wrapper::before {
  content: '"';
  position: absolute;
  top: -5rem;
  left: 2rem;
  font-size: 30rem;
  font-family: 'Playfair Display', serif;
  color: rgba(85, 197, 122, 0.05);
  line-height: 1;
}

.vision-content {
  display: grid;
  grid-template-columns: 1.2fr 0.8fr;
  gap: 8rem;
  align-items: center;
}

.vision-label {
  display: block;
  font-size: 1.4rem;
  font-weight: 800;
  color: #55c57a;
  text-transform: uppercase;
  letter-spacing: 2px;
  margin-bottom: 2rem;
}

.vision-title {
  font-family: 'Playfair Display', serif;
  font-size: 4rem;
  line-height: 1.3;
  color: #2d3436;
  font-weight: 700;
}

.mission-list {
  list-style: none;
  font-size: 1.8rem;
  color: #636e72;
}

.mission-list li {
  margin-bottom: 2rem;
  display: flex;
  gap: 1.5rem;
  align-items: flex-start;
}

.mission-list i {
  color: #55c57a;
  margin-top: 0.5rem;
}

/* Values */
.section-values { padding: 10rem 0; }

.values-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 3rem;
}

.value-item {
  text-align: center;
  padding: 5rem 3rem;
  border-radius: 25px;
  background: #fff;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.03);
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  border: 1px solid #f9fbf9;
}

.value-item:hover { 
  transform: translateY(-1.5rem); 
  box-shadow: 0 2rem 5rem rgba(85, 197, 122, 0.1);
  border-color: rgba(85, 197, 122, 0.2);
}

.value-item__icon {
  width: 8rem;
  height: 8rem;
  background: #f9fbf9;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3.5rem;
  color: #55c57a;
  margin: 0 auto 3rem;
  transition: all 0.4s;
}

.value-item:hover .value-item__icon {
  background: #55c57a;
  color: #fff;
}

.value-item__title {
  font-size: 2.2rem;
  font-weight: 800;
  margin-bottom: 1.5rem;
  color: #2d3436;
}

.value-item__desc {
  font-size: 1.6rem;
  color: #636e72;
  line-height: 1.6;
}

/* History Detail */
.section-history-detail { padding: 10rem 0; background: #fff; }

.history-list {
  max-width: 90rem;
  margin: 0 auto;
}

.history-item {
  display: flex;
  gap: 4rem;
  margin-bottom: 4rem;
}

.history-item__year {
  font-size: 3rem;
  font-weight: 900;
  color: #55c57a;
  flex: 0 0 12rem;
}

.history-item__content {
  background: #f9fbf9;
  padding: 3rem;
  border-radius: 15px;
  flex: 1;
  position: relative;
}

.history-item__icon {
  position: absolute;
  top: -2rem;
  right: 2rem;
  width: 5rem;
  height: 5rem;
  background: #55c57a;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  box-shadow: 0 1rem 2rem rgba(85, 197, 122, 0.3);
}

.history-item__title { font-size: 2.2rem; font-weight: 800; margin-bottom: 1rem; color: #333; }
.history-item__desc { font-size: 1.6rem; color: #666; line-height: 1.7; }

/* Accreditation */
.section-accreditation { padding: 10rem 0; background: #f9fbf9; }

.accreditation-box {
  background: #fff;
  padding: 8rem;
  border-radius: 40px;
  display: flex;
  align-items: center;
  gap: 8rem;
  box-shadow: 0 3rem 6rem rgba(0,0,0,0.06);
  border: 1px solid #f1f1f1;
}

.accreditation-box__img img {
  height: 18rem;
  filter: drop-shadow(0 1rem 2rem rgba(85, 197, 122, 0.2));
}

.accreditation-box__content {
  flex: 1;
}

.accreditation-box__content h2 {
  font-family: 'Playfair Display', serif;
  font-size: 4rem;
  font-weight: 800;
  color: #2d3436;
}

.accreditation-box__content p {
  font-size: 2rem;
  line-height: 1.8;
  color: #636e72;
  margin-top: 2rem;
  font-weight: 300;
}

@media (max-width: 1000px) {
  .vision-grid, .values-grid { grid-template-columns: 1fr; }
  .accreditation-box { flex-direction: column; text-align: center; }
}

/* Management Section */
.section-management {
  padding: 10rem 0;
  background-color: #fff;
}

.mgmt-layout {
  display: flex;
  flex-direction: column;
  gap: 8rem;
  max-width: 140rem;
  margin: 0 auto;
}

.mgmt-group {
  display: flex;
  flex-direction: column;
  gap: 3rem;
}

.mgmt-group-title {
  font-size: 2.2rem;
  font-weight: 800;
  color: #333;
  text-align: center;
  position: relative;
  padding-bottom: 1.5rem;
}

.mgmt-group-title::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 5rem;
  height: 3px;
  background-color: #55c57a;
}

.mgmt-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(30rem, 1fr));
  gap: 3rem;
  justify-content: center;
}

.mgmt-row--center {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 3rem;
}

.mgmt-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(28rem, 1fr));
  gap: 3rem;
}

.mgmt-card--large {
  width: 40rem;
}

.mgmt-card--small {
  width: 28rem;
}

/* Penyesuaian khusus untuk mini card agar tidak terlalu tinggi */
:deep(.mgmt-card--mini .mgmt-card__img-box) {
  height: 25rem;
}

@media (max-width: 1000px) {
  .mgmt-card--large, .mgmt-card--small {
    width: 100%;
    max-width: 40rem;
  }
}

@media (max-width: 800px) {
  .mgmt-row {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 600px) {
  .heading-primary { font-size: 3.5rem; }
  .profil-hero { height: 50vh; }
  .mgmt-grid {
    grid-template-columns: 1fr;
  }
}

/* Strategic Plan */
.section-renstra {
  padding: 12rem 0;
  background: linear-gradient(135deg, #f9fbf9 0%, #ffffff 100%);
  position: relative;
  overflow: hidden;
}

.section-renstra::before {
  content: "RENSTRA";
  position: absolute;
  top: 5rem;
  left: 50%;
  transform: translateX(-50%);
  font-size: 15rem;
  font-weight: 900;
  color: rgba(85, 197, 122, 0.03);
  z-index: 0;
}

.u-margin-top-big { margin-top: 8rem; }

.renstra-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 3rem;
  position: relative;
  z-index: 1;
}

.btn-download {
  display: inline-flex;
  align-items: center;
  gap: 1.5rem;
  padding: 1.8rem 4rem;
  background-color: #fff;
  color: #333;
  text-decoration: none;
  font-size: 1.6rem;
  font-weight: 700;
  border-radius: 50px;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.1);
  transition: all 0.3s;
  border: 1px solid #eee;
}

.btn-download i {
  font-size: 2.2rem;
  color: #e74c3c;
}

.btn-download:hover {
  transform: translateY(-3px);
  box-shadow: 0 1.5rem 4rem rgba(0,0,0,0.15);
  background-color: #fcfcfc;
}

@media (max-width: 1200px) {
  .renstra-grid { grid-template-columns: repeat(2, 1fr); }
}

@media (max-width: 600px) {
  .renstra-grid { grid-template-columns: 1fr; }
}

</style>
