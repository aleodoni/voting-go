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

  const res = http.get(`${BASE_URL}/api/v1/me`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  check(res, {
    "status is 200": (r) => r.status === 200,
    "has username": (r) => {
      const username = JSON.parse(r.body).Username;

      return username === "usuario.vereador" || username === "usuario.admin";
    },
  });
}