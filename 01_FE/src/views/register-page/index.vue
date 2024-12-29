<template>
    <div id="register-page" class="w-screen h-screen bg-slate-200 flex justify-center items-center">
        <div class="w-[30%] bg-slate-50 p-6 rounded-lg shadow-lg">
            <h2 class="text-2xl font-semibold mb-4">Đăng ký</h2>
            <div class="flex flex-col">
                <base-input v-model="data.email" type="email" id="email" placeholder="Nhập email" required>
                    Email
                </base-input>
                <base-input v-model="data.password" type="password" id="password" placeholder="Nhập mật khẩu" required>
                    Mật khẩu
                </base-input>
                <base-input v-model="data.fullname" type="text" id="fullname" placeholder="Nhập họ và tên" required>
                    Họ và tên
                </base-input>
                <base-select v-model="data.gender" id="gender" placeholder="Chọn giới tính" :options="genderOptions">
                    Giới tính
                </base-select>
                <base-date v-model="data.date_of_birth" id="date_of_birth">
                    Ngày sinh
                </base-date>

                <button type="button" @click="clickRegister"
                    class="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-500">
                    Đăng ký
                </button>
            </div>
            <p class="mt-4 text-center text-sm text-gray-600">
                Đã có tài khoản? <router-link to="/login" class="text-blue-500 hover:underline">Bấm vào đây để đăng
                    nhập</router-link>
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
import messages from '@/helpers/message'
import useRegister from './store'
import BaseInput from '@/components/BaseInput.vue'
import BaseSelect from '@/components/BaseSelect.vue'
import BaseDate from '@/components/BaseDate.vue'

const { data, genderOptions, register } = useRegister()

const clickRegister = async () => {
    const response = await register() // Gọi hàm register
    if (!response.STATUS) {
        alert(messages['E500'])
    } else {
        alert(messages['S200'])
    }
}
</script>