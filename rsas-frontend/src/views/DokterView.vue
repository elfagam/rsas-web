<script setup>
import { ref, computed, onMounted } from 'vue';
import SectionHeader from '../components/ui/SectionHeader.vue';
import DoctorCard from '../components/ui/DoctorCard.vue';
import { ASSETS } from '../config/assets';

const searchQuery = ref('');
const activeSpecialty = ref('Semua');

const specialties = [
  'Semua', 
  'Spesialis Jantung', 
  'Spesialis Anak', 
  'Spesialis Penyakit Dalam', 
  'Spesialis Bedah', 
  'Spesialis Kandungan'
];

const allDoctors = ref([
  { id: 1, name: 'dr. Ahmad Fachri, Sp.JP(K)', specialty: 'Spesialis Jantung', expertise: 'Konsultan Jantung Intervensi', time: '08:00 - 12:00', image: 'direktur-aloeisaboe.jpg', isAvailable: true, quota: 20, remaining: 5 },
  { id: 2, name: 'dr. Sarah Az-Zahra, Sp.A', specialty: 'Spesialis Anak', expertise: 'Konsultan Tumbuh Kembang', time: '09:00 - 13:00', image: 'direktur-aloeisaboe.jpg', isAvailable: true, quota: 25, remaining: 10 },
  { id: 3, name: 'dr. Budi Santoso, Sp.PD', specialty: 'Spesialis Penyakit Dalam', expertise: 'Konsultan Gastroentero-Hepatologi', time: '13:00 - 16:00', image: 'direktur-aloeisaboe.jpg', isAvailable: false, quota: 15, remaining: 0 },
  { id: 4, name: 'dr. Linda Wijaya, Sp.OG', specialty: 'Spesialis Kandungan', expertise: 'Konsultan Fertilitas Endokrinologi', time: '08:00 - 11:00', image: 'direktur-aloeisaboe.jpg', isAvailable: true, quota: 20, remaining: 18 },
  { id: 5, name: 'dr. Hendra Kusuma, Sp.B', specialty: 'Spesialis Bedah', expertise: 'Spesialis Bedah Onkologi', time: '10:00 - 14:00', image: 'direktur-aloeisaboe.jpg', isAvailable: true, quota: 10, remaining: 2 },
  { id: 6, name: 'dr. Siti Aminah, Sp.A', specialty: 'Spesialis Anak', expertise: 'Spesialis Alergi Imunologi', time: '14:00 - 17:00', image: 'direktur-aloeisaboe.jpg', isAvailable: true, quota: 20, remaining: 15 },
  { id: 7, name: 'dr. Robert Chen, Sp.JP', specialty: 'Spesialis Jantung', expertise: 'Konsultan Jantung Anak', time: '15:00 - 19:00', image: 'direktur-aloeisaboe.jpg', isAvailable: true, quota: 15, remaining: 8 },
]);

const getDoctorImage = (imgName) => {
  if (!imgName) return null;
  try {
    return new URL(`../assets/img/${imgName}`, import.meta.url).href;
  } catch (e) {
    console.error("Error loading image:", imgName, e);
    return null;
  }
};

const filteredDoctors = computed(() => {
  return allDoctors.value.filter(doc => {
    const matchesSearch = doc.name.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchesSpecialty = activeSpecialty.value === 'Semua' || doc.specialty === activeSpecialty.value;
    return matchesSearch && matchesSpecialty;
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
  <div class="dokter-view">
    <!-- Hero Section -->
    <header 
      class="dokter-hero"
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.dokter})` }"
    >
      <div class="container">
        <div class="dokter-hero__content" data-aos="fade-up">
          <div class="dokter-hero__badge">Tim Medis RSAS</div>
          <h1 class="heading-primary">
            <span class="heading-primary__main">Temukan Dokter</span>
            <span class="heading-primary__sub">Pilihan Anda</span>
          </h1>
          <p class="dokter-hero__sub">Pelayanan kesehatan profesional yang didedikasikan untuk kesembuhan dan kenyamanan Anda sekeluarga.</p>
        </div>
      </div>
    </header>

    <!-- Filter Section -->
    <section class="section-filter">
      <div class="container">
        <div class="filter-box">
          <div class="filter-row">
            <div class="search-group">
              <i class="fa-solid fa-magnifying-glass search-icon"></i>
              <input 
                type="text" 
                v-model="searchQuery" 
                placeholder="Cari nama dokter..." 
                class="search-input"
              />
            </div>
            
            <div class="select-group">
              <select v-model="activeSpecialty" class="specialty-select">
                <option v-for="spec in specialties" :key="spec" :value="spec">
                  {{ spec }}
                </option>
              </select>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Doctors Grid -->
    <section class="section-doctors-grid">
      <div class="container">
        <div v-if="filteredDoctors.length > 0" class="doctors-grid">
          <DoctorCard 
            v-for="doc in filteredDoctors" 
            :key="doc.id"
            v-bind="doc"
            :image="getDoctorImage(doc.image)"
          />
        </div>

        <!-- Empty State -->
        <div v-else class="empty-state" data-aos="fade-up">
          <div class="empty-state__icon">
            <i class="fa-solid fa-user-slash"></i>
          </div>
          <h3>Dokter tidak ditemukan</h3>
          <p>Maaf, dokter dengan kriteria tersebut tidak ditemukan. Silakan coba kata kunci lain.</p>
          <button @click="searchQuery = ''; activeSpecialty = 'Semua'" class="btn-reset">Lihat Semua Dokter</button>
        </div>
      </div>
    </section>

    <!-- Appointment Info CTA -->
    <section class="section-info">
      <div class="container">
        <div class="info-card">
          <div class="info-card__icon"><i class="fa-solid fa-circle-info"></i></div>
          <div class="info-card__content">
            <h3>Informasi Pendaftaran</h3>
            <p>Pendaftaran online dapat dilakukan h-1 sebelum jadwal praktik dokter. Pastikan Anda membawa kartu identitas dan kartu berobat saat datang ke rumah sakit.</p>
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

.dokter-view {
  /* Padding top dihapus agar Hero menyentuh bagian paling atas layar */
  min-height: 100vh;
}

/* Hero */
.dokter-hero {
  height: 60vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
  text-align: center;
}

.dokter-hero__badge {
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

.dokter-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  margin: 0 auto;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
  line-height: 1.6;
}

/* Filter Box */
.section-filter {
  margin-top: -4rem;
  position: relative;
  z-index: 10;
}

.filter-box {
  background: #fff;
  padding: 3rem;
  border-radius: 15px;
  box-shadow: 0 2rem 5rem rgba(0,0,0,0.1);
}

.filter-row {
  display: flex;
  gap: 2rem;
}

.search-group {
  position: relative;
  flex: 2;
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
  padding: 1.5rem 2rem 1.5rem 5.5rem;
  border: 1px solid #eee;
  border-radius: 10px;
  font-size: 1.6rem;
  transition: all 0.3s;
}

.search-input:focus {
  outline: none;
  border-color: #55c57a;
}

.select-group {
  flex: 1;
}

.specialty-select {
  width: 100%;
  padding: 1.5rem 2rem;
  border: 1px solid #eee;
  border-radius: 10px;
  font-size: 1.6rem;
  color: #333;
  cursor: pointer;
  background-color: #fff;
}

/* Grid */
.section-doctors-grid {
  padding: 8rem 0;
  background-color: #f9fbf9;
}

.doctors-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(35rem, 1fr));
  gap: 3rem;
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 6rem 0;
}

.empty-state__icon {
  font-size: 6rem;
  color: #ddd;
  margin-bottom: 2rem;
}

.btn-reset {
  margin-top: 2rem;
  padding: 1.2rem 3rem;
  background: #55c57a;
  color: #fff;
  border: none;
  border-radius: 50px;
  font-weight: 700;
  cursor: pointer;
}

/* Info Card */
.section-info {
  padding: 6rem 0;
}

.info-card {
  display: flex;
  align-items: center;
  gap: 3rem;
  padding: 3rem;
  background: #fff;
  border-left: 5px solid #55c57a;
  border-radius: 10px;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
}

.info-card__icon {
  font-size: 3rem;
  color: #55c57a;
}

.info-card__content h3 { font-size: 2rem; margin-bottom: 0.5rem; }
.info-card__content p { font-size: 1.5rem; color: #777; }

@media (max-width: 800px) {
  .filter-row { flex-direction: column; }
  .heading-primary { font-size: 3rem; }
}
</style>
