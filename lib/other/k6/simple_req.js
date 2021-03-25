import http from 'k6/http';
import { sleep, check } from 'k6';




export function get() {
    var params = {
        headers: {
          'Content-Type': 'application/json',
        },
    };
      
    let res = http.get('http://test.k6.io', null, params);
    check(res, {
        'has status 200': (r) => r.status === 200,
    });


    sleep(1);  
};