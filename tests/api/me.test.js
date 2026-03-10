import http from "k6/http";
import { check, sleep } from "k6";

const BASE_URL = "http://localhost:8080";

export const options = {
  scenarios: {
    smoke_test: {
      executor: "shared-iterations",
      vus: 1,
      iterations: 10,
    }  
  }
};

export default function () {
  const token = __ENV.TOKEN;

  const res = http.get(`${BASE_URL}/api/v1/me`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  check(res, {
    "status is 200": (r) => r.status === 200,
    "has username": (r) => JSON.parse(r.body).Username == "usuario.vereador",
  });
}