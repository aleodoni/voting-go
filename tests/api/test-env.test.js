import http from "k6/http";
import { check } from "k6";

export default function () {
  console.log("PUBLIC_KEY =", __ENV.PUBLIC_KEY);
  console.log("USER_ACCESS =", __ENV.USER_ACCESS);

  // Só para garantir que o script roda
  const res = http.get("https://httpbin.org/get"); 
  check(res, { "status is 200": (r) => r.status === 200 });
}