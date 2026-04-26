<script setup>
defineProps({
  items: {
    type: Array,
    required: true,
    // Format: [{ label: 'Dashboard', link: '/admin/dashboard' }, { label: 'Berita' }]
  }
});
</script>

<template>
  <nav class="admin-breadcrumbs">
    <ul>
      <li>
        <router-link to="/admin/dashboard">
          <i class="fa-solid fa-house"></i>
        </router-link>
      </li>
      <li v-for="(item, index) in items" :key="index">
        <span class="separator">/</span>
        <router-link v-if="item.link" :to="item.link">{{ item.label }}</router-link>
        <span v-else class="current">{{ item.label }}</span>
      </li>
    </ul>
  </nav>
</template>

<style lang="scss" scoped>
.admin-breadcrumbs {
  margin-bottom: 1.5rem;

  ul {
    display: flex;
    align-items: center;
    list-style: none;
    padding: 0;
    gap: 1rem;
  }

  li {
    display: flex;
    align-items: center;
    gap: 1rem;
    font-size: 1.4rem;
    font-weight: 600;
    color: #aaa;

    a {
      text-decoration: none;
      color: #888;
      transition: all 0.3s;

      &:hover {
        color: #55c57a;
      }
    }

    i {
      font-size: 1.2rem;
    }

    .separator {
      color: #ddd;
      font-weight: 300;
    }

    .current {
      color: #55c57a;
      font-weight: 700;
    }
  }
}
</style>
