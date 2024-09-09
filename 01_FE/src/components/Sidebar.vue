<template>
    <div id="sidebar" class="h-full p-3 flex flex-col">
        <div v-if="showMainSidebarPage" class="user-info">
            <div class="flex flex-col items-center">
                <div class="avatar w-[100px] h-[100px] rounded-full bg-cover bg-center"
                    style="background-image: url('https://randomuser.me/api/portraits/women/2.jpg');">
                </div>
                <div class="user-name">
                    <span class="text-xl font-bold">Xuan Mai</span>
                </div>
            </div>
            <div class="list-functions select-none">
                <div class="item p-2 flex justify-between items-center rounded-lg cursor-pointer hover:bg-slate-200"
                    @click="openBoxMediaFiles">
                    <span class="block">Ảnh, videos & files</span>
                    <span v-if="!collapseMediaFiles" class="w-[36px] h-[36px] flex justify-center items-center">
                        <font-awesome-icon :icon="['fas', 'chevron-right']" />
                    </span>
                    <span v-if="collapseMediaFiles" class="w-[36px] h-[36px] flex justify-center items-center">
                        <font-awesome-icon :icon="['fas', 'chevron-down']" />
                    </span>
                </div>
                <div v-if="collapseMediaFiles" @click="goToMediaPage"
                    class="item p-2 flex items-center rounded-lg cursor-pointer hover:bg-slate-200">
                    <span class="mr-3 rounded-full bg-slate-300 w-[36px] h-[36px] flex justify-center items-center">
                        <font-awesome-icon :icon="['fas', 'photo-video']" />
                    </span>
                    <span>Ảnh, videos</span>
                </div>
                <div v-if="collapseMediaFiles"
                    class="item p-2 flex items-center rounded-lg cursor-pointer hover:bg-slate-200">
                    <span class="mr-3 rounded-full bg-slate-300 w-[36px] h-[36px] flex justify-center items-center">
                        <font-awesome-icon :icon="['fas', 'file-alt']" />
                    </span>
                    <span>Files</span>
                </div>
            </div>
        </div>
        <div v-if="!showMainSidebarPage && showMediaPage" class="media-page max-h-full overflow-auto">
            <div class="header flex items-center pb-3">
                <span @click="goToMainSidebarPage"
                    class="w-[36px] h-[36px] flex justify-center items-center rounded-full cursor-pointer hover:bg-slate-300">
                    <font-awesome-icon :icon="['fas', 'arrow-left']" />
                </span>
                <span class="pl-3">Ảnh, videos</span>
            </div>
            <div class="content grid gap-4 grid-cols-3">
                <div class="item w-full aspect-square bg-center bg-cover"
                    style="background-image: url('https://randomuser.me/api/portraits/men/1.jpg');"></div>
                <div class="item w-full aspect-square bg-center bg-cover"
                    style="background-image: url('https://randomuser.me/api/portraits/men/1.jpg');"></div>
                <div class="item w-full aspect-square bg-center bg-cover"
                    style="background-image: url('https://randomuser.me/api/portraits/men/1.jpg');"></div>
                <div class="item w-full aspect-square bg-center bg-cover"
                    style="background-image: url('https://randomuser.me/api/portraits/men/1.jpg');"></div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const collapseMediaFiles = ref(false)
const showMainSidebarPage = ref(true)
const showMediaPage = ref(false)
const openBoxMediaFiles = () => {
    collapseMediaFiles.value = !collapseMediaFiles.value
}
const goToMediaPage = () => {
    showMainSidebarPage.value = false
    showMediaPage.value = true
}
const goToMainSidebarPage = () => {
    showMainSidebarPage.value = true
    showMediaPage.value = false
}
</script>