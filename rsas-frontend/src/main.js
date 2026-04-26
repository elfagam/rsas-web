import { createApp } from "vue";
import "./assets/sass/style.scss";
import "aos/dist/aos.css"; // Import style AOS
import App from "./App.vue";
import router from "./router";

const app = createApp(App);

app.use(router);
app.mount("#app");
