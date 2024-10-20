<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateMessagesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('messages', function (Blueprint $table) {
            $table->id();
            $table->foreignId('conversation_id')->constrained('conversations')->onDelete('cascade'); // Cuộc trò chuyện mà tin nhắn thuộc về
            $table->foreignId('sender_id')->constrained('users')->onDelete('cascade'); // Người gửi tin nhắn
            $table->text('message')->nullable(); // Nội dung tin nhắn (có thể null nếu chỉ có tệp)
            $table->string('file_path')->nullable(); // Đường dẫn tệp hoặc ảnh
            // $table->string('file_type')->nullable(); // Loại tệp (vd: image/jpeg, application/pdf)
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('messages');
    }
}
