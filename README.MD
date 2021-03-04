# DClock - a distributed scheduler service
> This is a sample project to demonstrate the usage of Rony framework. [https://github.com/ronaksoft/rony]

---
### Simple Setup
> STEP 1. Bring up the first server: <br />
`cli-dclock server --serverID R1 --bootstrap --dataPath ./_hdd/node1 --gatewayListen 0.0.0.0:81 --tunnelListen 0.0.0.0:91 --gossipPort 7081 --raftPort 7082 --replicaSet 1` <br />
this brings up the first bootstrap server   

> STEP 2. Bring up the second server: <br />
`cli-dclock server --serverID R2 --bootstrap --dataPath ./_hdd/node2 --gatewayListen 0.0.0.0:82 --tunnelListen 0.0.0.0:92 --gossipPort 8081 --raftPort 8082 --replicaSet 2 
> --join 127.0.0.1:7081` <br/>
this brings up the second server and joins the first bootstrap server

> STEP 3. Bring up the third server: <br />
`cli-dclock server --serverID R3 --bootstrap --dataPath ./_hdd/node3 --gatewayListen 0.0.0.0:83 --tunnelListen 0.0.0.0:93 --gossipPort 9081 --raftPort 9082 --replicaSet 3 
 --join 127.0.0.1:7081` <br />
this brings up the third server and joins the cluster

> STEP 4. Go to [https://webhook.site] and copy a webhook address for test <br/>

> STEP 5. Run the client: <br />
` cli-dclock client --hostPort 127.0.0.1:83` <br/>
you can connect to any of the three servers by choosing ports 81, 82 or 83. 

> STEP 6. In the interactive shell, you can set a hook. <br />
` hook-set --uniqueID SOME_UNIQUE --timestamp 10 --hookUrl <YOUR-WEBHOOK>` <br />
this command with setup a hook to be triggered in next 10s. You can check it on your webhook.site panel.

### Again, this is not a real project. Only to demonstrate how to write a distributed service using Rony framework.


