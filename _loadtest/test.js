import http from 'k6/http';

export const options = {
    discardResponseBodies: true,
    scenarios: {
        contacts: {
            executor: 'per-vu-iterations',
            vus: 100,
            iterations: 500,
            maxDuration: '1h00m',
        },
    },
};

export default function () {
    http.get('http://localhost:12345/v1/features');
}