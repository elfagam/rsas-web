<script setup>
defineProps({
  name: String,
  unit: String,
  message: String,
  rating: {
    type: Number,
    default: 5
  },
  date: String
});

const getInitials = (name) => {
  return name.split(' ').map(n => n[0]).join('').toUpperCase().substring(0, 2);
};

const colors = ['#55c57a', '#2d3436', '#3498db', '#f1c40f', '#e67e22'];
const randomColor = colors[Math.floor(Math.random() * colors.length)];
</script>

<template>
  <div class="testimonial-card">
    <div class="testimonial-card__header">
      <div class="testimonial-card__avatar" :style="{ backgroundColor: randomColor }">
        {{ getInitials(name) }}
      </div>
      <div class="testimonial-card__info">
        <h4 class="testimonial-card__name">{{ name }}</h4>
        <p class="testimonial-card__unit">{{ unit }}</p>
      </div>
      <div class="testimonial-card__rating">
        <i v-for="i in 5" :key="i" 
           class="fa-star" 
           :class="i <= rating ? 'fa-solid' : 'fa-regular'"></i>
      </div>
    </div>
    
    <div class="testimonial-card__content">
      <i class="fa-solid fa-quote-left testimonial-card__quote"></i>
      <p class="testimonial-card__text">{{ message }}</p>
    </div>
    
    <div class="testimonial-card__footer">
      <span class="testimonial-card__date">{{ date }}</span>
      <span class="testimonial-card__verified"><i class="fa-solid fa-circle-check"></i> Pasien Terverifikasi</span>
    </div>
  </div>
</template>

<style scoped>
.testimonial-card {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  padding: 4.5rem;
  border-radius: 30px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 2rem 5rem rgba(0, 0, 0, 0.04);
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  display: flex;
  flex-direction: column;
  gap: 3rem;
  height: 100%;
  position: relative;
}

.testimonial-card:hover {
  transform: translateY(-1rem);
  box-shadow: 0 2.5rem 5rem rgba(85, 197, 122, 0.15);
  background: #fff;
}

.testimonial-card__header {
  display: flex;
  align-items: center;
  gap: 2rem;
  position: relative;
}

.testimonial-card__avatar {
  width: 6rem;
  height: 6rem;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-weight: 800;
  font-size: 1.8rem;
  box-shadow: 0 0.5rem 1.5rem rgba(0, 0, 0, 0.1);
}

.testimonial-card__info {
  flex: 1;
}

.testimonial-card__name {
  font-size: 1.8rem;
  font-weight: 800;
  color: #333;
  margin-bottom: 0.3rem;
}

.testimonial-card__unit {
  font-size: 1.3rem;
  color: #55c57a;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.testimonial-card__rating {
  display: flex;
  gap: 0.4rem;
  color: #f1c40f;
  font-size: 1.4rem;
}

.testimonial-card__content {
  position: relative;
  flex: 1;
}

.testimonial-card__quote {
  font-size: 3rem;
  color: rgba(85, 197, 122, 0.1);
  position: absolute;
  top: -1.5rem;
  left: -1.5rem;
}

.testimonial-card__text {
  font-size: 1.6rem;
  line-height: 1.9;
  color: #555;
  font-style: italic;
  position: relative;
  z-index: 1;
}

.testimonial-card__footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: auto;
  padding-top: 1.5rem;
  border-top: 1px solid #f0f0f0;
}

.testimonial-card__date {
  font-size: 1.3rem;
  color: #999;
}

.testimonial-card__verified {
  font-size: 1.2rem;
  color: #27ae60;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}
</style>
