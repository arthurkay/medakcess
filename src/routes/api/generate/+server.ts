import { env } from '$env/dynamic/private';
// import pool from '$lib/server/database/pg/init';
import type { RequestHandler } from './$types';
export const POST: RequestHandler = async ({ request }: { request: Request }) => {
        //const confidence = Number.parseFloat(env.AI_MODEL_CONFIDENCE) || 0.7;
        const payload = await request.json();
        console.log(payload)
        console.log(env.OLLAMA_HOST)

        const data = payload
        const ollamaRes = await fetch(`${env.OLLAMA_HOST}/api/generate`, {
                method: 'POST',
                headers: {
                        'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                        model: 'medllama2',
                        prompt: data.message
                })
        });

        return new Response(ollamaRes.body, {
                headers: {
                        'Content-Type': 'text/event-stream'
                }
        });
};