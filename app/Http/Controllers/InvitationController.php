<?php

namespace App\Http\Controllers;

use App\Http\Requests\CreateInvitationRequest;
use App\Invitation;
use Tymon\JWTAuth\Facades\JWTAuth;

class InvitationController extends Controller
{
    public function store(CreateInvitationRequest $request)
    {
        /** @var \App\User $user */
        $user = JWTAuth::parseToken()->authenticate();
        if (! $user) {
            return response()->json([
                'error' => [
                    'message' => "couldn't find the user",
                ],
            ]);
        }
        $path = $request->file->store('invitations');
        $invitation = $user->invitations()->create([
            'title'       => $request->title,
            'description' => $request->description,
            'file'        => env("APP_URL") . '/storage/' . $path,
        ]);

        return response()->json([
            'data' => [
                'invitation' => $invitation,
            ],
        ], 201);
    }
}
