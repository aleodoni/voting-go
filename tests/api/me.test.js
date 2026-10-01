import { check } from 'k6';
import http from 'k6/http';

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080';
// Usuário dono do token (o mesmo TEST_USER usado por `make token`).
const EXPECTED_USER = __ENV.TEST_USER || 'usuario.admin';

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

	const res = http.get(`${BASE_URL}/api/v1/me`, {
		headers: {
			Authorization: `Bearer ${token}`,
		},
	});

	const bodyParsed = JSON.parse(res.body);

	check(res, {
		'status is 200': (r) => r.status === 200,
		'has expected username': (_r) => bodyParsed.username === EXPECTED_USER,
	});
}
