const cluster = require('cluster');
const numCPUs = require('os').cpus().length;
const WebSocket = require('ws');

if (cluster.isMaster) {
    console.log(`Master ${process.pid} is running`);
  
    // Fork workers.
    for (let i = 0; i < numCPUs; i++) {
      cluster.fork();
    }
  
    cluster.on('exit', (worker, code, signal) => {
      console.log(`worker ${worker.process.pid} died`);
    });
  } else {
    const wss = new WebSocket.Server({ port: 8080 });
    wss.on('connection', function connection(ws, req) {
        ws.on('message', function incoming(data) {
            const ip = req.connection.remoteAddress;
            console.log(ip + ' sent: ' + data);

            ws.send(data);
        });
    });
  
    console.log(`Worker ${process.pid} started`);
  }



