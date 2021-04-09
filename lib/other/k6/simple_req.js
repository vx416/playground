import http from 'k6/http';
import { sleep, check } from 'k6';

import crypto from 'k6/crypto';
import encoding from 'k6/encoding';
import { Counter } from "k6/metrics";

let counter1 = new Counter('resp_time dist_1 t < 800ms');
let counter2 = new Counter('resp_time dist_2 800ms < t < 1200ms');
let counter3 = new Counter('resp_time dist_3 t > 1200ms');

export function listTrades() {
    let userIndex = __VU % users.length 
    let user = users[userIndex]
    let apiSecret= user.apiSecret
    let apikey = user.apikey
    let payload = getPayload({identity: user.email})
    let sign = getSign(payload, apiSecret)
    

    let params = {
        headers: {
          'X-BITOPRO-PAYLOAD': payload,
          'X-BITOPRO-SIGNATURE': sign,
          'X-BITOPRO-APIKEY': apikey,
          'Content-Encoding': 'gzip, deflate',
          'Accept': 'application/json'
        },
    };
      
    let res = http.get('https://hotfix-api.bitopro.com/v3/orders/trades/yfi_twd?startTimesatp=1602094434963&endTimestamp=1604321149937&limit=500'
                        , params);
    
    if (res.timings.duration < 800) {
        counter1.add(1)
    } else if (res.timings.duration < 1200) {
        counter2.add(1)
    } else if (res.timings.duration > 1200) {
        counter3.add(1)
    }

    check(res, {
        'has status 200': (r) => r.status === 200,
    });
    check(res, {
        'has status data': (r) => r.json().data != null,
    })

    if (res.status != 200) {
        console.log(res.body)
    }

    sleep(1);
};

export function getSign(payload, apiSecret) {
    let signature = crypto.hmac('sha384', apiSecret, payload, 'hex')
    return signature
}

export function getPayload(requestBody) {
    requestBody["nonce"] = Date.now()  
    let data = JSON.stringify(requestBody)
    return encoding.b64encode(data)
}


let users = [
    // {id: 96357, email: "edgars@gravityteam.com", apikey: "3b78a8c6d67217a06a7125a930d4fc3e789m", apiSecret: "$2a$12$SBBE2jsn19V123456789"},
    {id: 1035, email: "owen760831@gmail.com", apikey: "apikey_1035", apiSecret:  "b46a14cb-5761-476f-bb1b-756407b853b2"},
    {id: 1036, email: "reggie.escobar94@gmail.com", apikey: "apikey_1036", apiSecret:  "80a9b957-199e-4ae7-bd09-a8a9081b0ca6"},
    {id: 1037, email: "ceres0813@hotmail.com", apikey: "apikey_1037", apiSecret:"8f028790-8aa0-48c2-8a86-966512d0519f"},
    {id: 1038, email: "jawang192@gmail.com", apikey: "apikey_1038", apiSecret:"9eaa3d71-2687-4a23-9b1e-229097ad55c0"},
    {id: 1039, email: "lees8@ms12.hinet.net",apikey: "apikey_1039", apiSecret:"814ceaa9-96f6-4a63-a4f5-f964bb55546a"},
    {id: 1040, email: "kobe124900@gmail.com", apikey: "apikey_1040", apiSecret:"fcd50df6-d206-4244-8bd0-01bc57566756"},
    {id: 1041, email: "tinachen927@hotmail.com", apikey: "apikey_1041",apiSecret: "09f73e16-9ba4-41d6-9bee-11e5293b4213"},
    {id: 1042, email: "alanng22@hotmail.com", apikey: "apikey_1042", apiSecret:"fab815a8-7879-4ee3-8e93-ecec587cd9ca"},
    {id: 1043, email: "anniechang01@gmail.com", apikey: "apikey_1043", apiSecret:"bd6ad433-58d4-4635-9ce8-8eb31e674971"},
    {id: 1044, email: "hahahatao@gmail.com", apikey: "apikey_1044", apiSecret:"5f802cce-edee-4a04-8045-6230271bed5b"},
    {id: 1045, email: "siyu10112328@gmail.com", apikey: "apikey_1045", apiSecret:"da3504b6-b488-4883-ba93-841e3005f605"},
    {id: 1046, email: "sunrisedm4+bitopro@gmail.com", apikey: "apikey_1046", apiSecret:"2c7fcf4c-7873-4f76-9b07-1fde7b6f9c43"},
    {id: 1047, email: "fishos2@yahoo.com.tw", apikey: "apikey_1047", apiSecret:"463b2e06-cb66-44ff-84fe-acabfc8c0c41"},
    {id: 1048, email: "rncchen@gmail.com", apikey: "apikey_1048", apiSecret:"16e11a60-69a1-4fdc-9f37-ed0768a32e84"},
    {id: 1049, email: "sihsyuan19@gmail.com", apikey: "apikey_1049", apiSecret:"48a58a51-af35-4744-ad20-4d23c9cd0d46"},
    {id: 1050, email: "billpeng77@hotmail.com", apikey: "apikey_1050", apiSecret:"2bb7b65f-8aa1-4142-8cbe-ce80415b6ad8"},
    {id: 1051, email: "mummysurf@gmail.com", apikey: "apikey_1051", apiSecret:"d7c90ad6-016e-437e-b76c-fda691e69288"},
    {id: 1052, email: "open3397@gmail.com", apikey: "apikey_1052", apiSecret:"c40312ee-4704-4878-ad1e-580eb5231dca"},
    {id: 1053, email: "sunboy5117@gmail.com", apikey: "apikey_1053", apiSecret:"655a86e0-4580-4166-b790-7e5c051e43b4"},
    {id: 1054, email: "gn01991495@hotmail.com", apikey: "apikey_1054", apiSecret:"94356fa6-be8f-4d7f-a725-b38d80edfddf"},
    {id: 1055, email: "gn01991495@gmail.com", apikey: "apikey_1055",apiSecret: "185b1d4d-9a7d-419e-9bbe-66a343bfe480"},
    {id: 1056, email: "menghuantsai@gmail.com", apikey: "apikey_1056", apiSecret:"41b7b78a-59f8-4048-8788-9c8962b3b96b"},
    {id: 1057, email: "tony8139@gmail.com", apikey: "apikey_1057", apiSecret:"d4542cbc-e865-4fd2-80f2-050ae854e69e"},
    {id: 1058, email: "g1991221@gmail.com", apikey: "apikey_1058",apiSecret: "04a0fbeb-7bde-438a-935d-c546221b85f9"},
    {id: 1059, email: "dengfong01@gmail.com", apikey: "apikey_1059", apiSecret:"932f6f90-57ad-4fb2-b9fc-a757013ad3b2"},
    {id: 1060, email: "mybaby1601777@gmail.con", apikey: "apikey_1060",apiSecret: "f54b61f2-9ff8-4b20-899a-78c5b3dc59cc"},
    {id: 1061, email: "jimmilk7995@gmail.com", apikey: "apikey_1061",apiSecret: "15d7901f-eca8-4da7-add0-b608d485c673"},
    {id: 1062, email: "giorno1220@gmail.com", apikey: "apikey_1062", apiSecret:"93157516-7b5e-44bc-8623-d7f5afbdcbd4"},
    {id: 1063, email: "stoms9903@gmail.com", apikey: "apikey_1063", apiSecret:"0a92748c-cb9f-4e20-8b34-a9f0a66a91dd"},
    {id: 1064, email: "pp092701244849@gmail.com", apikey: "apikey_1064", apiSecret:"8f1d7117-c0bf-497a-88f7-bc2f14540c4c"},
    {id: 1065, email: "mybaby1601777@gmail.com", apikey: "apikey_1065", apiSecret:"8008fab2-a9df-4f7d-af1c-f985e36648ef"},
    {id: 1066, email: "aq0218aq@gmail.com", apikey: "apikey_1066", apiSecret:"fddd8c5f-57a1-46b4-9629-f24e2bb0808f"},
    {id: 1067, email: "ppilkimo@hotmail.com", apikey: "apikey_1067", apiSecret:"217c4b20-ea94-4508-887e-056ef7735831"},
    {id: 1068, email: "pobby11221@gmail.com", apikey: "apikey_1068", apiSecret:"c89ca11b-1517-4482-906a-2f43684fa3fd"},
    {id: 1069, email: "angelchung131@gmail.com", apikey: "apikey_1069",apiSecret: "4b9610d9-ab7f-4160-962f-e168abd1e219"},
    {id: 1070, email: "a0916648984@gmail.com", apikey: "apikey_1070", apiSecret:"df2b89cc-4c67-4aeb-84ba-96f641c1e142"},
    {id: 1071, email: "garmgoon@gmail.com", apikey: "apikey_1071", apiSecret:"365e7297-c126-4e2a-8b20-77ecfd2e18af"},
    {id: 1072, email: "adsl147456@yahoo.com.tw", apikey: "apikey_1072", apiSecret:"6f96d097-0fb6-4e0f-92c1-d62fed835b4b"},
    {id: 1073, email: "ktt823@gmail.com", apikey: "apikey_1073", apiSecret:"0fbe8169-474a-4549-8098-3ab695640e3c"},
    {id: 1074, email: "rmbtw334@gmail.com", apikey: "apikey_1074", apiSecret:"c10d9d31-90cc-458d-b78e-d239083c4620"},
    {id: 1075, email: "brandyhigheroffice@gmail.com", apikey: "apikey_1075", apiSecret: "5d5369e4-6596-4139-a009-7cd31fc90807"},
    {id: 1076, email: "kingqk02.kk@gmail.com", apikey: "apikey_1076", apiSecret:"fee7399f-098e-4f39-960e-9754dc2591dd"},
    {id: 1077, email: "choundavid@gmail.com", apikey: "apikey_1077",apiSecret: "6d1d819d-df86-4931-8d9d-ad23cd19c37b"},
    {id: 1078, email: "sam690529@gmail.com", apikey: "apikey_1078", apiSecret:"db121b7d-04b8-4465-ae93-633630a7669c"},
    {id: 1079, email: "opp556687@gmail.com", apikey: "apikey_1079", apiSecret:"b9b3ce9a-9117-42ee-b768-3f58c42e2332"},
    {id: 1080, email: "crystalchang01@gmail.com", apikey: "apikey_1080", apiSecret:"6a97daca-b3eb-453a-9603-0bf498ad5889"},
    {id: 1081, email: "jolin6015@gmail.com", apikey: "apikey_1081",apiSecret: "44b6fd96-4041-44f7-871f-1595709b56f3"},
    {id: 1082, email: "k06120612@hotmail.com", apikey: "apikey_1082", apiSecret:"ff638017-d29e-43c1-9060-087199e2fe5d"},
    {id: 1083, email: "columnbb@gmail.com", apikey: "c187b96d93e20ecf1fdce2a32502696c", apiSecret:"$2a$10$nLtpNGXUw518NepzieA.Zeu5.eqqUr5fBdohz6"},
    {id: 1084, email: "aspire.wesley@gmail.com", apikey: "apikey_1084", apiSecret:"7ff7a920-322f-49e5-af4b-88be9ced12a4"},
    {id: 1085, email: "fubon1018@gmail.com", apikey: "apikey_1085", apiSecret: "196cf9f1-fd94-4a3b-8868-a774b7ab9c93"},
    {id: 1086, email: "s7013065789@gmail.com", apikey: "apikey_1086", apiSecret: "70430c49-082a-4775-84c6-cea87707f5f1"},
    {id: 1087, email: "makemoney13331@gmail.com", apikey: "apikey_1087",apiSecret: "ac3c70e6-e2b5-4c2c-a8da-560dd8c39c3a"},
    {id: 1088, email: "sonyjaihao@gmail.com", apikey: "apikey_1088", apiSecret:"1985211d-a2fa-43f9-a793-eca83436a152"},
    {id: 1089, email: "wei1227184@gmail.com", apikey: "apikey_1089", apiSecret:"912cd2a8-1239-4956-9291-4fad10e48984"},
    {id: 1090, email: "melodyaudiolee@gmail.com", apikey: "apikey_1090", apiSecret:"d9a188a5-8efa-4de6-a764-5298646fa8d4"},
    {id: 1091, email: "manpower.jerry71@gmail.com", apikey: "apikey_1091", apiSecret:"0cbee6e0-3a0e-4e11-baf8-f5b02780045e"},
    {id: 1092, email: "kk8383@msn.com", apikey: "apikey_1092", apiSecret:"4be30c50-5ba7-4e7d-bccc-8858f0b54c77"},
    {id: 1093, email: "mingnhsu@gmail.com", apikey: "apikey_1093",apiSecret: "890f3cdd-f04e-4380-8354-553d9f42a2f6"},
    {id: 1094, email: "dicky05932003@gmail.com", apikey: "apikey_1094", apiSecret:"61bf019b-779d-462a-b621-e2e13090d14d"},
    {id: 1095, email: "w24321123@gmail.com", apikey: "apikey_1095",apiSecret: "e76bd3ce-9254-4cf1-a942-80b57d9617ed"},
    {id: 1096, email: "x9454vup4@gmail.com", apikey: "apikey_1096", apiSecret:"dd40cbe6-9f24-411f-902a-3ca15e135a0a"},
]