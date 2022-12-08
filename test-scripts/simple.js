import http from 'k6/http';

export const options = {
  vus: 10,
  duration: '10s',
  thresholds: {
    'http_reqs{expected_response:true}': ['rate>10'],
  },
};

const url = `https://${__ENV.MY_HOSTNAME}/`

export default function () {
  http.get(url);
}
