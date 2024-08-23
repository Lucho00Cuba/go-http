import http from 'k6/http';
import { check, group } from 'k6';

const SERVER_HOST = "http://localhost:3000"
const SERVER_TYPE = "container"
const SERVER_CODE = 201

export const options = {
    noConnectionReuse: true,
    userAgent: "K6 Client/1.0",
    insecureSkipTLSVerify: true,
    tlsAuth: [
    ],
    scenarios: {
        jwtScenario: {
           exec: 'generalScenario',
            executor: 'shared-iterations',           
            vus: 100,           
            iterations: 1000,
            maxDuration: '1m'
        }
    }
};

export function generalScenario() {
    /**
     * go-server request
     */
    group('go-server request', () => {
        const params = {
            headers: {
                'Content-Type': 'application/json'
            },
        };
        // timestamp now
        const start = Date.now();
        const now = new Date().getTime();
        const res = http.get(`${SERVER_HOST}/${SERVER_CODE}/${SERVER_TYPE}-${now}`, params);
        console.log(`[${SERVER_TYPE}-${now}] ` + 'Response time was ' + String(res.timings.duration) + ' ms');
        const end = Date.now();
        console.log(`[${SERVER_TYPE}-${now}] Duration: ${end - start} ms`)

        // Verify response
        check(res, {
            'status is XXX': (r) => r.status === SERVER_CODE
        });
    })
}

