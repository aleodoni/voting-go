import { check } from 'k6';
import http from 'k6/http';

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080';

export const options = {
	scenarios: {
		smoke_test: {
			executor: 'shared-iterations',
			vus: 1,
			iterations: 1,
		},
	},
	// Sem threshold o k6 sai com sucesso mesmo quando um check falha.
	thresholds: { checks: ['rate==1'] },
};

export default function () {
	const token = __ENV.TOKEN;

	const res = http.get(`${BASE_URL}/api/v1/sincronia`, {
		headers: {
			Authorization: `Bearer ${token}`,
		},
	});

	const ok = check(res, {
		'status is 200': (r) => r.status === 200,
	});

	// Esta rota exige admin: com um TEST_USER que não seja admin a resposta é 403.
	if (!ok) {
		console.error(`GET /sincronia retornou ${res.status}: ${res.body}`);
	}
}
