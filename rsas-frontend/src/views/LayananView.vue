<script setup>
import { ref, computed, onMounted } from 'vue';
import SectionHeader from '../components/ui/SectionHeader.vue';
import ServiceCard from '../components/ui/ServiceCard.vue';
import { ASSETS } from '../config/assets';

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

const searchQuery = ref('');
const activeCategory = ref('Semua');

const categories = ['Semua', 'Unggulan', 'Rawat Jalan', 'Rawat Inap', 'Penunjang'];

const allServices = ref([
  { 
    id: 'jantung', 
    title: 'Pusat Jantung Terpadu', 
    category: 'Unggulan',
    icon: 'fa-solid fa-heart-pulse', 
    text: 'Layanan kardiovaskular komprehensif mulai dari diagnostik non-invasif hingga kateterisasi jantung (Cath Lab).',
    link: '/layanan/jantung'
  },
  { 
    id: 'onkologi', 
    title: 'Pusat Onkologi (Cancer Center)', 
    category: 'Unggulan',
    icon: 'fa-solid fa-ribbon', 
    text: 'Penanganan kanker terpadu dengan filosofi "Keluarga Melayani Keluarga", mencakup bedah onkologi, kemoterapi, dan dukungan psikososial.',
    link: '/layanan/onkologi',
    specialClass: 'service-card--oncology'
  },
  { 
    id: 'stroke', 
    title: 'Neuro-Vascular & Stroke Center', 
    category: 'Unggulan',
    icon: 'fa-solid fa-brain', 
    text: 'Layanan darurat stroke 24 jam dengan penanganan cepat untuk meminimalisir risiko kecacatan permanen.',
    link: '/layanan/stroke'
  },
  { 
    id: 'hemodialisa', 
    title: 'Unit Hemodialisa Terpadu', 
    category: 'Unggulan',
    icon: 'fa-solid fa-droplet', 
    text: 'Layanan cuci darah modern dengan peralatan steril dan pemantauan ketat oleh dokter spesialis ginjal.',
    link: '/layanan/hemodialisa'
  },
  { 
    id: 'igd', 
    title: 'IGD & Trauma Center 24 Jam', 
    category: 'Unggulan',
    icon: 'fa-solid fa-truck-medical', 
    text: 'Pusat penanganan gawat darurat dan trauma dengan tim medis bersertifikat ATLS & ACLS.',
    link: '/layanan/igd'
  },
  { 
    id: 'diagnostik', 
    title: 'Pusat Diagnostik Canggih', 
    category: 'Penunjang',
    icon: 'fa-solid fa-microscope', 
    text: 'Radiologi lanjut (CT-Scan 128 Slice, MRI 1.5 Tesla) dan laboratorium patologi klinik otomatis.',
    link: '/layanan/diagnostik'
  },
  { 
    id: 'mcu', 
    title: 'Executive Medical Check Up', 
    category: 'Rawat Jalan',
    icon: 'fa-solid fa-user-check', 
    text: 'Paket pemeriksaan kesehatan menyeluruh dengan fasilitas lounge eksklusif dan hasil yang cepat.',
    link: '/layanan/mcu'
  },
  { 
    id: 'poli-anak', 
    title: 'Poli Anak & Tumbuh Kembang', 
    category: 'Rawat Jalan',
    icon: 'fa-solid fa-baby', 
    text: 'Layanan kesehatan anak holistik termasuk imunisasi dan konsultasi psikologi anak.',
    link: '/layanan/poli-anak'
  },
  { 
    id: 'poli-dalam', 
    title: 'Poli Penyakit Dalam', 
    category: 'Rawat Jalan',
    icon: 'fa-solid fa-stethoscope', 
    text: 'Diagnosis dan penanganan penyakit metabolik, infeksi, dan imunologi secara mendalam.',
    link: '/layanan/poli-dalam'
  },
  { 
    id: 'icu', 
    title: 'ICU, NICU & PICU', 
    category: 'Rawat Inap',
    icon: 'fa-solid fa-bed-pulse', 
    text: 'Unit perawatan intensif untuk dewasa, bayi, dan anak dengan monitor hemodinamik invasif.',
    link: '/layanan/icu'
  },
  { 
    id: 'vip', 
    title: 'Rawat Inap VIP & VVIP', 
    category: 'Rawat Inap',
    icon: 'fa-solid fa-hotel', 
    text: 'Fasilitas perawatan premium dengan pelayanan sekelas hotel berbintang dan privasi maksimal.',
    link: '/layanan/vip'
  },
  { 
    id: 'fisioterapi', 
    title: 'Rehabilitasi Medik', 
    category: 'Penunjang',
    icon: 'fa-solid fa-person-walking-with-cane', 
    text: 'Terapi fisik, okupasi, dan wicara untuk pemulihan fungsi motorik dan kognitif pasca sakit.',
    link: '/layanan/fisioterapi'
  },
]);

const filteredServices = computed(() => {
  return allServices.value.filter(service => {
    const matchesSearch = service.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchesCategory = activeCategory.value === 'Semua' || service.category === activeCategory.value;
    return matchesSearch && matchesCategory;
  });
});

onMounted(() => {
  window.scrollTo(0, 0);
  if (window.AOS) {
    window.AOS.refresh();
  }
});
</script>

<template>
  <div class="layanan-view">
    <!-- Hero Section -->
    <header 
      class="layanan-hero"
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.layanan})` }"
    >
      <div class="container">
        <div class="layanan-hero__content" data-aos="fade-up">
          <h1 class="heading-primary">
            <span class="heading-primary__main">Layanan Medis</span>
            <span class="heading-primary__sub">Paripurna & Berempati</span>
          </h1>
          <p class="layanan-hero__sub">Menghadirkan standar keunggulan klinis yang dipadukan dengan ketulusan hati untuk setiap langkah kesembuhan Anda.</p>
        </div>
      </div>
    </header>

    <!-- Filter & Search Section -->
    <section class="section-filter">
      <div class="container">
        <div class="filter-box">
          <div class="search-group">
            <i class="fa-solid fa-magnifying-glass search-icon"></i>
            <input 
              type="text" 
              v-model="searchQuery" 
              placeholder="Cari layanan (misal: Jantung, Anak)..." 
              class="search-input"
            />
          </div>
          
          <div class="category-list">
            <button 
              v-for="cat in categories" 
              :key="cat"
              @click="activeCategory = cat"
              class="category-btn"
              :class="{ 'category-btn--active': activeCategory === cat }"
            >
              {{ cat }}
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- Services Grid -->
    <section class="section-services-grid">
      <div class="container">
        <div v-if="filteredServices.length > 0" class="services-grid">
          <ServiceCard 
            v-for="(service, index) in filteredServices" 
            :key="service.id"
            v-bind="service"
            :delay="index * 50"
            :class="service.specialClass"
          />
        </div>
        
        <!-- Empty State -->
        <div v-else class="empty-state" data-aos="fade-up">
          <div class="empty-state__icon">
            <i class="fa-solid fa-folder-open"></i>
          </div>
          <h3>Layanan tidak ditemukan</h3>
          <p>Coba gunakan kata kunci lain atau periksa kategori yang Anda pilih.</p>
          <button @click="searchQuery = ''; activeCategory = 'Semua'" class="btn-reset">Lihat Semua Layanan</button>
        </div>
      </div>
    </section>

    <!-- Sanctuary of Care Section (Soulful Philosophy) -->
    <section class="section-philosophy">
      <div class="container">
        <div class="philosophy-box" data-aos="fade-up">
          <div class="philosophy-content">
            <h2 class="heading-secondary">Sanctuary of Care: <span>Melayani Melebihi Tindakan Medis</span></h2>
            <p class="philosophy-text">
              Kami percaya bahwa kesembuhan sejati dimulai dari ketenangan batin. Di RSAS, kami tidak hanya mengobati penyakit, tetapi juga merawat harapan. Setiap pasien diperlakukan seperti keluarga sendiri, dengan penuh martabat, empati, dan doa yang tulus.
            </p>
            <div class="philosophy-features">
              <div class="p-feat">
                <i class="fa-solid fa-heart-circle-check"></i>
                <span>Pelayanan Berbasis Kekeluargaan</span>
              </div>
              <div class="p-feat">
                <i class="fa-solid fa-dove"></i>
                <span>Lingkungan yang Menenangkan</span>
              </div>
              <div class="p-feat">
                <i class="fa-solid fa-hands-holding-child"></i>
                <span>Dukungan Emosional bagi Keluarga</span>
              </div>
            </div>
          </div>
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

.layanan-view {
  /* Padding top dihapus agar Hero menyentuh bagian paling atas layar */
  min-height: 100vh;
}

/* Hero */
.layanan-hero {
  height: 60vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
  text-align: center;
}

.heading-academic {
  font-family: 'Playfair Display', serif;
  font-size: 6rem;
  font-weight: 900;
  line-height: 1.1;
  margin-bottom: 2.5rem;
}

.heading-academic span {
  font-style: italic;
  font-weight: 400;
  display: block;
  font-size: 4.5rem;
  margin-top: 1rem;
}

.layanan-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  margin: 0 auto;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
  line-height: 1.6;
}

/* Filter Box */
.section-filter {
  margin-top: -5rem;
  position: relative;
  z-index: 10;
}

.filter-box {
  background: #fff;
  padding: 3rem;
  border-radius: 20px;
  box-shadow: 0 2rem 5rem rgba(0,0,0,0.1);
  display: flex;
  flex-direction: column;
  gap: 2.5rem;
}

.search-group {
  position: relative;
  width: 100%;
}

.search-icon {
  position: absolute;
  left: 2rem;
  top: 50%;
  transform: translateY(-50%);
  color: #55c57a;
  font-size: 1.8rem;
}

.search-input {
  width: 100%;
  padding: 1.8rem 2rem 1.8rem 5.5rem;
  border: 2px solid #f1f1f1;
  border-radius: 12px;
  font-size: 1.6rem;
  transition: all 0.3s;
}

.search-input:focus {
  outline: none;
  border-color: #55c57a;
  box-shadow: 0 0 0 4px rgba(85, 197, 122, 0.1);
}

.category-list {
  display: flex;
  flex-wrap: wrap;
  gap: 1.2rem;
}

.category-btn {
  padding: 1rem 2.2rem;
  border-radius: 50px;
  border: 1px solid #eee;
  background: #fff;
  color: #666;
  font-size: 1.4rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
}

.category-btn:hover {
  background: #f9fbf9;
  border-color: #55c57a;
  color: #55c57a;
}

.category-btn--active {
  background: #55c57a;
  color: #fff;
  border-color: #55c57a;
  box-shadow: 0 0.5rem 1.5rem rgba(85, 197, 122, 0.3);
}

/* Grid */
.section-services-grid {
  padding: 8rem 0;
  background-color: #f9fbf9;
}

.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(35rem, 1fr));
  gap: 4rem;
}

/* Special Case: Oncology */
:deep(.service-card--oncology) {
  border-bottom: 5px solid #a29bfe !important;
}

:deep(.service-card--oncology:hover) {
  box-shadow: 0 2rem 5rem rgba(162, 155, 254, 0.2) !important;
}

:deep(.service-card--oncology .service-card__icon) {
  color: #6c5ce7;
}

/* Philosophy Section */
.section-philosophy {
  padding: 12rem 0;
  background-color: #fff;
}

.philosophy-box {
  background: #fcfdfe;
  padding: 8rem;
  border-radius: 40px;
  border: 1px solid #f0f0f0;
  box-shadow: 0 2rem 5rem rgba(0,0,0,0.02);
  text-align: center;
}

.philosophy-title {
  font-family: 'Playfair Display', serif;
  font-size: 4rem;
  font-weight: 900;
  color: #333;
  margin-bottom: 3rem;
}

.philosophy-title span {
  display: block;
  font-style: italic;
  font-weight: 400;
  font-size: 3rem;
  color: #55c57a;
  margin-top: 1rem;
}

.philosophy-text {
  font-size: 2rem;
  color: #666;
  line-height: 1.8;
  max-width: 90rem;
  margin: 0 auto 5rem;
}

.philosophy-features {
  display: flex;
  justify-content: center;
  gap: 6rem;
  flex-wrap: wrap;
}

.p-feat {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
  max-width: 20rem;
}

.p-feat i {
  font-size: 4rem;
  color: #55c57a;
  margin-bottom: 1rem;
}

.p-feat span {
  font-size: 1.5rem;
  font-weight: 700;
  color: #444;
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 10rem 0;
}

.empty-state__icon {
  font-size: 8rem;
  color: #ddd;
  margin-bottom: 2rem;
}

.empty-state h3 {
  font-size: 2.4rem;
  color: #333;
  margin-bottom: 1rem;
}

.empty-state p {
  font-size: 1.6rem;
  color: #777;
  margin-bottom: 3rem;
}

.btn-reset {
  padding: 1.2rem 3rem;
  background: #55c57a;
  color: #fff;
  border: none;
  border-radius: 50px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s;
}

.btn-reset:hover { background: #40945c; }

@media (max-width: 600px) {
  .heading-primary { font-size: 3.5rem; }
  .filter-box { padding: 2rem; }
  .category-list { justify-content: center; }
}
</style>
