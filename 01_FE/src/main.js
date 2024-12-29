import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from '@/App.vue'
import router from '@/router'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import '@/icons/index'
import { useApp } from './stores/app'

const app = createApp(App)
const pinia = createPinia()
const ENVIRONMENT = import.meta.env.VITE_APP_ENV
const DEVELOP_URL = import.meta.env.VITE_APP_API_URL_DEVELOP
const PRODUCTION_URL = import.meta.env.VITE_APP_API_URL_PRODUCTION

app.use(router)
app.use(pinia)
app.component('font-awesome-icon', FontAwesomeIcon)

const appStore = useApp()
appStore.setApiUrl(ENVIRONMENT === 'develop' ? DEVELOP_URL : PRODUCTION_URL)

app.mount('#app')
