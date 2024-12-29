<template>
    <div class="m-0">
        <label :for="id" class="block text-sm font-medium text-gray-700">
            <slot></slot> <!-- Để người dùng có thể truyền label tùy chỉnh -->
        </label>
        <select :id="id" v-model="selectedValue" @change="updateValue"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            required>
            <option value="" disabled selected>{{ placeholder }}</option>
            <option v-for="option in options" :key="option.value" :value="option.value">
                {{ option.text }}
            </option>
        </select>
        <span class="text-red-500 text-xs block min-h-[16px]"></span>
    </div>
</template>

<script setup>
import { ref, defineProps, defineEmits, watch } from 'vue'

// Định nghĩa props
const props = defineProps({
    id: {
        type: String
    },
    modelValue: {
        type: String,
        default: ''
    },
    placeholder: {
        type: String,
        default: 'Please select'
    },
    options: {
        type: Array,
        required: true, // Đảm bảo rằng options luôn được cung cấp
        validator: (value) => {
            // Kiểm tra từng phần tử của mảng options có đúng định dạng không
            return value.every(option => option.hasOwnProperty('value') && option.hasOwnProperty('text'))
        }
    }
});

// Định nghĩa emit
const emit = defineEmits(['update:modelValue'])

// Khai báo biến reactive cho giá trị đã chọn
const selectedValue = ref(props.modelValue)

// Hàm cập nhật giá trị đã chọn
const updateValue = () => {
    emit('update:modelValue', selectedValue.value)
};

// Theo dõi sự thay đổi của modelValue và cập nhật selectedValue
watch(() => props.modelValue, (newValue) => {
    selectedValue.value = newValue
});
</script>
