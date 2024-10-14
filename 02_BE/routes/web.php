<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return 123;
});

Route::get('/users', function () {
    // Lấy tất cả dữ liệu từ bảng users
    $users = DB::table('users')->get();

    // Trả về dữ liệu dưới dạng JSON
    return response()->json($users);
});
