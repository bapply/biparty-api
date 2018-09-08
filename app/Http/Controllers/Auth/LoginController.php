<?php

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use App\User;
use http\Env\Request;
use Illuminate\Foundation\Auth\AuthenticatesUsers;
use Tymon\JWTAuth\JWTAuth;

class LoginController extends Controller
{
    public function index(Request $request)
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
