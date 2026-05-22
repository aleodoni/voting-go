import http from 'k6/http';

const BASE_URL = 'http://localhost:8080';

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

	console.log(`Using token: ${token}`);

	const res = http.get(`${BASE_URL}/api/v1/sincronia`, {
		headers: {
			Authorization: `Bearer ${token}`,
		},
	});

	console.log(res.status);
	console.log(res.body);
}
