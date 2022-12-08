import http from 'k6/http';

export const options = {
  vus: 10,
  iterations: 100,
  thresholds: {
    'http_reqs{expected_response:true}': ['rate>10'],
  },
};

export default function () {
  http.get(`https://test.k6.io/`);
}
