<script setup>
const currentYear = new Date().getFullYear();
</script>

<template>
  <footer class="footer">
    <div class="footer__grid">
      <!-- Hospital Identity -->
      <div class="footer__col footer__col--brand">
        <div class="footer__logo-box">
          <div class="footer__medical-icon"><i class="fa-solid fa-square-h"></i></div>
          <h3 class="footer__hospital-name">
            RSUD PROF. DR. H. ALOEI SABOE<br />
            <span class="footer__hospital-city">KOTA GORONTALO</span>
          </h3>
        </div>
        <p class="footer__tagline">"Menjadi Lebih Baik Untuk Semua"</p>
        <p class="footer__address">
          <i class="fa-solid fa-location-dot"></i>
          Jl. Prof. Dr. H. Aloei Saboe No. 92, Kel. Wongkaditi Timur, Kota Gorontalo - 96122
        </p>
      </div>

      <!-- Quick Links -->
      <div class="footer__col">
        <h4 class="footer__title">Tautan Cepat</h4>
        <ul class="footer__list">
          <li><router-link to="/profil" class="footer__link">Profil Rumah Sakit</router-link></li>
          <li><router-link to="/layanan" class="footer__link">Layanan Unggulan</router-link></li>
          <li><router-link to="/dokter" class="footer__link">Jadwal Dokter</router-link></li>
          <li><router-link to="/berita" class="footer__link">Berita & Edukasi</router-link></li>
          <li><router-link to="/zona-integritas" class="footer__link">Zona Integritas</router-link></li>
          <li><router-link to="/faq" class="footer__link">Pusat Bantuan (FAQ)</router-link></li>
        </ul>
      </div>

      <!-- Contact Info -->
      <div class="footer__col">
        <h4 class="footer__title">Kontak Kami</h4>
        <ul class="footer__list">
          <li class="footer__contact-item">
            <i class="fa-solid fa-phone"></i>
            <span>(0435) 821010</span>
          </li>
          <li class="footer__contact-item">
            <i class="fa-brands fa-whatsapp"></i>
            <a href="https://wa.me/628114333446" target="_blank">0811-4333-446</a>
          </li>
          <li class="footer__contact-item">
            <i class="fa-solid fa-envelope"></i>
            <a href="mailto:rsudaloeisaboe@gorontalokota.go.id">Email RSAS</a>
          </li>
        </ul>
      </div>

      <!-- Social & Newsletter -->
      <div class="footer__col">
        <h4 class="footer__title">Media Sosial</h4>
        <div class="footer__socials">
          <a href="#" class="footer__social-link"><i class="fa-brands fa-instagram"></i></a>
          <a href="#" class="footer__social-link"><i class="fa-brands fa-facebook-f"></i></a>
          <a href="#" class="footer__social-link"><i class="fa-brands fa-youtube"></i></a>
          <a href="#" class="footer__social-link"><i class="fa-brands fa-tiktok"></i></a>
        </div>
        <div class="footer__accreditation">
          <span class="accreditation-tag">Terakreditasi KARS Paripurna</span>
        </div>
      </div>
    </div>

    <div class="footer__bottom">
      <div class="footer__container">
        <p class="footer__copyright">&copy; {{ currentYear }} RSUD Prof. Dr. H. Aloei Saboe. All Rights Reserved.</p>
        <div class="footer__bottom-links">
          <a href="#">Privacy Policy</a>
          <span class="divider">|</span>
          <a href="#">Terms of Use</a>
        </div>
      </div>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  background-color: #101413; /* Deep Dark Medical Theme */
  padding: 8.4rem 2.8rem 0; /* 6rem * 1.4, 2rem * 1.4 */
  color: #fff;
  position: relative;
}

.footer::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background: linear-gradient(to right, #55c57a, #28b485);
}

.footer__grid {
  max-width: 168rem; /* 120rem * 1.4 */
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1.5fr 1fr 1fr 1fr;
  gap: 5.6rem; /* 4rem * 1.4 */
  padding-bottom: 7rem; /* 5rem * 1.4 */
}

.footer__col--brand {
  padding-right: 2.8rem;
}

.footer__medical-icon {
  font-size: 3.9rem; /* 2.8rem * 1.4 */
  color: #55c57a;
  margin-bottom: 1.4rem;
}

.footer__hospital-name {
  font-size: 1.8rem; /* 1.3rem * 1.4 */
  font-weight: 800;
  line-height: 1.2;
  margin-bottom: 1.4rem;
  letter-spacing: 0.5px;
}

.footer__hospital-city {
  color: #55c57a;
  font-size: 1.4rem;
}

.footer__tagline {
  font-size: 1.5rem; /* 1.1rem * 1.4 */
  font-style: italic;
  color: #888;
  margin-bottom: 2.8rem;
}

.footer__address {
  font-size: 1.4rem; /* 1rem * 1.4 */
  color: #bbb;
  line-height: 1.6;
  display: flex;
  gap: 1.4rem;
}

.footer__address i { color: #55c57a; margin-top: 0.4rem; }

.footer__title {
  font-size: 1.7rem; /* 1.2rem * 1.4 */
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 2px;
  margin-bottom: 3.5rem;
  color: #55c57a;
}

.footer__list {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 1.7rem;
}

.footer__link {
  text-decoration: none;
  color: #bbb;
  font-size: 1.4rem;
  transition: all 0.3s;
}

.footer__link:hover {
  color: #fff;
  padding-left: 7px;
}

.footer__contact-item {
  display: flex;
  align-items: center;
  gap: 1.7rem;
  font-size: 1.4rem;
  color: #bbb;
  margin-bottom: 1.7rem;
}

.footer__contact-item i { color: #55c57a; width: 2.1rem; }
.footer__contact-item a { color: inherit; text-decoration: none; }
.footer__contact-item a:hover { color: #fff; }

.footer__socials {
  display: flex;
  gap: 1.4rem;
  margin-bottom: 3.5rem;
}

.footer__social-link {
  width: 4.9rem; /* 3.5rem * 1.4 */
  height: 4.9rem;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 2rem;
  transition: all 0.3s;
  text-decoration: none;
}

.footer__social-link:hover {
  background: #55c57a;
  transform: translateY(-3px);
}

.accreditation-tag {
  display: inline-block;
  padding: 0.6rem 1.2rem;
  background: rgba(85, 197, 122, 0.1);
  border: 1px solid rgba(85, 197, 122, 0.2);
  color: #55c57a;
  font-size: 0.9rem;
  font-weight: 700;
  border-radius: 5px;
}

.footer__bottom {
  border-top: 1px solid rgba(255, 255, 255, 0.05);
  padding: 2.5rem 0;
  background: rgba(0, 0, 0, 0.2);
}

.footer__container {
  max-width: 120rem;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 2rem;
}

.footer__copyright {
  font-size: 0.9rem;
  color: #666;
}

.footer__bottom-links {
  font-size: 0.9rem;
  color: #666;
  display: flex;
  gap: 1.5rem;
}

.footer__bottom-links a { color: inherit; text-decoration: none; transition: color 0.3s; }
.footer__bottom-links a:hover { color: #888; }

@media (max-width: 1024px) {
  .footer__grid { grid-template-columns: 1fr 1fr; gap: 4rem; }
}

@media (max-width: 600px) {
  .footer__grid { grid-template-columns: 1fr; }
  .footer__container { flex-direction: column; gap: 1.5rem; text-align: center; }
}
</style>
