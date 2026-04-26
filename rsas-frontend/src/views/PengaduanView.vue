<script setup>
import { onMounted } from 'vue';
import SectionHeader from '../components/ui/SectionHeader.vue';
import { ASSETS } from '../config/assets';

const channels = [
  { 
    title: 'SP4N-LAPOR!', 
    desc: 'Layanan aspirasi dan pengaduan online rakyat yang terhubung dengan pemerintah pusat.', 
    icon: 'fa-solid fa-flag',
    color: '#e74c3c',
    link: 'https://www.lapor.go.id'
  },
  { 
    title: 'WhatsApp Hotline', 
    desc: 'Layanan cepat melalui pesan instan WhatsApp untuk respon yang lebih personal.', 
    icon: 'fa-brands fa-whatsapp',
    color: '#25d366',
    link: 'https://wa.me/62811000000'
  },
  { 
    title: 'Email Pengaduan', 
    desc: 'Sampaikan keluhan secara mendalam melalui surat elektronik resmi kami.', 
    icon: 'fa-solid fa-envelope',
    color: '#3498db',
    link: 'mailto:pengaduan@rsas.go.id'
  }
];

const flowSteps = [
  { num: '01', title: 'Penyampaian Laporan', desc: 'Sampaikan keluhan Anda melalui kanal yang tersedia dengan bukti pendukung.' },
  { num: '02', title: 'Verifikasi & Validasi', desc: 'Tim pengaduan akan memeriksa kelengkapan data dalam waktu maksimal 24 jam.' },
  { num: '03', title: 'Tindak Lanjut', desc: 'Unit terkait akan memproses dan memberikan solusi atas keluhan Anda.' },
  { num: '04', title: 'Umpan Balik', desc: 'Anda akan menerima pemberitahuan mengenai status dan hasil penyelesaian.' }
];

onMounted(() => {
  window.scrollTo(0, 0);
});
</script>

<template>
  <div class="complaint-page">
    <!-- Hero Section -->
    <header 
      class="complaint-hero"
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.pengaduan})` }"
    >
      <div class="container">
        <div class="complaint-hero__content" data-aos="fade-up">
          <div class="complaint-hero__badge">Layanan Aspirasi & Pengaduan</div>
          <h1 class="heading-academic">Suara Anda <span>Adalah Amanah Kami</span></h1>
          <p class="complaint-hero__sub">Sampaikan aspirasi, saran, atau keluhan Anda. Setiap masukan adalah pijakan bagi kami untuk memberikan pelayanan yang lebih baik.</p>
        </div>
      </div>
    </header>

    <!-- Flow Section -->
    <section class="section-flow">
      <div class="container">
        <SectionHeader 
          title="Alur Penanganan Pengaduan" 
          subtitle="Proses transparan dan akuntabel dalam menindaklanjuti setiap aspirasi Anda."
          align="center"
        />

        <div class="flow-grid">
          <div v-for="step in flowSteps" :key="step.num" class="flow-card" data-aos="fade-up">
            <div class="flow-card__num">{{ step.num }}</div>
            <h3 class="flow-card__title">{{ step.title }}</h3>
            <p class="flow-card__desc">{{ step.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Complaint Channels -->
    <section class="section-channels">
      <div class="container">
        <div class="channels-grid">
          <div v-for="channel in channels" :key="channel.title" class="channel-card" data-aos="zoom-in">
            <div class="channel-card__icon" :style="{ backgroundColor: channel.color }">
              <i :class="channel.icon"></i>
            </div>
            <h3 class="channel-card__title">{{ channel.title }}</h3>
            <p class="channel-card__desc">{{ channel.desc }}</p>
            <a :href="channel.link" class="btn btn--outline-dark" target="_blank">Hubungi Sekarang</a>
          </div>
        </div>
      </div>
    </section>

    <!-- Complaint Form -->
    <section class="section-form-complaint">
      <div class="container">
        <div class="form-container">
          <div class="form-info" data-aos="fade-right">
            <h2 class="heading-secondary">Kirim Pengaduan Langsung</h2>
            <p>Gunakan formulir di bawah ini jika Anda ingin mengirimkan laporan langsung kepada tim manajemen RSAS.</p>
            <div class="form-features">
              <div class="f-item"><i class="fa-solid fa-shield-halved"></i> <span>Identitas Dirahasiakan</span></div>
              <div class="f-item"><i class="fa-solid fa-clock"></i> <span>Respon < 2x24 Jam</span></div>
              <div class="f-item"><i class="fa-solid fa-check-double"></i> <span>Penyelesaian Tuntas</span></div>
            </div>
          </div>
          <form class="complaint-form" data-aos="fade-left">
            <div class="form-group">
              <label>Nama Lengkap (Opsional)</label>
              <input type="text" placeholder="Masukkan nama Anda" />
            </div>
            <div class="form-group">
              <label>Nomor Telepon / WhatsApp</label>
              <input type="tel" placeholder="0812xxxx" required />
            </div>
            <div class="form-group">
              <label>Kategori Masalah</label>
              <select>
                <option>Fasilitas & Sarana</option>
                <option>Pelayanan Petugas</option>
                <option>Antrean & Waktu Tunggu</option>
                <option>Kebersihan</option>
                <option>Lainnya</option>
              </select>
            </div>
            <div class="form-group">
              <label>Isi Laporan</label>
              <textarea rows="5" placeholder="Jelaskan secara detail masalah yang Anda alami..."></textarea>
            </div>
            <button type="submit" class="btn btn--green u-margin-top-small">Kirim Laporan</button>
          </form>
        </div>
      </div>
    </section>

    <!-- Stats Counter -->
    <section class="section-complaint-stats">
      <div class="container">
        <div class="c-stats-grid">
          <div class="c-stat-item">
            <span class="c-stat-num">98%</span>
            <span class="c-stat-label">Aduan Terselesaikan</span>
          </div>
          <div class="c-stat-item">
            <span class="c-stat-num">24h</span>
            <span class="c-stat-label">Rata-rata Respon</span>
          </div>
          <div class="c-stat-item">
            <span class="c-stat-num">1.2k</span>
            <span class="c-stat-label">Aspirasi Diterima</span>
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

.complaint-page {
  padding-top: 8rem;
}

/* Hero */
.complaint-hero {
  height: 60vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
  position: relative;
}

.complaint-hero__badge {
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

.complaint-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  line-height: 1.6;
  margin-bottom: 5rem;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
}

/* Flow */
.section-flow { padding: 10rem 0; background-color: #f9fbf9; }

.flow-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 3rem;
  margin-top: 6rem;
}

.flow-card {
  text-align: center;
  position: relative;
}

.flow-card__num {
  font-size: 6rem;
  font-weight: 900;
  color: rgba(85, 197, 122, 0.1);
  margin-bottom: -3rem;
}

.flow-card__title {
  font-size: 2rem;
  font-weight: 800;
  color: #333;
  margin-bottom: 1.5rem;
}

.flow-card__desc { font-size: 1.4rem; color: #777; line-height: 1.7; }

/* Channels */
.section-channels { padding: 10rem 0; }

.channels-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 4rem;
}

.channel-card {
  background: #fff;
  padding: 5rem 4rem;
  border-radius: 30px;
  box-shadow: 0 1rem 4rem rgba(0,0,0,0.05);
  text-align: center;
  transition: all 0.3s;
}

.channel-card:hover { transform: translateY(-1rem); box-shadow: 0 2rem 6rem rgba(0,0,0,0.1); }

.channel-card__icon {
  width: 8rem;
  height: 8rem;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3.5rem;
  color: #fff;
  margin: 0 auto 3rem;
}

.channel-card__title { font-size: 2.2rem; font-weight: 800; margin-bottom: 1.5rem; }
.channel-card__desc { font-size: 1.6rem; color: #666; margin-bottom: 3rem; line-height: 1.6; }

/* Form */
.section-form-complaint { padding-bottom: 12rem; }

.form-container {
  background: #fff;
  border-radius: 40px;
  box-shadow: 0 3rem 8rem rgba(0,0,0,0.1);
  display: flex;
  overflow: hidden;
}

.form-info {
  flex: 1;
  background: #2d3436;
  color: #fff;
  padding: 8rem 6rem;
}

.form-info p { font-size: 1.8rem; opacity: 0.8; margin: 2rem 0 4rem; }

.form-features { display: grid; gap: 2rem; }
.f-item { display: flex; align-items: center; gap: 2rem; font-size: 1.8rem; font-weight: 700; }
.f-item i { color: #55c57a; font-size: 2.4rem; }

.complaint-form { flex: 1.5; padding: 8rem 6rem; }

.form-group { margin-bottom: 2.5rem; }
.form-group label { display: block; font-size: 1.4rem; font-weight: 800; color: #333; margin-bottom: 1rem; text-transform: uppercase; }
.form-group input, .form-group select, .form-group textarea {
  width: 100%;
  padding: 1.5rem 2rem;
  border: 1px solid #ddd;
  border-radius: 12px;
  font-size: 1.6rem;
  background: #f9fbf9;
}

/* Stats */
.section-complaint-stats { padding-bottom: 10rem; }
.c-stats-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 4rem; text-align: center; }
.c-stat-num { display: block; font-size: 5rem; font-weight: 900; color: #55c57a; }
.c-stat-label { font-size: 1.6rem; font-weight: 700; color: #777; text-transform: uppercase; }

/* Buttons */
.btn { padding: 1.5rem 4rem; border-radius: 50px; text-decoration: none; font-size: 1.6rem; font-weight: 700; transition: all 0.3s; display: inline-block; cursor: pointer; border: none; }
.btn--green { background: #55c57a; color: #fff; }
.btn--outline-dark { background: transparent; border: 2px solid #333; color: #333; }
.btn--outline-dark:hover { background: #333; color: #fff; }

@media (max-width: 1000px) {
  .flow-grid, .channels-grid, .c-stats-grid { grid-template-columns: 1fr; }
  .form-container { flex-direction: column; }
  .heading-primary { font-size: 3.5rem; }
}
</style>
