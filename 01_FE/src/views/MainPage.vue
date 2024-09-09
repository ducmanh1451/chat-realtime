<template>
    <div id="main-page" class="w-screen h-screen bg-slate-200 flex p-4">
        <div class="w-[25%] pr-2">
            <div class="bg-slate-50 h-full rounded-lg">
                <contact-list></contact-list>
            </div>
        </div>
        <div class="w-[50%] px-2">
            <div class="bg-slate-50 h-full rounded-lg">
                <chat-window></chat-window>
            </div>
        </div>
        <div class="w-[25%] pl-2">
            <div class="bg-slate-50 h-full rounded-lg">
                <sidebar></sidebar>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import ContactList from "@/components/ContactList.vue"
import ChatWindow from "@/components/ChatWindow.vue"
import Sidebar from "@/components/Sidebar.vue"
</script>