import { env } from '$env/dynamic/private';
// import pool from '$lib/server/database/pg/init';
import type { RequestHandler } from './$types';
export const POST: RequestHandler = async ({ request }: { request: Request }) => {
        //const confidence = Number.parseFloat(env.AI_MODEL_CONFIDENCE) || 0.7;
        const payload = await request.json();

        let message = payload.message;
        if (payload.duration != "" || payload.severity != "") {
                if (payload.duration != "") {
                        message += `\nAlso consider that the condition has been there for the duration of ${payload.duration}`;
                }
                if (payload.severity != "") {
                        message += `\nAnd the severity on a scale of 1 to 10 being ${payload.severity}`;
                }
        }

        const ollamaRes = await fetch(`${env.OLLAMA_HOST}/api/generate`, {
                method: 'POST',
                headers: {
                        'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                        model: env.AI_MODEL || 'medllama2',
                        prompt: message
                })
        });

        return new Response(ollamaRes.body, {
                headers: {
                        'Content-Type': 'text/event-stream'
                }
        });
};