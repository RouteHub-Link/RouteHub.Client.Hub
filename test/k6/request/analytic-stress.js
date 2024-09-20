import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  // Key configurations for Stress in this section
  stages: [
    { duration: '3m', target: 1000 },
    { duration: '2m', target: 500 }, 
    { duration: '1m', target: 0 },
  ],
  thresholds: { http_req_duration: ['avg<100', 'p(95)<200'] },
  noConnectionReuse: true,
  userAgent: 'MyK6UserAgentString/1.0',

  links: [
    'http://127.0.0.1:7331/test-Z56S',
    'http://127.0.0.1:7331/test-fTwy',
    'http://127.0.0.1:7331/test-Hs4X',
  ]
};

export default () => {
    const randomIndex = Math.floor(Math.random() * options.links.length);
    const url = options.links[randomIndex];
    const res = http.get(url);
    console.log(`Response status for ${url}: ${res.status}`);
    sleep(0.10); // sleep for 0.10 second between requests
  };