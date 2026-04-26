<script setup>
defineProps({
  name: String,
  position: String,
  nip: String,
  image: String,
  motto: String,
  delay: {
    type: Number,
    default: 0
  }
});
</script>

<template>
  <div class="mgmt-card" data-aos="fade-up" :data-aos-delay="delay">
    <div class="mgmt-card__img-box">
      <img v-if="image" :src="image" :alt="name" class="mgmt-card__img" />
      <div v-else class="mgmt-card__placeholder">
        <i class="fa-solid fa-user-tie"></i>
      </div>
      <div class="mgmt-card__overlay">
        <p class="mgmt-card__motto">"{{ motto || 'Melayani dengan Sepenuh Hati' }}"</p>
      </div>
    </div>
    <div class="mgmt-card__info">
      <h3 class="mgmt-card__name">{{ name }}</h3>
      <span class="mgmt-card__position">{{ position }}</span>
      <span v-if="nip" class="mgmt-card__nip">NIP. {{ nip }}</span>
    </div>
  </div>
</template>

<style scoped>
.mgmt-card {
  background: #fff;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 1.5rem 4rem rgba(0, 0, 0, 0.05);
  transition: all 0.4s ease;
  text-align: center;
  border: 1px solid #f0f0f0;
}

.mgmt-card:hover {
  transform: translateY(-1rem);
  box-shadow: 0 3rem 6rem rgba(0, 0, 0, 0.1);
  border-color: #55c57a;
}

.mgmt-card__img-box {
  position: relative;
  height: 33.6rem; /* 24rem * 1.4 */
  overflow: hidden;
  background-color: #f9fbf9;
}

.mgmt-card__img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: top;
  transition: transform 0.6s;
}

.mgmt-card:hover .mgmt-card__img {
  transform: scale(1.1);
}

.mgmt-card__placeholder {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 11.2rem; /* 8rem * 1.4 */
  color: #ddd;
}

.mgmt-card__overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  padding: 2.8rem;
  background: linear-gradient(to top, rgba(0,0,0,0.8), transparent);
  transform: translateY(100%);
  transition: all 0.4s;
  opacity: 0;
}

.mgmt-card:hover .mgmt-card__overlay {
  transform: translateY(0);
  opacity: 1;
}

.mgmt-card__motto {
  color: #fff;
  font-size: 1.8rem; /* 1.3rem * 1.4 */
  font-style: italic;
  font-weight: 500;
}

.mgmt-card__info {
  padding: 2.5rem 2.1rem; /* 1.8rem * 1.4, 1.5rem * 1.4 */
}

.mgmt-card__name {
  font-size: 2.1rem; /* 1.5rem * 1.4 */
  font-weight: 800;
  color: #333;
  margin-bottom: 0.7rem;
  line-height: 1.2;
}

.mgmt-card__position {
  display: block;
  font-size: 1.5rem; /* 1.1rem * 1.4 */
  color: #55c57a;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 1px;
  margin-bottom: 1.1rem;
}

.mgmt-card__nip {
  font-size: 1.4rem; /* 1rem * 1.4 */
  color: #999;
  font-weight: 500;
}
</style>
