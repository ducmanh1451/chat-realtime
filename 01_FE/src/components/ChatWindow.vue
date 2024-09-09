<template>
    <div id="chat-window" class="h-full p-3 flex flex-col">
        <div class="header min-h-[60px] flex items-center border-b border-slate-200 pb-1">
            <div class="user-info flex items-center rounded-lg p-2 cursor-pointer hover:bg-slate-200">
                <div class="avatar w-[36px] h-[36px] rounded-full bg-cover bg-center"
                    style="background-image: url('https://randomuser.me/api/portraits/men/1.jpg');">
                </div>
                <div class="flex flex-col pl-2">
                    <span class="text-lg">Bảo Nguyên</span>
                    <span class="text-xs">Hoạt động 3 phút trước</span>
                </div>
            </div>
        </div>
        <!-- <div class="content h-full max-h-[calc(100%-60px-54px)] overflow-auto flex flex-col">
            Content
        </div> -->

        <div class="content h-full max-h-[calc(100%-60px-54px)] overflow-auto flex flex-col space-y-1 p-4">
            <!-- Tin nhắn của người khác (bên trái) -->
            <div class="message-container flex items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Chờ tí
                </div>
            </div>
            <div class="message-container flex items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Tớ vào fb trên messenger
                </div>
            </div>
            <div class="message-container flex items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Lên lao
                </div>
            </div>

            <!-- Tin nhắn của người dùng (bên phải) -->
            <div class="message-container flex justify-end items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Okee
                </div>
            </div>
            <div class="message-container flex justify-end items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Vô nha
                </div>
            </div>

            <!-- Tin nhắn của người khác (bên trái) -->
            <div class="message-container flex items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Được rồi, đang xử lý nè
                </div>
            </div>
            <div class="message-container flex items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Đợi chút nha!
                </div>
            </div>

            <!-- Tin nhắn của người dùng (bên phải) -->
            <div class="message-container flex justify-end items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Ok, không vội đâu
                </div>
            </div>
            <div class="message-container flex justify-end items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Cứ từ từ làm nha
                </div>
            </div>

            <!-- Tin nhắn của người khác (bên trái) -->
            <div class="message-container flex items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Cảm ơn nhiều nha
                </div>
            </div>
            <div class="message-container flex items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Mới mua cái tai nghe này, nghe ổn phết.
                </div>
            </div>

            <!-- Tin nhắn của người dùng (bên phải) -->
            <div class="message-container flex justify-end items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Haha, thế để hôm nào thử
                </div>
            </div>
            <div class="message-container flex justify-end items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Nghe gì hay không?
                </div>
            </div>

            <!-- Tin nhắn của người khác (bên trái) -->
            <div class="message-container flex items-start">
                <div class="message bg-indigo-500 text-white rounded-xl p-2 max-w-xs">
                    Nhạc EDM với chút acoustic thôi, thích lắm.
                </div>
            </div>
        </div>

        <div class="footer w-full flex py-3 justify-end">
            <div class="w-[5%] flex">
                <div
                    class="w-[30px] h-[30px] mx-2 flex justify-center items-center rounded-full cursor-pointer text-blue-500 hover:bg-slate-200">
                    <font-awesome-icon :icon="['fas', 'image']" />
                </div>
            </div>
            <div class="w-[95%] flex">
                <input type="text" class="w-full h-full px-3 py-1 rounded-xl focus:outline-none bg-gray-200"
                    placeholder="Aa">
                <div
                    class="w-[30px] h-[30px] mx-2 flex justify-center items-center rounded-full cursor-pointer text-blue-500 hover:bg-slate-200">
                    <font-awesome-icon :icon="['fas', 'paper-plane']" />
                </div>
            </div>
        </div>
    </div>
</template>