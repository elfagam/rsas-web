<script setup>
defineProps({
  kategori: String,
  tanggal: String,
  judul: String,
  ringkasan: String,
  gambar: String,
  tautan: {
    type: String,
    default: "#",
  },
});
</script>

<template>
  <div class="news-card">
    <div class="news-card__img-container">
      <img :src="gambar" :alt="judul" class="news-card__img" loading="lazy" />
    </div>

    <div class="news-card__content">
      <div class="news-card__meta">
        <span class="news-card__category">{{ kategori }}</span>
        <span class="news-card__divider">&bull;</span>
        <span class="news-card__date">{{ tanggal }}</span>
      </div>

      <h3 class="news-card__title">{{ judul }}</h3>
      <p class="news-card__text">{{ ringkasan }}</p>

      <router-link :to="tautan" class="news-card__link">
        Baca Selengkapnya <span>&rarr;</span>
      </router-link>
    </div>
  </div>
</template>

<style scoped>

.news-card {
  background-color: #fff;
  border-radius: 17px;
  overflow: hidden;
  box-shadow: 0 1.4rem 4.2rem rgba(0, 0, 0, 0.08);
  transition: all 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  display: flex;
  flex-direction: column;
  height: 100%;
}

.news-card:hover {
  transform: translateY(-1.4rem);
  box-shadow: 0 2.8rem 5.6rem rgba(0, 0, 0, 0.15);
}

.news-card__img-container {
  height: 25.2rem;
  overflow: hidden;
}

.news-card__img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s ease;
}

.news-card:hover .news-card__img {
  transform: scale(1.1);
}

.news-card__content {
  padding: 3.5rem;
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}

.news-card__meta {
  font-size: 1.5rem;
  color: #888;
  margin-bottom: 1.1rem;
  text-transform: uppercase;
  letter-spacing: 1.4px;
}

.news-card__category {
  color: #55c57a;
  font-weight: 700;
}

.news-card__divider {
  margin: 0 0.5rem;
}

.news-card__title {
  font-size: 2.2rem; /* 1.6rem * 1.4 */
  font-weight: 800;
  color: #333;
  line-height: 1.3;
  margin-bottom: 1.2rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.news-card__text {
  font-size: 1.5rem; /* 1.1rem * 1.4 */
  font-weight: 400;
  color: #777;
  line-height: 1.7;
  margin-bottom: 2.8rem;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  flex-grow: 1;
}

.news-card__link {
  font-size: 1.4rem;
  color: #55c57a;
  text-decoration: none;
  font-weight: 600;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  transition: all 0.3s;
}

.news-card__link span {
  transition: transform 0.3s ease;
}

.news-card__link:hover {
  color: #40945c;
}

.news-card__link:hover span {
  transform: translateX(5px);
}
</style>
