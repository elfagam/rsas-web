<script setup>
import { ref, onMounted } from 'vue';
import AOS from 'aos';

const fontSize = ref('normal'); // 'normal', 'lg', 'xl'

const setFontSize = (size) => {
  fontSize.value = size;
  const root = document.documentElement;
  
  // Direct Style Injection with 10px base (62.5% Trick)
  if (size === 'normal') {
    root.style.fontSize = '10px';
  } else if (size === 'lg') {
    root.style.fontSize = '12.5px';
  } else if (size === 'xl') {
    root.style.fontSize = '15px';
  }
  
  // Also keep classes for CSS-specific overrides if needed
  root.classList.remove('font-size-lg', 'font-size-xl');
  if (size === 'lg') root.classList.add('font-size-lg');
  if (size === 'xl') root.classList.add('font-size-xl');
  
  // Refresh AOS to recalculate positions after font scaling
  setTimeout(() => {
    AOS.refresh();
  }, 200);
  
  // Save preference
  localStorage.setItem('rsas-font-size', size);
};

onMounted(() => {
  const savedSize = localStorage.getItem('rsas-font-size');
  if (savedSize) {
    setFontSize(savedSize);
  }
});
</script>

<template>
  <div class="a11y-toolbar">
    <div class="a11y-toolbar__buttons">
      <button 
        @click="setFontSize('normal')" 
        class="a11y-btn" 
        :class="{ 'a11y-btn--active': fontSize === 'normal' }"
        title="Ukuran Normal"
      >
        A
      </button>
      <button 
        @click="setFontSize('lg')" 
        class="a11y-btn" 
        :class="{ 'a11y-btn--active': fontSize === 'lg' }"
        title="Ukuran Besar"
      >
        A+
      </button>
      <button 
        @click="setFontSize('xl')" 
        class="a11y-btn" 
        :class="{ 'a11y-btn--active': fontSize === 'xl' }"
        title="Ukuran Sangat Besar"
      >
        A++
      </button>
    </div>
  </div>
</template>

<style scoped>
.a11y-toolbar {
  position: fixed;
  bottom: 4rem;
  left: 3rem;
  z-index: 9999;
  pointer-events: none; /* Klik tembus ke belakang jika bukan pada tombol */
}

.a11y-toolbar__buttons {
  display: flex;
  flex-direction: column; /* Dibuat vertikal agar lebih hemat ruang */
  gap: 1rem;
  pointer-events: auto;
}

.a11y-btn {
  width: 3.2rem;
  height: 3.2rem;
  border-radius: 50%;
  border: 1px solid rgba(0, 0, 0, 0.05);
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(5px);
  cursor: pointer;
  font-weight: 800;
  font-size: 1rem;
  color: #555;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 0.5rem 1.5rem rgba(0, 0, 0, 0.1);
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

/* Efek Membesar saat Hover */
.a11y-btn:hover {
  transform: scale(1.4);
  background: #fff;
  border-color: #55c57a;
  color: #55c57a;
  z-index: 10;
}

/* Mode Aktif (Tetap Kecil tapi berwarna) */
.a11y-btn--active {
  background: #55c57a;
  color: #fff !important;
  border-color: #55c57a;
  box-shadow: 0 0.5rem 1.5rem rgba(85, 197, 122, 0.4);
}

/* Penyesuaian ukuran teks dalam tombol */
.a11y-btn--active {
  transform: scale(1); /* Kembali mengecil saat aktif */
}

@media (max-width: 768px) {
  .a11y-toolbar {
    bottom: 10rem; /* Naik sedikit agar tidak tertutup nav mobile */
    left: 2rem;
  }
}
</style>
