<script setup>
import { ref } from 'vue';
import TestimonialCard from './TestimonialCard.vue';

const testimonials = ref([
  {
    name: 'Siti Aminah',
    unit: 'Poli Jantung',
    message: 'Pelayanan di RSAS sangat profesional. Dokter menjelaskan kondisi saya dengan sangat detail dan menenangkan. Terima kasih RSAS!',
    rating: 5,
    date: '24 April 2026'
  },
  {
    name: 'Budi Santoso',
    unit: 'IGD',
    message: 'Respon cepat saat keadaan darurat. Perawat sangat sigap dan fasilitas peralatan medisnya sangat modern. Sangat terbantu.',
    rating: 5,
    date: '20 April 2026'
  },
  {
    name: 'Hj. Fatimah',
    unit: 'Rawat Inap Interna',
    message: 'Kamar bersih dan nyaman seperti di rumah sendiri. Menu makanannya pun enak dan terjaga gizinya. Pelayanan perawat ramah sekali.',
    rating: 5,
    date: '15 April 2026'
  },
  {
    name: 'Andi Pratama',
    unit: 'Radiologi',
    message: 'Proses antrean sekarang jauh lebih cepat dengan sistem online. Hasil lab juga bisa diakses dengan mudah. Mantap!',
    rating: 4,
    date: '10 April 2026'
  }
]);

const emit = defineEmits(['open-modal']);
</script>

<template>
  <section class="section-testimonials">
    <div class="container">
      <div class="testimonial-header" data-aos="fade-up">
        <div class="u-center-text u-margin-bottom-big">
          <span class="testimonial-header__badge">Suara Hati Pasien</span>
          <h2 class="heading-secondary">Dinding Apresiasi & Harapan</h2>
          <div class="heart-divider"><i class="fa-solid fa-heart"></i></div>
          <p class="section-subtitle">
            "Di balik setiap kalimat di sini, ada harapan keluarga yang kami jaga. <br>
            Setiap senyum kesembuhan mereka adalah detak jantung pengabdian kami. <br>
            Inilah alasan kami terus melayani dengan hati."
          </p>
        </div>
        
        <button class="btn-testimonial" @click="$emit('open-modal')" data-aos="fade-left">
          <i class="fa-solid fa-pen-nib"></i> Tulis Kesan Anda
        </button>
      </div>

      <div class="testimonial-grid">
        <div v-for="(item, index) in testimonials" :key="index" 
             class="testimonial-grid__item"
             :data-aos="index % 2 === 0 ? 'fade-up' : 'fade-down'"
             :data-aos-delay="index * 100">
          <TestimonialCard v-bind="item" />
        </div>
      </div>
      
      <div class="testimonial-footer u-center-text u-margin-top-big" data-aos="zoom-in">
        <p class="testimonial-footer__text radiant-text">
          <i class="fa-solid fa-wand-magic-sparkles"></i> 
          "Setiap kata baik Anda adalah cahaya bagi sesama. Bagikan kisah Anda di sini, <br>
          agar energi positif ini terus mengalir dan menguatkan hati mereka yang sedang berjuang untuk sembuh."
        </p>
        <div class="testimonial-stats">
          <div class="stat-item">
            <span class="stat-item__num">4.9/5</span>
            <span class="stat-item__label">Rating Kepuasan</span>
          </div>
          <div class="stat-divider"></div>
          <div class="stat-item">
            <span class="stat-item__num">1,200+</span>
            <span class="stat-item__label">Testimoni Masuk</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.section-testimonials {
  padding: 18rem 0;
  background: linear-gradient(135deg, #f9fbf9 0%, #fff 100%);
  position: relative;
  overflow: hidden;
}

.section-testimonials::before {
  content: "APRESIASI";
  position: absolute;
  top: 5rem;
  right: -5rem;
  font-size: 15rem;
  font-weight: 950;
  color: rgba(85, 197, 122, 0.03);
  z-index: 0;
  pointer-events: none;
}

.container {
  max-width: 144rem;
  margin: 0 auto;
  padding: 0 3rem;
  position: relative;
  z-index: 1;
}

.testimonial-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  margin-bottom: 6rem;
}

.section-subtitle {
  font-family: 'Playfair Display', serif;
  font-size: 2.2rem;
  font-style: italic;
  color: #4a4a4a;
  margin-top: 2rem;
  line-height: 1.8;
  max-width: 90rem;
  margin-left: auto;
  margin-right: auto;
  font-weight: 400;
}

.heart-divider {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1.5rem;
  margin: 1rem 0;
  color: #55c57a;
  font-size: 1.8rem;
  opacity: 0.6;
}

.heart-divider::before,
.heart-divider::after {
  content: "";
  height: 1px;
  width: 5rem;
  background: linear-gradient(to right, transparent, #55c57a, transparent);
}

.testimonial-header__badge {
  display: inline-block;
  padding: 0.8rem 2rem;
  background: rgba(85, 197, 122, 0.1);
  color: #55c57a;
  border-radius: 50px;
  font-size: 1.4rem;
  font-weight: 800;
  text-transform: uppercase;
  margin-bottom: 1.5rem;
}

.btn-testimonial {
  background: #55c57a;
  color: #fff;
  border: none;
  padding: 1.8rem 3.5rem;
  border-radius: 50px;
  font-size: 1.6rem;
  font-weight: 700;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 1.2rem;
  box-shadow: 0 1rem 3rem rgba(85, 197, 122, 0.3);
  transition: all 0.3s;
}

.btn-testimonial:hover {
  transform: translateY(-3px);
  box-shadow: 0 1.5rem 4rem rgba(85, 197, 122, 0.4);
  background: #28b485;
}

.testimonial-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(35rem, 1fr));
  gap: 6rem;
}

.testimonial-footer__text {
  font-size: 1.8rem;
  color: #777;
  margin-bottom: 3rem;
  line-height: 1.8;
}

.radiant-text {
  color: #444;
  font-style: italic;
  font-weight: 500;
  text-shadow: 0 0 20px rgba(85, 197, 122, 0.2);
  animation: breathe 4s ease-in-out infinite;
}

.radiant-text i {
  color: #f1c40f;
  margin-bottom: 1.5rem;
  display: block;
  font-size: 2.5rem;
}

@keyframes breathe {
  0%, 100% { transform: scale(1); opacity: 0.8; }
  50% { transform: scale(1.02); opacity: 1; text-shadow: 0 0 30px rgba(85, 197, 122, 0.4); }
}

.testimonial-stats {
  display: inline-flex;
  align-items: center;
  gap: 5rem;
  background: #fff;
  padding: 2.5rem 6rem;
  border-radius: 100px;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.05);
}

.stat-item {
  display: flex;
  flex-direction: column;
}

.stat-item__num {
  font-size: 3rem;
  font-weight: 900;
  color: #333;
}

.stat-item__label {
  font-size: 1.4rem;
  color: #55c57a;
  font-weight: 700;
  text-transform: uppercase;
}

.stat-divider {
  width: 1px;
  height: 5rem;
  background: #eee;
}

@media (max-width: 1024px) {
  .testimonial-header {
    flex-direction: column;
    align-items: center;
    text-align: center;
    gap: 3rem;
  }
}

@media (max-width: 768px) {
  .testimonial-stats {
    flex-direction: column;
    gap: 2rem;
    padding: 3rem;
    border-radius: 20px;
  }
  .stat-divider {
    width: 50%;
    height: 1px;
  }
}
</style>
