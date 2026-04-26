<script setup>
defineProps({
  title: String,
  desc: String,
  icon: String,
  target: String,
  delay: {
    type: Number,
    default: 0
  }
});
</script>

<template>
  <div class="renstra-card" data-aos="fade-up" :data-aos-delay="delay">
    <div class="renstra-card__header">
      <span class="renstra-card__target">Target {{ target }}</span>
      <div class="renstra-card__icon"><i :class="icon"></i></div>
    </div>
    <h3 class="renstra-card__title">{{ title }}</h3>
    <p class="renstra-card__desc">{{ desc }}</p>
  </div>
</template>

<style scoped>
.renstra-card {
  background: #fff;
  padding: 3.5rem 2.8rem; /* 2.5rem * 1.4, 2rem * 1.4 */
  border-radius: 20px;
  box-shadow: 0 1rem 3rem rgba(0,0,0,0.04);
  border: 1px solid rgba(85, 197, 122, 0.1);
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  height: 100%;
}

.renstra-card:hover {
  transform: translateY(-1.4rem); /* 1rem * 1.4 */
  box-shadow: 0 2rem 4rem rgba(85, 197, 122, 0.12);
  border-color: #55c57a;
}

.renstra-card__header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 2.8rem;
}

.renstra-card__target {
  font-size: 1.6rem; /* 1.2rem * 1.4 */
  font-weight: 800;
  color: #55c57a;
  background: rgba(85, 197, 122, 0.1);
  padding: 0.7rem 1.6rem;
  border-radius: 50px;
  text-transform: uppercase;
}

.renstra-card__icon {
  font-size: 4.2rem; /* 3rem * 1.4 */
  color: #55c57a;
}

.renstra-card__title {
  font-size: 2.2rem; /* 1.6rem * 1.4 */
  font-weight: 800;
  color: #1a1a1a;
  margin-bottom: 1.6rem;
  line-height: 1.3;
}

.renstra-card__desc {
  font-size: 1.5rem; /* 1.1rem * 1.4 */
  color: #666;
  line-height: 1.7;
}
</style>
