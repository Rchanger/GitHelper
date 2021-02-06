// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import axiosApi from "axios";

Vue.use(axiosApi);

Vue.config.productionTip = false;

const axios = axiosApi.create({});

//Use the window object to make it available globally.
window.axios = axios;

axios.interceptors.response.use(
  response => {
    return response;
  },
  function(error) {
    if (error.response.status === 401) {
      alert(error.response.data);
      router.push("/");
    }
    return Promise.reject(error.response);
  }
);

/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  components: { App },
  template: "<App/>"
});
