<script setup>
import TheNavbar from "./components/layout/TheNavbar.vue";
import TheFooter from "./components/layout/TheFooter.vue";
import AccessibilityToolbar from "./components/ui/AccessibilityToolbar.vue";
import FloatingEmergency from "./components/ui/FloatingEmergency.vue";
import AdminSidebar from "./components/admin/AdminSidebar.vue";
import { onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import AOS from 'aos';

const route = useRoute();
const isAuthOrAdmin = computed(() => {
  return route.path.startsWith('/admin') || route.path === '/login';
});

const isAdminArea = computed(() => {
  return route.path.startsWith('/admin');
});

onMounted(() => {
  AOS.init({
    duration: 1000,
    once: true,
    offset: 100,
  });
});
</script>

<template>
  <div class="app-container" :class="{ 'is-admin': isAdminArea }">
    <TheNavbar v-if="!isAuthOrAdmin" />
    <AdminSidebar v-if="isAdminArea" />
    
    <main class="main-content">
      <router-view />
    </main>
    
    <TheFooter v-if="!isAuthOrAdmin" />
    
    <!-- Floating Tools -->
    <AccessibilityToolbar v-if="!isAuthOrAdmin" />
    <FloatingEmergency v-if="!isAuthOrAdmin" />
  </div>
</template>

<style>
/* Global Layout styles */
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;

  &.is-admin {
    flex-direction: row;

    .main-content {
      margin-left: 28rem; 
      width: calc(100% - 28rem);
      background: #f8f9fa;
      min-height: 100vh;
    }
  }
}

.main-content {
  flex: 1;
}

/* Ensure smooth scrolling */
html {
  scroll-behavior: smooth;
}
</style>
