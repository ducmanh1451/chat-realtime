<template>
    <div id="login-page" class="w-screen h-screen bg-slate-200 flex justify-center items-center">
        <div class="w-[30%] bg-slate-50 p-6 rounded-lg shadow-lg">
            <h2 class="text-2xl font-semibold mb-4">Đăng nhập</h2>
            <div class="flex flex-col space-y-4">
                <div>
                    <label for="username" class="block text-sm font-medium text-gray-700">Tên đăng nhập</label>
                    <input type="text" id="username" name="username"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                        placeholder="Nhập tên đăng nhập" />
                </div>
                <div>
                    <label for="password" class="block text-sm font-medium text-gray-700">Mật khẩu</label>
                    <input type="password" id="password" name="password"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                        placeholder="Nhập mật khẩu" />
                </div>
                <button type="button"
                    class="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-500">
                    <router-link to="/">Đăng nhập</router-link>
                </button>
            </div>
            <p class="mt-4 text-center text-sm text-gray-600">
                Bạn chưa có tài khoản? <router-link to="/register" class="text-blue-500 hover:underline">Bấm vào đây để
                    đăng
                    ký</router-link>
            </p>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
</script>