<script setup>
defineProps({
  title: String,
  subtitle: String,
  white: {
    type: Boolean,
    default: false
  }
});
</script>

<template>
  <div class="section-header u-center-text u-margin-bottom-big">
    <h2 class="heading-secondary" :class="{ 'heading-secondary--white': white }">{{ title }}</h2>
    <p class="section-subtitle" v-if="subtitle" :class="{ 'section-subtitle--white': white }">{{ subtitle }}</p>
  </div>
</template>

<style scoped>
.u-center-text { text-align: center; }
.u-margin-bottom-big { margin-bottom: 8rem; }

.heading-secondary {
  font-size: 3.5rem;
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 1px;
  display: inline-block;
  position: relative;
  color: #333;
}

.heading-secondary--white { color: #fff; }

.section-subtitle {
  font-size: 1.8rem;
  color: #777;
  margin-top: 1.5rem;
  max-width: 70rem;
  margin-left: auto;
  margin-right: auto;
}

.section-subtitle--white { color: rgba(255, 255, 255, 0.8); }

@media (max-width: 600px) {
  .heading-secondary { font-size: 2.8rem; }
  .section-subtitle { font-size: 1.6rem; }
}
</style>
