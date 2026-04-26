<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import SectionHeader from '../components/ui/SectionHeader.vue';
import StepFlow from '../components/ui/StepFlow.vue';
import DoctorCard from '../components/ui/DoctorCard.vue';
import { ASSETS } from '../config/assets';

// Helper untuk resolusi gambar dinamis
const getImageUrl = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

const route = useRoute();
const isLoading = ref(true);

const servicesData = {
  'jantung': {
    name: 'Pusat Jantung Terpadu',
    heroImg: ASSETS.hero.layanan,
    desc: 'Layanan kardiovaskular komprehensif mulai dari diagnostik non-invasif hingga kateterisasi jantung (Cath Lab) dengan standar internasional.',
    features: [
      { title: 'Catheterization Lab', icon: 'fa-solid fa-heart-pulse', desc: 'Laboratorium kateterisasi untuk tindakan intervensi jantung non-bedah.' },
      { title: 'Echocardiography', icon: 'fa-solid fa-wave-square', desc: 'USG Jantung 4D untuk melihat struktur dan fungsi jantung secara detail.' },
      { title: 'Treadmill Test', icon: 'fa-solid fa-person-running', desc: 'Uji beban jantung untuk mendeteksi gangguan irama dan aliran darah.' }
    ],
    steps: [
      { title: 'Pendaftaran', desc: 'Pendaftaran di loket pusat jantung atau via Mobile JKN.' },
      { title: 'Pemeriksaan', desc: 'Pemeriksaan tanda vital dan konsultasi dokter spesialis.' },
      { title: 'Tindakan/Resep', desc: 'Tindakan diagnostik atau pemberian resep obat.' }
    ],
    doctors: [
      { id: 1, name: 'dr. Ahmad Sp.JP(K)', specialty: 'Spesialis Jantung', time: '08:00 - 12:00', isAvailable: true, quota: 20, remaining: 5 },
      { id: 2, name: 'dr. Sarah Sp.JP', specialty: 'Spesialis Jantung', time: '13:00 - 16:00', isAvailable: true, quota: 15, remaining: 12 }
    ]
  },
  'onkologi': {
    name: 'Pusat Onkologi (Cancer Center)',
    heroImg: ASSETS.hero.layanan,
    desc: 'Memberikan harapan baru melalui penanganan kanker yang terintegrasi, mulai dari deteksi dini, bedah onkologi, hingga kemoterapi dalam satu atap.',
    features: [
      { title: 'Chemotherapy Suite', icon: 'fa-solid fa-house-medical', desc: 'Ruang kemoterapi yang nyaman dengan pemantauan tim medis khusus.' },
      { title: 'Oncology Surgery', icon: 'fa-solid fa-scalpel', desc: 'Tim bedah onkologi berpengalaman untuk berbagai jenis tumor dan kanker.' },
      { title: 'Biopsy Lab', icon: 'fa-solid fa-microscope', desc: 'Laboratorium patologi anatomi untuk diagnosis jenis sel kanker secara akurat.' }
    ],
    steps: [
      { title: 'Konsultasi', desc: 'Konsultasi dengan tim onkologi multidisiplin.' },
      { title: 'Staging', desc: 'Penentuan stadium kanker melalui berbagai tes diagnostik.' },
      { title: 'Treatment', desc: 'Pelaksanaan protokol pengobatan (Bedah/Kemoterapi).' }
    ],
    doctors: [
      { id: 3, name: 'dr. Budi Sp.B(K)Onk', specialty: 'Bedah Onkologi', time: '09:00 - 13:00', isAvailable: true, quota: 10, remaining: 2 }
    ]
  },
  'stroke': {
    name: 'Neuro-Vascular & Stroke Center',
    heroImg: ASSETS.hero.layanan,
    desc: 'Layanan "Golden Hour" untuk penanganan cepat serangan stroke guna meminimalisir kerusakan otak dan mempercepat pemulihan pasien.',
    features: [
      { title: 'Stroke Unit', icon: 'fa-solid fa-brain', desc: 'Unit perawatan khusus pasien stroke dengan tim perawat terlatih.' },
      { title: 'CT-Scan 128 Slice', icon: 'fa-solid fa-x-ray', desc: 'Pencitraan otak super cepat untuk membedakan jenis stroke.' },
      { title: 'Rehabilitasi Medik', icon: 'fa-solid fa-person-walking-with-cane', desc: 'Terapi pemulihan fungsi gerak dan bicara pasca stroke.' }
    ],
    steps: [
      { title: 'Emergency Room', desc: 'Penanganan awal di IGD khusus Stroke.' },
      { title: 'Stabilisasi', desc: 'Tindakan trombolisis atau manajemen tekanan otak.' },
      { title: 'Recovery', desc: 'Perawatan di Stroke Unit dan fisioterapi dini.' }
    ],
    doctors: [
      { id: 4, name: 'dr. Linda Sp.S', specialty: 'Spesialis Saraf', time: '08:00 - 14:00', isAvailable: true, quota: 25, remaining: 10 }
    ]
  }
};

const service = ref(servicesData[route.params.id] || servicesData['jantung']);

onMounted(() => {
  window.scrollTo(0, 0);
  // Re-load data if ID changes
  service.value = servicesData[route.params.id] || servicesData['jantung'];
  setTimeout(() => {
    isLoading.value = false;
  }, 800);
});
</script>

<template>
  <div class="service-detail">
    <!-- Hero Section -->
    <header 
      class="service-hero" 
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${service.heroImg})` }"
    >
      <div class="container">
        <div class="service-hero__content" data-aos="fade-up">
          <div class="service-hero__badge">Pusat Layanan Unggulan</div>
          <h1 class="heading-academic">{{ service.name }}</h1>
          <p class="service-hero__desc">{{ service.desc }}</p>
          <div class="service-hero__actions">
            <button class="btn btn--green">Daftar Online</button>
            <button class="btn btn--white">Lihat Jadwal Dokter</button>
          </div>
        </div>
      </div>
    </header>

    <!-- Features Section -->
    <section class="section-features">
      <div class="container">
        <SectionHeader title="Fasilitas & Kemampuan" subtitle="Dukungan teknologi medis mutakhir untuk akurasi diagnosis dan efektivitas pengobatan." />
        
        <div class="features-grid">
          <div v-for="(feat, index) in service.features" :key="index" class="feat-card" data-aos="fade-up" :data-aos-delay="index * 100">
            <div class="feat-card__icon"><i :class="feat.icon"></i></div>
            <h3 class="feat-card__title">{{ feat.title }}</h3>
            <p class="feat-card__desc">{{ feat.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Workflow Section -->
    <section class="section-workflow">
      <div class="container">
        <SectionHeader title="Alur Pelayanan" subtitle="Panduan langkah demi langkah untuk mendapatkan layanan di Pusat Jantung Terpadu." />
        <StepFlow :steps="service.steps" />
      </div>
    </section>

    <!-- Doctors Section -->
    <section class="section-doctors">
      <div class="container">
        <SectionHeader title="Tim Dokter Spesialis" subtitle="Para ahli kardiovaskular yang siap melayani Anda dengan sepenuh hati." />
        
        <div class="doctors-grid">
          <DoctorCard v-for="doc in service.doctors" :key="doc.id" v-bind="doc" />
        </div>
      </div>
    </section>

    <!-- Contact CTA -->
    <section class="section-cta">
      <div class="container">
        <div class="cta-box" data-aos="zoom-in">
          <div class="cta-box__content">
            <h2>Butuh Bantuan atau Informasi Lebih Lanjut?</h2>
            <p>Tim kami siap menjawab pertanyaan Anda mengenai layanan Pusat Jantung Terpadu.</p>
          </div>
          <div class="cta-box__actions">
            <a href="https://wa.me/628123456789" class="btn-wa">
              <i class="fa-brands fa-whatsapp"></i> Chat via WhatsApp
            </a>
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

.service-detail {
  padding-top: 8rem;
}

/* Hero */
.service-hero {
  height: 70vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
}

.service-hero__badge {
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

.service-hero__desc {
  font-size: 1.8rem;
  max-width: 70rem;
  line-height: 1.8;
  margin-bottom: 4rem;
  color: rgba(255, 255, 255, 0.9);
}

.service-hero__actions {
  display: flex;
  gap: 2rem;
}

.btn {
  padding: 1.5rem 3.5rem;
  border-radius: 50px;
  border: none;
  font-size: 1.5rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s;
}

.btn--green { background: #55c57a; color: #fff; }
.btn--green:hover { background: #40945c; transform: translateY(-3px); }

.btn--white { background: #fff; color: #333; }
.btn--white:hover { background: #f1f1f1; transform: translateY(-3px); }

/* Features */
.section-features, .section-workflow, .section-doctors {
  padding: 12rem 0;
}

.section-workflow { background-color: #f9fbf9; }

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 4rem;
}

.feat-card {
  padding: 4rem;
  background: #fff;
  border-radius: 20px;
  box-shadow: 0 1.5rem 4rem rgba(0, 0, 0, 0.05);
  text-align: center;
  transition: all 0.3s;
}

.feat-card:hover { transform: translateY(-1rem); }

.feat-card__icon {
  font-size: 4rem;
  color: #55c57a;
  margin-bottom: 2.5rem;
}

.feat-card__title {
  font-size: 2.2rem;
  font-weight: 800;
  margin-bottom: 1.5rem;
  color: #333;
}

.feat-card__desc {
  font-size: 1.5rem;
  color: #777;
  line-height: 1.7;
}

/* Doctors */
.doctors-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(35rem, 1fr));
  gap: 4rem;
}

/* CTA */
.section-cta { padding: 10rem 0; }

.cta-box {
  background: linear-gradient(135deg, #55c57a 0%, #28b485 100%);
  padding: 6rem;
  border-radius: 30px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #fff;
  box-shadow: 0 3rem 6rem rgba(85, 197, 122, 0.3);
}

.cta-box__content h2 { font-size: 3rem; margin-bottom: 1.5rem; font-weight: 900; }
.cta-box__content p { font-size: 1.8rem; color: rgba(255, 255, 255, 0.9); }

.btn-wa {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  padding: 2rem 4rem;
  background: #fff;
  color: #25d366;
  text-decoration: none;
  font-size: 1.8rem;
  font-weight: 800;
  border-radius: 50px;
  transition: all 0.3s;
}

.btn-wa:hover { transform: scale(1.05); box-shadow: 0 1rem 3rem rgba(0,0,0,0.2); }

@media (max-width: 1000px) {
  .cta-box { flex-direction: column; text-align: center; gap: 4rem; }
  .features-grid { grid-template-columns: 1fr; }
  .service-hero__actions { flex-direction: column; }
  .heading-primary { font-size: 3.5rem; }
}
</style>
