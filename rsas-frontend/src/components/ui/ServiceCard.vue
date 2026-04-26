<script setup>
defineProps({
  icon: String,
  title: String,
  text: String,
  link: {
    type: String,
    default: "#"
  },
  delay: {
    type: Number,
    default: 0
  }
});
</script>

<template>
  <div class="service-card" data-aos="fade-up" :data-aos-delay="delay">
    <div class="service-card__icon-box">
      <i :class="icon" class="service-card__icon"></i>
    </div>
    <h3 class="service-card__title">{{ title }}</h3>
    <p class="service-card__text">{{ text }}</p>
    <router-link :to="link" class="service-card__link">Selengkapnya &rarr;</router-link>
  </div>
</template>

<style scoped>
.service-card {
  background-color: #fff;
  padding: 3.5rem 2.8rem; /* 2.5rem * 1.4, 2rem * 1.4 */
  text-align: center;
  border-radius: 15px;
  box-shadow: 0 1rem 3rem rgba(0, 0, 0, 0.05);
  transition: all 0.4s;
  border-bottom: 4px solid transparent;
}

.service-card:hover {
  transform: translateY(-1.4rem); /* 1rem * 1.4 */
  box-shadow: 0 2rem 4rem rgba(0, 0, 0, 0.1);
  border-bottom: 4px solid #55c57a;
}

.service-card__icon-box {
  width: 8.4rem; /* 6rem * 1.4 */
  height: 8.4rem;
  background-color: rgba(85, 197, 122, 0.1);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 2.8rem;
  transition: all 0.4s;
}

.service-card:hover .service-card__icon-box {
  background-color: #55c57a;
  transform: rotateY(180deg);
}

.service-card__icon {
  font-size: 3.5rem; /* 2.5rem * 1.4 */
  color: #55c57a;
  transition: all 0.4s;
}

.service-card:hover .service-card__icon {
  color: #fff;
}

.service-card__title {
  font-size: 2.2rem; /* 1.6rem * 1.4 */
  font-weight: 700;
  margin-bottom: 1.4rem;
  color: #333;
}

.service-card__text {
  font-size: 1.5rem; /* 1.1rem * 1.4 */
  line-height: 1.7;
  color: #777;
  margin-bottom: 2.1rem;
}

.service-card__link {
  font-size: 1.3rem;
  font-weight: 700;
  color: #55c57a;
  text-decoration: none;
  text-transform: uppercase;
  letter-spacing: 1px;
  transition: all 0.3s;
}

.service-card__link:hover {
  color: #40945c;
  gap: 1rem;
}
</style>
