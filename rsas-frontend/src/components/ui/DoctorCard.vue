<script setup>
import { ref } from 'vue';

const props = defineProps({
  name: String,
  specialist: String, // Contoh: Spesialis Penyakit Dalam
  specialistType: String, // Contoh: Spesialistik Dasar / Spesialistik Penunjang
  polyclinic: String, // Contoh: Poliklinik Penyakit Dalam
  image: String,
  education: String,
  schedules: {
    type: Array,
    default: () => [] 
  },
  isOnline: Boolean
});

const isFlipped = ref(false);
const days = ['Sen', 'Sel', 'Rab', 'Kam', 'Jum', 'Sab', 'Min'];

const toggleFlip = () => {
  isFlipped.value = !isFlipped.value;
};

const getScheduleByDay = (dayShort) => {
  const dayMap = { 'Sen': 'Senin', 'Sel': 'Selasa', 'Rab': 'Rabu', 'Kam': 'Kamis', 'Jum': 'Jumat', 'Sab': 'Sabtu', 'Min': 'Minggu' };
  return props.schedules.find(s => s.day === dayMap[dayShort]);
};
</script>

<template>
  <div class="doctor-card-container" @click="toggleFlip">
    <div class="doctor-card-inner" :class="{ 'is-flipped': isFlipped }">
      
      <!-- FRONT SIDE: The Profile -->
      <div class="doctor-card-front">
        <div class="doctor-card-front__img-wrapper">
          <img :src="image || 'https://images.unsplash.com/photo-1612349317150-e413f6a5b16d?auto=format&fit=crop&q=80&w=600'" :alt="name">
          <div class="overlay-info">
            <span class="specialist-badge">{{ specialist }}</span>
          </div>
          <div v-if="isOnline" class="live-indicator">● Sedang Praktek</div>
        </div>
        
        <div class="doctor-card-front__body">
          <h3 class="name">{{ name }}</h3>
          <div class="polyclinic-info">
            <i class="fa-solid fa-location-dot"></i> 
            {{ polyclinic || (specialist ? 'Poliklinik ' + specialist.replace('Spesialis ', '') : 'Poliklinik RSAS') }}
          </div>
          <p class="click-hint"><i class="fa-solid fa-rotate"></i> Klik untuk Detail</p>
        </div>
      </div>

      <!-- BACK SIDE: The Details -->
      <div class="doctor-card-back">
        <div class="doctor-card-back__header">
          <div class="header-top">
            <span class="type-tag">{{ specialistType || 'Spesialistik' }}</span>
            <h4>Jadwal & Detail</h4>
          </div>
          <span class="back-name">{{ name }}</span>
          <div class="back-poly">{{ polyclinic }}</div>
        </div>

        <div class="doctor-card-back__content">
          <div class="schedule-section">
            <p class="label">Jadwal Praktek Mingguan:</p>
            <div class="schedule-grid">
              <div 
                v-for="day in days" 
                :key="day" 
                class="day-box"
                :class="{ 'day-box--active': getScheduleByDay(day) }"
              >
                <span class="day">{{ day }}</span>
                <div v-if="getScheduleByDay(day)" class="time-info">
                  {{ getScheduleByDay(day).time }}
                </div>
                <div v-else class="off-mark">-</div>
              </div>
            </div>
          </div>

          <div class="info-section">
            <p class="label">Informasi Tambahan:</p>
            <p class="info-text">{{ education || 'Dokter berpengalaman dengan komitmen tinggi pada kesehatan pasien RSAS.' }}</p>
          </div>
        </div>

        <div class="doctor-card-back__footer">
          <button class="btn-action">Daftar Antrian Online</button>
          <span class="tap-back">Klik untuk kembali &larr;</span>
        </div>
      </div>

    </div>
  </div>
</template>

<style lang="scss" scoped>
.doctor-card-container {
  background-color: transparent;
  width: 100%;
  height: 50rem;
  perspective: 150rem; /* Memberikan efek 3D */
  cursor: pointer;
}

.doctor-card-inner {
  position: relative;
  width: 100%;
  height: 100%;
  text-align: center;
  transition: transform 0.8s cubic-bezier(0.2, 0.8, 0.2, 1);
  transform-style: preserve-3d;
  box-shadow: 0 2rem 5rem rgba(0, 0, 0, 0.1);
  border-radius: 30px;

  &.is-flipped {
    transform: rotateY(180deg);
  }
}

/* Common sides styling */
.doctor-card-front, .doctor-card-back {
  position: absolute;
  width: 100%;
  height: 100%;
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
  border-radius: 30px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* --- FRONT SIDE --- */
.doctor-card-front {
  background: #fff;
  
  &__img-wrapper {
    position: relative;
    height: 75%;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      transition: transform 0.5s;
    }

    .overlay-info {
      position: absolute;
      bottom: 0;
      left: 0;
      width: 100%;
      padding: 3rem 2rem 1.5rem;
      background: linear-gradient(to top, rgba(0,0,0,0.7), transparent);
      text-align: left;

      .specialist-badge {
        background: #55c57a;
        color: #fff;
        padding: 0.6rem 1.5rem;
        border-radius: 50px;
        font-size: 1.1rem;
        font-weight: 800;
        text-transform: uppercase;
      }
    }

    .live-indicator {
      position: absolute;
      top: 2rem;
      right: 2rem;
      background: rgba(255,255,255,0.9);
      color: #55c57a;
      padding: 0.5rem 1.2rem;
      border-radius: 50px;
      font-size: 1rem;
      font-weight: 800;
      box-shadow: 0 0.5rem 1.5rem rgba(0,0,0,0.1);
      animation: pulse-green 2s infinite;
    }
  }

  &__body {
    height: 25%;
    padding: 2rem;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

      .name {
      font-family: 'Outfit', sans-serif;
      font-size: 2rem;
      font-weight: 900;
      color: #333;
      margin-bottom: 0.5rem;
    }

    .polyclinic-info {
      font-size: 1.3rem;
      font-weight: 700;
      color: #55c57a;
      margin-bottom: 1.5rem;
      display: flex;
      align-items: center;
      gap: 0.8rem;
      i { font-size: 1.2rem; opacity: 0.8; }
    }

    .click-hint {
      font-size: 1.2rem;
      color: #999;
      font-weight: 600;
      i { margin-right: 0.5rem; color: #55c57a; }
    }
  }

  &:hover &__img-wrapper img {
    transform: scale(1.1);
  }
}

/* --- BACK SIDE --- */
.doctor-card-back {
  background: #fff;
  transform: rotateY(180deg);
  padding: 3.5rem;
  background: linear-gradient(135deg, #f8f9fa 0%, #ffffff 100%);
  text-align: left;

  &__header {
    margin-bottom: 3rem;
    
    .header-top {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 1.5rem;

      .type-tag {
        font-size: 1rem;
        background: #eee;
        color: #777;
        padding: 0.4rem 1rem;
        border-radius: 6px;
        font-weight: 800;
        text-transform: uppercase;
      }

      h4 { color: #55c57a; font-size: 1.1rem; text-transform: uppercase; font-weight: 900; margin: 0; }
    }
    
    .back-name { font-family: 'Outfit', sans-serif; font-size: 2.1rem; font-weight: 900; color: #333; display: block; margin-bottom: 0.5rem; }
    .back-poly { font-size: 1.4rem; font-weight: 700; color: #55c57a; }
  }

  &__content {
    flex-grow: 1;

    .label { font-size: 1.3rem; font-weight: 800; color: #444; margin-bottom: 1.5rem; }

    .schedule-grid {
      display: grid;
      grid-template-columns: repeat(2, 1fr);
      gap: 1rem;
      margin-bottom: 3rem;

      .day-box {
        background: #fff;
        padding: 1.2rem;
        border-radius: 12px;
        border: 1px solid #eee;
        display: flex;
        justify-content: space-between;
        align-items: center;

        .day { font-weight: 800; color: #777; font-size: 1.2rem; }
        .time-info { color: #55c57a; font-weight: 800; font-size: 1.1rem; }
        .off-mark { color: #ddd; }

        &--active {
          border-color: #55c57a;
          background: rgba(85, 197, 122, 0.05);
          .day { color: #333; }
        }
      }
    }

    .info-text {
      font-size: 1.3rem;
      color: #666;
      line-height: 1.6;
      font-style: italic;
    }
  }

  &__footer {
    margin-top: 2rem;
    display: flex;
    flex-direction: column;
    gap: 1.5rem;

    .btn-action {
      background: #55c57a;
      color: #fff;
      border: none;
      padding: 1.5rem;
      border-radius: 15px;
      font-weight: 800;
      font-size: 1.5rem;
      cursor: pointer;
      box-shadow: 0 1rem 2rem rgba(85, 197, 122, 0.3);
      transition: 0.3s;
      &:hover { background: #2d8a4d; transform: scale(1.02); }
    }

    .tap-back {
      font-size: 1.1rem;
      color: #bbb;
      text-align: center;
      font-weight: 700;
    }
  }
}

@keyframes pulse-green {
  0% { box-shadow: 0 0 0 0 rgba(85, 197, 122, 0.4); }
  70% { box-shadow: 0 0 0 10px rgba(85, 197, 122, 0); }
  100% { box-shadow: 0 0 0 0 rgba(85, 197, 122, 0); }
}
</style>
