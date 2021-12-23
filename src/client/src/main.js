import Vue from 'vue';
import App from './App.vue';
import store from './store/index';
import axios from 'axios';

Vue.config.productionTip = false;
Vue.prototype.$http = axios;

new Vue({
  render: (h) => h(App),
  store: store,
}).$mount('#app');
