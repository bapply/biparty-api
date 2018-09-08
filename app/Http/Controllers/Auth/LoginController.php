<?php

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use App\User;
use Tymon\JWTAuth\Facades\JWTAuth;

class LoginController extends Controller
{
    public function index()
    {
        $user = User::create();
        $token = JWTAuth::fromUser($user);

        return response()->json([
            'data' => [
                'token' => $token,
            ],
        ])->header('Authorization', 'Bearer ' . $token);
    }
}
