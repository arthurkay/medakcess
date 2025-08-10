import { env } from '$env/dynamic/private';
import type { RequestHandler } from './$types';
export const POST: RequestHandler = async ({ request }: { request: Request }) => {
       // const confidence = Number.parseFloat(env.AI_MODEL_CONFIDENCE) || 0.7;
        const payload = await request.json();
        // This is the conversation history
        const history = payload;

        // The chat is the last message in the history
        /* const chat = history.pop();

        const ollamaEmbed = await fetch(`${env.OLLAMA_HOST}/api/embed`, {
                method: 'POST',
                headers: {
                        'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                        model: 'nomic-embed-text',
                        input: chat.content
                })
        });

        if (!ollamaEmbed.ok) {
                return new Response(ollamaEmbed.statusText, {
                        status: ollamaEmbed.status
                });
        }

        const ollamaEmbedJson = await ollamaEmbed.json();

        const embedding = JSON.stringify(ollamaEmbedJson.embeddings[0]);
        const res = await pool.query(
                `SELECT
                id,
                chunk,
                1 - (embedding <=> $1) AS similarity
         FROM tutorials_contents_embeddings
         WHERE
                1 - (embedding <=> $1)  >= ${confidence}
         ORDER BY embedding <=> $1
         LIMIT 10`,
                [embedding]
        );

        console.log(res.rows);

        let context: string = '';
        for (let i = 0; i < res.rows.length; i++) {
                context += res.rows[i].chunk + '\n\n';
        }

        let chatHistory = payload; */
        // If a matching entry has been found in the db, add it to the chontext to use when answering the question
        /* if (res.rows.length > 0) {
                chatHistory = history;
                const template = `Answer this question based only on this Context: \n ${context} \nQuestion: ${chat.content}`;
                chatHistory.push({ role: 'user', content: template });
        } else {
                const template = `ONLY answer this question using the following context: \nYou are an open source trainer focusing on Linux, Docker and Kubernetes.
                Only respond as an open source trainer.
                Do not respond to any question that is not related to Linux, Docker or Kubernetes.
                In your responses, provide some commands or code snippets where needed..
                When asked about something that has nothing to do with technology, you must politely decline to answer. \nQuestion: ${chat.content}`;
                chatHistory.push({ role: 'user', content: template });
        } */

        const ollamaRes = await fetch(`${env.OLLAMA_HOST}/api/chat`, {
                method: 'POST',
                headers: {
                        'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                        model: env.AI_MODEL || 'medllama2',
                        stream: true,
                        messages: history
                })
        });

        return new Response(ollamaRes.body, {
                headers: {
                        'Content-Type': 'text/event-stream'
                }
        });
};