<script setup>
import { ref } from 'vue';

defineProps({
  name: String,
  specialty: String,
  expertise: String,
  time: String,
  image: String,
  isAvailable: Boolean,
  quota: Number,
  remaining: Number,
});

const showDetails = ref(false);

const toggleDetails = () => {
  showDetails.value = !showDetails.value;
};
</script>

<template>
  <div class="doctor-card" @click="toggleDetails">
    <!-- Main Content: Always Visible -->
    <div class="doctor-card__main">
      <div class="doctor-card__img-box">
        <img v-if="image" :src="image" :alt="name" class="doctor-card__img" />
        <div v-else class="doctor-card__placeholder">
          <i class="fa-solid fa-user-doctor"></i>
        </div>
        
        <!-- Quick Status Badge on Image -->
        <div class="doctor-card__status-badge" :class="isAvailable ? 'status--available' : 'status--full'">
          {{ isAvailable ? 'Tersedia' : 'Penuh' }}
        </div>
      </div>

      <div class="doctor-card__info">
        <h3 class="doctor-card__name">{{ name }}</h3>
        <span class="doctor-card__specialty">{{ specialty }}</span>
        <div class="doctor-card__hint">
          <span>Klik untuk jadwal</span>
          <i class="fa-solid fa-chevron-up"></i>
        </div>
      </div>
    </div>

    <!-- Slide-Up Details Overlay -->
    <div class="doctor-card__overlay" :class="{ 'is-active': showDetails }">
      <div class="doctor-card__overlay-header">
        <div class="doctor-card__close-btn">
          <i class="fa-solid fa-chevron-down"></i>
        </div>
        <h3 class="overlay-name">{{ name }}</h3>
        <p class="overlay-specialty">{{ specialty }}</p>
      </div>

      <div class="doctor-card__overlay-body">
        <div class="detail-item">
          <label>Keahlian Khusus</label>
          <p>{{ expertise || 'Umum' }}</p>
        </div>

        <div class="detail-item">
          <label>Jadwal Praktik</label>
          <div class="detail-time">
            <i class="fa-regular fa-clock"></i> {{ time }}
          </div>
        </div>

        <div class="quota-container" v-if="isAvailable">
          <div class="quota-header">
            <span>Sisa Kuota</span>
            <span><strong>{{ remaining }}</strong> / {{ quota }}</span>
          </div>
          <div class="quota-bar-bg">
            <div class="quota-bar-fill" :style="{ width: (remaining / quota * 100) + '%' }"></div>
          </div>
        </div>

        <button class="btn-booking" :disabled="!isAvailable" @click.stop>
          {{ isAvailable ? 'Daftar Sekarang' : 'Kuota Penuh' }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.doctor-card {
  position: relative;
  background: #fff;
  border-radius: 24px;
  overflow: hidden;
  box-shadow: 0 1.5rem 4rem rgba(0, 0, 0, 0.08);
  height: 48rem;
  cursor: pointer;
  border: 1px solid #f1f1f1;
  transition: all 0.3s ease;
}

.doctor-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 2rem 5rem rgba(0, 0, 0, 0.12);
}

.doctor-card__main {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.doctor-card__img-box {
  height: 70%;
  position: relative;
  overflow: hidden;
  background: #f9fbf9;
}

.doctor-card__img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: top;
  transition: transform 0.5s ease;
}

.doctor-card:hover .doctor-card__img {
  transform: scale(1.05);
}

.doctor-card__placeholder {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 8rem;
  color: #e0e0e0;
}

.doctor-card__status-badge {
  position: absolute;
  top: 1.5rem;
  right: 1.5rem;
  padding: 0.5rem 1.5rem;
  border-radius: 50px;
  font-size: 1.1rem;
  font-weight: 800;
  text-transform: uppercase;
  backdrop-filter: blur(8px);
  box-shadow: 0 5px 15px rgba(0,0,0,0.1);
}

.status--available { background: rgba(85, 197, 122, 0.9); color: #fff; }
.status--full { background: rgba(235, 77, 75, 0.9); color: #fff; }

.doctor-card__info {
  padding: 2.5rem;
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.doctor-card__name {
  font-size: 2rem;
  font-weight: 900;
  color: #2d3436;
  margin-bottom: 0.5rem;
  line-height: 1.2;
}

.doctor-card__specialty {
  font-size: 1.4rem;
  color: #55c57a;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.doctor-card__hint {
  margin-top: 1.5rem;
  font-size: 1.2rem;
  color: #b2bec3;
  display: flex;
  align-items: center;
  gap: 0.8rem;
  font-weight: 600;
}

/* Slide-Up Overlay */
.doctor-card__overlay {
  position: absolute;
  top: 100%; /* Hidden below */
  left: 0;
  width: 100%;
  height: 100%;
  background: #fff;
  z-index: 10;
  transition: top 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  padding: 3rem;
  display: flex;
  flex-direction: column;
}

.doctor-card__overlay.is-active {
  top: 0; /* Slide up to cover */
}

.doctor-card__overlay-header {
  text-align: center;
  margin-bottom: 3rem;
  padding-bottom: 2rem;
  border-bottom: 1px solid #f1f1f1;
}

.doctor-card__close-btn {
  width: 4rem;
  height: 4rem;
  background: #f9fbf9;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: -1.5rem auto 1.5rem;
  color: #55c57a;
  font-size: 1.6rem;
}

.overlay-name {
  font-size: 1.8rem;
  font-weight: 900;
  color: #2d3436;
}

.overlay-specialty {
  font-size: 1.2rem;
  color: #55c57a;
  font-weight: 700;
  text-transform: uppercase;
}

.doctor-card__overlay-body {
  flex: 1;
}

.detail-item {
  margin-bottom: 2rem;
}

.detail-item label {
  display: block;
  font-size: 1rem;
  text-transform: uppercase;
  color: #b2bec3;
  font-weight: 800;
  margin-bottom: 0.5rem;
  letter-spacing: 1px;
}

.detail-item p {
  font-size: 1.5rem;
  color: #2d3436;
  font-weight: 600;
}

.detail-time {
  font-size: 1.6rem;
  color: #55c57a;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 1rem;
}

.quota-container {
  background: #f9fbf9;
  padding: 1.5rem;
  border-radius: 12px;
  margin-bottom: 3rem;
}

.quota-header {
  display: flex;
  justify-content: space-between;
  font-size: 1.3rem;
  margin-bottom: 1rem;
  color: #636e72;
}

.quota-header strong { color: #55c57a; }

.quota-bar-bg {
  height: 6px;
  background: #e0e0e0;
  border-radius: 10px;
  overflow: hidden;
}

.quota-bar-fill {
  height: 100%;
  background: linear-gradient(to right, #55c57a, #33ab5a);
}

.btn-booking {
  width: 100%;
  padding: 1.5rem;
  border-radius: 12px;
  border: none;
  background: linear-gradient(135deg, #55c57a 0%, #33ab5a 100%);
  color: #fff;
  font-size: 1.5rem;
  font-weight: 800;
  cursor: pointer;
  box-shadow: 0 1rem 2rem rgba(85, 197, 122, 0.2);
}

.btn-booking:disabled {
  background: #dfe6e9;
  box-shadow: none;
  color: #b2bec3;
}
</style>
