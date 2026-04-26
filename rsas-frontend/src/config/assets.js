/**
 * RSAS Web Assets Configuration
 * 
 * Pusat pengelolaan URL gambar dan dokumen untuk seluruh website.
 * Saat migrasi ke Google Cloud Storage (GCP), cukup ubah BASE_URL 
 * ke URL bucket Anda (misal: https://storage.googleapis.com/rsas-assets)
 */

// Helper untuk mendapatkan path aset lokal yang kompatibel dengan Vite
const getLocalAsset = (name) => {
  return new URL(`../assets/img/${name}`, import.meta.url).href;
};

export const ASSETS = {
  // Hero Section Backgrounds
  hero: {
    home: getLocalAsset('hero-home.jpg'),
    profil: getLocalAsset('hero-profil.jpg'),
    layanan: getLocalAsset('hero-layanan.jpg'),
    pendidikan: getLocalAsset('hero-profil.jpg'),
    zonaIntegritas: getLocalAsset('hero-home.jpg'),
    pengaduan: getLocalAsset('hero-home.jpg'),
    galeri: getLocalAsset('hero-home.jpg'),
    faq: getLocalAsset('hero-home.jpg'),
    dokter: getLocalAsset('hero-layanan.jpg'),
    berita: getLocalAsset('hero-profil.jpg'),
  },

  // Layanan Khusus & Detail
  layanan: {
    jantung: getLocalAsset('hero-home.jpg'),
    onkologi: getLocalAsset('hero-home.jpg'),
    stroke: getLocalAsset('hero-home.jpg'),
  },

  // Dokumentasi & Galeri
  galeri: {
    gedung: getLocalAsset('hero-home.jpg'),
    operasi: getLocalAsset('hero-home.jpg'),
    lab: getLocalAsset('hero-home.jpg'),
  }
};
