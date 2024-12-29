import { reactive } from 'vue'
import route from './route'

export default function useRegister() {
  const data = reactive({
    email: '',
    password: '',
    fullname: '',
    gender: '',
    date_of_birth: ''
  })

  const genderOptions = [
    { value: '0', text: 'Nữ' },
    { value: '1', text: 'Nam' }
  ]

  const register = async () => {
    try {
      const response = await route.register(data)
      return response
    } catch (error) {
      return error
    }
  }

  return {
    data,
    genderOptions,
    register
  }
}
