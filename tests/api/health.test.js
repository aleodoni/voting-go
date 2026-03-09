import http from "k6/http";
import { check } from "k6";

const BASE_URL = "http://localhost:8080";

export default function () {

  const token = __ENV.TOKEN;

  const res = http.get(`${BASE_URL}/api/v1/health`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  console.log(`Response body: ${res.body}`);

  check(res, {
    "status is 200": (r) => r.status === 200,
    "status is ok": (r) => JSON.parse(r.body).status === "ok",
  });
}