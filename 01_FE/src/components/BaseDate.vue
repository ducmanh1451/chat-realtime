<template>
    <div class="m-0">
        <label :for="id" class="block text-sm font-medium text-gray-700">
            <slot></slot> <!-- Để người dùng có thể truyền label tùy chỉnh -->
        </label>
        <input :id="id" type="date" v-model="inputValue"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm" />
        <span class="text-red-500 text-xs block min-h-[16px]"></span>
    </div>
</template>

<script setup>
import { ref, watch } from 'vue';

// Định nghĩa các props cho component
const props = defineProps({
    id: {
        type: String
    },
    modelValue: {
        type: String,
        required: true,
    }
});

// Sử dụng `emit` để gửi sự kiện ra ngoài
const emit = defineEmits(['update:modelValue'])

// Khai báo các biến reactive
const inputValue = ref(props.modelValue)

// Theo dõi sự thay đổi của `inputValue` và emit sự kiện `update:modelValue`
watch(inputValue, (newValue) => {
    emit('update:modelValue', newValue)
});
</script>