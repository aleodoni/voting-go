import http from "k6/http";
import { check, sleep } from "k6";

const BASE_URL = "http://localhost:8080";

export const options = {
  scenarios: {
    smoke_test: {
      executor: "shared-iterations",
      vus: 1,
      iterations: 1,
    }  
  }
};

export function setup() {
  const tokenUser = __ENV.TOKEN_USER;

  const res = http.get(`${BASE_URL}/api/v1/me`, {
      headers: {
        Authorization: `Bearer ${tokenUser}`,
      },
    });
  
  const bodyParsed = JSON.parse(res.body);


  return {userId: bodyParsed.id};
}

export default function (data) {
  const token = __ENV.TOKEN;

  console.log('DATA', data)

  const payload = JSON.stringify({
    user_id: data.userId,
    display_name: "Vereador Um",
    is_active: true,
    can_admin: false,
    can_vote: true
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  };

  const res = http.put(
    `${BASE_URL}/api/v1/usuarios/fantasia-credenciais`,
    payload,
    params
  );

  check(res, {
    "status is 204": (r) => r.status === 204,
  });
}