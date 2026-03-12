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

export default function () {
  const token = __ENV.TOKEN;

  const payload = JSON.stringify({
    user_id: "gpmy9p0tflapqf35leco52w2",
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