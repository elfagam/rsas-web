<script setup>
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const isHome = computed(() => route.name === "home" || route.path === "/");
</script>

<template>
  <nav class="navbar" :class="{ 'navbar--solid': !isHome }">
    <div class="navbar__logo-box">
      <img
        v-if="isHome"
        src="../../assets/img/logo-white.png"
        alt="Logo"
        class="navbar__logo"
      />
      <img
        v-else
        src="../../assets/img/logo-green-1x.png"
        alt="Logo"
        class="navbar__logo"
      />
    </div>

    <input type="checkbox" class="navbar__checkbox" id="navi-toggle" />
    <label for="navi-toggle" class="navbar__button">
      <span class="navbar__icon">&nbsp;</span>
    </label>

    <ul class="navbar__list">
      <li class="navbar__item">
        <router-link to="/" class="navbar__link">Home</router-link>
      </li>
      <li class="navbar__item">
        <router-link to="/berita" class="navbar__link">Berita</router-link>
      </li>
      <li class="navbar__item"><a href="#" class="navbar__link">Profil</a></li>
      <li class="navbar__item"><a href="#" class="navbar__link">Layanan</a></li>
      <li class="navbar__item">
        <a href="#" class="navbar__link">Informasi</a>
      </li>
      <li class="navbar__item">
        <a href="#" class="navbar__link">Pendidikan dan Pelatihan</a>
      </li>
      <li class="navbar__item">
        <a href="#" class="navbar__link">Zona Integritas</a>
      </li>
      <li class="navbar__item">
        <a href="#" class="navbar__link">Pengumuman</a>
      </li>
      <li class="navbar__item">
        <a href="#" class="navbar__link">Hubungi Kami</a>
      </li>
      <li class="navbar__item"><a href="#" class="navbar__link">FAQ</a></li>
    </ul>
  </nav>
</template>

<style scoped>
.navbar {
  position: absolute !important;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 999 !important; /* lapisan teratas */
  background-color: transparent !important;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 2rem 5rem;
}

.navbar__link {
  color: #fff !important;
  text-decoration: none;
  font-weight: 700;
  margin-left: 3rem;
  font-size: 1.4rem;
}

/* Pastikan ini memiliki dua tanda hubung agar sesuai dengan template di atas */
.navbar.navbar--solid {
  position: relative !important;
  background-color: #f8f3e1 !important;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  /* padding-bottom: 1.5rem; */
}

/* Teks menjadi gelap saat mode solid aktif */
.navbar.navbar--solid .navbar__link {
  color: #333 !important;
}
</style>
