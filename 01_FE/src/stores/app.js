import { defineStore } from 'pinia'

export const useApp = defineStore('app', {
  state: () => ({
    apiUrl: ''
  }),
  actions: {
    setApiUrl(url) {
      this.apiUrl = url
    }
  }
})
