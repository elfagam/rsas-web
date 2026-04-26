<script setup>
import { ref, onMounted, onUnmounted } from 'vue';

const isVisible = ref(true);
const lastScrollY = ref(0);

const handleScroll = () => {
  // Sembunyikan FAB saat scroll ke bawah, tunjukkan saat scroll ke atas (opsional)
  // Untuk RS, biasanya lebih baik selalu muncul (isVisible = true)
  isVisible.value = true;
};

onMounted(() => {
  window.addEventListener('scroll', handleScroll);
});

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll);
});
</script>

<template>
  <transition name="bounce">
    <div v-if="isVisible" class="fab-emergency">
      <a href="tel:0435821010" class="fab-emergency__link">
        <div class="fab-emergency__pulse"></div>
        <i class="fa-solid fa-phone-volume"></i>
        <span class="fab-emergency__label">IGD</span>
      </a>
    </div>
  </transition>
</template>

<style scoped>
.fab-emergency {
  position: fixed;
  bottom: 3rem;
  right: 3rem;
  z-index: 999;
  display: none; /* Default sembunyi di desktop */
}

@media (max-width: 1024px) {
  .fab-emergency {
    display: block; /* Muncul di tablet & mobile */
  }
}

.fab-emergency__link {
  width: 7rem;
  height: 7rem;
  background: linear-gradient(135deg, #ff5252 0%, #b33939 100%);
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #fff;
  text-decoration: none;
  box-shadow: 0 1rem 3rem rgba(255, 82, 82, 0.4);
  position: relative;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.fab-emergency__link:active {
  transform: scale(0.9);
}

.fab-emergency__link i {
  font-size: 2.2rem;
  margin-bottom: 0.2rem;
}

.fab-emergency__label {
  font-size: 1rem;
  font-weight: 900;
  letter-spacing: 1px;
}

/* Pulse Animation */
.fab-emergency__pulse {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: #ff5252;
  border-radius: 50%;
  z-index: -1;
  animation: fab-pulse 2s infinite;
}

@keyframes fab-pulse {
  0% { transform: scale(1); opacity: 0.6; }
  100% { transform: scale(1.8); opacity: 0; }
}

/* Transition */
.bounce-enter-active {
  animation: bounce-in 0.5s;
}
.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}

@keyframes bounce-in {
  0% { transform: scale(0); opacity: 0; }
  50% { transform: scale(1.1); }
  100% { transform: scale(1); opacity: 1; }
}

@media (max-width: 600px) {
  .fab-emergency {
    bottom: 2rem;
    right: 2rem;
  }
  .fab-emergency__link {
    width: 6rem;
    height: 6rem;
  }
}
</style>
