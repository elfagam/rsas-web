<script setup>
defineProps({
  steps: {
    type: Array,
    required: true
  }
});
</script>

<template>
  <div class="step-flow">
    <div v-for="(step, index) in steps" :key="index" class="step-item" data-aos="fade-up" :data-aos-delay="index * 100">
      <div class="step-item__number">0{{ index + 1 }}</div>
      <div class="step-item__content">
        <h3 class="step-item__title">{{ step.title }}</h3>
        <p class="step-item__desc">{{ step.desc }}</p>
      </div>
      <div v-if="index < steps.length - 1" class="step-item__arrow">
        <i class="fa-solid fa-chevron-right"></i>
      </div>
    </div>
  </div>
</template>

<style scoped>
.step-flow {
  display: flex;
  justify-content: space-between;
  gap: 2rem;
  margin-top: 4rem;
}

.step-item {
  flex: 1;
  background: #fff;
  padding: 4rem 3rem;
  border-radius: 20px;
  box-shadow: 0 1.5rem 4rem rgba(0, 0, 0, 0.05);
  position: relative;
  text-align: center;
  transition: all 0.3s;
  border: 1px solid #f1f1f1;
}

.step-item:hover {
  transform: translateY(-1rem);
  border-color: #55c57a;
}

.step-item__number {
  font-size: 5rem;
  font-weight: 900;
  color: rgba(85, 197, 122, 0.15);
  position: absolute;
  top: 1rem;
  right: 2rem;
}

.step-item__title {
  font-size: 2rem;
  font-weight: 800;
  color: #333;
  margin-bottom: 1.5rem;
  position: relative;
}

.step-item__desc {
  font-size: 1.5rem;
  color: #777;
  line-height: 1.6;
}

.step-item__arrow {
  position: absolute;
  top: 50%;
  right: -3rem;
  transform: translateY(-50%);
  font-size: 2rem;
  color: #55c57a;
}

@media (max-width: 1000px) {
  .step-flow { flex-direction: column; gap: 4rem; }
  .step-item__arrow {
    top: auto;
    bottom: -3rem;
    right: 50%;
    transform: translateX(50%) rotate(90deg);
  }
}
</style>
