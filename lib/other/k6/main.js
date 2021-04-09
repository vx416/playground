import * as req from  './simple_req.js'


export let options = {
    vus: 100,
    duration: '50s',
    // iterations: 30,
};

// main function
export default function(){
    req.listTrades()

};