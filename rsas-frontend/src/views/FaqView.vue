<script setup>
import { ref, computed, onMounted } from 'vue';
import { ASSETS } from '../config/assets';

const searchQuery = ref('');
const activeCategory = ref('Semua');

const categories = ['Semua', 'Layanan & Fasilitas', 'Pendaftaran', 'BPJS & Asuransi', 'Pendidikan'];

const faqs = [
  { 
    category: 'Pendaftaran',
    question: 'Bagaimana cara mendaftar antrean online di RSAS?', 
    answer: 'Anda dapat mendaftar melalui aplikasi mobile RSAS yang tersedia di Play Store, atau melalui menu "Layanan" di website ini. Pastikan Anda memiliki nomor RM (Rekam Medis) untuk pendaftaran online.' 
  },
  { 
    category: 'BPJS & Asuransi',
    question: 'Apa saja syarat berobat menggunakan BPJS Kesehatan?', 
    answer: 'Syarat utama adalah membawa Kartu BPJS (aktif), Surat Rujukan dari Faskes Tingkat 1 (Puskesmas/Klinik), dan KTP. Untuk pasien rawat jalan, rujukan berlaku selama 3 bulan.' 
  },
  { 
    category: 'Layanan & Fasilitas',
    question: 'Apakah RSAS menyediakan layanan gawat darurat 24 jam?', 
    answer: 'Ya, Instalasi Gawat Darurat (IGD) kami beroperasi 24 jam sehari, 7 hari seminggu dengan tim medis yang selalu siap siaga.' 
  },
  { 
    category: 'Pendaftaran',
    question: 'Berapa lama sebelum jadwal praktik saya harus datang?', 
    answer: 'Disarankan datang minimal 30-60 menit sebelum jadwal praktik dokter dimulai untuk proses verifikasi administrasi dan pemeriksaan tanda vital di poli.' 
  },
  { 
    category: 'BPJS & Asuransi',
    question: 'Apakah pasien umum bisa mendapatkan layanan di RSAS?', 
    answer: 'Tentu saja. Kami melayani pasien umum, pasien dengan asuransi swasta, maupun pasien jaminan perusahaan selain BPJS Kesehatan.' 
  },
  { 
    category: 'Pendidikan',
    question: 'Bagaimana prosedur pengajuan izin penelitian di RSAS?', 
    answer: 'Silakan kunjungi Portal Pendidikan di website kami, unduh formulir izin penelitian, dan ajukan melalui sekretariat Komkordik atau secara online melalui sistem KEPK kami.' 
  }
];

const filteredFaqs = computed(() => {
  return faqs.filter(faq => {
    const matchesSearch = faq.question.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
                         faq.answer.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchesCategory = activeCategory.value === 'Semua' || faq.category === activeCategory.value;
    return matchesSearch && matchesCategory;
  });
});

const openIndex = ref(null);
const toggleFaq = (index, event) => {
  const isOpening = openIndex.value !== index;
  openIndex.value = openIndex.value === index ? null : index;

  if (isOpening && event) {
    // Memberikan jeda singkat agar transisi pembukaan dimulai
    setTimeout(() => {
      const element = event.currentTarget.closest('.faq-item');
      if (element) {
        // Scroll secara halus agar item berada di area pandang yang nyaman
        element.scrollIntoView({ 
          behavior: 'smooth', 
          block: 'nearest',
          inline: 'nearest' 
        });
      }
    }, 250);
  }
};

// JS Hooks for smooth zero-delay accordion
const beforeEnter = (el) => {
  el.style.height = '0';
  el.style.opacity = '0';
};

const enter = (el) => {
  el.style.height = el.scrollHeight + 'px';
  el.style.opacity = '1';
};

const afterEnter = (el) => {
  el.style.height = 'auto';
};

const beforeLeave = (el) => {
  el.style.height = el.scrollHeight + 'px';
};

const leave = (el) => {
  // Trigger reflow
  el.offsetHeight; 
  el.style.height = '0';
  el.style.opacity = '0';
};

onMounted(() => {
  window.scrollTo(0, 0);
  if (window.AOS) {
    window.AOS.refresh();
  }
});
</script>

<template>
  <div class="faq-page">
    <!-- Hero Section -->
    <header 
      class="faq-hero"
      :style="{ backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.8)), url(${ASSETS.hero.faq})` }"
    >
      <div class="container">
        <div class="faq-hero__content" data-aos="fade-up">
          <div class="faq-hero__badge">Pusat Informasi & Bantuan</div>
          <h1 class="heading-primary">
            <span class="heading-primary__main">Apa yang Bisa</span>
            <span class="heading-primary__sub">Kami Bantu?</span>
          </h1>
          <p class="faq-hero__sub">Temukan jawaban cepat dan terpercaya untuk setiap pertanyaan Anda mengenai layanan kesehatan kami.</p>
          
          <div class="search-box">
            <i class="fa-solid fa-magnifying-glass"></i>
            <input 
              type="text" 
              v-model="searchQuery" 
              placeholder="Cari pertanyaan atau kata kunci..." 
            />
          </div>
        </div>
      </div>
    </header>

    <!-- Content Section -->
    <section class="section-faq-content">
      <div class="container">
        <!-- Categories -->
        <div class="faq-categories" data-aos="fade-up">
          <button 
            v-for="cat in categories" 
            :key="cat"
            class="cat-btn"
            :class="{ 'cat-btn--active': activeCategory === cat }"
            @click="activeCategory = cat"
          >
            {{ cat }}
          </button>
        </div>

        <!-- FAQ List -->
        <div class="faq-list">
          <div 
            v-for="(faq, index) in filteredFaqs" 
            :key="index"
            class="faq-item"
            :class="{ 'faq-item--open': openIndex === index }"
            data-aos="fade-up"
          >
            <div class="faq-item__question" @click="toggleFaq(index, $event)">
              <div class="faq-item__q-text">
                <span class="faq-category-tag">{{ faq.category }}</span>
                <h3>{{ faq.question }}</h3>
              </div>
              <div class="faq-item__icon">
                <i class="fa-solid fa-chevron-down"></i>
              </div>
            </div>
            
            <transition 
              name="faq-slide"
              @before-enter="beforeEnter"
              @enter="enter"
              @after-enter="afterEnter"
              @before-leave="beforeLeave"
              @leave="leave"
            >
              <div class="faq-item__answer" v-if="openIndex === index">
                <div class="answer-content">
                  <p>{{ faq.answer }}</p>
                </div>
              </div>
            </transition>
          </div>

          <div v-if="filteredFaqs.length === 0" class="faq-empty">
            <i class="fa-solid fa-face-frown"></i>
            <p>Maaf, tidak ada pertanyaan yang sesuai dengan kata kunci Anda.</p>
            <button @click="searchQuery = ''; activeCategory = 'Semua'" class="btn-text">Tampilkan Semua FAQ</button>
          </div>
        </div>
      </div>
    </section>

    <!-- Support CTA -->
    <section class="section-faq-support">
      <div class="container">
        <div class="support-card" data-aos="zoom-in">
          <div class="support-card__icon"><i class="fa-solid fa-headset"></i></div>
          <div class="support-card__info">
            <h2>Masih Punya Pertanyaan?</h2>
            <p>Tim layanan pelanggan kami siap membantu Anda 24/7 melalui berbagai kanal komunikasi.</p>
          </div>
          <div class="support-card__actions">
            <router-link to="/pengaduan" class="btn btn--green">Hubungi Kami</router-link>
            <a href="#" class="btn btn--outline-dark">Tanya di WhatsApp</a>
          </div>
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
  padding: 0 4rem;

  @include respond("tablet") {
    padding: 0 2rem;
  }
}

.faq-page {
  /* Padding top dihapus agar Hero menyentuh bagian paling atas layar */
  min-height: 100vh;
}

/* Hero */
.faq-hero {
  height: 60vh;
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  color: #fff;
  text-align: center;
}

.faq-hero__badge {
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

.faq-hero__sub {
  font-size: 2.2rem;
  max-width: 80rem;
  margin: 0 auto;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 300;
  line-height: 1.6;
  margin-bottom: 5rem;
}

.search-box {
  max-width: 70rem;
  margin: 0 auto;
  position: relative;
}

.search-box i {
  position: absolute;
  left: 2.5rem;
  top: 50%;
  transform: translateY(-50%);
  font-size: 2rem;
  color: #55c57a;
}

.search-box input {
  width: 100%;
  padding: 2rem 3rem 2rem 6.5rem;
  border: 2px solid #eee;
  border-radius: 50px;
  font-size: 1.6rem;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
  transition: all 0.3s;
}

.search-box input:focus {
  border-color: #55c57a;
  outline: none;
  box-shadow: 0 1.5rem 4rem rgba(85, 197, 122, 0.15);
}

/* Categories */
.section-faq-content { padding: 8rem 0; background-color: #f9fbf9; }

.faq-categories {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 1.5rem;
  margin-bottom: 6rem;
}

.cat-btn {
  padding: 1.2rem 2.5rem;
  border: 1px solid #ddd;
  border-radius: 50px;
  background: #fff;
  font-size: 1.4rem;
  font-weight: 700;
  color: #555;
  cursor: pointer;
  transition: all 0.3s;
}

.cat-btn:hover { border-color: #55c57a; color: #55c57a; }
.cat-btn--active { background: #55c57a; border-color: #55c57a; color: #fff; }

/* FAQ Items */
.faq-list { display: flex; flex-direction: column; gap: 2rem; }

.faq-item {
  background: #fff;
  border-radius: 25px;
  overflow: hidden;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.04);
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  border: 1px solid #f0f0f0;
}

.faq-item:hover { transform: translateY(-5px); box-shadow: 0 2rem 5rem rgba(0,0,0,0.08); }

.faq-item--open {
  border-color: rgba(85, 197, 122, 0.3);
  box-shadow: 0 2rem 6rem rgba(85, 197, 122, 0.1);
}

.faq-item__question {
  padding: 3rem 4rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 3rem;
  cursor: pointer;
}

.faq-item__q-text { flex: 1; text-align: left; }

.faq-category-tag {
  display: inline-block;
  font-size: 1.1rem;
  font-weight: 800;
  color: #55c57a;
  text-transform: uppercase;
  background: rgba(85, 197, 122, 0.08);
  padding: 0.4rem 1.2rem;
  border-radius: 5px;
  margin-bottom: 1rem;
}

.faq-item__question h3 {
  font-size: 2.2rem;
  font-weight: 800;
  color: #2d3436;
  line-height: 1.3;
}

.faq-item__icon {
  width: 4rem;
  height: 4rem;
  background: #f9fbf9;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.faq-item__icon i { font-size: 1.6rem; color: #55c57a; }

.faq-item--open .faq-item__icon {
  background: #55c57a;
  transform: rotate(180deg);
}

.faq-item--open .faq-item__icon i { color: #fff; }

/* Zero-Delay JS Transition Styles */
.faq-item__answer {
  overflow: hidden;
  transition: height 0.2s cubic-bezier(0, 0, 0.2, 1), opacity 0.2s linear;
}

.answer-content {
  padding: 0 4rem 4rem 4rem;
  font-size: 1.7rem;
  color: #636e72;
  line-height: 1.8;
}

.faq-empty { text-align: center; padding: 6rem 0; }
.faq-empty i { font-size: 6rem; color: #ddd; margin-bottom: 2rem; }
.faq-empty p { font-size: 1.8rem; color: #777; margin-bottom: 2rem; }

/* Support Card */
.section-faq-support { padding: 10rem 0; }

.support-card {
  background: #fff;
  padding: 6rem;
  border-radius: 30px;
  box-shadow: 0 3rem 6rem rgba(0,0,0,0.08);
  display: flex;
  align-items: center;
  gap: 5rem;
}

.support-card__icon {
  width: 10rem;
  height: 10rem;
  background: rgba(85, 197, 122, 0.1);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 4.5rem;
  color: #55c57a;
}

.support-card__info { flex: 1; }
.support-card__info h2 { font-size: 3rem; font-weight: 900; margin-bottom: 1rem; }
.support-card__info p { font-size: 1.6rem; color: #666; }

.support-card__actions { display: flex; gap: 2rem; }

/* Buttons */
.btn { padding: 1.5rem 3.5rem; border-radius: 50px; text-decoration: none; font-size: 1.5rem; font-weight: 700; transition: all 0.3s; display: inline-block; cursor: pointer; border: none; }
.btn--green { background: #55c57a; color: #fff; }
.btn--outline-dark { background: transparent; border: 2px solid #333; color: #333; }

@media (max-width: 1000px) {
  .support-card { flex-direction: column; text-align: center; }
  .support-card__actions { justify-content: center; }
  .heading-primary { font-size: 3.5rem; }
}
</style>
