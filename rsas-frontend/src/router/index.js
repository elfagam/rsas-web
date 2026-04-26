import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import BeritaView from "../views/BeritaView.vue";
import BeritaDetailView from "../views/BeritaDetailView.vue";
import ProfilView from "../views/ProfilView.vue";
import LayananDetailView from "../views/LayananDetailView.vue";
import LayananView from "../views/LayananView.vue";
import DokterView from "../views/DokterView.vue";
import PendidikanView from "../views/PendidikanView.vue";
import ZonaIntegritasView from "../views/ZonaIntegritasView.vue";
import PengaduanView from "../views/PengaduanView.vue";
import FaqView from "../views/FaqView.vue";
import GaleriView from "../views/GaleriView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/profil",
      name: "profil",
      component: ProfilView,
    },
    {
      path: "/layanan",
      name: "layanan",
      component: LayananView,
    },
    {
      path: "/dokter",
      name: "dokter",
      component: DokterView,
    },
    {
      path: "/pendidikan",
      name: "pendidikan",
      component: PendidikanView,
    },
    {
      path: "/zona-integritas",
      name: "zona-integritas",
      component: ZonaIntegritasView,
    },
    {
      path: "/pengaduan",
      name: "pengaduan",
      component: PengaduanView,
    },
    {
      path: "/faq",
      name: "faq",
      component: FaqView,
    },
    {
      path: "/galeri",
      name: "galeri",
      component: GaleriView,
    },
    {
      path: "/berita",
      name: "berita",
      component: BeritaView,
    },
    {
      path: "/berita/:id",
      name: "berita-detail",
      component: BeritaDetailView,
    },
    {
      path: "/layanan/:id",
      name: "layanan-detail",
      component: LayananDetailView,
    },
  ],
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return { top: 0, behavior: 'smooth' };
    }
  },
});

export default router;
