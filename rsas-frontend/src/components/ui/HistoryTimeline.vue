<script setup>
defineProps({
  historyData: {
    type: Array,
    required: true
  }
});
</script>

<template>
  <div class="history-list">
    <div v-for="(item, index) in historyData" :key="index" class="history-item" data-aos="fade-up">
      <div class="history-item__year">{{ item.year }}</div>
      <div class="history-item__content">
        <div class="history-item__icon"><i :class="item.icon"></i></div>
        <h3 class="history-item__title">{{ item.title }}</h3>
        <p class="history-item__desc">{{ item.desc }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.history-list {
  max-width: 90rem;
  margin: 0 auto;
}

.history-item {
  display: flex;
  gap: 4rem;
  margin-bottom: 4rem;
}

.history-item__year {
  font-size: 3rem;
  font-weight: 900;
  color: #55c57a;
  flex: 0 0 12rem;
}

.history-item__content {
  background: #f9fbf9;
  padding: 3rem;
  border-radius: 15px;
  flex: 1;
  position: relative;
}

.history-item__icon {
  position: absolute;
  top: -2rem;
  right: 2rem;
  width: 5rem;
  height: 5rem;
  background: #55c57a;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  box-shadow: 0 1rem 2rem rgba(85, 197, 122, 0.3);
}

.history-item__title { font-size: 2.2rem; font-weight: 800; margin-bottom: 1rem; color: #333; }
.history-item__desc { font-size: 1.6rem; color: #666; line-height: 1.7; }

@media (max-width: 1000px) {
  .history-item { flex-direction: column; gap: 1rem; }
  .history-item__year { flex: none; }
}
</style>
