<template>
    <div class="m-0">
        <label :for="id" class="block text-sm font-medium text-gray-700">
            <slot></slot> <!-- Để người dùng có thể truyền label tùy chỉnh -->
        </label>
        <input :id="id" :type="type" v-model="inputValue" @focusout="validateInput" :class="inputClass"
            class="mt-1 block w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none sm:text-sm"
            :placeholder="placeholder" :required="required" />
        <span class="text-red-500 text-xs block min-h-[16px]">{{ currentErrorMessage }}</span>
    </div>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import messages from '@/helpers/message'

// Props nhận vào từ component cha
const props = defineProps({
    id: {
        type: String
    },
    type: {
        type: String,
        default: 'text',
        validator: (value) => ['text', 'email', 'password', 'number'].includes(value),
    },
    modelValue: {
        type: String,
        required: true,
    },
    placeholder: {
        type: String,
        default: '',
    },
    required: {
        type: Boolean,
        default: false,
    },
    errorMessage: {
        type: String,
        default: '', // Nếu không truyền vào, mặc định không có lỗi
    },
})

// Sử dụng `emit` để gửi sự kiện ra ngoài
const emit = defineEmits(['update:modelValue'])

// Khai báo các biến reactive
const inputValue = ref(props.modelValue)
const internalErrorMessage = ref(props.errorMessage) // Set lỗi từ props khi khởi tạo
const currentErrorMessage = computed(() => internalErrorMessage.value)

// Tính toán class của input dựa trên trạng thái của lỗi
const inputClass = computed(() => {
    if (currentErrorMessage.value) {
        return 'border-red-500 focus:ring-red-500 focus:border-red-500'
    }
    return 'border-gray-300 focus:ring-blue-500 focus:border-blue-500'
});

// Hàm validate email và password khi mất focus
const validateInput = () => {
    // Khi blur thì luôn luôn clear lỗi dù có từ ngoài hay không
    internalErrorMessage.value = ''

    if (props.type === 'email') {
        if (inputValue.value === '') {
            internalErrorMessage.value = '' // Nếu giá trị trống, clear lỗi
        } else {
            const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
            if (!emailPattern.test(inputValue.value)) {
                internalErrorMessage.value = messages['E001']
            } else {
                internalErrorMessage.value = ''
            }
        }
    }

    if (props.type === 'password') {
        if (inputValue.value === '') {
            internalErrorMessage.value = '' // Nếu giá trị trống, clear lỗi
        } else {
            const passwordPattern = /^[A-Za-z0-9]{8,}$/ // Chỉ cho phép các ký tự chữ và số, độ dài tối thiểu là 8 ký tự
            if (!passwordPattern.test(inputValue.value)) {
                internalErrorMessage.value = messages['E002']
            } else {
                internalErrorMessage.value = ''
            }
        }
    }
};

// Theo dõi sự thay đổi của `inputValue` và emit sự kiện `update:modelValue`
watch(inputValue, (newValue) => {
    emit('update:modelValue', newValue)
});
</script>
