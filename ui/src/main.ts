import Vue from 'vue';
import Buefy from 'buefy';
import 'buefy/dist/buefy.css';
import dayjs from 'dayjs';
import advancedFormat from 'dayjs/plugin/advancedFormat';
import { capitalize } from 'lodash-es';
import App from './App.vue';
import router from './router';
import store from './store';
import SvgIcon from './assets/SvgIcon.vue';

dayjs.extend(advancedFormat);
Vue.config.productionTip = false;
Vue.use(Buefy, {
  defaultIconPack: 'svg',
  defaultIconComponent: SvgIcon,
  customIconPacks: {
    svg: {
      sizes: {
        default: '1x',
        'is-small': '1x',
        'is-medium': '2x',
        'is-large': '3x',
      },
    },
  },
});
Vue.filter('time', (d: string) => dayjs(d).format('Do MMM YYYY'));
Vue.filter('caps', (s: string) => capitalize(s));
Vue.filter('publisherLabel', (id: string) => store.getters.getPublisherLabel(id));

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
