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

  const reunioesRes = http.get(`${BASE_URL}/api/v1/reunioes-dia`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  const reunioes = JSON.parse(reunioesRes.body);

  // supondo que retorna um array
  const reuniaoId = reunioes[0].id;

  console.log(`Reunião escolhida: ${reuniaoId}`);

  check(reunioesRes, {
    "status is 200": (r) => r.status === 200,
  });

  // Buscar projetos da reunião
  const projetosRes = http.get(
    `${BASE_URL}/api/v1/reunioes/${reuniaoId}/projetos`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  const projetos = JSON.parse(projetosRes.body);
  const projetoId = projetos[0].id;
  console.log(`Projeto escolhido: ${projetoId}`);

  // Abrir votação para o projeto
  const abreVotacaoRes = http.post(
  `${BASE_URL}/api/v1/projetos/${projetoId}/votacao/abrir`,
  null, 
  {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }
);

  console.log(`Abre votação status: ${abreVotacaoRes.status}`);

  check(abreVotacaoRes, {
    "abre-votacao status is 204": (r) => r.status === 204,
  });

}