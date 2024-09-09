<template>
    <div id="contact-list" class="h-full p-3 flex flex-col">
        <div class="header min-h-[60px] flex items-center">
            <span class="text-3xl font-bold">Chats</span>
        </div>
        <div class="input-search mt-2 relative">
            <div class="absolute left-[10px] top-[5px]">
                <font-awesome-icon :icon="['fas', 'search']" />
            </div>
            <input type="text" class="w-full pl-8 pr-3 py-1 rounded-xl focus:outline-none bg-gray-200"
                placeholder="Tìm kiếm tin nhắn">
        </div>
        <div class="list-users my-3 pr-3 max-h-[730px] overflow-auto">
            <div v-for="(user, index) in users" :key="user.userID" @click="selectItem(index)"
                class="item rounded-lg p-2 flex cursor-pointer hover:bg-slate-200" :class="itemClass(index)">
                <div class="avatar w-[48px] h-[48px] rounded-full bg-cover bg-center" :style="{ backgroundImage: `url(${user.avatar})` }">
                </div>
                <div class="user-info flex flex-col pl-3 w-[calc(100%-48px)] max-w-[calc(100%-48px)]">
                    <span class="block truncate">{{ user.userName }}</span>
                    <div class="w-full flex">
                        <span class="block max-w-[calc(100%-10px)] truncate">{{ user.lastMessage }}</span>
                        <span class="pl-1">{{ user.lastMessageTime }}</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="btn-logout">
            <button type="button"
                class="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-500">
                <router-link to="/login">Đăng xuất</router-link>
            </button>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const selectedIndex = ref(0)
const users = reactive([
    { userID: 'user1', userName: 'John Doe', avatar: 'https://randomuser.me/api/portraits/men/1.jpg', lastMessage: 'Hey there!', lastMessageTime: '3m' },
    { userID: 'user2', userName: 'Jane Smith', avatar: 'https://randomuser.me/api/portraits/women/2.jpg', lastMessage: 'How are you?', lastMessageTime: '15m' },
    { userID: 'user3', userName: 'Alice Johnson', avatar: 'https://randomuser.me/api/portraits/women/3.jpg', lastMessage: 'Let\'s meet up!', lastMessageTime: '26s' },
    { userID: 'user4', userName: 'Bob Brown', avatar: 'https://randomuser.me/api/portraits/men/4.jpg', lastMessage: 'Can we talk?', lastMessageTime: '2h' },
    { userID: 'user5', userName: 'Charlie Davis', avatar: 'https://randomuser.me/api/portraits/men/5.jpg', lastMessage: 'Good morning!', lastMessageTime: '1d' },
    { userID: 'user6', userName: 'Emily Wilson', avatar: 'https://randomuser.me/api/portraits/women/6.jpg', lastMessage: 'Don\'t forget our meeting.', lastMessageTime: '3m' },
    { userID: 'user7', userName: 'Michael Harris', avatar: 'https://randomuser.me/api/portraits/men/7.jpg', lastMessage: 'What\'s up?', lastMessageTime: '15m' },
    { userID: 'user8', userName: 'Sarah Lewis', avatar: 'https://randomuser.me/api/portraits/women/8.jpg', lastMessage: 'See you soon!', lastMessageTime: '26s' },
    { userID: 'user9', userName: 'David Clark', avatar: 'https://randomuser.me/api/portraits/men/9.jpg', lastMessage: 'Hi!', lastMessageTime: '2h' },
    { userID: 'user10', userName: 'Laura Martinez', avatar: 'https://randomuser.me/api/portraits/women/10.jpg', lastMessage: 'Call me when free.', lastMessageTime: '1d' },
    { userID: 'user11', userName: 'James Walker', avatar: 'https://randomuser.me/api/portraits/men/11.jpg', lastMessage: 'Have a great day!', lastMessageTime: '3m' },
    { userID: 'user12', userName: 'Olivia Brown', avatar: 'https://randomuser.me/api/portraits/women/12.jpg', lastMessage: 'Miss you!', lastMessageTime: '15m' },
    { userID: 'user13', userName: 'Daniel Evans', avatar: 'https://randomuser.me/api/portraits/men/13.jpg', lastMessage: 'Let\'s catch up!', lastMessageTime: '26s' },
    { userID: 'user14', userName: 'Sophia Robinson', avatar: 'https://randomuser.me/api/portraits/women/14.jpg', lastMessage: 'Can we reschedule?', lastMessageTime: '2h' },
    { userID: 'user15', userName: 'Matthew Hall', avatar: 'https://randomuser.me/api/portraits/men/15.jpg', lastMessage: 'How was your day?', lastMessageTime: '1d' },
    { userID: 'user16', userName: 'Isabella Young', avatar: 'https://randomuser.me/api/portraits/women/16.jpg', lastMessage: 'Let\'s grab lunch!', lastMessageTime: '3m' },
    { userID: 'user17', userName: 'Ethan King', avatar: 'https://randomuser.me/api/portraits/men/17.jpg', lastMessage: 'See you later.', lastMessageTime: '15m' },
    { userID: 'user18', userName: 'Mia Scott', avatar: 'https://randomuser.me/api/portraits/women/18.jpg', lastMessage: 'Are you free tonight?', lastMessageTime: '26s' },
    { userID: 'user19', userName: 'Alexander Adams', avatar: 'https://randomuser.me/api/portraits/men/19.jpg', lastMessage: 'Talk to you soon!', lastMessageTime: '2h' },
    { userID: 'user20', userName: 'Ava Martinez', avatar: 'https://randomuser.me/api/portraits/women/20.jpg', lastMessage: 'Thanks for your help.', lastMessageTime: '1d' },
    { userID: 'user21', userName: 'Lucas Nelson', avatar: 'https://randomuser.me/api/portraits/men/21.jpg', lastMessage: 'How\'s everything?', lastMessageTime: '3m' },
    { userID: 'user22', userName: 'Megan Carter', avatar: 'https://randomuser.me/api/portraits/women/22.jpg', lastMessage: 'Good to hear from you!', lastMessageTime: '15m' },
    { userID: 'user23', userName: 'Ryan Mitchell', avatar: 'https://randomuser.me/api/portraits/men/23.jpg', lastMessage: 'Check your email.', lastMessageTime: '26s' },
    { userID: 'user24', userName: 'Emily Moore', avatar: 'https://randomuser.me/api/portraits/women/24.jpg', lastMessage: 'See you tomorrow!', lastMessageTime: '2h' },
    { userID: 'user25', userName: 'Liam Turner', avatar: 'https://randomuser.me/api/portraits/men/25.jpg', lastMessage: 'Want to chat?', lastMessageTime: '1d' },
    { userID: 'user26', userName: 'Charlotte Lee', avatar: 'https://randomuser.me/api/portraits/women/26.jpg', lastMessage: 'Let\'s do this!', lastMessageTime: '3m' },
    { userID: 'user27', userName: 'Noah Young', avatar: 'https://randomuser.me/api/portraits/men/27.jpg', lastMessage: 'Talk soon.', lastMessageTime: '15m' },
    { userID: 'user28', userName: 'Amelia Perez', avatar: 'https://randomuser.me/api/portraits/women/28.jpg', lastMessage: 'What time works for you?', lastMessageTime: '26s' },
    { userID: 'user29', userName: 'Jacob Phillips', avatar: 'https://randomuser.me/api/portraits/men/29.jpg', lastMessage: 'Catch you later!', lastMessageTime: '2h' },
    { userID: 'user30', userName: 'Harper Davis', avatar: 'https://randomuser.me/api/portraits/women/30.jpg', lastMessage: 'Let me know!', lastMessageTime: '1d' },
])
const selectItem = (index) => {
    selectedIndex.value = index
}
const itemClass = (index) => {
    let classes = ''
    if (selectedIndex.value == index) {
        classes += 'bg-slate-300'
    }
    return classes
}
</script>