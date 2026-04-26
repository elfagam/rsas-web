<script setup>
import { ref } from 'vue';

const props = defineProps({
  show: Boolean
});

const emit = defineEmits(['close']);

const form = ref({
  name: '',
  unit: '',
  message: '',
  rating: 5
});

const isSubmitting = ref(false);
const showSuccess = ref(false);

const submitTestimonial = () => {
  isSubmitting.value = true;
  
  // Simulate API call
  setTimeout(() => {
    isSubmitting.value = false;
    showSuccess.value = true;
    
    // Reset form
    setTimeout(() => {
      emit('close');
      showSuccess.value = false;
      form.value = { name: '', unit: '', message: '', rating: 5 };
    }, 2000);
  }, 1500);
};
</script>

<template>
  <Transition name="modal">
    <div v-if="show" class="modal-mask" @click.self="$emit('close')">
      <div class="modal-wrapper">
        <div class="modal-container">
          <button class="modal-close" @click="$emit('close')">
            <i class="fa-solid fa-xmark"></i>
          </button>
          
          <div v-if="!showSuccess" class="modal-content">
            <div class="modal-header u-center-text">
              <i class="fa-solid fa-heart-pulse modal-icon"></i>
              <h2 class="modal-title">Tulis Kesan & Pesan Anda</h2>
              <p class="modal-subtitle">Terima kasih telah mempercayakan RSAS sebagai mitra kesehatan Anda.</p>
            </div>
            
            <form @submit.prevent="submitTestimonial" class="modal-form">
              <div class="form-group">
                <label class="form-label">Nama Lengkap (atau Inisial)</label>
                <input v-model="form.name" type="text" class="form-input" placeholder="Contoh: Siti A." required>
              </div>
              
              <div class="form-group">
                <label class="form-label">Unit Layanan yang Dikunjungi</label>
                <select v-model="form.unit" class="form-input" required>
                  <option value="" disabled>Pilih Unit Layanan</option>
                  <option value="IGD">IGD (Gawat Darurat)</option>
                  <option value="Poli Jantung">Poli Jantung</option>
                  <option value="Poli Interna">Poli Interna</option>
                  <option value="Poli Anak">Poli Anak</option>
                  <option value="Rawat Inap">Rawat Inap</option>
                  <option value="Radiologi">Radiologi / Lab</option>
                </select>
              </div>
              
              <div class="form-group">
                <label class="form-label">Berikan Rating</label>
                <div class="rating-input">
                  <i v-for="i in 5" :key="i" 
                     @click="form.rating = i"
                     class="fa-star" 
                     :class="i <= form.rating ? 'fa-solid active' : 'fa-regular'"></i>
                </div>
              </div>
              
              <div class="form-group">
                <label class="form-label">Kesan & Pengalaman Anda</label>
                <textarea v-model="form.message" class="form-input form-textarea" placeholder="Tuliskan pengalaman Anda melayani kami..." required></textarea>
              </div>
              
              <button type="submit" class="btn-submit" :disabled="isSubmitting">
                <span v-if="!isSubmitting">Kirim Testimoni</span>
                <span v-else><i class="fa-solid fa-spinner fa-spin"></i> Mengirim...</span>
              </button>
            </form>
          </div>
          
          <div v-else class="modal-success u-center-text">
            <div class="success-icon"><i class="fa-solid fa-check"></i></div>
            <h2>Terima Kasih!</h2>
            <p>Testimoni Anda telah terkirim dan akan segera kami tampilkan setelah proses moderasi.</p>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.modal-mask {
  position: fixed;
  z-index: 10000;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(5px);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
  padding: 2rem;
}

.modal-container {
  width: 100%;
  max-width: 55rem;
  margin: 0px auto;
  background-color: #fff;
  border-radius: 25px;
  box-shadow: 0 3rem 6rem rgba(0, 0, 0, 0.2);
  transition: all 0.3s ease;
  padding: 4rem;
  position: relative;
  overflow: hidden;
}

.modal-close {
  position: absolute;
  top: 2rem;
  right: 2.5rem;
  background: none;
  border: none;
  font-size: 2.5rem;
  color: #bbb;
  cursor: pointer;
  transition: color 0.3s;
}

.modal-close:hover { color: #55c57a; }

.modal-icon {
  font-size: 5rem;
  color: #55c57a;
  margin-bottom: 2rem;
}

.modal-title {
  font-size: 2.5rem;
  font-weight: 800;
  color: #333;
  margin-bottom: 1rem;
}

.modal-subtitle {
  font-size: 1.5rem;
  color: #777;
  margin-bottom: 3.5rem;
  line-height: 1.6;
}

.modal-form {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.8rem;
}

.form-label {
  font-size: 1.4rem;
  font-weight: 700;
  color: #555;
  margin-left: 0.5rem;
}

.form-input {
  padding: 1.5rem 2rem;
  border: 2px solid #f0f0f0;
  border-radius: 12px;
  font-size: 1.5rem;
  font-family: inherit;
  transition: all 0.3s;
}

.form-input:focus {
  outline: none;
  border-color: #55c57a;
  background-color: #f9fbf9;
}

.form-textarea {
  height: 12rem;
  resize: none;
}

.rating-input {
  display: flex;
  gap: 1rem;
  font-size: 2.5rem;
  color: #ddd;
}

.rating-input i { cursor: pointer; transition: transform 0.2s; }
.rating-input i:hover { transform: scale(1.2); }
.rating-input i.active { color: #f1c40f; }

.btn-submit {
  background: #55c57a;
  color: #fff;
  border: none;
  padding: 1.8rem;
  border-radius: 12px;
  font-size: 1.7rem;
  font-weight: 800;
  cursor: pointer;
  margin-top: 1rem;
  transition: all 0.3s;
}

.btn-submit:hover:not(:disabled) {
  background: #28b485;
  box-shadow: 0 1rem 2rem rgba(85, 197, 122, 0.3);
}

.btn-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* Success State */
.modal-success {
  padding: 4rem 0;
}

.success-icon {
  width: 8rem;
  height: 8rem;
  background: #55c57a;
  color: #fff;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 4rem;
  margin: 0 auto 3rem;
  animation: scaleUp 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

@keyframes scaleUp {
  from { transform: scale(0); }
  to { transform: scale(1); }
}

/* Modal Transition */
.modal-enter-from { opacity: 0; }
.modal-leave-to { opacity: 0; }

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
  transform: scale(1.1);
}

@media (max-width: 600px) {
  .modal-container { padding: 3rem 2rem; }
}
</style>
