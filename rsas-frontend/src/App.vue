<script setup>
import TheNavbar from "./components/layout/TheNavbar.vue";
import TheFooter from "./components/layout/TheFooter.vue";
import AccessibilityToolbar from "./components/ui/AccessibilityToolbar.vue";
import FloatingEmergency from "./components/ui/FloatingEmergency.vue";
import { onMounted } from 'vue';
import AOS from 'aos';

onMounted(() => {
  AOS.init({
    duration: 1000,
    once: true,
    offset: 100,
  });
});
</script>

<template>
  <div class="app-container">
    <TheNavbar />
    <main class="main-content">
      <router-view />
    </main>
    <TheFooter />
    
    <!-- Floating Tools -->
    <AccessibilityToolbar />
    <FloatingEmergency />
  </div>
</template>

<style>
/* Global Layout styles */
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.main-content {
  flex: 1;
}

/* Ensure smooth scrolling */
html {
  scroll-behavior: smooth;
}
</style>
