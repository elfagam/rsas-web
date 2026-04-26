<script setup>
import { ref, onMounted, computed } from 'vue';
import HeroSection from '../components/ui/HeroSection.vue';
import ServiceCard from '../components/ui/ServiceCard.vue';
import DoctorCard from '../components/ui/DoctorCard.vue';
import NewsCard from '../components/ui/NewsCard.vue';
import TestimonialSection from '../components/ui/TestimonialSection.vue';
import TestimonialModal from '../components/ui/TestimonialModal.vue';

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

// Helper untuk resolusi gambar aman
const getDoctorImage = (name) => {
  try {
    if (!name) return new URL('../assets/img/direktur-aloeisaboe.jpg', import.meta.url).href;
    return new URL(`../assets/img/${name}`, import.meta.url).href;
  } catch (e) {
    return new URL('../assets/img/direktur-aloeisaboe.jpg', import.meta.url).href;
  }
};

// Data Dokter (Bisa dipindah ke file JSON nanti)
const daftarDokter = ref([
  {
    id: 1,
    name: 'dr. H. Mohammad Ali, Sp.PD',
    specialty: 'Penyakit Dalam',
    image: 'direktur-aloeisaboe.jpg',
    time: '08:00 - 12:00 WITA',
    isAvailable: true,
    availableDays: ['Senin', 'Rabu', 'Jumat'],
    quota: 20,
    remaining: 8
  },
  {
    id: 2,
    name: 'dr. Siti Aminah, Sp.A',
    specialty: 'Klinik Anak',
    image: 'direktur-aloeisaboe.jpg',
    time: '13:00 - 16:00 WITA',
    isAvailable: false,
    availableDays: ['Selasa', 'Kamis', 'Sabtu'],
    quota: 15,
    remaining: 0
  },
  {
    id: 3,
    name: 'dr. Bambang S., Sp.B',
    specialty: 'Bedah Umum',
    image: 'direktur-aloeisaboe.jpg',
    time: '09:00 - 13:00 WITA',
    isAvailable: true,
    availableDays: ['Senin', 'Selasa', 'Kamis'],
    quota: 10,
    remaining: 4
  }
]);

// Daftar Poli untuk Dropdown
const daftarPoli = [
  'Penyakit Dalam',
  'Klinik Anak',
  'Bedah Umum',
  'Kandungan',
  'Mata'
];

// Logika Filter Dokter
const searchDokter = ref('');
const selectedDate = ref('');
const selectedPoli = ref('');

const filteredDokter = computed(() => {
  let filtered = daftarDokter.value;

  // Filter berdasarkan Nama
  if (searchDokter.value) {
    const search = searchDokter.value.toLowerCase();
    filtered = filtered.filter(dokter => 
      dokter.name.toLowerCase().includes(search)
    );
  }

  // Filter berdasarkan Poli
  if (selectedPoli.value) {
    filtered = filtered.filter(dokter => 
      dokter.specialty === selectedPoli.value
    );
  }

  // Filter berdasarkan Tanggal (Hari)
  if (selectedDate.value) {
    const days = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu'];
    const date = new Date(selectedDate.value);
    const dayName = days[date.getDay()];
    
    filtered = filtered.filter(dokter => 
      dokter.availableDays.includes(dayName)
    );
  }

  return filtered;
});

// Data Berita (Simulasi)
const beritaTerbaru = ref([
  {
    id: 1,
    kategori: 'PENGUMUMAN',
    tanggal: '26 April 2026',
    judul: 'Lapor Gratifikasi: Kenali dan Hindari',
    ringkasan: 'RSUD Prof. Dr. H. Aloei Saboe berkomitmen penuh dalam pemberantasan korupsi...',
    gambar: getImageUrl('nat-1.jpg'),
    tautan: '/berita/1'
  },
  {
    id: 2,
    kategori: 'LAYANAN',
    tanggal: '25 April 2026',
    judul: 'Peresmian Gedung Rawat Inap Baru',
    ringkasan: 'Untuk meningkatkan kenyamanan pasien, manajemen RSAS telah meresmikan fasilitas rawat inap VIP...',
    gambar: getImageUrl('nat-2.jpg'),
    tautan: '/berita/2'
  },
  {
    id: 3,
    kategori: 'EDUKASI',
    tanggal: '24 April 2026',
    judul: 'Waspada Hipertensi: Si Pembunuh Senyap',
    ringkasan: 'Kenali gejala dini hipertensi dan cara mengelolanya melalui pola makan sehat dan olahraga rutin untuk mencegah risiko stroke...',
    gambar: getImageUrl('nat-3.jpg'),
    tautan: '/berita/3'
  }
]);

// Data Sejarah
const sejarahRSAS = ref([
  { 
    tahun: '1926', 
    deskripsi: 'Pembangunan RSU Kotamadya Gorontalo sebagai langkah awal pelayanan kesehatan modern di wilayah ini.',
    posisi: 'fade-right'
  },
  { 
    tahun: '1929', 
    deskripsi: 'Pemanfaatan secara resmi dengan nama RSU Kotamadya Gorontalo untuk melayani masyarakat umum.',
    posisi: 'fade-left'
  },
  { 
    tahun: '1979', 
    deskripsi: 'Berubah nama menjadi RSU Prof. Dr. H. Aloei Saboe Gorontalo sebagai penghormatan terhadap tokoh kesehatan daerah.',
    posisi: 'fade-right'
  },
  { 
    tahun: '2005', 
    deskripsi: 'Ditetapkan sebagai Rumah Sakit Kelas B dan menjadi Badan Layanan Umum Daerah (BLUD).',
    posisi: 'fade-left'
  }
]);

const isModalOpen = ref(false);

onMounted(() => {
  // Inisialisasi swiper jika diperlukan
});
</script>

<template>
  <div class="home-page">
    <!-- Hero Section -->
    <HeroSection />

    <!-- Director's Message Section -->
    <section class="section-director">
      <div class="container">
        <div class="director-box" data-aos="fade-up">
          <div class="director-box__img-container">
            <img 
              :src="getImageUrl('direktur-aloeisaboe.jpg')" 
              alt="Direktur RSAS" 
              class="director-box__img"
            >
            <div class="director-box__badge">Direktur Utama</div>
          </div>
          <div class="director-box__content">
            <h4 class="director-box__subtitle">Sambutan Pimpinan</h4>
            <h2 class="director-box__title">Melayani dengan Hati, <br>Membangun Gorontalo Sehat</h2>
            <div class="director-box__divider"></div>
            <p class="director-box__text">
              "Selamat datang di portal digital RSUD Prof. Dr. H. Aloei Saboe. Sebagai rumah sakit rujukan utama dan rumah sakit pendidikan, kami berkomitmen untuk selalu menghadirkan inovasi medis terbaik yang dibalut dengan keramahan dan empati. Teknologi hanyalah alat, namun kemanusiaan adalah esensi dari kesembuhan."
            </p>
            <div class="director-box__info">
              <span class="director-box__name">dr. H. Ahmad Lihu, M.Si</span>
              <span class="director-box__nip">NIP. 19700101 199503 1 001</span>
            </div>
            <router-link to="/profil" class="btn-text">Selengkapnya tentang kami &rarr;</router-link>
          </div>
        </div>
      </div>
    </section>

    <!-- Services Section -->
    <section class="section-services" id="layanan">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big">
          <h2 class="heading-secondary">Layanan <span>Unggulan Kami</span></h2>
        </div>
        <div class="services-grid">
          <ServiceCard 
            icon="fa-solid fa-hospital-user" 
            title="IGD 24 Jam Terpadu" 
            text="Pelayanan gawat darurat cepat tanggap dengan fasilitas medis terkini dan tenaga ahli bersertifikat."
            :delay="0"
          />
          <ServiceCard 
            icon="fa-solid fa-heart-pulse" 
            title="Pusat Jantung Terpadu" 
            text="Penanganan komprehensif kardiovaskular didukung oleh dokter spesialis dan alat kateterisasi jantung."
            link="/layanan/jantung"
            :delay="150"
          />
          <ServiceCard 
            icon="fa-solid fa-microscope" 
            title="Laboratorium Canggih" 
            text="Hasil diagnosa cepat dan akurat sebagai penunjang utama kesembuhan pasien, beroperasi penuh 24 jam."
            :delay="300"
          />
          <ServiceCard 
            icon="fa-solid fa-hand-holding-heart" 
            title="Pusat Onkologi (Kanker)" 
            text="Kami mendampingi setiap langkah perjuangan pasien dengan filosofi 'Keluarga Melayani Keluarga' dan kasih sayang tanpa batas."
            :delay="450"
            class="service-card--oncology"
          />
        </div>
      </div>
    </section>

    <!-- Schedule Section -->
    <section class="section-schedule" id="jadwal-dokter">
      <div class="container">
        <div class="u-center-text u-margin-bottom-medium">
          <h2 class="heading-secondary">Jadwal <span>Dokter Spesialis</span></h2>
          <p class="heading-tertiary">
            Pilih tanggal atau cari nama dokter untuk melihat jadwal praktik.
          </p>
        </div>

        <div class="schedule-filter" data-aos="fade-up">
          <div class="filter-container">
            <div class="filter-group filter-group--search">
              <i class="fa-solid fa-magnifying-glass filter-icon"></i>
              <input 
                v-model="searchDokter"
                type="text" 
                class="filter-input" 
                placeholder="Cari Nama Dokter..." 
              />
            </div>

            <div class="filter-group filter-group--select">
              <i class="fa-solid fa-stethoscope filter-icon"></i>
              <select v-model="selectedPoli" class="filter-input">
                <option value="">Semua Poli</option>
                <option v-for="poli in daftarPoli" :key="poli" :value="poli">{{ poli }}</option>
              </select>
            </div>

            <div class="filter-group filter-group--date">
              <i class="fa-solid fa-calendar-days filter-icon"></i>
              <input 
                v-model="selectedDate"
                type="date" 
                class="filter-input" 
              />
            </div>
            <button class="btn-reset" v-if="searchDokter || selectedDate || selectedPoli" @click="searchDokter = ''; selectedDate = ''; selectedPoli = ''">
               Reset
            </button>
          </div>
        </div>

        <div class="schedule-grid" v-if="filteredDokter.length > 0">
          <DoctorCard 
            v-for="dokter in filteredDokter" 
            :key="dokter.id"
            v-bind="dokter"
            :image="getDoctorImage(dokter.image)"
            data-aos="fade-up"
          />
        </div>

        <div class="u-center-text u-margin-top-medium" v-else>
           <p style="font-size: 1.8rem; color: #777;">
             <i class="fa-solid fa-user-slash"></i> Maaf, Dokter atau Poli tidak ditemukan.
           </p>
        </div>
      </div>
    </section>

    <!-- Information Posters Section -->
    <section class="section-info" id="informasi">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big">
          <h2 class="heading-secondary">Informasi Pasien</h2>
        </div>
        <div class="info-grid">
          <div class="info-card" data-aos="zoom-in">
            <div class="info-card__badge">Antrean Online</div>
            <div class="info-card__content">
              <h3>Mudahnya Pakai Mobile JKN</h3>
              <p>Download aplikasi Mobile JKN untuk pendaftaran online, cek kepesertaan, dan kemudahan akses layanan kesehatan lainnya.</p>
              <ul class="info-card__list">
                <li><i class="fa-solid fa-check"></i> Ambil nomor antrean dari rumah</li>
                <li><i class="fa-solid fa-check"></i> Cek ketersediaan tempat tidur</li>
                <li><i class="fa-solid fa-check"></i> Konsultasi dokter online</li>
              </ul>
              <button class="btn btn-green">Panduan Download</button>
            </div>
          </div>
          <div class="info-card info-card--blue" data-aos="zoom-in" data-delay="100">
            <div class="info-card__badge info-card__badge--blue">Pendaftaran</div>
            <div class="info-card__content">
              <h3>Syarat Pendaftaran Pasien</h3>
              <p>Pastikan Anda membawa dokumen berikut untuk mempercepat proses administrasi di loket pendaftaran:</p>
              <div class="info-card__tabs">
                <div class="tab"><strong>BPJS:</strong> Kartu JKN/KIS, KTP, Surat Rujukan (jika ada)</div>
                <div class="tab"><strong>UMUM:</strong> KTP / Kartu Keluarga</div>
              </div>
              <button class="btn btn-blue">Lihat Alur Layanan</button>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- IKM Section -->
    <section class="section-ikm">
      <div class="ikm-wave">
        <svg viewBox="0 0 1440 320" preserveAspectRatio="none"><path fill="#f7fbf8" fill-opacity="1" d="M0,192L80,176C160,160,320,128,480,144C640,160,800,224,960,229.3C1120,235,1280,181,1360,154.7L1440,128L1440,0L1360,0C1280,0,1120,0,960,0C800,0,640,0,480,0C320,0,160,0,80,0L0,0Z"></path></svg>
      </div>
      <div class="container u-center-text">
        <h2 class="ikm-title" data-aos="fade-up">Indikator Kepuasan Masyarakat (IKM)</h2>
        <p class="ikm-year" data-aos="fade-up">Periode Tahun 2025</p>
        <div class="ikm-number" data-aos="zoom-in">87,38<span>%</span></div>
        <div class="ikm-status" data-aos="fade-up">Sangat Baik</div>
      </div>
    </section>

    <!-- Mitra Slider Section -->
    <section class="section-mitra">
       <div class="container container--full">
         <div class="u-center-text u-margin-bottom-small">
           <p class="mitra-subtitle">RSUD Prof. Dr. H. Aloei Saboe telah bekerjasama dengan:</p>
         </div>
         <div class="mitra-slider">
           <div class="mitra-track">
             <!-- Kita ulangi seluruh rangkaian logo 2 kali untuk efek loop tanpa putus -->
             <template v-for="loop in 2" :key="loop">
               <img :src="getImageUrl('bpjs.png')" alt="BPJS Kesehatan" class="mitra-logo" />
               <img :src="getImageUrl('mandiri-inhealth.png')" alt="Mandiri Inhealth" class="mitra-logo" />
               <img :src="getImageUrl('taspen.png')" alt="Taspen" class="mitra-logo" />
               <img :src="getImageUrl('jasa-raharja.png')" alt="Jasa Raharja" class="mitra-logo" />
               <img :src="getImageUrl('bpsj-ketenagakerjaan.png')" alt="BPJS Ketenagakerjaan" class="mitra-logo" />
             </template>
           </div>
         </div>
       </div>
    </section>

    <!-- Visitor Stats Section -->
    <section class="section-stats">
      <div class="container">
        <div class="u-center-text u-margin-bottom-medium">
          <h2 class="heading-secondary">Statistik Pengunjung</h2>
        </div>
        <div class="stats-grid">
          <div class="stat-box" data-aos="fade-up">
            <span class="stat-label">Hari Ini</span>
            <span class="stat-value">48</span>
          </div>
          <div class="stat-box" data-aos="fade-up" data-aos-delay="50">
            <span class="stat-label">Minggu Ini</span>
            <span class="stat-value">535</span>
          </div>
          <div class="stat-box" data-aos="fade-up" data-aos-delay="100">
            <span class="stat-label">Bulan Ini</span>
            <span class="stat-value">584</span>
          </div>
          <div class="stat-box" data-aos="fade-up" data-aos-delay="150">
            <span class="stat-label">Tahun Ini</span>
            <span class="stat-value">5931</span>
          </div>
        </div>
      </div>
    </section>

    <!-- History Section -->
    <section class="section-history" id="profil">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big">
          <h2 class="heading-secondary">Sejarah Singkat RSAS</h2>
        </div>
        
        <div class="timeline">
          <div 
            v-for="(item, index) in sejarahRSAS" 
            :key="index"
            class="timeline-item"
            :class="index % 2 === 0 ? 'timeline-item--left' : 'timeline-item--right'"
            :data-aos="item.posisi"
          >
            <div class="timeline-content">
              <span class="year">{{ item.tahun }}</span>
              <p class="timeline-desc">{{ item.deskripsi }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Latest News Section -->
    <section class="section-news" id="berita">
      <div class="container">
        <div class="u-center-text u-margin-bottom-big">
          <h2 class="heading-secondary">Berita Terkini</h2>
        </div>
        <div class="news-home-grid">
          <NewsCard 
            v-for="berita in beritaTerbaru" 
            :key="berita.id"
            v-bind="berita"
          />
        </div>
        <div class="u-center-text u-margin-top-medium">
          <router-link to="/berita" class="btn-text">Lihat Semua Berita &rarr;</router-link>
        </div>
      </div>
    </section>

    <!-- Testimonials Section -->
    <TestimonialSection @open-modal="isModalOpen = true" />

    <!-- Testimonial Modal Form -->
    <TestimonialModal :show="isModalOpen" @close="isModalOpen = false" />
  </div>
</template>

<style scoped>
.container {
  max-width: 1240px;
  margin: 0 auto;
  padding: 0 2rem;
}

.container--full {
  max-width: 100%;
  padding: 0;
}

.u-center-text { text-align: center; }
.u-margin-bottom-big { margin-bottom: 8rem; }
.u-margin-bottom-medium { margin-bottom: 4rem; }
.u-margin-bottom-small { margin-bottom: 2rem; }
.u-margin-top-medium { margin-top: 4rem; }

.section-director {
  padding: 18rem 0 10rem;
  background-color: #fff;
  overflow: hidden;
}

.director-box {
  display: flex;
  align-items: center;
  gap: 8rem;
  max-width: 110rem;
  margin: 0 auto;
}

.director-box__img-container {
  flex: 0 0 40rem;
  position: relative;
}

.director-box__img {
  width: 100%;
  height: 50rem;
  object-fit: cover;
  border-radius: 30px;
  box-shadow: 0 3rem 6rem rgba(0, 0, 0, 0.1);
}

.director-box__badge {
  position: absolute;
  bottom: 3rem;
  right: -3rem;
  background-color: #55c57a;
  color: #fff;
  padding: 1.5rem 3rem;
  border-radius: 15px;
  font-weight: 800;
  font-size: 1.4rem;
  text-transform: uppercase;
  letter-spacing: 2px;
  box-shadow: 0 1.5rem 3rem rgba(85, 197, 122, 0.3);
}

.director-box__content {
  flex: 1;
}

.director-box__subtitle {
  font-size: 1.4rem;
  color: #55c57a;
  text-transform: uppercase;
  font-weight: 800;
  letter-spacing: 3px;
  margin-bottom: 1rem;
}

.director-box__title {
  font-size: 2.8rem;
  line-height: 1.2;
  color: #333;
  margin-bottom: 3rem;
}

.director-box__divider {
  width: 6rem;
  height: 4px;
  background-color: #55c57a;
  margin-bottom: 3rem;
}

.director-box__text {
  font-size: 1.3rem;
  font-style: italic;
  color: #666;
  line-height: 1.8;
  margin-bottom: 4rem;
  position: relative;
}

.director-box__info {
  display: flex;
  flex-direction: column;
  margin-bottom: 4rem;
}

.director-box__name {
  font-size: 2.2rem;
  font-weight: 800;
  color: #333;
}

.director-box__nip {
  font-size: 1.4rem;
  color: #999;
  font-weight: 600;
}

@media (max-width: 1000px) {
  .director-box { flex-direction: column; gap: 6rem; text-align: center; }
  .director-box__img-container { flex: 0 0 auto; width: 100%; max-width: 40rem; }
  .director-box__badge { right: 50%; transform: translateX(50%) translateY(50%); bottom: 0; }
  .director-box__divider { margin: 0 auto 3rem; }
}

/* Khusus Kartu Onkologi */
.service-card--oncology {
  border-bottom: 4px solid #a29bfe !important; /* Warna Lavender Lembut */
  transition: all 0.4s ease;
}

.service-card--oncology:hover {
  box-shadow: 0 2rem 5rem rgba(162, 155, 254, 0.2) !important;
  transform: translateY(-1.5rem) !important;
}

:deep(.service-card--oncology .service-card__icon) {
  color: #6c5ce7; /* Warna Ungu Onkologi */
}

.section-services {
  padding: 10rem 0 10rem; /* Dikurangi dari 15rem agar lebih mengalir */
}

.section-schedule, .section-history, .section-news, .section-info, .section-stats {
  padding: 10rem 0;
}

/* Info Cards */
.info-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
}

.info-card {
  background: white;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 2rem 5rem rgba(0, 0, 0, 0.05);
  position: relative;
  border-bottom: 5px solid #55c57a;
  transition: transform 0.3s;
}

.info-card:hover { transform: translateY(-10px); }

.info-card--blue { border-bottom-color: #2980b9; }

.info-card__badge {
  position: absolute;
  top: 2rem;
  right: 2rem;
  background: #55c57a;
  color: white;
  padding: 0.5rem 1.5rem;
  border-radius: 50px;
  font-size: 1.2rem;
  font-weight: 700;
}

.info-card__badge--blue { background: #2980b9; }

.info-card__content {
  padding: 5rem 4rem;
}

.info-card__content h3 {
  font-size: 2.8rem;
  margin-bottom: 2rem;
  color: #333;
}

.info-card__content p {
  font-size: 1.6rem;
  margin-bottom: 2.5rem;
  line-height: 1.7;
  color: #666;
}

.info-card__list {
  list-style: none;
  margin-bottom: 3rem;
}

.info-card__list li {
  font-size: 1.5rem;
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  color: #555;
}

.info-card__list i { color: #55c57a; }

.info-card__tabs {
  background: #f9f9f9;
  padding: 2rem;
  border-radius: 10px;
  margin-bottom: 3rem;
}

.tab { font-size: 1.4rem; margin-bottom: 1rem; }

.btn {
  padding: 1.5rem 3rem;
  border-radius: 10px;
  border: none;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s;
}

.btn-green { background: #55c57a; color: white; }
.btn-green:hover { background: #409d61; }

.btn-blue { background: #2980b9; color: white; }
.btn-blue:hover { background: #1f6391; }

/* IKM Section */
.section-ikm {
  background-color: #f7fbf8;
  padding: 5rem 0 10rem;
  position: relative;
}

.ikm-wave {
  position: absolute;
  top: -10rem;
  left: 0;
  width: 100%;
  height: 10rem;
}

.ikm-title { font-size: 3rem; margin-bottom: 1rem; color: #333; font-weight: 900; }
.ikm-year { font-size: 1.8rem; color: #55c57a; font-weight: 700; margin-bottom: 3rem; }
.ikm-number { font-size: 12rem; font-weight: 950; color: #333; line-height: 1; margin-bottom: 1rem; letter-spacing: -5px; }
.ikm-number span { font-size: 5rem; color: #55c57a; margin-left: -1rem; }
.ikm-status { font-size: 2.5rem; text-transform: uppercase; letter-spacing: 5px; color: #55c57a; font-weight: 800; }

/* Mitra Slider */
.section-mitra { background: #fff; padding: 6rem 0; overflow: hidden; }

.mitra-subtitle { font-size: 1.4rem; color: #888; text-transform: uppercase; letter-spacing: 2px; font-weight: 700; }

.mitra-slider {
  display: flex;
  width: 100%;
  padding: 2rem 0;
}

.mitra-track {
  display: flex;
  width: calc(250px * 20);
  animation: scroll 40s linear infinite;
  align-items: center;
  gap: 8rem;
}

.mitra-logo {
  height: 60px;
  filter: grayscale(1) opacity(0.5);
  transition: all 0.3s;
}

.mitra-logo:hover { filter: grayscale(0) opacity(1); transform: scale(1.1); }

@keyframes scroll {
  0% { transform: translateX(0); }
  100% { transform: translateX(calc(-250px * 10)); }
}

/* Stats Grid */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 3rem;
}

.stat-box {
  background: #f7f7f7;
  padding: 4rem 2rem;
  border-radius: 15px;
  text-align: center;
}

.stat-label { display: block; font-size: 1.4rem; color: #888; text-transform: uppercase; margin-bottom: 1rem; font-weight: 700; }
.stat-value { font-size: 4rem; font-weight: 900; color: #333; }

/* Existing Home Styles */
.section-schedule { background-color: #f7f7f7; }

.schedule-filter {
  max-width: 100rem;
  margin: 0 auto 6rem;
  width: 100%;
}

.filter-container {
  display: flex;
  gap: 0;
  align-items: center;
  background: white;
  padding: 0.5rem 1rem;
  border-radius: 100px;
  box-shadow: 0 1.5rem 4rem rgba(0, 0, 0, 0.1);
  border: 1px solid #f1f1f1;
}

.filter-group { display: flex; align-items: center; padding: 1.5rem 3rem; height: 100%; transition: all 0.3s; }
.filter-group:hover { background-color: #fcfcfc; }
.filter-group:first-child { border-radius: 100px 0 0 100px; }
.filter-icon { font-size: 2rem; color: #55c57a; margin-right: 1.5rem; }
.filter-input { width: 100%; border: none; font-size: 1.5rem; outline: none; background: transparent; font-family: inherit; color: #444; font-weight: 500; }

.filter-group--search { flex: 1.5; border-right: 1px solid #eee; }
.filter-group--select { flex: 1.2; border-right: 1px solid #eee; }
.filter-group--date { flex: 1; }

.btn-reset { margin-right: 1.5rem; background: #f1f2f6; border: none; padding: 1rem 2.5rem; border-radius: 50px; font-size: 1.2rem; font-weight: 800; color: #eb4d4b; cursor: pointer; transition: all 0.3s; white-space: nowrap; }
.btn-reset:hover { background: #eb4d4b; color: white; box-shadow: 0 0.5rem 1.5rem rgba(235, 77, 75, 0.3); }

.schedule-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(30rem, 1fr)); gap: 4rem; }

/* Timeline Styles */
.timeline { position: relative; max-width: 100rem; margin: 0 auto; }
.timeline::after { content: ''; position: absolute; width: 2px; background-color: #55c57a; top: 0; bottom: 0; left: 50%; margin-left: -1px; }
.timeline-item { padding: 1rem 4rem; position: relative; background-color: inherit; width: 50%; }
.timeline-item::after { content: ''; position: absolute; width: 2rem; height: 2rem; right: -1rem; background-color: white; border: 4px solid #55c57a; top: 1.5rem; border-radius: 50%; z-index: 1; }
.timeline-item--left { left: 0; }
.timeline-item--right { left: 50%; }
.timeline-item--right::after { left: -1rem; }
.timeline-content { padding: 2rem 3rem; background-color: white; position: relative; border-radius: 12px; box-shadow: 0 1rem 3rem rgba(0, 0, 0, 0.1); }
.year { display: block; font-size: 2.2rem; font-weight: 900; color: #55c57a; margin-bottom: 1rem; }
.timeline-desc { font-size: 1.5rem; line-height: 1.6; color: #777; }

/* Latest News */
.news-home-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 4rem; }
.btn-text { font-size: 1.6rem; color: #55c57a; text-decoration: none; border-bottom: 1px solid #55c57a; padding: 3px; transition: all 0.2s; font-weight: 700; }
.btn-text:hover { background-color: #55c57a; color: #fff; box-shadow: 0 1rem 2rem rgba(0, 0, 0, 0.15); transform: translateY(-2px); }

/* Responsive */
@media (max-width: 1000px) {
  .info-grid { grid-template-columns: 1fr; }
  .stats-grid { grid-template-columns: 1fr 1fr; }
  .news-home-grid { grid-template-columns: 1fr; }
}

@media (max-width: 768px) {
  .ikm-number { font-size: 8rem; }
  .timeline::after { left: 3.1rem; }
  .timeline-item { width: 100%; padding-left: 7rem; padding-right: 2.5rem; }
  .timeline-item::after { left: 2.1rem; }
  .timeline-item--right { left: 0%; }
}
</style>
