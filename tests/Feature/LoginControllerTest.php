<?php

namespace Tests\Feature;

use Tests\TestCase;
use Illuminate\Foundation\Testing\WithFaker;
use Illuminate\Foundation\Testing\RefreshDatabase;

class LoginControllerTest extends TestCase
{
    use RefreshDatabase;
    /**
     * A basic test example.
     *
     * @return void
     * @test
     */
    public function it_should_get_token_sucessfully()
    {
        $response = $this->get('/api/v1/login');
        $response->assertHeader('Authorization');
    }
}
