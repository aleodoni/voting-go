import http from "k6/http";
import { check } from "k6";

const BASE_URL = "https://folha.suite.betha.cloud";

export default function () {

  const cpf = __ENV.CPF;
  const publicKey = __ENV.PUBLIC_KEY;
  const userAccess = __ENV.USER_ACCESS;

  const filter = encodeURIComponent(`pessoa.cpf = '${cpf}'`);
  const url = `${BASE_URL}/dados/v2/matriculas?filter=${filter}&limit=1`;

  const res = http.get(url, {
    headers: {
      Authorization: `Bearer ${publicKey}`,
      "user-access": userAccess
    },
  });

  const body = JSON.parse(res.body);

  // Mostra o corpo da resposta para facilitar a depuração, se necessário
  // console.log(JSON.stringify(body, null, 2));

  check(res, {
    "status is 200": (r) => r.status === 200,
    "status is ok": () => body.content[0].pessoa.cpf === cpf,
  });
}