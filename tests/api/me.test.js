import { check } from 'k6';
import http from 'k6/http';

// const BASE_URL = 'http://localhost:8080';
const BASE_URL = 'http://192.168.1.61:8080';

export const options = {
	scenarios: {
		smoke_test: {
			executor: 'shared-iterations',
			vus: 1,
			iterations: 1,
		},
	},
};

export default function () {
	const token = __ENV.TOKEN;

	const res = http.get(`${BASE_URL}/api/v1/me`, {
		headers: {
			Authorization: `Bearer ${token}`,
		},
	});

	console.log(res.body);

	const bodyParsed = JSON.parse(res.body);

	check(res, {
		'status is 200': (r) => r.status === 200,
		'has username': (_r) => {
			const username = bodyParsed.username;

			return username === 'usuario.vereador' || username === 'usuario.admin';
		},
	});
}
