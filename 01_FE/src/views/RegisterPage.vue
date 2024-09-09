<template>
    <div id="register-page" class="w-screen h-screen bg-slate-200 flex justify-center items-center">
        <div class="w-[30%] bg-slate-50 p-6 rounded-lg shadow-lg">
            <h2 class="text-2xl font-semibold mb-4">Đăng ký</h2>
            <form class="flex flex-col space-y-4">
                <div>
                    <label for="username" class="block text-sm font-medium text-gray-700">Tên đăng nhập</label>
                    <input type="text" id="username" name="username"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                        placeholder="Nhập tên đăng nhập" required />
                </div>
                <div>
                    <label for="password" class="block text-sm font-medium text-gray-700">Mật khẩu</label>
                    <input type="password" id="password" name="password"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                        placeholder="Nhập mật khẩu" required />
                </div>
                <div>
                    <label for="full-name" class="block text-sm font-medium text-gray-700">Họ và tên</label>
                    <input type="text" id="full-name" name="full-name"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                        placeholder="Nhập họ và tên" required />
                </div>
                <div>
                    <label for="gender" class="block text-sm font-medium text-gray-700">Giới tính</label>
                    <select id="gender" name="gender"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                        required>
                        <option value="" disabled selected>Chọn giới tính</option>
                        <option value="male">Nam</option>
                        <option value="female">Nữ</option>
                    </select>
                </div>
                <div>
                    <label for="dob" class="block text-sm font-medium text-gray-700">Ngày sinh</label>
                    <input type="date" id="dob" name="dob"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                        required />
                </div>
                <button type="submit"
                    class="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-500">
                    Đăng ký
                </button>
            </form>
            <p class="mt-4 text-center text-sm text-gray-600">
                Đã có tài khoản? <router-link to="/login" class="text-blue-500 hover:underline">Bấm vào đây để đăng
                    nhập</router-link>
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
</script>