import axios from 'axios'
import { useApp } from '@/stores/app'

const ENDPOINT = 'register'

const route = {
  register: async (data) => {
    try {
      const appStore = useApp()
      const API_URL = `${appStore.apiUrl}/${ENDPOINT}`

      // call axios
      const response = await axios.post(API_URL, data, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
      return response.data
    } catch (error) {
      throw error
    }
  }
}

export default route
